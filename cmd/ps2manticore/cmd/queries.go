package cmd

const (
	/*
			#### Manticore RT Schema:
			rt_attr_bigint = id_product
			rt_attr_bigint = price
			rt_attr_bigint = id_shop_default
			rt_attr_bigint = is_virtual
			rt_attr_string = link_rewrite
			rt_attr_bigint = active
			rt_attr_string = shopname
			rt_attr_bigint = shop_id
			rt_attr_bigint = id_image
			rt_attr_bigint = price_final
			rt_attr_bigint = nb_downloadable
			rt_attr_bigint = sav_quantity
			rt_attr_bigint = badge_danger
			rt_attr_multi  = features
			rt_attr_multi  = feature_values

		    rt_field          = name
		    rt_field          = reference
		    rt_field          = description
		    rt_field          = name_category
		    stored_fields     = name,reference,description,name_category
	*/
	sqlAdminCatalogProducts = `SELECT p.id_product AS id_product,
			 p.reference  AS reference,
			 sa.price  AS price,
			 p.id_shop_default  AS id_shop_default,
			 p.is_virtual  AS is_virtual,
			 pl.name  AS name,
			 pl.description  AS description,
			 pl.description_short  AS description_short,
			 pl.link_rewrite  AS link_rewrite,
			 sa.active  AS active,
	         m.id_manufacturer AS manufacturer_id, 
	         m.name AS manufacturer,
			 shop.id_shop  AS shop_id,
			 shop.name  AS shop_name,
			 image_shop.id_image  AS id_image,
			 cl.name  AS name_category,
			 0 AS price_final,
			 pd.nb_downloadable  AS nb_downloadable,
			 sav.quantity  AS sav_quantity,
			 IF(sav.quantity<=0, 1, 0) AS badge_danger,
			 GROUP_CONCAT(DISTINCT(fp.id_feature) SEPARATOR ",") as features,
			 GROUP_CONCAT(DISTINCT(fp.id_feature_value) SEPARATOR ",") as feature_values

			FROM  eg_product p
			 LEFT JOIN eg_product_lang pl ON (pl.id_product = p.id_product AND pl.id_lang = 1 AND pl.id_shop = 1)
			 LEFT JOIN eg_stock_available sav ON (sav.id_product = p.id_product AND sav.id_product_attribute = 0 AND sav.id_shop = 1  AND sav.id_shop_group = 0 )
			 JOIN eg_product_shop sa ON (p.id_product = sa.id_product AND sa.id_shop = 1)
			 LEFT JOIN eg_category_lang cl ON (sa.id_category_default = cl.id_category AND cl.id_lang = 1 AND cl.id_shop = 1)
			 LEFT JOIN eg_category c ON (c.id_category = cl.id_category)
			 LEFT JOIN eg_shop shop ON (shop.id_shop = 1)
			 LEFT JOIN eg_image_shop image_shop ON (image_shop.id_product = p.id_product AND image_shop.cover = 1 AND image_shop.id_shop = 1)
			 LEFT JOIN eg_image i ON (i.id_image = image_shop.id_image)
			 LEFT JOIN eg_product_download pd ON (pd.id_product = p.id_product)
             LEFT JOIN eg_feature_product fp ON p.id_product = fp.id_product
	         LEFT JOIN eg_manufacturer AS m ON p.id_manufacturer=m.id_manufacturer 
			WHERE (1 AND state = 1)
			GROUP BY p.id_product
			ORDER BY  id_product desc` // LIMIT {offset},{limit}
)

const (
	/*
			// todo. set mapping for columns
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

		    rt_field          = product
		    rt_field          = reference
		    rt_field          = description
		    rt_field          = name_category
		    stored_fields     = name,reference,description,name_category
	*/
	sqlProduct = `SELECT  pl.id_product, 
        pl.name AS product, 
        pl.description AS description, 
        ROUND(AVG(eg.price)) AS price, 
        eg.id_category_default AS category_id, 
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
	/*
			// todo. set mapping for columns
			#### Manticore RT Schema:
			rt_attr_timestamp = date_add
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

		    rt_field          = manufacturer_name
		    rt_field          = name
		    rt_field          = reference
		    rt_field          = description
		    rt_field          = description_short
		    rt_field          = categories
		    stored_fields     = name,reference,description,description_short,categories,manufacturer_name
	*/
	sqlProductAllWithCategories = `
SELECT p.id_product, pl.name, GROUP_CONCAT(DISTINCT(cl.name) SEPARATOR ",") as categories, p.price, p.reference, pm.name manufacturer_name, pl.description_short, pl.description, p.date_add
FROM eg_product p
LEFT JOIN eg_product_lang pl ON (p.id_product = pl.id_product)
LEFT JOIN eg_category_product cp ON (p.id_product = cp.id_product)
LEFT JOIN eg_category_lang cl ON (cp.id_category = cl.id_category)
LEFT JOIN eg_category c ON (cp.id_category = c.id_category)
LEFT JOIN eg_manufacturer pm ON (p.id_manufacturer = pm.id_manufacturer)
WHERE pl.id_lang = {id_lang}
AND cl.id_lang = {id_lang}
AND p.id_shop_default = {id_shop}
AND c.id_shop_default = {id_shop}
GROUP BY p.id_product` // LIMIT {offset},{limit}
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
