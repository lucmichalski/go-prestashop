package cmd

import (
    "time"
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
    GROUP BY pl.id_product, pl.name`  // LIMIT {offset},{limit}

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

type Product struct {
    DateAdd           time.Time `json:"date_add"`
    DateUpd           time.Time `json:"date_upd"`
    ShopId            string    `json:"shop_id"`
    LangId            string    `json:"lang_id"`
    CategoryId        uint      `json:"category_id"`
    CategoryLink      string    `json:"category_link"`
    ManufacturerId    uint      `json:"manufacturer_id"`
    Afo               string    `json:"afo"`
    AvailableForOrder string    `json:"available_for_order"`
    Price             int       `json:"price"`
    Product           string    `json:"product"`
    Category          string    `json:"category"`
    Manufacturer      string    `json:"manufacturer"`
    IdProduct         uint      `json:"id_product"`
    LinkRewrite       string    `json:"link_rewrite"`
    Name              string    `json:"name"`
    Description       string    `json:"description"`
}

type Category struct {
    ShopId       uint   `json:"shop_id"`
    LangId       uint   `json:"lang_id"`
    ParentId     uint   `json:"parent_id"`
    CategoryLink string `json:"category_link"`
    Category     string `json:"category"`
    Description  string `json:"description"`
    MetaTitle    string `json:"meta_title"`
}
