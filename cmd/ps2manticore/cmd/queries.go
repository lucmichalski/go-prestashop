package cmd

const (
	/*
			#### Based on SQL query:

			SELECT SQL_CALC_FOUND_ROWS p.`id_product`  AS `id_product`,
			 p.`reference`  AS `reference`,
			 sa.`price`  AS `price`,
			 p.`id_shop_default`  AS `id_shop_default`,
			 p.`is_virtual`  AS `is_virtual`,
			 pl.`name`  AS `name`,
			 pl.`link_rewrite`  AS `link_rewrite`,
			 sa.`active`  AS `active`,
			 shop.`name`  AS `shopname`,
			 image_shop.`id_image`  AS `id_image`,
			 cl.`name`  AS `name_category`,
			 0 AS `price_final`,
			 pd.`nb_downloadable`  AS `nb_downloadable`,
			 sav.`quantity`  AS `sav_quantity`,
			 IF(sav.`quantity`<=0, 1, 0) AS `badge_danger`
			FROM  `eg_product` p
			 LEFT JOIN `eg_product_lang` pl ON (pl.`id_product` = p.`id_product` AND pl.`id_lang` = 1 AND pl.`id_shop` = 1)
			 LEFT JOIN `eg_stock_available` sav ON (sav.`id_product` = p.`id_product` AND sav.`id_product_attribute` = 0 AND sav.id_shop = 1  AND sav.id_shop_group = 0 )
			 JOIN `eg_product_shop` sa ON (p.`id_product` = sa.`id_product` AND sa.id_shop = 1)
			 LEFT JOIN `eg_category_lang` cl ON (sa.`id_category_default` = cl.`id_category` AND cl.`id_lang` = 1 AND cl.id_shop = 1)
			 LEFT JOIN `eg_category` c ON (c.`id_category` = cl.`id_category`)
			 LEFT JOIN `eg_shop` shop ON (shop.id_shop = 1)
			 LEFT JOIN `eg_image_shop` image_shop ON (image_shop.`id_product` = p.`id_product` AND image_shop.`cover` = 1 AND image_shop.id_shop = 1)
			 LEFT JOIN `eg_image` i ON (i.`id_image` = image_shop.`id_image`)
			 LEFT JOIN `eg_product_download` pd ON (pd.`id_product` = p.`id_product`)
			WHERE (1 AND state = 1)
			ORDER BY  `id_product` desc
			LIMIT 0, 20;

			#### Fields:
			- id_product
			- reference
			- price
			- id_shop_default
			- is_virtual
			- name
			- description (extra)
			- link_rewrite
			- active
			- shopname
			- id_image
			- name_category
			- price_final
			- nb_downloadable
			- quantity

			#### Manticore RT Schema:
			rt_attr_bigint = id_product
			rt_attr_bigint = price
			rt_attr_bigint = id_shop_default
			rt_attr_bigint = is_virtual
			rt_attr_string = link_rewrite
			rt_attr_bigint = active
			rt_attr_string = shopname
			rt_attr_bigint = id_image
			rt_attr_bigint = price_final
			rt_attr_bigint = nb_downloadable
			rt_attr_bigint = quantity

		    rt_field          = name
		    rt_field          = reference
		    rt_field          = description
		    rt_field          = name_category
		    stored_fields     = name,reference,description,name_category

			### Escaped backtick "`" with gowrap because reference is a mysql command
	*/
	sqlAdminCatalogProducts = `SELECT SQL_CALC_FOUND_ROWS 
 p.` +
		"`" +
		`id_product` +
		"`" +
		`  AS ` +
		"`" +
		`id_product` +
		"`" +
		`,
 p.` +
		"`" +
		`reference` +
		"`" +
		`  AS ` +
		"`" +
		`reference` +
		"`" +
		`,
 sa.` +
		"`" +
		`price` +
		"`" +
		`  AS ` +
		"`" +
		`price` +
		"`" +
		`,
 p.` +
		"`" +
		`id_shop_default` +
		"`" +
		`  AS ` +
		"`" +
		`id_shop_default` +
		"`" +
		`,
 p.` +
		"`" +
		`is_virtual` +
		"`" +
		`  AS ` +
		"`" +
		`is_virtual` +
		"`" +
		`,
 pl.` +
		"`" +
		`name` +
		"`" +
		`  AS ` +
		"`" +
		`name` +
		"`" +
		`,
 pl.` +
		"`" +
		`description` +
		"`" +
		`  AS ` +
		"`" +
		`description` +
		"`" +
		`,
 pl.` +
		"`" +
		`link_rewrite` +
		"`" +
		`  AS ` +
		"`" +
		`link_rewrite` +
		"`" +
		`,
 sa.` +
		"`" +
		`active` +
		"`" +
		`  AS ` +
		"`" +
		`active` +
		"`" +
		`,
 shop.` +
		"`" +
		`name` +
		"`" +
		`  AS ` +
		"`" +
		`shopname` +
		"`" +
		`,
 image_shop.` +
		"`" +
		`id_image` +
		"`" +
		`  AS ` +
		"`" +
		`id_image` +
		"`" +
		`,
 cl.` +
		"`" +
		`name` +
		"`" +
		`  AS ` +
		"`" +
		`name_category` +
		"`" +
		`,
 0 AS ` +
		"`" +
		`price_final` +
		"`" +
		`,
 pd.` +
		"`" +
		`nb_downloadable` +
		"`" +
		`  AS ` +
		"`" +
		`nb_downloadable` +
		"`" +
		`,
 sav.` +
		"`" +
		`quantity` +
		"`" +
		`  AS ` +
		"`" +
		`sav_quantity` +
		"`" +
		`,
 IF(sav.` +
		"`" +
		`quantity` +
		"`" +
		`<=0, 1, 0) AS ` +
		"`" +
		`badge_danger` +
		"`" +
		`
FROM  ` +
		"`" +
		`eg_product` +
		"`" +
		` p
 LEFT JOIN ` +
		"`" +
		`eg_product_lang` +
		"`" +
		` pl ON (pl.` +
		"`" +
		`id_product` +
		"`" +
		` = p.` +
		"`" +
		`id_product` +
		"`" +
		` AND pl.` +
		"`" +
		`id_lang` +
		"`" +
		` = {id_lang} AND pl.` +
		"`" +
		`id_shop` +
		"`" +
		` = {id_shop})
 LEFT JOIN ` +
		"`" +
		`eg_stock_available` +
		"`" +
		` sav ON (sav.` +
		"`" +
		`id_product` +
		"`" +
		` = p.` +
		"`" +
		`id_product` +
		"`" +
		` AND sav.` +
		"`" +
		`id_product_attribute` +
		"`" +
		` = 0 AND sav.id_shop = {id_shop}  AND sav.id_shop_group = 0 )
 JOIN ` +
		"`" +
		`eg_product_shop` +
		"`" +
		` sa ON (p.` +
		"`" +
		`id_product` +
		"`" +
		` = sa.` +
		"`" +
		`id_product` +
		"`" +
		` AND sa.id_shop = {id_shop})
 LEFT JOIN ` +
		"`" +
		`eg_category_lang` +
		"`" +
		` cl ON (sa.` +
		"`" +
		`id_category_default` +
		"`" +
		` = cl.` +
		"`" +
		`id_category` +
		"`" +
		` AND cl.` +
		"`" +
		`id_lang` +
		"`" +
		` = 1 AND cl.id_shop = {id_shop})
 LEFT JOIN ` +
		"`" +
		`eg_category` +
		"`" +
		` c ON (c.` +
		"`" +
		`id_category` +
		"`" +
		` = cl.` +
		"`" +
		`id_category` +
		"`" +
		`)
 LEFT JOIN ` +
		"`" +
		`eg_shop` +
		"`" +
		` shop ON (shop.id_shop = {id_shop})
 LEFT JOIN ` +
		"`" +
		`eg_image_shop` +
		"`" +
		` image_shop ON (image_shop.` +
		"`" +
		`id_product` +
		"`" +
		` = p.` +
		"`" +
		`id_product` +
		"`" +
		` AND image_shop.` +
		"`" +
		`cover` +
		"`" +
		` = 1 AND image_shop.id_shop = {id_shop})
 LEFT JOIN ` +
		"`" +
		`eg_image` +
		"`" +
		` i ON (i.` +
		"`" +
		`id_image` +
		"`" +
		` = image_shop.` +
		"`" +
		`id_image` +
		"`" +
		`)
 LEFT JOIN ` +
		"`" +
		`eg_product_download` +
		"`" +
		` pd ON (pd.` +
		"`" +
		`id_product` +
		"`" +
		` = p.` +
		"`" +
		`id_product` +
		"`" +
		`)
WHERE (1 AND state = 1)

ORDER BY  ` +
		"`" +
		`id_product` +
		"`" +
		` desc

LIMIT {offset}, {limit};`
)

const (
	sqlProduct = `SELECT  pl.id_product, 
        pl.name AS product, 
        pl.description AS description, 
        ROUND(AVG(eg.price)) AS price, 
        eg.id_category_default AS categoryId, 
        GROUP_CONCAT(eg.available_for_order ORDER BY eg.id_shop ASC SEPARATOR ',') AS afo, 
        CONV(REVERSE(GROUP_CONCAT(eg.available_for_order ORDER BY eg.id_shop ASC SEPARATOR '')),2,10) AS available_for_order, 
        pl.id_lang AS langId_attr, 
        cl.link_rewrite AS categoryLink, 
        cl.name AS category, 
        m.id_manufacturer AS manufacturerId, m.name AS manufacturer, 
        GROUP_CONCAT(DISTINCT pl.id_shop) AS shopId_attr,
        p.date_add,
        p.date_upd
    FROM eg_product_lang AS pl 
        INNER JOIN eg_product_shop AS eg ON pl.id_product=eg.id_product AND pl.id_shop=eg.id_shop AND eg.active=1
        INNER JOIN eg_category c ON eg.id_category_default=c.id_category AND c.active=1 
        LEFT JOIN eg_category_lang AS cl ON eg.id_category_default=cl.id_category AND pl.id_shop=cl.id_shop AND pl.id_lang=cl.id_lang 
        LEFT JOIN eg_product AS p ON pl.id_product=p.id_product 
        LEFT JOIN eg_manufacturer AS m ON p.id_manufacturer=m.id_manufacturer 
    WHERE pl.id_lang=1 AND eg.indexed=1 AND eg.visibility IN ('both','search') 
    GROUP BY pl.id_product, pl.name` // LIMIT {offset},{limit}
)

const (
	sqlCategory = `SELECT  cl.id_category, c.id_parent AS parentId,
        GROUP_CONCAT(DISTINCT cl.id_lang ORDER BY cl.id_lang ASC) AS langId_attr,
        GROUP_CONCAT(DISTINCT cl.link_rewrite ORDER BY cl.id_lang ASC SEPARATOR '|') AS categoryLink,
        GROUP_CONCAT(distinct cl.name ORDER BY cl.id_lang ASC SEPARATOR '|') AS category,
        GROUP_CONCAT(DISTINCT cl.id_shop) AS shopId_attr
    FROM eg_category_lang AS cl
    LEFT JOIN eg_category AS c ON cl.id_category=c.id_category
    LEFT JOIN eg_lang l ON l.id_lang = cl.id_lang
    WHERE c.id_parent>1 AND c.active=1
    GROUP BY cl.id_category`
)
