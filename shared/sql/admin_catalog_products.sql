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
 
LIMIT 0, 20
;
