package cmd

const (
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
			 cl.id_category AS id_category,
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
			ORDER BY  id_product desc`
) // LIMIT {{ .offset }},{{ .limit }}

const (
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
    GROUP BY pl.id_product, pl.name`
) // LIMIT {{ .offset }},{{ .limit }}

const (
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
GROUP BY p.id_product`
) // LIMIT {{ .offset }},{{ .limit }}

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
) // LIMIT {{ .offset }},{{ .limit }}

const (
	sqlProductExtended = `SELECT 	p.reference,
									pl.name,
									GROUP_CONCAT(DISTINCT(al.name) SEPARATOR ", ") AS "combination",
									s.quantity,
									p.price,
									pa.price,
									p.wholesale_price,
									GROUP_CONCAT(DISTINCT(cl.name) SEPARATOR ",") AS "product_groups",
									p.weight,
									p.id_tax_rules_group,
									pa.reference,
									pl.description_short,
									pl.description,
									pl.meta_title,
									pl.meta_keywords,
									pl.meta_description,
									pl.link_rewrite,
									pl.available_now,
									pl.available_later,
									p.available_for_order,
									p.date_add,
									p.show_price,
									p.online_only,
									pa.ean13 AS ean13,
									GROUP_CONCAT(DISTINCT(CONCAT("https://prestanish.evolutive.group", /*This part get the images of the product and display all of them in 1 cell, separated by ´|´ the vertical bar*/
									    "/img/p/",
									    IF(CHAR_LENGTH(pi.id_image) >= 5, 
									        CONCAT(
									        SUBSTRING(pi.id_image, -5, 1),
									        "/"),
									        ""),
									    IF(CHAR_LENGTH(pi.id_image) >= 4, CONCAT(SUBSTRING(pi.id_image, -4, 1), "/"), ""),
									    IF(CHAR_LENGTH(pi.id_image) >= 3, CONCAT(SUBSTRING(pi.id_image, -3, 1), "/"), ""),
									    IF(CHAR_LENGTH(pi.id_image) >= 2, CONCAT(SUBSTRING(pi.id_image, -2, 1), "/"), ""),
									    IF(CHAR_LENGTH(pi.id_image) >= 1, CONCAT(SUBSTRING(pi.id_image, -1, 1), "/"), ""),
									    pi.id_image,
									    ".jpg")) SEPARATOR " | ") AS "Images"
									    FROM
									eg_product p
									LEFT JOIN eg_product_lang pl ON (p.id_product = pl.id_product and pl.id_lang=1)
									LEFT JOIN eg_manufacturer m ON (p.id_manufacturer = m.id_manufacturer)
									LEFT JOIN eg_category_product cp ON (p.id_product = cp.id_product)
									LEFT JOIN eg_category c ON (cp.id_category = c.id_category)
									LEFT JOIN eg_category_lang cl ON (cp.id_category = cl.id_category and cl.id_lang=1)
									LEFT JOIN eg_product_attribute pa ON (p.id_product = pa.id_product)
									LEFT JOIN eg_stock_available s ON (p.id_product = s.id_product and (pa.id_product_attribute=s.id_product_attribute or pa.id_product_attribute is null))
									LEFT JOIN eg_product_tag pt ON (p.id_product = pt.id_product)
									LEFT JOIN eg_product_attribute_combination pac ON (pac.id_product_attribute = pa.id_product_attribute)
									LEFT JOIN eg_attribute_lang al ON (al.id_attribute = pac.id_attribute and al.id_lang=1)
									LEFT JOIN eg_image pi ON (p.id_product = pi.id_product)
									GROUP BY p.id_product,pac.id_product_attribute order by p.id_product`
) // LIMIT {{ .offset }},{{ .limit }}

const (
	sqlOrdersExtended = `SELECT  
	    a.id_order AS id_order,
	    a.id_currency AS id_currency,
		a.id_order AS id_pdf,
		a.id_shop AS id_shop,
		a.id_lang AS id_lang,
		a.payment AS order_payment,
		a.module AS module_payment,
		a.total_paid AS total_paid,
		a.total_discounts AS total_discounts,
		a.total_discounts_tax_incl AS total_discounts_tax_incl,
		a.total_discounts_tax_excl AS total_discounts_tax_excl,
		a.total_paid_tax_incl AS total_paid_tax_incl,
		a.total_paid_tax_excl AS total_paid_tax_excl,
		a.total_paid_real AS total_paid_real,
		a.total_products AS total_products,
		a.total_products_wt AS total_products_wt,
		a.total_shipping AS total_shipping,
		a.total_shipping_tax_incl AS total_shipping_tax_incl,
		a.total_shipping_tax_excl AS total_shipping_tax_excl,
		a.carrier_tax_rate AS carrier_tax_rate,
		a.total_wrapping AS total_wrapping,
		a.total_wrapping_tax_incl AS total_wrapping_tax_incl,
		a.total_wrapping_tax_excl AS total_wrapping_tax_excl,
		a.round_mode AS round_mode,
		a.round_type AS round_type,
		a.invoice_number AS invoice_number,
		a.delivery_number AS delivery_number,
		a.invoice_date AS invoice_date,
		a.delivery_date AS delivery_date,
		a.valid AS valid,
		a.date_add AS date_add,
		a.date_upd AS date_upd,
		a.carrier_tax_rate,
		a.id_address_delivery AS id_address_delivery,
		CONCAT(LEFT(c.firstname, 1), '. ', c.lastname) AS customer_name,
		osl.name AS order_status_name,
		os.color AS order_status_color,
		IF((SELECT so.id_order FROM eg_orders so WHERE so.id_customer = a.id_customer AND so.id_order < a.id_order LIMIT 1) > 0, 0, 1) AS order_new,
		country.id_country AS id_country,
		country_lang.name AS country_name,
		address.company AS company,
		address.vat_number AS vat_number,
		address.phone_mobile AS phone_mobile,
		address.phone AS phone,
		address.lastname AS delivery_lastname,
		address.firstname AS delivery_firstname,
		address.address1 AS address_1,	
		address.address2 AS address_2,			
		address.postcode AS postcode,
		address.city AS city,
		IF(a.valid, 1, 0) AS badge_success
		FROM eg_orders a
		LEFT JOIN eg_customer c ON (c.id_customer = a.id_customer)
		INNER JOIN eg_address address ON address.id_address = a.id_address_delivery
		INNER JOIN eg_country country ON address.id_country = country.id_country
		INNER JOIN eg_country_lang country_lang ON (country.id_country = country_lang.id_country AND country_lang.id_lang = 1)
		LEFT JOIN eg_order_state os ON (os.id_order_state = a.current_state)
		LEFT JOIN eg_order_state_lang osl ON (os.id_order_state = osl.id_order_state AND osl.id_lang = 1)
		ORDER BY a.id_order ASC  
		LIMIT {{ .offset }},{{ .limit }}`
)

const (
	sqlCustomerExtended = `
		SELECT 	c.id_customer AS id_customer, 
				c.firstname AS firstname, 
				c.lastname AS lastname, 
				CONCAT(c.firstname, ' ', c.lastname) AS fullname,				
				c.email AS email, 
				c.active AS active, 
				c.deleted AS deleted,
				c.note as note,
				c.newsletter AS newsletter, 
				c.optin AS optin, 
				c.ip_registration_newsletter AS ip_registration_newsletter,
				c.birthday AS birthday, 
				c.date_add AS date_add, 
				gl.name as social_title,
				ecg.id_group as id_group,
				egl.name as group_name,
				s.name as shop_name, 
				c.company as company, 
				(SELECT SUM(total_paid_real / conversion_rate) FROM eg_orders o WHERE (o.id_customer = c.id_customer) AND (o.id_shop IN ('1')) AND (o.valid = 1)) as total_spent, 
				(SELECT con.date_add FROM eg_guest g LEFT JOIN eg_connections con ON con.id_guest = g.id_guest WHERE g.id_customer = c.id_customer ORDER BY con.date_add DESC LIMIT 1) as connect 
			FROM eg_customer c 
				LEFT JOIN eg_customer_group ecg ON ecg.id_customer = c.id_customer
				LEFT JOIN eg_group_lang egl ON egl.id_group = ecg.id_group
				LEFT JOIN eg_gender_lang gl ON c.id_gender = gl.id_gender AND gl.id_lang = 1 
				LEFT JOIN eg_shop s ON c.id_shop = s.id_shop 
			WHERE (c.deleted = 0) AND (c.id_shop IN ('1')
		) 
		ORDER BY c.date_add DESC
		LIMIT {{ .offset }},{{ .limit }}`
)
