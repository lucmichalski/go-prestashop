DROP PROCEDURE IF EXISTS repair_nested_tree;
DELIMITER //
CREATE PROCEDURE repair_nested_tree ()
MODIFIES SQL DATA
BEGIN
    DECLARE currentId, currentParentId  INT;
    DECLARE currentLeft                 INT;
    DECLARE startId                     INT DEFAULT 1;
    # Determines the max size for MEMORY tables.
    SET max_heap_table_size = 1024 * 1024 * 512;
    START TRANSACTION;
    # Temporary MEMORY table to do all the heavy lifting in,
    # otherwise performance is simply abysmal.
    CREATE TABLE `tmp_category_tree` (
        `id_category` int(10) unsigned NOT NULL DEFAULT 0,
        `id_parent`   int(10)          DEFAULT NULL,
        `nleft`       int(10) unsigned DEFAULT NULL,
        `nright`      int(10) unsigned DEFAULT NULL,
        PRIMARY KEY      (`id_category`),
        INDEX USING HASH (`id_parent`),
        INDEX USING HASH (`nleft`),
        INDEX USING HASH (`nright`)
    ) ENGINE = MEMORY
    SELECT `id_category`,
           `id_parent`,
           `nleft`,
           `nright`
    FROM   `eg_category`;
    # Leveling the playing field.
    UPDATE  `tmp_category_tree`
    SET     `nleft`  = NULL,
            `nright` = NULL;
    # Establishing starting numbers for all root elements.
    WHILE EXISTS (SELECT * FROM `tmp_category_tree` WHERE `id_parent` = 0 AND `nleft` IS NULL AND `nright` IS NULL LIMIT 1) DO
        UPDATE `tmp_category_tree`
        SET    `nleft`  = startId,
               `nright` = startId + 1
        WHERE  `id_parent` = 0
          AND  `nleft`       IS NULL
          AND  `nright`      IS NULL
        LIMIT  1;
        SET startId = startId + 2;
    END WHILE;
    # Switching the indexes for the lft/rght columns to B-Trees to speed up the next section, which uses range queries.
    DROP INDEX `nleft`  ON `tmp_category_tree`;
    DROP INDEX `nright` ON `tmp_category_tree`;
    CREATE INDEX `nleft`  USING BTREE ON `tmp_category_tree` (`nleft`);
    CREATE INDEX `nright` USING BTREE ON `tmp_category_tree` (`nright`);
    # Numbering all child elements
    WHILE EXISTS (SELECT * FROM `tmp_category_tree` WHERE `nleft` IS NULL LIMIT 1) DO
        # Picking an unprocessed element which has a processed parent.
        SELECT     `tmp_category_tree`.`id_category`
          INTO     currentId
        FROM       `tmp_category_tree`
        INNER JOIN `tmp_category_tree` AS `parents`
                ON `tmp_category_tree`.`id_parent` = `parents`.`id_category`
        WHERE      `tmp_category_tree`.`nleft` IS NULL
          AND      `parents`.`nleft`  IS NOT NULL
        LIMIT      1;
        # Finding the element's parent.
        SELECT  `id_parent`
          INTO  currentParentId
        FROM    `tmp_category_tree`
        WHERE   `id_category` = currentId;
        # Finding the parent's nleft value.
        SELECT  `nleft`
          INTO  currentLeft
        FROM    `tmp_category_tree`
        WHERE   `id_category` = currentParentId;
        # Shifting all elements to the right of the current element 2 to the right.
        UPDATE `tmp_category_tree`
        SET    `nright` = `nright` + 2
        WHERE  `nright` > currentLeft;
        UPDATE `tmp_category_tree`
        SET    `nleft` = `nleft` + 2
        WHERE  `nleft` > currentLeft;
        # Setting nleft and nright values for current element.
        UPDATE `tmp_category_tree`
        SET    `nleft`  = currentLeft + 1,
               `nright` = currentLeft + 2
        WHERE  `id_category`   = currentId;
    END WHILE;
    # Writing calculated values back to physical table.
    UPDATE `eg_category`, `tmp_category_tree`
    SET    `eg_category`.`nleft`  = `tmp_category_tree`.`nleft`,
           `eg_category`.`nright` = `tmp_category_tree`.`nright`
    WHERE  `eg_category`.`id_category`   = `tmp_category_tree`.`id_category`;
    COMMIT;
    DROP TABLE `tmp_category_tree`;
END//
DELIMITER ;
