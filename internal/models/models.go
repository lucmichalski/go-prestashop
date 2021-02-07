package models

import (
	"time"

	"gorm.io/datatypes"
)

type Access struct {
	IDProfile           uint32 `gorm:"primaryKey;column:id_profile;type:int(10) unsigned;not null"`
	IDAuthorizationRole uint32 `gorm:"primaryKey;column:id_authorization_role;type:int(10) unsigned;not null"`
}

type Accessory struct {
	IDProduct1 uint32 `gorm:"index:accessory_product;column:id_product_1;type:int(10) unsigned;not null"`
	IDProduct2 uint32 `gorm:"index:accessory_product;column:id_product_2;type:int(10) unsigned;not null"`
}

type Address struct {
	IDAddress      uint32    `gorm:"primaryKey;column:id_address;type:int(10) unsigned;not null"`
	IDCountry      uint32    `gorm:"index:id_country;column:id_country;type:int(10) unsigned;not null"`
	IDState        uint32    `gorm:"index:id_state;column:id_state;type:int(10) unsigned"`
	IDCustomer     uint32    `gorm:"index:address_customer;column:id_customer;type:int(10) unsigned;not null;default:0"`
	IDManufacturer uint32    `gorm:"index:id_manufacturer;column:id_manufacturer;type:int(10) unsigned;not null;default:0"`
	IDSupplier     uint32    `gorm:"index:id_supplier;column:id_supplier;type:int(10) unsigned;not null;default:0"`
	IDWarehouse    uint32    `gorm:"index:id_warehouse;column:id_warehouse;type:int(10) unsigned;not null;default:0"`
	Alias          string    `gorm:"column:alias;type:varchar(32);not null"`
	Company        string    `gorm:"column:company;type:varchar(255)"`
	Lastname       string    `gorm:"column:lastname;type:varchar(255);not null"`
	Firstname      string    `gorm:"column:firstname;type:varchar(255);not null"`
	Address1       string    `gorm:"column:address1;type:varchar(128);not null"`
	Address2       string    `gorm:"column:address2;type:varchar(128)"`
	Postcode       string    `gorm:"column:postcode;type:varchar(12)"`
	City           string    `gorm:"column:city;type:varchar(64);not null"`
	Other          string    `gorm:"column:other;type:text"`
	Phone          string    `gorm:"column:phone;type:varchar(32)"`
	PhoneMobile    string    `gorm:"column:phone_mobile;type:varchar(32)"`
	VatNumber      string    `gorm:"column:vat_number;type:varchar(32)"`
	Dni            string    `gorm:"column:dni;type:varchar(16)"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd        time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Active         bool      `gorm:"column:active;type:tinyint(1) unsigned;not null;default:1"`
	Deleted        bool      `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type AddressFormat struct {
	IDCountry uint32 `gorm:"primaryKey;column:id_country;type:int(10) unsigned;not null"`
	Format    string `gorm:"column:format;type:varchar(255);not null;default:''"`
}

type AdminFilter struct {
	ID         int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	Employee   int    `gorm:"uniqueIndex:admin_filter_search_id_idx;column:employee;type:int(11);not null"`
	Shop       int    `gorm:"uniqueIndex:admin_filter_search_id_idx;column:shop;type:int(11);not null"`
	Controller string `gorm:"uniqueIndex:admin_filter_search_id_idx;column:controller;type:varchar(60);not null"`
	Action     string `gorm:"uniqueIndex:admin_filter_search_id_idx;column:action;type:varchar(100);not null"`
	Filter     string `gorm:"column:filter;type:longtext;not null"`
	FilterID   string `gorm:"uniqueIndex:admin_filter_search_id_idx;column:filter_id;type:varchar(191);not null"`
}

type Advice struct {
	IDAdvice   int    `gorm:"primaryKey;column:id_advice;type:int(11);not null"`
	IDPsAdvice int    `gorm:"column:id_ps_advice;type:int(11);not null"`
	IDTab      int    `gorm:"column:id_tab;type:int(11);not null"`
	IDsTab     string `gorm:"column:ids_tab;type:text"`
	Validated  bool   `gorm:"column:validated;type:tinyint(1) unsigned;not null;default:0"`
	Hide       bool   `gorm:"column:hide;type:tinyint(1);not null;default:0"`
	Location   string `gorm:"column:location;type:enum('after','before');not null"`
	Selector   string `gorm:"column:selector;type:varchar(255)"`
	StartDay   int    `gorm:"column:start_day;type:int(11);not null;default:0"`
	StopDay    int    `gorm:"column:stop_day;type:int(11);not null;default:0"`
	Weight     int    `gorm:"column:weight;type:int(11);default:1"`
}

type AdviceLang struct {
	IDAdvice int    `gorm:"primaryKey;column:id_advice;type:int(11);not null"`
	IDLang   int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	HTML     string `gorm:"column:html;type:text"`
}

type Alias struct {
	IDAlias uint32 `gorm:"primaryKey;column:id_alias;type:int(10) unsigned;not null"`
	Alias   string `gorm:"unique;column:alias;type:varchar(191);not null"`
	Search  string `gorm:"column:search;type:varchar(255);not null"`
	Active  bool   `gorm:"column:active;type:tinyint(1);not null;default:1"`
}

type Attachment struct {
	IDAttachment uint32 `gorm:"primaryKey;column:id_attachment;type:int(10) unsigned;not null"`
	File         string `gorm:"column:file;type:varchar(40);not null"`
	FileName     string `gorm:"column:file_name;type:varchar(128);not null"`
	FileSize     uint64 `gorm:"column:file_size;type:bigint(10) unsigned;not null;default:0"`
	Mime         string `gorm:"column:mime;type:varchar(128);not null"`
}

type AttachmentLang struct {
	IDAttachment uint32 `gorm:"primaryKey;column:id_attachment;type:int(10) unsigned;not null"`
	IDLang       uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name         string `gorm:"column:name;type:varchar(32)"`
	Description  string `gorm:"column:description;type:text"`
}

type Attribute struct {
	IDAttribute      int    `gorm:"primaryKey;column:id_attribute;type:int(11);not null"`
	IDAttributeGroup int    `gorm:"index:attribute_group;column:id_attribute_group;type:int(11);not null"`
	Color            string `gorm:"column:color;type:varchar(32);not null"`
	Position         int    `gorm:"column:position;type:int(11);not null"`
}

type AttributeGroup struct {
	IDAttributeGroup int    `gorm:"primaryKey;column:id_attribute_group;type:int(11);not null"`
	IsColorGroup     bool   `gorm:"column:is_color_group;type:tinyint(1);not null"`
	GroupType        string `gorm:"column:group_type;type:varchar(255);not null"`
	Position         int    `gorm:"column:position;type:int(11);not null"`
}

type AttributeGroupLang struct {
	IDAttributeGroup int    `gorm:"primaryKey;index:IDX_2958662667A664FB;column:id_attribute_group;type:int(11);not null"`
	IDLang           int    `gorm:"primaryKey;index:IDX_29586626BA299860;column:id_lang;type:int(11);not null"`
	Name             string `gorm:"column:name;type:varchar(128);not null"`
	PublicName       string `gorm:"column:public_name;type:varchar(64);not null"`
}

type AttributeGroupShop struct {
	IDAttributeGroup int `gorm:"primaryKey;index:IDX_B43BAEE667A664FB;column:id_attribute_group;type:int(11);not null"`
	IDShop           int `gorm:"primaryKey;index:IDX_B43BAEE6274A50A0;column:id_shop;type:int(11);not null"`
}

type AttributeImpact struct {
	IDAttributeImpact uint32  `gorm:"primaryKey;column:id_attribute_impact;type:int(10) unsigned;not null"`
	IDProduct         uint32  `gorm:"uniqueIndex:id_product;column:id_product;type:int(11) unsigned;not null"`
	IDAttribute       uint32  `gorm:"uniqueIndex:id_product;column:id_attribute;type:int(11) unsigned;not null"`
	Weight            float64 `gorm:"column:weight;type:decimal(20,6);not null"`
	Price             float64 `gorm:"column:price;type:decimal(20,6);not null"`
}

type AttributeLang struct {
	IDAttribute int    `gorm:"primaryKey;index:IDX_EB57414F7A4F53DC;column:id_attribute;type:int(11);not null"`
	IDLang      int    `gorm:"primaryKey;index:IDX_EB57414FBA299860;column:id_lang;type:int(11);not null"`
	Name        string `gorm:"column:name;type:varchar(128);not null"`
}

type AttributeShop struct {
	IDAttribute int `gorm:"primaryKey;index:IDX_7634898F7A4F53DC;column:id_attribute;type:int(11);not null"`
	IDShop      int `gorm:"primaryKey;index:IDX_7634898F274A50A0;column:id_shop;type:int(11);not null"`
}

type AuthorizationRole struct {
	IDAuthorizationRole uint32 `gorm:"primaryKey;column:id_authorization_role;type:int(10) unsigned;not null"`
	Slug                string `gorm:"unique;column:slug;type:varchar(191);not null"`
}

type Badge struct {
	IDBadge       int    `gorm:"primaryKey;column:id_badge;type:int(11);not null"`
	IDPsBadge     int    `gorm:"column:id_ps_badge;type:int(11);not null"`
	Type          string `gorm:"column:type;type:varchar(32);not null"`
	IDGroup       int    `gorm:"column:id_group;type:int(11);not null"`
	GroupPosition int    `gorm:"column:group_position;type:int(11);not null"`
	Scoring       int    `gorm:"column:scoring;type:int(11);not null"`
	Awb           int    `gorm:"column:awb;type:int(11);default:0"`
	Validated     bool   `gorm:"column:validated;type:tinyint(1) unsigned;not null;default:0"`
}

type BadgeLang struct {
	IDBadge     int    `gorm:"primaryKey;column:id_badge;type:int(11);not null"`
	IDLang      int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	Name        string `gorm:"column:name;type:varchar(64)"`
	Description string `gorm:"column:description;type:varchar(255)"`
	GroupName   string `gorm:"column:group_name;type:varchar(255)"`
}

type Carrier struct {
	IDCarrier          uint32  `gorm:"primaryKey;column:id_carrier;type:int(10) unsigned;not null"`
	IDReference        uint32  `gorm:"index:reference;column:id_reference;type:int(10) unsigned;not null"`
	IDTaxRulesGroup    uint32  `gorm:"index:id_tax_rules_group;column:id_tax_rules_group;type:int(10) unsigned;default:0"`
	Name               string  `gorm:"column:name;type:varchar(64);not null"`
	URL                string  `gorm:"column:url;type:varchar(255)"`
	Active             bool    `gorm:"index:deleted;index:reference;column:active;type:tinyint(1) unsigned;not null;default:0"`
	Deleted            bool    `gorm:"index:deleted;index:reference;column:deleted;type:tinyint(1) unsigned;not null;default:0"`
	ShippingHandling   bool    `gorm:"column:shipping_handling;type:tinyint(1) unsigned;not null;default:1"`
	RangeBehavior      bool    `gorm:"column:range_behavior;type:tinyint(1) unsigned;not null;default:0"`
	IsModule           bool    `gorm:"column:is_module;type:tinyint(1) unsigned;not null;default:0"`
	IsFree             bool    `gorm:"column:is_free;type:tinyint(1) unsigned;not null;default:0"`
	ShippingExternal   bool    `gorm:"column:shipping_external;type:tinyint(1) unsigned;not null;default:0"`
	NeedRange          bool    `gorm:"column:need_range;type:tinyint(1) unsigned;not null;default:0"`
	ExternalModuleName string  `gorm:"column:external_module_name;type:varchar(64)"`
	ShippingMethod     int     `gorm:"column:shipping_method;type:int(2);not null;default:0"`
	Position           uint32  `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
	MaxWidth           int     `gorm:"column:max_width;type:int(10);default:0"`
	MaxHeight          int     `gorm:"column:max_height;type:int(10);default:0"`
	MaxDepth           int     `gorm:"column:max_depth;type:int(10);default:0"`
	MaxWeight          float64 `gorm:"column:max_weight;type:decimal(20,6);default:0.000000"`
	Grade              int     `gorm:"column:grade;type:int(10);default:0"`
}

type CarrierGroup struct {
	IDCarrier uint32 `gorm:"primaryKey;column:id_carrier;type:int(10) unsigned;not null"`
	IDGroup   uint32 `gorm:"primaryKey;column:id_group;type:int(10) unsigned;not null"`
}

type CarrierLang struct {
	IDCarrier uint32 `gorm:"primaryKey;column:id_carrier;type:int(10) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang    uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Delay     string `gorm:"column:delay;type:varchar(512)"`
}

type CarrierShop struct {
	IDCarrier uint32 `gorm:"primaryKey;column:id_carrier;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type CarrierTaxRulesGroupShop struct {
	IDCarrier       uint32 `gorm:"primaryKey;column:id_carrier;type:int(11) unsigned;not null"`
	IDTaxRulesGroup uint32 `gorm:"primaryKey;column:id_tax_rules_group;type:int(11) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null"`
}

type CarrierZone struct {
	IDCarrier uint32 `gorm:"primaryKey;column:id_carrier;type:int(10) unsigned;not null"`
	IDZone    uint32 `gorm:"primaryKey;column:id_zone;type:int(10) unsigned;not null"`
}

type Cart struct {
	IDCart                uint32    `gorm:"primaryKey;column:id_cart;type:int(10) unsigned;not null"`
	IDShopGroup           uint32    `gorm:"index:id_shop_group;column:id_shop_group;type:int(11) unsigned;not null;default:1"`
	IDShop                uint32    `gorm:"index:id_shop_2;index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDCarrier             uint32    `gorm:"index:id_carrier;column:id_carrier;type:int(10) unsigned;not null"`
	DeliveryOption        string    `gorm:"column:delivery_option;type:text;not null"`
	IDLang                uint32    `gorm:"index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	IDAddressDelivery     uint32    `gorm:"index:id_address_delivery;column:id_address_delivery;type:int(10) unsigned;not null"`
	IDAddressInvoice      uint32    `gorm:"index:id_address_invoice;column:id_address_invoice;type:int(10) unsigned;not null"`
	IDCurrency            uint32    `gorm:"index:id_currency;column:id_currency;type:int(10) unsigned;not null"`
	IDCustomer            uint32    `gorm:"index:cart_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDGuest               uint32    `gorm:"index:id_guest;column:id_guest;type:int(10) unsigned;not null"`
	SecureKey             string    `gorm:"column:secure_key;type:varchar(32);not null;default:-1"`
	Recyclable            bool      `gorm:"column:recyclable;type:tinyint(1) unsigned;not null;default:1"`
	Gift                  bool      `gorm:"column:gift;type:tinyint(1) unsigned;not null;default:0"`
	GiftMessage           string    `gorm:"column:gift_message;type:text"`
	MobileTheme           bool      `gorm:"column:mobile_theme;type:tinyint(1);not null;default:0"`
	AllowSeperatedPackage bool      `gorm:"column:allow_seperated_package;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd               time.Time `gorm:"index:id_shop;column:date_add;type:datetime;not null"`
	DateUpd               time.Time `gorm:"index:id_shop_2;column:date_upd;type:datetime;not null"`
	CheckoutSessionData   string    `gorm:"column:checkout_session_data;type:mediumtext"`
}

type CartCartRule struct {
	IDCart     uint32 `gorm:"primaryKey;column:id_cart;type:int(10) unsigned;not null"`
	IDCartRule uint32 `gorm:"primaryKey;index:id_cart_rule;column:id_cart_rule;type:int(10) unsigned;not null"`
}

type CartProduct struct {
	IDCart             uint32    `gorm:"primaryKey;index:id_cart_order;column:id_cart;type:int(10) unsigned;not null"`
	IDProduct          uint32    `gorm:"primaryKey;index:id_cart_order;column:id_product;type:int(10) unsigned;not null"`
	IDAddressDelivery  uint32    `gorm:"primaryKey;column:id_address_delivery;type:int(10) unsigned;not null;default:0"`
	IDShop             uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDProductAttribute uint32    `gorm:"primaryKey;index:id_product_attribute;index:id_cart_order;column:id_product_attribute;type:int(10) unsigned;not null;default:0"`
	IDCustomization    uint32    `gorm:"primaryKey;column:id_customization;type:int(10) unsigned;not null;default:0"`
	Quantity           uint32    `gorm:"column:quantity;type:int(10) unsigned;not null;default:0"`
	DateAdd            time.Time `gorm:"index:id_cart_order;column:date_add;type:datetime;not null"`
}

type CartRule struct {
	IDCartRule              uint32    `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDCustomer              uint32    `gorm:"index:id_customer;index:id_customer_2;column:id_customer;type:int(10) unsigned;not null;default:0"`
	DateFrom                time.Time `gorm:"index:date_from;column:date_from;type:datetime;not null"`
	DateTo                  time.Time `gorm:"index:id_customer;index:group_restriction;index:id_customer_2;index:group_restriction_2;index:date_to;column:date_to;type:datetime;not null"`
	Description             string    `gorm:"column:description;type:text"`
	Quantity                uint32    `gorm:"column:quantity;type:int(10) unsigned;not null;default:0"`
	QuantityPerUser         uint32    `gorm:"column:quantity_per_user;type:int(10) unsigned;not null;default:0"`
	Priority                uint32    `gorm:"column:priority;type:int(10) unsigned;not null;default:1"`
	PartialUse              bool      `gorm:"column:partial_use;type:tinyint(1) unsigned;not null;default:0"`
	Code                    string    `gorm:"column:code;type:varchar(254);not null"`
	MinimumAmount           float64   `gorm:"column:minimum_amount;type:decimal(20,6);not null;default:0.000000"`
	MinimumAmountTax        bool      `gorm:"column:minimum_amount_tax;type:tinyint(1);not null;default:0"`
	MinimumAmountCurrency   uint32    `gorm:"column:minimum_amount_currency;type:int(10) unsigned;not null;default:0"`
	MinimumAmountShipping   bool      `gorm:"column:minimum_amount_shipping;type:tinyint(1);not null;default:0"`
	CountryRestriction      bool      `gorm:"column:country_restriction;type:tinyint(1) unsigned;not null;default:0"`
	CarrierRestriction      bool      `gorm:"column:carrier_restriction;type:tinyint(1) unsigned;not null;default:0"`
	GroupRestriction        bool      `gorm:"index:group_restriction;index:group_restriction_2;column:group_restriction;type:tinyint(1) unsigned;not null;default:0"`
	CartRuleRestriction     bool      `gorm:"column:cart_rule_restriction;type:tinyint(1) unsigned;not null;default:0"`
	ProductRestriction      bool      `gorm:"column:product_restriction;type:tinyint(1) unsigned;not null;default:0"`
	ShopRestriction         bool      `gorm:"column:shop_restriction;type:tinyint(1) unsigned;not null;default:0"`
	FreeShipping            bool      `gorm:"column:free_shipping;type:tinyint(1);not null;default:0"`
	ReductionPercent        float64   `gorm:"column:reduction_percent;type:decimal(5,2);not null;default:0.00"`
	ReductionAmount         float64   `gorm:"column:reduction_amount;type:decimal(20,6);not null;default:0.000000"`
	ReductionTax            bool      `gorm:"column:reduction_tax;type:tinyint(1) unsigned;not null;default:0"`
	ReductionCurrency       uint32    `gorm:"column:reduction_currency;type:int(10) unsigned;not null;default:0"`
	ReductionProduct        int       `gorm:"column:reduction_product;type:int(10);not null;default:0"`
	ReductionExcludeSpecial bool      `gorm:"column:reduction_exclude_special;type:tinyint(1) unsigned;not null;default:0"`
	GiftProduct             uint32    `gorm:"column:gift_product;type:int(10) unsigned;not null;default:0"`
	GiftProductAttribute    uint32    `gorm:"column:gift_product_attribute;type:int(10) unsigned;not null;default:0"`
	Highlight               bool      `gorm:"index:id_customer_2;index:group_restriction_2;column:highlight;type:tinyint(1) unsigned;not null;default:0"`
	Active                  bool      `gorm:"index:id_customer;index:group_restriction;index:id_customer_2;index:group_restriction_2;column:active;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd                 time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd                 time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type CartRuleCarrier struct {
	IDCartRule uint32 `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDCarrier  uint32 `gorm:"primaryKey;column:id_carrier;type:int(10) unsigned;not null"`
}

type CartRuleCombination struct {
	IDCartRule1 uint32 `gorm:"primaryKey;index:id_cart_rule_1;column:id_cart_rule_1;type:int(10) unsigned;not null"`
	IDCartRule2 uint32 `gorm:"primaryKey;index:id_cart_rule_2;column:id_cart_rule_2;type:int(10) unsigned;not null"`
}

type CartRuleCountry struct {
	IDCartRule uint32 `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDCountry  uint32 `gorm:"primaryKey;column:id_country;type:int(10) unsigned;not null"`
}

type CartRuleGroup struct {
	IDCartRule uint32 `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDGroup    uint32 `gorm:"primaryKey;column:id_group;type:int(10) unsigned;not null"`
}

type CartRuleLang struct {
	IDCartRule uint32 `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDLang     uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name       string `gorm:"column:name;type:varchar(254);not null"`
}

type CartRuleProductRule struct {
	IDProductRule      uint32 `gorm:"primaryKey;column:id_product_rule;type:int(10) unsigned;not null"`
	IDProductRuleGroup uint32 `gorm:"column:id_product_rule_group;type:int(10) unsigned;not null"`
	Type               string `gorm:"column:type;type:enum('products','categories','attributes','manufacturers','suppliers');not null"`
}

type CartRuleProductRuleGroup struct {
	IDProductRuleGroup uint32 `gorm:"primaryKey;column:id_product_rule_group;type:int(10) unsigned;not null"`
	IDCartRule         uint32 `gorm:"column:id_cart_rule;type:int(10) unsigned;not null"`
	Quantity           uint32 `gorm:"column:quantity;type:int(10) unsigned;not null;default:1"`
}

type CartRuleProductRuleValue struct {
	IDProductRule uint32 `gorm:"primaryKey;column:id_product_rule;type:int(10) unsigned;not null"`
	IDItem        uint32 `gorm:"primaryKey;column:id_item;type:int(10) unsigned;not null"`
}

type CartRuleShop struct {
	IDCartRule uint32 `gorm:"primaryKey;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDShop     uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type Category struct {
	IDCategory     uint32    `gorm:"primaryKey;column:id_category;type:int(10) unsigned;not null"`
	IDParent       uint32    `gorm:"index:category_parent;column:id_parent;type:int(10) unsigned;not null"`
	IDShopDefault  uint32    `gorm:"column:id_shop_default;type:int(10) unsigned;not null;default:1"`
	LevelDepth     uint8     `gorm:"index:level_depth;column:level_depth;type:tinyint(3) unsigned;not null;default:0"`
	Nleft          uint32    `gorm:"index:nleftrightactive;index:activenleft;column:nleft;type:int(10) unsigned;not null;default:0"`
	Nright         uint32    `gorm:"index:nleftrightactive;index:nright;index:activenright;column:nright;type:int(10) unsigned;not null;default:0"`
	Active         bool      `gorm:"index:nleftrightactive;index:activenleft;index:activenright;column:active;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd        time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Position       uint32    `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
	IsRootCategory bool      `gorm:"column:is_root_category;type:tinyint(1);not null;default:0"`
}

type CategoryGroup struct {
	IDCategory uint32 `gorm:"primaryKey;index:id_category;column:id_category;type:int(10) unsigned;not null"`
	IDGroup    uint32 `gorm:"primaryKey;index:id_group;column:id_group;type:int(10) unsigned;not null"`
}

type CategoryLang struct {
	IDCategory      uint32 `gorm:"primaryKey;column:id_category;type:int(10) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang          uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name            string `gorm:"index:category_name;column:name;type:varchar(128);not null"`
	Description     string `gorm:"column:description;type:text"`
	LinkRewrite     string `gorm:"column:link_rewrite;type:varchar(128);not null"`
	MetaTitle       string `gorm:"column:meta_title;type:varchar(255)"`
	MetaKeywords    string `gorm:"column:meta_keywords;type:varchar(255)"`
	MetaDescription string `gorm:"column:meta_description;type:varchar(512)"`
}

type CategoryProduct struct {
	IDCategory uint32 `gorm:"primaryKey;index:id_category;column:id_category;type:int(10) unsigned;not null"`
	IDProduct  uint32 `gorm:"primaryKey;index:id_product;column:id_product;type:int(10) unsigned;not null"`
	Position   uint32 `gorm:"index:id_category;column:position;type:int(10) unsigned;not null;default:0"`
}

type CategoryShop struct {
	IDCategory int    `gorm:"primaryKey;column:id_category;type:int(11);not null"`
	IDShop     int    `gorm:"primaryKey;column:id_shop;type:int(11);not null"`
	Position   uint32 `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
}

type Cms struct {
	IDCms         uint32 `gorm:"primaryKey;column:id_cms;type:int(10) unsigned;not null"`
	IDCmsCategory uint32 `gorm:"column:id_cms_category;type:int(10) unsigned;not null"`
	Position      uint32 `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
	Active        bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	Indexation    bool   `gorm:"column:indexation;type:tinyint(1) unsigned;not null;default:1"`
}

type CmsCategory struct {
	IDCmsCategory uint32    `gorm:"primaryKey;column:id_cms_category;type:int(10) unsigned;not null"`
	IDParent      uint32    `gorm:"index:category_parent;column:id_parent;type:int(10) unsigned;not null"`
	LevelDepth    uint8     `gorm:"column:level_depth;type:tinyint(3) unsigned;not null;default:0"`
	Active        bool      `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd       time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd       time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Position      uint32    `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
}

type CmsCategoryLang struct {
	IDCmsCategory   uint32 `gorm:"primaryKey;column:id_cms_category;type:int(10) unsigned;not null"`
	IDLang          uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null;default:1"`
	Name            string `gorm:"index:category_name;column:name;type:varchar(128);not null"`
	Description     string `gorm:"column:description;type:text"`
	LinkRewrite     string `gorm:"column:link_rewrite;type:varchar(128);not null"`
	MetaTitle       string `gorm:"column:meta_title;type:varchar(255)"`
	MetaKeywords    string `gorm:"column:meta_keywords;type:varchar(255)"`
	MetaDescription string `gorm:"column:meta_description;type:varchar(512)"`
}

type CmsCategoryShop struct {
	IDCmsCategory uint32 `gorm:"primaryKey;column:id_cms_category;type:int(10) unsigned;not null"`
	IDShop        uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type CmsLang struct {
	IDCms           uint32 `gorm:"primaryKey;column:id_cms;type:int(10) unsigned;not null"`
	IDLang          uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null;default:1"`
	MetaTitle       string `gorm:"column:meta_title;type:varchar(255);not null"`
	HeadSeoTitle    string `gorm:"column:head_seo_title;type:varchar(255)"`
	MetaDescription string `gorm:"column:meta_description;type:varchar(512)"`
	MetaKeywords    string `gorm:"column:meta_keywords;type:varchar(255)"`
	Content         string `gorm:"column:content;type:longtext"`
	LinkRewrite     string `gorm:"column:link_rewrite;type:varchar(128);not null"`
}

type CmsRole struct {
	IDCmsRole uint32 `gorm:"primaryKey;column:id_cms_role;type:int(11) unsigned;not null"`
	Name      string `gorm:"unique;column:name;type:varchar(50);not null"`
	IDCms     uint32 `gorm:"primaryKey;column:id_cms;type:int(11) unsigned;not null"`
}

type CmsRoleLang struct {
	IDCmsRole uint32 `gorm:"primaryKey;column:id_cms_role;type:int(11) unsigned;not null"`
	IDLang    uint32 `gorm:"primaryKey;column:id_lang;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null"`
	Name      string `gorm:"column:name;type:varchar(128)"`
}

type CmsShop struct {
	IDCms  uint32 `gorm:"primaryKey;column:id_cms;type:int(11) unsigned;not null"`
	IDShop uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Condition struct {
	IDCondition       int       `gorm:"primaryKey;column:id_condition;type:int(11);not null"`
	IDPsCondition     int       `gorm:"primaryKey;column:id_ps_condition;type:int(11);not null"`
	Type              string    `gorm:"column:type;type:enum('configuration','install','sql');not null"`
	Request           string    `gorm:"column:request;type:text"`
	Operator          string    `gorm:"column:operator;type:varchar(32)"`
	Value             string    `gorm:"column:value;type:varchar(64)"`
	Result            string    `gorm:"column:result;type:varchar(64)"`
	CalculationType   string    `gorm:"column:calculation_type;type:enum('hook','time')"`
	CalculationDetail string    `gorm:"column:calculation_detail;type:varchar(64)"`
	Validated         bool      `gorm:"column:validated;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd           time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd           time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type ConditionAdvice struct {
	IDCondition int  `gorm:"primaryKey;column:id_condition;type:int(11);not null"`
	IDAdvice    int  `gorm:"primaryKey;column:id_advice;type:int(11);not null"`
	Display     bool `gorm:"column:display;type:tinyint(1) unsigned;not null;default:0"`
}

type ConditionBadge struct {
	IDCondition int `gorm:"primaryKey;column:id_condition;type:int(11);not null"`
	IDBadge     int `gorm:"primaryKey;column:id_badge;type:int(11);not null"`
}

type Configuration struct {
	IDConfiguration uint32    `gorm:"primaryKey;column:id_configuration;type:int(10) unsigned;not null"`
	IDShopGroup     uint32    `gorm:"index:id_shop_group;column:id_shop_group;type:int(11) unsigned"`
	IDShop          uint32    `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned"`
	Name            string    `gorm:"index:name;column:name;type:varchar(254);not null"`
	Value           string    `gorm:"column:value;type:text"`
	DateAdd         time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd         time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type ConfigurationKpi struct {
	IDConfigurationKpi uint32    `gorm:"primaryKey;column:id_configuration_kpi;type:int(10) unsigned;not null"`
	IDShopGroup        uint32    `gorm:"index:id_shop_group;column:id_shop_group;type:int(11) unsigned"`
	IDShop             uint32    `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned"`
	Name               string    `gorm:"index:name;column:name;type:varchar(64);not null"`
	Value              string    `gorm:"column:value;type:text"`
	DateAdd            time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd            time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type ConfigurationKpiLang struct {
	IDConfigurationKpi uint32    `gorm:"primaryKey;column:id_configuration_kpi;type:int(10) unsigned;not null"`
	IDLang             uint32    `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Value              string    `gorm:"column:value;type:text"`
	DateUpd            time.Time `gorm:"column:date_upd;type:datetime"`
}

type ConfigurationLang struct {
	IDConfiguration uint32    `gorm:"primaryKey;column:id_configuration;type:int(10) unsigned;not null"`
	IDLang          uint32    `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Value           string    `gorm:"column:value;type:text"`
	DateUpd         time.Time `gorm:"column:date_upd;type:datetime"`
}

type Connections struct {
	IDConnections uint32    `gorm:"primaryKey;column:id_connections;type:int(10) unsigned;not null"`
	IDShopGroup   uint32    `gorm:"column:id_shop_group;type:int(11) unsigned;not null;default:1"`
	IDShop        uint32    `gorm:"column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDGuest       uint32    `gorm:"index:id_guest;column:id_guest;type:int(10) unsigned;not null"`
	IDPage        uint32    `gorm:"index:id_page;column:id_page;type:int(10) unsigned;not null"`
	IPAddress     int64     `gorm:"column:ip_address;type:bigint(20)"`
	DateAdd       time.Time `gorm:"index:date_add;column:date_add;type:datetime;not null"`
	HTTPReferer   string    `gorm:"column:http_referer;type:varchar(255)"`
}

type ConnectionsPage struct {
	IDConnections uint32    `gorm:"primaryKey;column:id_connections;type:int(10) unsigned;not null"`
	IDPage        uint32    `gorm:"primaryKey;column:id_page;type:int(10) unsigned;not null"`
	TimeStart     time.Time `gorm:"primaryKey;column:time_start;type:datetime;not null"`
	TimeEnd       time.Time `gorm:"column:time_end;type:datetime"`
}

type ConnectionsSource struct {
	IDConnectionsSource uint32    `gorm:"primaryKey;column:id_connections_source;type:int(10) unsigned;not null"`
	IDConnections       uint32    `gorm:"index:connections;column:id_connections;type:int(10) unsigned;not null"`
	HTTPReferer         string    `gorm:"index:http_referer;column:http_referer;type:varchar(255)"`
	RequestURI          string    `gorm:"index:request_uri;column:request_uri;type:varchar(255)"`
	Keywords            string    `gorm:"column:keywords;type:varchar(255)"`
	DateAdd             time.Time `gorm:"index:orderby;column:date_add;type:datetime;not null"`
}

type Contact struct {
	IDContact       uint32 `gorm:"primaryKey;column:id_contact;type:int(10) unsigned;not null"`
	Email           string `gorm:"column:email;type:varchar(255);not null"`
	CustomerService bool   `gorm:"column:customer_service;type:tinyint(1);not null;default:0"`
	Position        uint8  `gorm:"column:position;type:tinyint(2) unsigned;not null;default:0"`
}

type ContactLang struct {
	IDContact   uint32 `gorm:"primaryKey;column:id_contact;type:int(10) unsigned;not null"`
	IDLang      uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text"`
}

type ContactShop struct {
	IDContact uint32 `gorm:"primaryKey;column:id_contact;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Country struct {
	IDCountry                uint32 `gorm:"primaryKey;column:id_country;type:int(10) unsigned;not null"`
	IDZone                   uint32 `gorm:"index:country_;column:id_zone;type:int(10) unsigned;not null"`
	IDCurrency               uint32 `gorm:"column:id_currency;type:int(10) unsigned;not null;default:0"`
	IsoCode                  string `gorm:"index:country_iso_code;column:iso_code;type:varchar(3);not null"`
	CallPrefix               int    `gorm:"column:call_prefix;type:int(10);not null;default:0"`
	Active                   bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	ContainsStates           bool   `gorm:"column:contains_states;type:tinyint(1);not null;default:0"`
	NeedIDentificationNumber bool   `gorm:"column:need_identification_number;type:tinyint(1);not null;default:0"`
	NeedZipCode              bool   `gorm:"column:need_zip_code;type:tinyint(1);not null;default:1"`
	ZipCodeFormat            string `gorm:"column:zip_code_format;type:varchar(12);not null;default:''"`
	DisplayTaxLabel          bool   `gorm:"column:display_tax_label;type:tinyint(1);not null"`
}

type CountryLang struct {
	IDCountry uint32 `gorm:"primaryKey;column:id_country;type:int(10) unsigned;not null"`
	IDLang    uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name      string `gorm:"column:name;type:varchar(64);not null"`
}

type CountryShop struct {
	IDCountry uint32 `gorm:"primaryKey;column:id_country;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Currency struct {
	IDCurrency     uint32  `gorm:"primaryKey;column:id_currency;type:int(10) unsigned;not null"`
	Name           string  `gorm:"column:name;type:varchar(64);not null"`
	IsoCode        string  `gorm:"index:currency_iso_code;column:iso_code;type:varchar(3);not null;default:0"`
	NumericIsoCode string  `gorm:"column:numeric_iso_code;type:varchar(3)"`
	Precision      int     `gorm:"column:precision;type:int(2);not null;default:6"`
	ConversionRate float64 `gorm:"column:conversion_rate;type:decimal(13,6);not null"`
	Deleted        bool    `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
	Active         bool    `gorm:"column:active;type:tinyint(1) unsigned;not null;default:1"`
	Unofficial     bool    `gorm:"column:unofficial;type:tinyint(1) unsigned;not null;default:0"`
	Modified       bool    `gorm:"column:modified;type:tinyint(1) unsigned;not null;default:0"`
}

type CurrencyLang struct {
	IDCurrency uint32 `gorm:"primaryKey;column:id_currency;type:int(10) unsigned;not null"`
	IDLang     uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name       string `gorm:"column:name;type:varchar(255);not null"`
	Symbol     string `gorm:"column:symbol;type:varchar(255);not null"`
	Pattern    string `gorm:"column:pattern;type:varchar(255)"`
}

type CurrencyShop struct {
	IDCurrency     uint32  `gorm:"primaryKey;column:id_currency;type:int(11) unsigned;not null"`
	IDShop         uint32  `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	ConversionRate float64 `gorm:"column:conversion_rate;type:decimal(13,6);not null"`
}

type Customer struct {
	IDCustomer               uint32         `gorm:"primaryKey;index:id_customer_passwd;column:id_customer;type:int(10) unsigned;not null"`
	IDShopGroup              uint32         `gorm:"index:id_shop_group;column:id_shop_group;type:int(11) unsigned;not null;default:1"`
	IDShop                   uint32         `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDGender                 uint32         `gorm:"index:id_gender;column:id_gender;type:int(10) unsigned;not null"`
	IDDefaultGroup           uint32         `gorm:"column:id_default_group;type:int(10) unsigned;not null;default:1"`
	IDLang                   uint32         `gorm:"column:id_lang;type:int(10) unsigned"`
	IDRisk                   uint32         `gorm:"column:id_risk;type:int(10) unsigned;not null;default:1"`
	Company                  string         `gorm:"column:company;type:varchar(255)"`
	Siret                    string         `gorm:"column:siret;type:varchar(14)"`
	Ape                      string         `gorm:"column:ape;type:varchar(5)"`
	Firstname                string         `gorm:"column:firstname;type:varchar(255);not null"`
	Lastname                 string         `gorm:"column:lastname;type:varchar(255);not null"`
	Email                    string         `gorm:"index:customer_email;index:customer_login;column:email;type:varchar(255);not null"`
	Passwd                   string         `gorm:"index:customer_login;index:id_customer_passwd;column:passwd;type:varchar(255);not null"`
	LastPasswdGen            time.Time      `gorm:"column:last_passwd_gen;type:timestamp;not null;default:current_timestamp()"`
	Birthday                 datatypes.Date `gorm:"column:birthday;type:date"`
	Newsletter               bool           `gorm:"column:newsletter;type:tinyint(1) unsigned;not null;default:0"`
	IPRegistrationNewsletter string         `gorm:"column:ip_registration_newsletter;type:varchar(15)"`
	NewsletterDateAdd        time.Time      `gorm:"column:newsletter_date_add;type:datetime"`
	Optin                    bool           `gorm:"column:optin;type:tinyint(1) unsigned;not null;default:0"`
	Website                  string         `gorm:"column:website;type:varchar(128)"`
	OutstandingAllowAmount   float64        `gorm:"column:outstanding_allow_amount;type:decimal(20,6);not null;default:0.000000"`
	ShowPublicPrices         bool           `gorm:"column:show_public_prices;type:tinyint(1) unsigned;not null;default:0"`
	MaxPaymentDays           uint32         `gorm:"column:max_payment_days;type:int(10) unsigned;not null;default:60"`
	SecureKey                string         `gorm:"column:secure_key;type:varchar(32);not null;default:-1"`
	Note                     string         `gorm:"column:note;type:text"`
	Active                   bool           `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	IsGuest                  bool           `gorm:"column:is_guest;type:tinyint(1);not null;default:0"`
	Deleted                  bool           `gorm:"column:deleted;type:tinyint(1);not null;default:0"`
	DateAdd                  time.Time      `gorm:"index:id_shop;column:date_add;type:datetime;not null"`
	DateUpd                  time.Time      `gorm:"column:date_upd;type:datetime;not null"`
	ResetPasswordToken       string         `gorm:"column:reset_password_token;type:varchar(40)"`
	ResetPasswordValidity    time.Time      `gorm:"column:reset_password_validity;type:datetime"`
}

type CustomerGroup struct {
	IDCustomer uint32 `gorm:"primaryKey;index:id_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDGroup    uint32 `gorm:"primaryKey;index:customer_login;column:id_group;type:int(10) unsigned;not null"`
}

type CustomerMessage struct {
	IDCustomerMessage uint32    `gorm:"primaryKey;column:id_customer_message;type:int(10) unsigned;not null"`
	IDCustomerThread  int       `gorm:"index:id_customer_thread;column:id_customer_thread;type:int(11)"`
	IDEmployee        uint32    `gorm:"index:id_employee;column:id_employee;type:int(10) unsigned"`
	Message           string    `gorm:"column:message;type:mediumtext;not null"`
	FileName          string    `gorm:"column:file_name;type:varchar(18)"`
	IPAddress         string    `gorm:"column:ip_address;type:varchar(16)"`
	UserAgent         string    `gorm:"column:user_agent;type:varchar(128)"`
	DateAdd           time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd           time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Private           int8      `gorm:"column:private;type:tinyint(4);not null;default:0"`
	Read              bool      `gorm:"column:read;type:tinyint(1);not null;default:0"`
}

type CustomerMessageSyncImap struct {
	Md5Header []byte `gorm:"index:md5_header_index;column:md5_header;type:varbinary(32);not null"`
}

type CustomerSession struct {
	IDCustomerSession uint32 `gorm:"primaryKey;column:id_customer_session;type:int(11) unsigned;not null"`
	IDCustomer        uint32 `gorm:"column:id_customer;type:int(10) unsigned"`
	Token             string `gorm:"column:token;type:varchar(40)"`
}

type CustomerThread struct {
	IDCustomerThread uint32    `gorm:"primaryKey;column:id_customer_thread;type:int(11) unsigned;not null"`
	IDShop           uint32    `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang           uint32    `gorm:"index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	IDContact        uint32    `gorm:"index:id_contact;column:id_contact;type:int(10) unsigned;not null"`
	IDCustomer       uint32    `gorm:"index:id_customer;column:id_customer;type:int(10) unsigned"`
	IDOrder          uint32    `gorm:"index:id_order;column:id_order;type:int(10) unsigned"`
	IDProduct        uint32    `gorm:"index:id_product;column:id_product;type:int(10) unsigned"`
	Status           string    `gorm:"column:status;type:enum('open','closed','pending1','pending2');not null;default:open"`
	Email            string    `gorm:"column:email;type:varchar(255);not null"`
	Token            string    `gorm:"column:token;type:varchar(12)"`
	DateAdd          time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd          time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type Customization struct {
	IDCustomization    uint32 `gorm:"primaryKey;column:id_customization;type:int(10) unsigned;not null"`
	IDProductAttribute uint32 `gorm:"index:id_product_attribute;index:id_cart_product;column:id_product_attribute;type:int(10) unsigned;not null;default:0"`
	IDAddressDelivery  uint32 `gorm:"primaryKey;column:id_address_delivery;type:int(10) unsigned;not null;default:0"`
	IDCart             uint32 `gorm:"primaryKey;index:id_cart_product;column:id_cart;type:int(10) unsigned;not null"`
	IDProduct          int    `gorm:"primaryKey;index:id_cart_product;column:id_product;type:int(10);not null"`
	Quantity           int    `gorm:"column:quantity;type:int(10);not null"`
	QuantityRefunded   int    `gorm:"column:quantity_refunded;type:int(11);not null;default:0"`
	QuantityReturned   int    `gorm:"column:quantity_returned;type:int(11);not null;default:0"`
	InCart             bool   `gorm:"column:in_cart;type:tinyint(1) unsigned;not null;default:0"`
}

type CustomizationField struct {
	IDCustomizationField uint32 `gorm:"primaryKey;column:id_customization_field;type:int(10) unsigned;not null"`
	IDProduct            uint32 `gorm:"index:id_product;column:id_product;type:int(10) unsigned;not null"`
	Type                 bool   `gorm:"column:type;type:tinyint(1);not null"`
	Required             bool   `gorm:"column:required;type:tinyint(1);not null"`
	IsModule             bool   `gorm:"column:is_module;type:tinyint(1);not null;default:0"`
	IsDeleted            bool   `gorm:"column:is_deleted;type:tinyint(1);not null;default:0"`
}

type CustomizationFieldLang struct {
	IDCustomizationField uint32 `gorm:"primaryKey;column:id_customization_field;type:int(10) unsigned;not null"`
	IDLang               uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	IDShop               uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null;default:1"`
	Name                 string `gorm:"column:name;type:varchar(255);not null"`
}

type CustomizedData struct {
	IDCustomization uint32  `gorm:"primaryKey;column:id_customization;type:int(10) unsigned;not null"`
	Type            bool    `gorm:"primaryKey;column:type;type:tinyint(1);not null"`
	Index           int     `gorm:"primaryKey;column:index;type:int(3);not null"`
	Value           string  `gorm:"column:value;type:varchar(255);not null"`
	IDModule        int     `gorm:"column:id_module;type:int(10);not null;default:0"`
	Price           float64 `gorm:"column:price;type:decimal(20,6);not null;default:0.000000"`
	Weight          float64 `gorm:"column:weight;type:decimal(20,6);not null;default:0.000000"`
}

type DateRange struct {
	IDDateRange uint32    `gorm:"primaryKey;column:id_date_range;type:int(10) unsigned;not null"`
	TimeStart   time.Time `gorm:"column:time_start;type:datetime;not null"`
	TimeEnd     time.Time `gorm:"column:time_end;type:datetime;not null"`
}

type Delivery struct {
	IDDelivery    uint32  `gorm:"primaryKey;column:id_delivery;type:int(10) unsigned;not null"`
	IDShop        uint32  `gorm:"column:id_shop;type:int(10) unsigned"`
	IDShopGroup   uint32  `gorm:"column:id_shop_group;type:int(10) unsigned"`
	IDCarrier     uint32  `gorm:"index:id_carrier;column:id_carrier;type:int(10) unsigned;not null"`
	IDRangePrice  uint32  `gorm:"index:id_range_price;column:id_range_price;type:int(10) unsigned"`
	IDRangeWeight uint32  `gorm:"index:id_range_weight;column:id_range_weight;type:int(10) unsigned"`
	IDZone        uint32  `gorm:"index:id_zone;index:id_carrier;column:id_zone;type:int(10) unsigned;not null"`
	Price         float64 `gorm:"column:price;type:decimal(20,6);not null"`
}

type Emailsubscription struct {
	ID                       int       `gorm:"primaryKey;column:id;type:int(6);not null"`
	IDShop                   uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDShopGroup              uint32    `gorm:"column:id_shop_group;type:int(10) unsigned;not null;default:1"`
	Email                    string    `gorm:"column:email;type:varchar(255);not null"`
	NewsletterDateAdd        time.Time `gorm:"column:newsletter_date_add;type:datetime"`
	IPRegistrationNewsletter string    `gorm:"column:ip_registration_newsletter;type:varchar(15);not null"`
	HTTPReferer              string    `gorm:"column:http_referer;type:varchar(255)"`
	Active                   bool      `gorm:"column:active;type:tinyint(1);not null;default:0"`
	IDLang                   int       `gorm:"column:id_lang;type:int(10);not null;default:0"`
}

type Employee struct {
	IDEmployee            uint32         `gorm:"primaryKey;index:id_employee_passwd;column:id_employee;type:int(10) unsigned;not null"`
	IDProfile             uint32         `gorm:"index:id_profile;column:id_profile;type:int(10) unsigned;not null"`
	IDLang                uint32         `gorm:"column:id_lang;type:int(10) unsigned;not null;default:0"`
	Lastname              string         `gorm:"column:lastname;type:varchar(255);not null"`
	Firstname             string         `gorm:"column:firstname;type:varchar(255);not null"`
	Email                 string         `gorm:"index:employee_login;column:email;type:varchar(255);not null"`
	Passwd                string         `gorm:"index:employee_login;index:id_employee_passwd;column:passwd;type:varchar(255);not null"`
	LastPasswdGen         time.Time      `gorm:"column:last_passwd_gen;type:timestamp;not null;default:current_timestamp()"`
	StatsDateFrom         datatypes.Date `gorm:"column:stats_date_from;type:date"`
	StatsDateTo           datatypes.Date `gorm:"column:stats_date_to;type:date"`
	StatsCompareFrom      datatypes.Date `gorm:"column:stats_compare_from;type:date"`
	StatsCompareTo        datatypes.Date `gorm:"column:stats_compare_to;type:date"`
	StatsCompareOption    uint32         `gorm:"column:stats_compare_option;type:int(1) unsigned;not null;default:1"`
	PreselectDateRange    string         `gorm:"column:preselect_date_range;type:varchar(32)"`
	BoColor               string         `gorm:"column:bo_color;type:varchar(32)"`
	BoTheme               string         `gorm:"column:bo_theme;type:varchar(32)"`
	BoCSS                 string         `gorm:"column:bo_css;type:varchar(64)"`
	DefaultTab            uint32         `gorm:"column:default_tab;type:int(10) unsigned;not null;default:0"`
	BoWidth               uint32         `gorm:"column:bo_width;type:int(10) unsigned;not null;default:0"`
	BoMenu                bool           `gorm:"column:bo_menu;type:tinyint(1);not null;default:1"`
	Active                bool           `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	Optin                 bool           `gorm:"column:optin;type:tinyint(1) unsigned"`
	IDLastOrder           uint32         `gorm:"column:id_last_order;type:int(10) unsigned;not null;default:0"`
	IDLastCustomerMessage uint32         `gorm:"column:id_last_customer_message;type:int(10) unsigned;not null;default:0"`
	IDLastCustomer        uint32         `gorm:"column:id_last_customer;type:int(10) unsigned;not null;default:0"`
	LastConnectionDate    datatypes.Date `gorm:"column:last_connection_date;type:date"`
	ResetPasswordToken    string         `gorm:"column:reset_password_token;type:varchar(40)"`
	ResetPasswordValidity time.Time      `gorm:"column:reset_password_validity;type:datetime"`
}

type EmployeeSession struct {
	IDEmployeeSession uint32 `gorm:"primaryKey;column:id_employee_session;type:int(11) unsigned;not null"`
	IDEmployee        uint32 `gorm:"column:id_employee;type:int(10) unsigned"`
	Token             string `gorm:"column:token;type:varchar(40)"`
}

type EmployeeShop struct {
	IDEmployee uint32 `gorm:"primaryKey;column:id_employee;type:int(11) unsigned;not null"`
	IDShop     uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Feature struct {
	IDFeature uint32 `gorm:"primaryKey;column:id_feature;type:int(10) unsigned;not null"`
	Position  uint32 `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
}

type FeatureLang struct {
	IDFeature uint32 `gorm:"primaryKey;column:id_feature;type:int(10) unsigned;not null"`
	IDLang    uint32 `gorm:"primaryKey;index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	Name      string `gorm:"index:id_lang;column:name;type:varchar(128)"`
}

type FeatureProduct struct {
	IDFeature      uint32 `gorm:"primaryKey;column:id_feature;type:int(10) unsigned;not null"`
	IDProduct      uint32 `gorm:"primaryKey;index:id_product;column:id_product;type:int(10) unsigned;not null"`
	IDFeatureValue uint32 `gorm:"primaryKey;index:id_feature_value;column:id_feature_value;type:int(10) unsigned;not null"`
}

type FeatureShop struct {
	IDFeature uint32 `gorm:"primaryKey;column:id_feature;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type FeatureValue struct {
	IDFeatureValue uint32 `gorm:"primaryKey;column:id_feature_value;type:int(10) unsigned;not null"`
	IDFeature      uint32 `gorm:"index:feature;column:id_feature;type:int(10) unsigned;not null"`
	Custom         uint8  `gorm:"column:custom;type:tinyint(3) unsigned"`
}

type FeatureValueLang struct {
	IDFeatureValue uint32 `gorm:"primaryKey;column:id_feature_value;type:int(10) unsigned;not null"`
	IDLang         uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Value          string `gorm:"column:value;type:varchar(255)"`
}

type Gender struct {
	IDGender int  `gorm:"primaryKey;column:id_gender;type:int(11);not null"`
	Type     bool `gorm:"column:type;type:tinyint(1);not null"`
}

type GenderLang struct {
	IDGender uint32 `gorm:"primaryKey;index:id_gender;column:id_gender;type:int(10) unsigned;not null"`
	IDLang   uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name     string `gorm:"column:name;type:varchar(20);not null"`
}

type Group struct {
	IDGroup            uint32    `gorm:"primaryKey;column:id_group;type:int(10) unsigned;not null"`
	Reduction          float64   `gorm:"column:reduction;type:decimal(5,2);not null;default:0.00"`
	PriceDisplayMethod int8      `gorm:"column:price_display_method;type:tinyint(4);not null;default:0"`
	ShowPrices         bool      `gorm:"column:show_prices;type:tinyint(1) unsigned;not null;default:1"`
	DateAdd            time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd            time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type GroupLang struct {
	IDGroup uint32 `gorm:"primaryKey;column:id_group;type:int(10) unsigned;not null"`
	IDLang  uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name    string `gorm:"column:name;type:varchar(32);not null"`
}

type GroupReduction struct {
	IDGroupReduction string  `gorm:"primaryKey;column:id_group_reduction;type:mediumint(8) unsigned;not null"`
	IDGroup          uint32  `gorm:"uniqueIndex:id_group;column:id_group;type:int(10) unsigned;not null"`
	IDCategory       uint32  `gorm:"uniqueIndex:id_group;column:id_category;type:int(10) unsigned;not null"`
	Reduction        float64 `gorm:"column:reduction;type:decimal(5,4);not null"`
}

type GroupShop struct {
	IDGroup uint32 `gorm:"primaryKey;column:id_group;type:int(11) unsigned;not null"`
	IDShop  uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type GsitemapSitemap struct {
	Link   string `gorm:"column:link;type:varchar(255)"`
	IDShop int    `gorm:"column:id_shop;type:int(11);default:0"`
}

type Guest struct {
	IDGuest           uint32 `gorm:"primaryKey;column:id_guest;type:int(10) unsigned;not null"`
	IDOperatingSystem uint32 `gorm:"index:id_operating_system;column:id_operating_system;type:int(10) unsigned"`
	IDWebBrowser      uint32 `gorm:"index:id_web_browser;column:id_web_browser;type:int(10) unsigned"`
	IDCustomer        uint32 `gorm:"index:id_customer;column:id_customer;type:int(10) unsigned"`
	Javascript        bool   `gorm:"column:javascript;type:tinyint(1);default:0"`
	ScreenResolutionX uint16 `gorm:"column:screen_resolution_x;type:smallint(5) unsigned"`
	ScreenResolutionY uint16 `gorm:"column:screen_resolution_y;type:smallint(5) unsigned"`
	ScreenColor       uint8  `gorm:"column:screen_color;type:tinyint(3) unsigned"`
	SunJava           bool   `gorm:"column:sun_java;type:tinyint(1)"`
	AdobeFlash        bool   `gorm:"column:adobe_flash;type:tinyint(1)"`
	AdobeDirector     bool   `gorm:"column:adobe_director;type:tinyint(1)"`
	AppleQuicktime    bool   `gorm:"column:apple_quicktime;type:tinyint(1)"`
	RealPlayer        bool   `gorm:"column:real_player;type:tinyint(1)"`
	WindowsMedia      bool   `gorm:"column:windows_media;type:tinyint(1)"`
	AcceptLanguage    string `gorm:"column:accept_language;type:varchar(8)"`
	MobileTheme       bool   `gorm:"column:mobile_theme;type:tinyint(1);not null;default:0"`
}

type Homeslider struct {
	IDHomesliderSlides uint32 `gorm:"primaryKey;column:id_homeslider_slides;type:int(10) unsigned;not null"`
	IDShop             uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type HomesliderSlides struct {
	IDHomesliderSlides uint32 `gorm:"primaryKey;column:id_homeslider_slides;type:int(10) unsigned;not null"`
	Position           uint32 `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
	Active             bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
}

type HomesliderSlidesLang struct {
	IDHomesliderSlides uint32 `gorm:"primaryKey;column:id_homeslider_slides;type:int(10) unsigned;not null"`
	IDLang             uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Title              string `gorm:"column:title;type:varchar(255);not null"`
	Description        string `gorm:"column:description;type:text;not null"`
	Legend             string `gorm:"column:legend;type:varchar(255);not null"`
	URL                string `gorm:"column:url;type:varchar(255);not null"`
	Image              string `gorm:"column:image;type:varchar(255);not null"`
}

type Hook struct {
	IDHook      uint32 `gorm:"primaryKey;column:id_hook;type:int(10) unsigned;not null"`
	Name        string `gorm:"unique;column:name;type:varchar(191);not null"`
	Title       string `gorm:"column:title;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text"`
	Position    bool   `gorm:"column:position;type:tinyint(1);not null;default:1"`
}

type HookAlias struct {
	IDHookAlias uint32 `gorm:"primaryKey;column:id_hook_alias;type:int(10) unsigned;not null"`
	Alias       string `gorm:"unique;column:alias;type:varchar(191);not null"`
	Name        string `gorm:"column:name;type:varchar(191);not null"`
}

type HookModule struct {
	IDModule uint32 `gorm:"primaryKey;index:id_module;column:id_module;type:int(10) unsigned;not null"`
	IDShop   uint32 `gorm:"primaryKey;index:position;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDHook   uint32 `gorm:"primaryKey;index:id_hook;column:id_hook;type:int(10) unsigned;not null"`
	Position uint8  `gorm:"index:position;column:position;type:tinyint(2) unsigned;not null"`
}

type HookModuleExceptions struct {
	IDHookModuleExceptions uint32 `gorm:"primaryKey;column:id_hook_module_exceptions;type:int(10) unsigned;not null"`
	IDShop                 uint32 `gorm:"column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDModule               uint32 `gorm:"index:id_module;column:id_module;type:int(10) unsigned;not null"`
	IDHook                 uint32 `gorm:"index:id_hook;column:id_hook;type:int(10) unsigned;not null"`
	FileName               string `gorm:"column:file_name;type:varchar(255)"`
}

type Image struct {
	IDImage   uint32 `gorm:"primaryKey;uniqueIndex:idx_product_image;column:id_image;type:int(10) unsigned;not null"`
	IDProduct uint32 `gorm:"uniqueIndex:id_product_cover;uniqueIndex:idx_product_image;index:image_product;column:id_product;type:int(10) unsigned;not null"`
	Position  uint16 `gorm:"column:position;type:smallint(2) unsigned;not null;default:0"`
	Cover     bool   `gorm:"uniqueIndex:id_product_cover;uniqueIndex:idx_product_image;column:cover;type:tinyint(1) unsigned"`
}

type ImageLang struct {
	IDImage uint32 `gorm:"primaryKey;index:id_image;column:id_image;type:int(10) unsigned;not null"`
	IDLang  uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Legend  string `gorm:"column:legend;type:varchar(128)"`
}

type ImageShop struct {
	IDProduct uint32 `gorm:"uniqueIndex:id_product;column:id_product;type:int(10) unsigned;not null"`
	IDImage   uint32 `gorm:"primaryKey;column:id_image;type:int(11) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;uniqueIndex:id_product;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	Cover     bool   `gorm:"uniqueIndex:id_product;column:cover;type:tinyint(1) unsigned"`
}

type ImageType struct {
	IDImageType   uint32 `gorm:"primaryKey;column:id_image_type;type:int(10) unsigned;not null"`
	Name          string `gorm:"index:image_type_name;column:name;type:varchar(64);not null"`
	Width         uint32 `gorm:"column:width;type:int(10) unsigned;not null"`
	Height        uint32 `gorm:"column:height;type:int(10) unsigned;not null"`
	Products      bool   `gorm:"column:products;type:tinyint(1);not null;default:1"`
	Categories    bool   `gorm:"column:categories;type:tinyint(1);not null;default:1"`
	Manufacturers bool   `gorm:"column:manufacturers;type:tinyint(1);not null;default:1"`
	Suppliers     bool   `gorm:"column:suppliers;type:tinyint(1);not null;default:1"`
	Stores        bool   `gorm:"column:stores;type:tinyint(1);not null;default:1"`
}

type ImportMatch struct {
	IDImportMatch int    `gorm:"primaryKey;column:id_import_match;type:int(10);not null"`
	Name          string `gorm:"column:name;type:varchar(32);not null"`
	Match         string `gorm:"column:match;type:text;not null"`
	Skip          int    `gorm:"column:skip;type:int(2);not null"`
}

type Info struct {
	IDInfo uint32 `gorm:"primaryKey;column:id_info;type:int(10) unsigned;not null"`
}

type InfoLang struct {
	IDInfo uint32 `gorm:"primaryKey;column:id_info;type:int(10) unsigned;not null"`
	IDShop uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
	IDLang uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Text   string `gorm:"column:text;type:text;not null"`
}

type InfoShop struct {
	IDInfo uint32 `gorm:"primaryKey;column:id_info;type:int(10) unsigned;not null"`
	IDShop uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type Lang struct {
	IDLang         int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	Name           string `gorm:"column:name;type:varchar(32);not null"`
	Active         bool   `gorm:"column:active;type:tinyint(1);not null"`
	IsoCode        string `gorm:"column:iso_code;type:varchar(2);not null"`
	LanguageCode   string `gorm:"column:language_code;type:varchar(5);not null"`
	Locale         string `gorm:"column:locale;type:varchar(5);not null"`
	DateFormatLite string `gorm:"column:date_format_lite;type:varchar(32);not null"`
	DateFormatFull string `gorm:"column:date_format_full;type:varchar(32);not null"`
	IsRtl          bool   `gorm:"column:is_rtl;type:tinyint(1);not null"`
}

type LangShop struct {
	IDLang int `gorm:"primaryKey;index:IDX_A5D79262BA299860;column:id_lang;type:int(11);not null"`
	IDShop int `gorm:"primaryKey;index:IDX_A5D79262274A50A0;column:id_shop;type:int(11);not null"`
}

type LayeredCategory struct {
	IDLayeredCategory uint32 `gorm:"primaryKey;column:id_layered_category;type:int(10) unsigned;not null"`
	IDShop            uint32 `gorm:"index:id_category_shop;column:id_shop;type:int(11) unsigned;not null"`
	IDCategory        uint32 `gorm:"index:id_category_shop;index:id_category;column:id_category;type:int(10) unsigned;not null"`
	IDValue           uint32 `gorm:"index:id_category_shop;column:id_value;type:int(10) unsigned;default:0"`
	Type              string `gorm:"index:id_category_shop;index:id_category;column:type;type:enum('category','id_feature','id_attribute_group','quantity','condition','manufacturer','weight','price');not null"`
	Position          uint32 `gorm:"index:id_category_shop;column:position;type:int(10) unsigned;not null"`
	FilterType        uint32 `gorm:"column:filter_type;type:int(10) unsigned;not null;default:0"`
	FilterShowLimit   uint32 `gorm:"column:filter_show_limit;type:int(10) unsigned;not null;default:0"`
}

type LayeredFilter struct {
	IDLayeredFilter uint32    `gorm:"primaryKey;column:id_layered_filter;type:int(10) unsigned;not null"`
	Name            string    `gorm:"column:name;type:varchar(64);not null"`
	Filters         string    `gorm:"column:filters;type:longtext"`
	NCategories     uint32    `gorm:"column:n_categories;type:int(10) unsigned;not null"`
	DateAdd         time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type LayeredFilterBlock struct {
	Hash string `gorm:"primaryKey;column:hash;type:char(32);not null;default:''"`
	Data string `gorm:"column:data;type:text"`
}

type LayeredFilterShop struct {
	IDLayeredFilter uint32 `gorm:"primaryKey;column:id_layered_filter;type:int(10) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type LayeredIndexableAttributeGroup struct {
	IDAttributeGroup int  `gorm:"primaryKey;column:id_attribute_group;type:int(11);not null"`
	Indexable        bool `gorm:"column:indexable;type:tinyint(1);not null;default:0"`
}

type LayeredIndexableAttributeGroupLangValue struct {
	IDAttributeGroup int    `gorm:"primaryKey;column:id_attribute_group;type:int(11);not null"`
	IDLang           int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	URLName          string `gorm:"column:url_name;type:varchar(128)"`
	MetaTitle        string `gorm:"column:meta_title;type:varchar(128)"`
}

type LayeredIndexableAttributeLangValue struct {
	IDAttribute int    `gorm:"primaryKey;column:id_attribute;type:int(11);not null"`
	IDLang      int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	URLName     string `gorm:"column:url_name;type:varchar(128)"`
	MetaTitle   string `gorm:"column:meta_title;type:varchar(128)"`
}

type LayeredIndexableFeature struct {
	IDFeature int  `gorm:"primaryKey;column:id_feature;type:int(11);not null"`
	Indexable bool `gorm:"column:indexable;type:tinyint(1);not null;default:0"`
}

type LayeredIndexableFeatureLangValue struct {
	IDFeature int    `gorm:"primaryKey;column:id_feature;type:int(11);not null"`
	IDLang    int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	URLName   string `gorm:"column:url_name;type:varchar(128);not null"`
	MetaTitle string `gorm:"column:meta_title;type:varchar(128)"`
}

type LayeredIndexableFeatureValueLangValue struct {
	IDFeatureValue int    `gorm:"primaryKey;column:id_feature_value;type:int(11);not null"`
	IDLang         int    `gorm:"primaryKey;column:id_lang;type:int(11);not null"`
	URLName        string `gorm:"column:url_name;type:varchar(128)"`
	MetaTitle      string `gorm:"column:meta_title;type:varchar(128)"`
}

type LayeredPriceIndex struct {
	IDProduct  int     `gorm:"primaryKey;column:id_product;type:int(11);not null"`
	IDCurrency int     `gorm:"primaryKey;index:id_currency;column:id_currency;type:int(11);not null"`
	IDShop     int     `gorm:"primaryKey;column:id_shop;type:int(11);not null"`
	PriceMin   float64 `gorm:"index:price_min;column:price_min;type:decimal(11,5);not null"`
	PriceMax   float64 `gorm:"index:price_max;column:price_max;type:decimal(11,5);not null"`
	IDCountry  int     `gorm:"primaryKey;column:id_country;type:int(11);not null"`
}

type LayeredProductAttribute struct {
	IDAttribute      uint32 `gorm:"primaryKey;uniqueIndex:id_attribute_group;column:id_attribute;type:int(10) unsigned;not null"`
	IDProduct        uint32 `gorm:"primaryKey;uniqueIndex:id_attribute_group;column:id_product;type:int(10) unsigned;not null"`
	IDAttributeGroup uint32 `gorm:"uniqueIndex:id_attribute_group;column:id_attribute_group;type:int(10) unsigned;not null;default:0"`
	IDShop           uint32 `gorm:"primaryKey;uniqueIndex:id_attribute_group;column:id_shop;type:int(10) unsigned;not null;default:1"`
}

type LinkBlock struct {
	IDLinkBlock uint32 `gorm:"primaryKey;column:id_link_block;type:int(10) unsigned;not null"`
	IDHook      uint32 `gorm:"column:id_hook;type:int(1) unsigned"`
	Position    uint32 `gorm:"column:position;type:int(10) unsigned;not null;default:0"`
	Content     string `gorm:"column:content;type:text"`
}

type LinkBlockLang struct {
	IDLinkBlock   uint32 `gorm:"primaryKey;column:id_link_block;type:int(10) unsigned;not null"`
	IDLang        uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name          string `gorm:"column:name;type:varchar(40);not null;default:''"`
	CustomContent string `gorm:"column:custom_content;type:text"`
}

type LinkBlockShop struct {
	IDLinkBlock uint32 `gorm:"primaryKey;column:id_link_block;type:int(10) unsigned;not null"`
	IDShop      uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type Linksmenutop struct {
	IDLinksmenutop uint32 `gorm:"primaryKey;column:id_linksmenutop;type:int(10) unsigned;not null"`
	IDShop         uint32 `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	NewWindow      bool   `gorm:"column:new_window;type:tinyint(1);not null"`
}

type LinksmenutopLang struct {
	IDLinksmenutop uint32 `gorm:"index:id_linksmenutop;column:id_linksmenutop;type:int(11) unsigned;not null"`
	IDLang         uint32 `gorm:"index:id_linksmenutop;column:id_lang;type:int(11) unsigned;not null"`
	IDShop         uint32 `gorm:"index:id_linksmenutop;column:id_shop;type:int(11) unsigned;not null"`
	Label          string `gorm:"column:label;type:varchar(128);not null"`
	Link           string `gorm:"column:link;type:varchar(128);not null"`
}

type Log struct {
	IDLog      uint32    `gorm:"primaryKey;column:id_log;type:int(10) unsigned;not null"`
	Severity   bool      `gorm:"column:severity;type:tinyint(1);not null"`
	ErrorCode  int       `gorm:"column:error_code;type:int(11)"`
	Message    string    `gorm:"column:message;type:text;not null"`
	ObjectType string    `gorm:"column:object_type;type:varchar(32)"`
	ObjectID   uint32    `gorm:"column:object_id;type:int(10) unsigned"`
	IDEmployee uint32    `gorm:"column:id_employee;type:int(10) unsigned"`
	DateAdd    time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd    time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type Mail struct {
	IDMail    uint32    `gorm:"primaryKey;column:id_mail;type:int(11) unsigned;not null"`
	Recipient string    `gorm:"index:recipient;column:recipient;type:varchar(126);not null"`
	Template  string    `gorm:"column:template;type:varchar(62);not null"`
	Subject   string    `gorm:"column:subject;type:varchar(254);not null"`
	IDLang    uint32    `gorm:"column:id_lang;type:int(11) unsigned;not null"`
	DateAdd   time.Time `gorm:"column:date_add;type:timestamp;not null;default:current_timestamp()"`
}

type Manufacturer struct {
	IDManufacturer uint32    `gorm:"primaryKey;column:id_manufacturer;type:int(10) unsigned;not null"`
	Name           string    `gorm:"column:name;type:varchar(64);not null"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd        time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Active         bool      `gorm:"column:active;type:tinyint(1);not null;default:0"`
}

type ManufacturerLang struct {
	IDManufacturer   uint32 `gorm:"primaryKey;column:id_manufacturer;type:int(10) unsigned;not null"`
	IDLang           uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Description      string `gorm:"column:description;type:text"`
	ShortDescription string `gorm:"column:short_description;type:text"`
	MetaTitle        string `gorm:"column:meta_title;type:varchar(255)"`
	MetaKeywords     string `gorm:"column:meta_keywords;type:varchar(255)"`
	MetaDescription  string `gorm:"column:meta_description;type:varchar(512)"`
}

type ManufacturerShop struct {
	IDManufacturer uint32 `gorm:"primaryKey;column:id_manufacturer;type:int(11) unsigned;not null"`
	IDShop         uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type MemcachedServers struct {
	IDMemcachedServer uint32 `gorm:"primaryKey;column:id_memcached_server;type:int(11) unsigned;not null"`
	IP                string `gorm:"column:ip;type:varchar(254);not null"`
	Port              uint32 `gorm:"column:port;type:int(11) unsigned;not null"`
	Weight            uint32 `gorm:"column:weight;type:int(11) unsigned;not null"`
}

type Message struct {
	IDMessage  uint32    `gorm:"primaryKey;column:id_message;type:int(10) unsigned;not null"`
	IDCart     uint32    `gorm:"index:id_cart;column:id_cart;type:int(10) unsigned"`
	IDCustomer uint32    `gorm:"index:id_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDEmployee uint32    `gorm:"index:id_employee;column:id_employee;type:int(10) unsigned"`
	IDOrder    uint32    `gorm:"index:message_order;column:id_order;type:int(10) unsigned;not null"`
	Message    string    `gorm:"column:message;type:text;not null"`
	Private    bool      `gorm:"column:private;type:tinyint(1) unsigned;not null;default:1"`
	DateAdd    time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type MessageReaded struct {
	IDMessage  uint32    `gorm:"primaryKey;column:id_message;type:int(10) unsigned;not null"`
	IDEmployee uint32    `gorm:"primaryKey;column:id_employee;type:int(10) unsigned;not null"`
	DateAdd    time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type Meta struct {
	IDMeta       uint32 `gorm:"primaryKey;column:id_meta;type:int(10) unsigned;not null"`
	Page         string `gorm:"unique;column:page;type:varchar(64);not null"`
	Configurable bool   `gorm:"column:configurable;type:tinyint(1) unsigned;not null;default:1"`
}

type MetaLang struct {
	IDMeta      uint32 `gorm:"primaryKey;column:id_meta;type:int(10) unsigned;not null"`
	IDShop      uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang      uint32 `gorm:"primaryKey;index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	Title       string `gorm:"column:title;type:varchar(128)"`
	Description string `gorm:"column:description;type:varchar(255)"`
	Keywords    string `gorm:"column:keywords;type:varchar(255)"`
	URLRewrite  string `gorm:"column:url_rewrite;type:varchar(254);not null"`
}

type Module struct {
	IDModule uint32 `gorm:"primaryKey;column:id_module;type:int(10) unsigned;not null"`
	Name     string `gorm:"unique;index:name;column:name;type:varchar(64);not null"`
	Active   bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	Version  string `gorm:"column:version;type:varchar(8);not null"`
}

type ModuleAccess struct {
	IDProfile           uint32 `gorm:"primaryKey;column:id_profile;type:int(10) unsigned;not null"`
	IDAuthorizationRole uint32 `gorm:"primaryKey;column:id_authorization_role;type:int(10) unsigned;not null"`
}

type ModuleCarrier struct {
	IDModule    uint32 `gorm:"primaryKey;column:id_module;type:int(10) unsigned;not null"`
	IDShop      uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDReference int    `gorm:"primaryKey;column:id_reference;type:int(11);not null"`
}

type ModuleCountry struct {
	IDModule  uint32 `gorm:"primaryKey;column:id_module;type:int(10) unsigned;not null"`
	IDShop    uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDCountry uint32 `gorm:"primaryKey;column:id_country;type:int(10) unsigned;not null"`
}

type ModuleCurrency struct {
	IDModule   uint32 `gorm:"primaryKey;index:id_module;column:id_module;type:int(10) unsigned;not null"`
	IDShop     uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDCurrency int    `gorm:"primaryKey;column:id_currency;type:int(11);not null"`
}

type ModuleGroup struct {
	IDModule uint32 `gorm:"primaryKey;column:id_module;type:int(10) unsigned;not null"`
	IDShop   uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDGroup  uint32 `gorm:"primaryKey;column:id_group;type:int(11) unsigned;not null"`
}

type ModuleHistory struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	IDEmployee int       `gorm:"column:id_employee;type:int(11);not null"`
	IDModule   int       `gorm:"column:id_module;type:int(11);not null"`
	DateAdd    time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd    time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type ModulePreference struct {
	IDModulePreference int    `gorm:"primaryKey;column:id_module_preference;type:int(11);not null"`
	IDEmployee         int    `gorm:"uniqueIndex:employee_module;column:id_employee;type:int(11);not null"`
	Module             string `gorm:"uniqueIndex:employee_module;column:module;type:varchar(191);not null"`
	Interest           bool   `gorm:"column:interest;type:tinyint(1)"`
	Favorite           bool   `gorm:"column:favorite;type:tinyint(1)"`
}

type ModuleShop struct {
	IDModule     uint32 `gorm:"primaryKey;column:id_module;type:int(11) unsigned;not null"`
	IDShop       uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	EnableDevice bool   `gorm:"column:enable_device;type:tinyint(1);not null;default:7"`
}

type OperatingSystem struct {
	IDOperatingSystem uint32 `gorm:"primaryKey;column:id_operating_system;type:int(10) unsigned;not null"`
	Name              string `gorm:"column:name;type:varchar(64)"`
}

type OrderCarrier struct {
	IDOrderCarrier      int       `gorm:"primaryKey;column:id_order_carrier;type:int(11);not null"`
	IDOrder             uint32    `gorm:"index:id_order;column:id_order;type:int(11) unsigned;not null"`
	IDCarrier           uint32    `gorm:"index:id_carrier;column:id_carrier;type:int(11) unsigned;not null"`
	IDOrderInvoice      uint32    `gorm:"index:id_order_invoice;column:id_order_invoice;type:int(11) unsigned"`
	Weight              float64   `gorm:"column:weight;type:decimal(20,6)"`
	ShippingCostTaxExcl float64   `gorm:"column:shipping_cost_tax_excl;type:decimal(20,6)"`
	ShippingCostTaxIncl float64   `gorm:"column:shipping_cost_tax_incl;type:decimal(20,6)"`
	TrackingNumber      string    `gorm:"column:tracking_number;type:varchar(64)"`
	DateAdd             time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type OrderCartRule struct {
	IDOrderCartRule uint32  `gorm:"primaryKey;column:id_order_cart_rule;type:int(10) unsigned;not null"`
	IDOrder         uint32  `gorm:"index:id_order;column:id_order;type:int(10) unsigned;not null"`
	IDCartRule      uint32  `gorm:"index:id_cart_rule;column:id_cart_rule;type:int(10) unsigned;not null"`
	IDOrderInvoice  uint32  `gorm:"column:id_order_invoice;type:int(10) unsigned;default:0"`
	Name            string  `gorm:"column:name;type:varchar(254);not null"`
	Value           float64 `gorm:"column:value;type:decimal(20,6);not null;default:0.000000"`
	ValueTaxExcl    float64 `gorm:"column:value_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	FreeShipping    bool    `gorm:"column:free_shipping;type:tinyint(1);not null;default:0"`
	Deleted         bool    `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type OrderDetail struct {
	IDOrderDetail             uint32    `gorm:"primaryKey;index:id_order_id_order_detail;column:id_order_detail;type:int(10) unsigned;not null"`
	IDOrder                   uint32    `gorm:"index:order_detail_order;index:id_order_id_order_detail;column:id_order;type:int(10) unsigned;not null"`
	IDOrderInvoice            int       `gorm:"column:id_order_invoice;type:int(11)"`
	IDWarehouse               uint32    `gorm:"column:id_warehouse;type:int(10) unsigned;default:0"`
	IDShop                    uint32    `gorm:"column:id_shop;type:int(11) unsigned;not null"`
	ProductID                 uint32    `gorm:"index:product_id;column:product_id;type:int(10) unsigned;not null"`
	ProductAttributeID        uint32    `gorm:"index:product_id;index:product_attribute_id;column:product_attribute_id;type:int(10) unsigned"`
	IDCustomization           uint32    `gorm:"column:id_customization;type:int(10) unsigned;default:0"`
	ProductName               string    `gorm:"column:product_name;type:varchar(255);not null"`
	ProductQuantity           uint32    `gorm:"column:product_quantity;type:int(10) unsigned;not null;default:0"`
	ProductQuantityInStock    int       `gorm:"column:product_quantity_in_stock;type:int(10);not null;default:0"`
	ProductQuantityRefunded   uint32    `gorm:"column:product_quantity_refunded;type:int(10) unsigned;not null;default:0"`
	ProductQuantityReturn     uint32    `gorm:"column:product_quantity_return;type:int(10) unsigned;not null;default:0"`
	ProductQuantityReinjected uint32    `gorm:"column:product_quantity_reinjected;type:int(10) unsigned;not null;default:0"`
	ProductPrice              float64   `gorm:"column:product_price;type:decimal(20,6);not null;default:0.000000"`
	ReductionPercent          float64   `gorm:"column:reduction_percent;type:decimal(5,2);not null;default:0.00"`
	ReductionAmount           float64   `gorm:"column:reduction_amount;type:decimal(20,6);not null;default:0.000000"`
	ReductionAmountTaxIncl    float64   `gorm:"column:reduction_amount_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	ReductionAmountTaxExcl    float64   `gorm:"column:reduction_amount_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	GroupReduction            float64   `gorm:"column:group_reduction;type:decimal(5,2);not null;default:0.00"`
	ProductQuantityDiscount   float64   `gorm:"column:product_quantity_discount;type:decimal(20,6);not null;default:0.000000"`
	ProductEan13              string    `gorm:"column:product_ean13;type:varchar(13)"`
	ProductIsbn               string    `gorm:"column:product_isbn;type:varchar(32)"`
	ProductUpc                string    `gorm:"column:product_upc;type:varchar(12)"`
	ProductMpn                string    `gorm:"column:product_mpn;type:varchar(40)"`
	ProductReference          string    `gorm:"column:product_reference;type:varchar(64)"`
	ProductSupplierReference  string    `gorm:"column:product_supplier_reference;type:varchar(64)"`
	ProductWeight             float64   `gorm:"column:product_weight;type:decimal(20,6);not null"`
	IDTaxRulesGroup           uint32    `gorm:"index:id_tax_rules_group;column:id_tax_rules_group;type:int(11) unsigned;default:0"`
	TaxComputationMethod      bool      `gorm:"column:tax_computation_method;type:tinyint(1) unsigned;not null;default:0"`
	TaxName                   string    `gorm:"column:tax_name;type:varchar(16);not null"`
	TaxRate                   float64   `gorm:"column:tax_rate;type:decimal(10,3);not null;default:0.000"`
	Ecotax                    float64   `gorm:"column:ecotax;type:decimal(17,6);not null;default:0.000000"`
	EcotaxTaxRate             float64   `gorm:"column:ecotax_tax_rate;type:decimal(5,3);not null;default:0.000"`
	DiscountQuantityApplied   bool      `gorm:"column:discount_quantity_applied;type:tinyint(1);not null;default:0"`
	DownloadHash              string    `gorm:"column:download_hash;type:varchar(255)"`
	DownloadNb                uint32    `gorm:"column:download_nb;type:int(10) unsigned;default:0"`
	DownloadDeadline          time.Time `gorm:"column:download_deadline;type:datetime"`
	TotalPriceTaxIncl         float64   `gorm:"column:total_price_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalPriceTaxExcl         float64   `gorm:"column:total_price_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	UnitPriceTaxIncl          float64   `gorm:"column:unit_price_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	UnitPriceTaxExcl          float64   `gorm:"column:unit_price_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingPriceTaxIncl float64   `gorm:"column:total_shipping_price_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingPriceTaxExcl float64   `gorm:"column:total_shipping_price_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	PurchaseSupplierPrice     float64   `gorm:"column:purchase_supplier_price;type:decimal(20,6);not null;default:0.000000"`
	OriginalProductPrice      float64   `gorm:"column:original_product_price;type:decimal(20,6);not null;default:0.000000"`
	OriginalWholesalePrice    float64   `gorm:"column:original_wholesale_price;type:decimal(20,6);not null;default:0.000000"`
	TotalRefundedTaxExcl      float64   `gorm:"column:total_refunded_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalRefundedTaxIncl      float64   `gorm:"column:total_refunded_tax_incl;type:decimal(20,6);not null;default:0.000000"`
}

type OrderDetailTax struct {
	IDOrderDetail int     `gorm:"index:id_order_detail;column:id_order_detail;type:int(11);not null"`
	IDTax         int     `gorm:"index:id_tax;column:id_tax;type:int(11);not null"`
	UnitAmount    float64 `gorm:"column:unit_amount;type:decimal(16,6);not null;default:0.000000"`
	TotalAmount   float64 `gorm:"column:total_amount;type:decimal(16,6);not null;default:0.000000"`
}

type OrderHistory struct {
	IDOrderHistory uint32    `gorm:"primaryKey;column:id_order_history;type:int(10) unsigned;not null"`
	IDEmployee     uint32    `gorm:"index:id_employee;column:id_employee;type:int(10) unsigned;not null"`
	IDOrder        uint32    `gorm:"index:order_history_order;column:id_order;type:int(10) unsigned;not null"`
	IDOrderState   uint32    `gorm:"index:id_order_state;column:id_order_state;type:int(10) unsigned;not null"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type OrderInvoice struct {
	IDOrderInvoice               uint32    `gorm:"primaryKey;column:id_order_invoice;type:int(11) unsigned;not null"`
	IDOrder                      int       `gorm:"index:id_order;column:id_order;type:int(11);not null"`
	Number                       int       `gorm:"column:number;type:int(11);not null"`
	DeliveryNumber               int       `gorm:"column:delivery_number;type:int(11);not null"`
	DeliveryDate                 time.Time `gorm:"column:delivery_date;type:datetime"`
	TotalDiscountTaxExcl         float64   `gorm:"column:total_discount_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalDiscountTaxIncl         float64   `gorm:"column:total_discount_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalPaidTaxExcl             float64   `gorm:"column:total_paid_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalPaidTaxIncl             float64   `gorm:"column:total_paid_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalProducts                float64   `gorm:"column:total_products;type:decimal(20,6);not null;default:0.000000"`
	TotalProductsWt              float64   `gorm:"column:total_products_wt;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingTaxExcl         float64   `gorm:"column:total_shipping_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingTaxIncl         float64   `gorm:"column:total_shipping_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	ShippingTaxComputationMethod uint32    `gorm:"column:shipping_tax_computation_method;type:int(10) unsigned;not null"`
	TotalWrappingTaxExcl         float64   `gorm:"column:total_wrapping_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalWrappingTaxIncl         float64   `gorm:"column:total_wrapping_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	ShopAddress                  string    `gorm:"column:shop_address;type:text"`
	Note                         string    `gorm:"column:note;type:text"`
	DateAdd                      time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type OrderInvoicePayment struct {
	IDOrderInvoice uint32 `gorm:"primaryKey;column:id_order_invoice;type:int(11) unsigned;not null"`
	IDOrderPayment uint32 `gorm:"primaryKey;index:order_payment;column:id_order_payment;type:int(11) unsigned;not null"`
	IDOrder        uint32 `gorm:"index:id_order;column:id_order;type:int(11) unsigned;not null"`
}

type OrderInvoiceTax struct {
	IDOrderInvoice int     `gorm:"column:id_order_invoice;type:int(11);not null"`
	Type           string  `gorm:"column:type;type:varchar(15);not null"`
	IDTax          int     `gorm:"index:id_tax;column:id_tax;type:int(11);not null"`
	Amount         float64 `gorm:"column:amount;type:decimal(10,6);not null;default:0.000000"`
}

type OrderMessage struct {
	IDOrderMessage uint32    `gorm:"primaryKey;column:id_order_message;type:int(10) unsigned;not null"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type OrderMessageLang struct {
	IDOrderMessage uint32 `gorm:"primaryKey;column:id_order_message;type:int(10) unsigned;not null"`
	IDLang         uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name           string `gorm:"column:name;type:varchar(128);not null"`
	Message        string `gorm:"column:message;type:text;not null"`
}

type OrderPayment struct {
	IDOrderPayment int       `gorm:"primaryKey;column:id_order_payment;type:int(11);not null"`
	OrderReference string    `gorm:"index:order_reference;column:order_reference;type:varchar(9)"`
	IDCurrency     uint32    `gorm:"column:id_currency;type:int(10) unsigned;not null"`
	Amount         float64   `gorm:"column:amount;type:decimal(20,6);not null"`
	PaymentMethod  string    `gorm:"column:payment_method;type:varchar(255);not null"`
	ConversionRate float64   `gorm:"column:conversion_rate;type:decimal(13,6);not null;default:1.000000"`
	TransactionID  string    `gorm:"column:transaction_id;type:varchar(254)"`
	CardNumber     string    `gorm:"column:card_number;type:varchar(254)"`
	CardBrand      string    `gorm:"column:card_brand;type:varchar(254)"`
	CardExpiration string    `gorm:"column:card_expiration;type:char(7)"`
	CardHolder     string    `gorm:"column:card_holder;type:varchar(254)"`
	DateAdd        time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type OrderReturn struct {
	IDOrderReturn uint32    `gorm:"primaryKey;column:id_order_return;type:int(10) unsigned;not null"`
	IDCustomer    uint32    `gorm:"index:order_return_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDOrder       uint32    `gorm:"index:id_order;column:id_order;type:int(10) unsigned;not null"`
	State         bool      `gorm:"column:state;type:tinyint(1) unsigned;not null;default:1"`
	Question      string    `gorm:"column:question;type:text;not null"`
	DateAdd       time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd       time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type OrderReturnDetail struct {
	IDOrderReturn   uint32 `gorm:"primaryKey;column:id_order_return;type:int(10) unsigned;not null"`
	IDOrderDetail   uint32 `gorm:"primaryKey;column:id_order_detail;type:int(10) unsigned;not null"`
	IDCustomization uint32 `gorm:"primaryKey;column:id_customization;type:int(10) unsigned;not null;default:0"`
	ProductQuantity uint32 `gorm:"column:product_quantity;type:int(10) unsigned;not null;default:0"`
}

type OrderReturnState struct {
	IDOrderReturnState uint32 `gorm:"primaryKey;column:id_order_return_state;type:int(10) unsigned;not null"`
	Color              string `gorm:"column:color;type:varchar(32)"`
}

type OrderReturnStateLang struct {
	IDOrderReturnState uint32 `gorm:"primaryKey;column:id_order_return_state;type:int(10) unsigned;not null"`
	IDLang             uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name               string `gorm:"column:name;type:varchar(64);not null"`
}

type OrderSlip struct {
	IDOrderSlip          uint32    `gorm:"primaryKey;column:id_order_slip;type:int(10) unsigned;not null"`
	ConversionRate       float64   `gorm:"column:conversion_rate;type:decimal(13,6);not null;default:1.000000"`
	IDCustomer           uint32    `gorm:"index:order_slip_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDOrder              uint32    `gorm:"index:id_order;column:id_order;type:int(10) unsigned;not null"`
	TotalProductsTaxExcl float64   `gorm:"column:total_products_tax_excl;type:decimal(20,6)"`
	TotalProductsTaxIncl float64   `gorm:"column:total_products_tax_incl;type:decimal(20,6)"`
	TotalShippingTaxExcl float64   `gorm:"column:total_shipping_tax_excl;type:decimal(20,6)"`
	TotalShippingTaxIncl float64   `gorm:"column:total_shipping_tax_incl;type:decimal(20,6)"`
	ShippingCost         uint8     `gorm:"column:shipping_cost;type:tinyint(3) unsigned;not null;default:0"`
	Amount               float64   `gorm:"column:amount;type:decimal(20,6);not null;default:0.000000"`
	ShippingCostAmount   float64   `gorm:"column:shipping_cost_amount;type:decimal(20,6);not null;default:0.000000"`
	Partial              bool      `gorm:"column:partial;type:tinyint(1);not null"`
	OrderSlipType        bool      `gorm:"column:order_slip_type;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd              time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd              time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type OrderSlipDetail struct {
	IDOrderSlip       uint32  `gorm:"primaryKey;column:id_order_slip;type:int(10) unsigned;not null"`
	IDOrderDetail     uint32  `gorm:"primaryKey;column:id_order_detail;type:int(10) unsigned;not null"`
	ProductQuantity   uint32  `gorm:"column:product_quantity;type:int(10) unsigned;not null;default:0"`
	UnitPriceTaxExcl  float64 `gorm:"column:unit_price_tax_excl;type:decimal(20,6)"`
	UnitPriceTaxIncl  float64 `gorm:"column:unit_price_tax_incl;type:decimal(20,6)"`
	TotalPriceTaxExcl float64 `gorm:"column:total_price_tax_excl;type:decimal(20,6)"`
	TotalPriceTaxIncl float64 `gorm:"column:total_price_tax_incl;type:decimal(20,6)"`
	AmountTaxExcl     float64 `gorm:"column:amount_tax_excl;type:decimal(20,6)"`
	AmountTaxIncl     float64 `gorm:"column:amount_tax_incl;type:decimal(20,6)"`
}

type OrderSlipDetailTax struct {
	IDOrderSlipDetail uint32  `gorm:"index:id_order_slip_detail;column:id_order_slip_detail;type:int(11) unsigned;not null"`
	IDTax             uint32  `gorm:"index:id_tax;column:id_tax;type:int(11) unsigned;not null"`
	UnitAmount        float64 `gorm:"column:unit_amount;type:decimal(16,6);not null;default:0.000000"`
	TotalAmount       float64 `gorm:"column:total_amount;type:decimal(16,6);not null;default:0.000000"`
}

type OrderState struct {
	IDOrderState uint32 `gorm:"primaryKey;column:id_order_state;type:int(10) unsigned;not null"`
	Invoice      bool   `gorm:"column:invoice;type:tinyint(1) unsigned;default:0"`
	SendEmail    bool   `gorm:"column:send_email;type:tinyint(1) unsigned;not null;default:0"`
	ModuleName   string `gorm:"index:module_name;column:module_name;type:varchar(255)"`
	Color        string `gorm:"column:color;type:varchar(32)"`
	Unremovable  bool   `gorm:"column:unremovable;type:tinyint(1) unsigned;not null"`
	Hidden       bool   `gorm:"column:hidden;type:tinyint(1) unsigned;not null;default:0"`
	Logable      bool   `gorm:"column:logable;type:tinyint(1);not null;default:0"`
	Delivery     bool   `gorm:"column:delivery;type:tinyint(1) unsigned;not null;default:0"`
	Shipped      bool   `gorm:"column:shipped;type:tinyint(1) unsigned;not null;default:0"`
	Paid         bool   `gorm:"column:paid;type:tinyint(1) unsigned;not null;default:0"`
	PdfInvoice   bool   `gorm:"column:pdf_invoice;type:tinyint(1) unsigned;not null;default:0"`
	PdfDelivery  bool   `gorm:"column:pdf_delivery;type:tinyint(1) unsigned;not null;default:0"`
	Deleted      bool   `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type OrderStateLang struct {
	IDOrderState uint32 `gorm:"primaryKey;column:id_order_state;type:int(10) unsigned;not null"`
	IDLang       uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name         string `gorm:"column:name;type:varchar(64);not null"`
	Template     string `gorm:"column:template;type:varchar(64);not null"`
}

type Orders struct {
	IDOrder               uint32    `gorm:"primaryKey;column:id_order;type:int(10) unsigned;not null"`
	Reference             string    `gorm:"index:reference;column:reference;type:varchar(9)"`
	IDShopGroup           uint32    `gorm:"index:id_shop_group;column:id_shop_group;type:int(11) unsigned;not null;default:1"`
	IDShop                uint32    `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDCarrier             uint32    `gorm:"index:id_carrier;column:id_carrier;type:int(10) unsigned;not null"`
	IDLang                uint32    `gorm:"index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	IDCustomer            uint32    `gorm:"index:id_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDCart                uint32    `gorm:"index:id_cart;column:id_cart;type:int(10) unsigned;not null"`
	IDCurrency            uint32    `gorm:"index:id_currency;column:id_currency;type:int(10) unsigned;not null"`
	IDAddressDelivery     uint32    `gorm:"index:id_address_delivery;column:id_address_delivery;type:int(10) unsigned;not null"`
	IDAddressInvoice      uint32    `gorm:"index:id_address_invoice;column:id_address_invoice;type:int(10) unsigned;not null"`
	CurrentState          uint32    `gorm:"index:current_state;column:current_state;type:int(10) unsigned;not null"`
	SecureKey             string    `gorm:"column:secure_key;type:varchar(32);not null;default:-1"`
	Payment               string    `gorm:"column:payment;type:varchar(255);not null"`
	ConversionRate        float64   `gorm:"column:conversion_rate;type:decimal(13,6);not null;default:1.000000"`
	Module                string    `gorm:"column:module;type:varchar(255)"`
	Recyclable            bool      `gorm:"column:recyclable;type:tinyint(1) unsigned;not null;default:0"`
	Gift                  bool      `gorm:"column:gift;type:tinyint(1) unsigned;not null;default:0"`
	GiftMessage           string    `gorm:"column:gift_message;type:text"`
	MobileTheme           bool      `gorm:"column:mobile_theme;type:tinyint(1);not null;default:0"`
	ShippingNumber        string    `gorm:"column:shipping_number;type:varchar(64)"`
	TotalDiscounts        float64   `gorm:"column:total_discounts;type:decimal(20,6);not null;default:0.000000"`
	TotalDiscountsTaxIncl float64   `gorm:"column:total_discounts_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalDiscountsTaxExcl float64   `gorm:"column:total_discounts_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalPaid             float64   `gorm:"column:total_paid;type:decimal(20,6);not null;default:0.000000"`
	TotalPaidTaxIncl      float64   `gorm:"column:total_paid_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalPaidTaxExcl      float64   `gorm:"column:total_paid_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	TotalPaidReal         float64   `gorm:"column:total_paid_real;type:decimal(20,6);not null;default:0.000000"`
	TotalProducts         float64   `gorm:"column:total_products;type:decimal(20,6);not null;default:0.000000"`
	TotalProductsWt       float64   `gorm:"column:total_products_wt;type:decimal(20,6);not null;default:0.000000"`
	TotalShipping         float64   `gorm:"column:total_shipping;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingTaxIncl  float64   `gorm:"column:total_shipping_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalShippingTaxExcl  float64   `gorm:"column:total_shipping_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	CarrierTaxRate        float64   `gorm:"column:carrier_tax_rate;type:decimal(10,3);not null;default:0.000"`
	TotalWrapping         float64   `gorm:"column:total_wrapping;type:decimal(20,6);not null;default:0.000000"`
	TotalWrappingTaxIncl  float64   `gorm:"column:total_wrapping_tax_incl;type:decimal(20,6);not null;default:0.000000"`
	TotalWrappingTaxExcl  float64   `gorm:"column:total_wrapping_tax_excl;type:decimal(20,6);not null;default:0.000000"`
	RoundMode             bool      `gorm:"column:round_mode;type:tinyint(1);not null;default:2"`
	RoundType             bool      `gorm:"column:round_type;type:tinyint(1);not null;default:1"`
	InvoiceNumber         uint32    `gorm:"index:invoice_number;column:invoice_number;type:int(10) unsigned;not null;default:0"`
	DeliveryNumber        uint32    `gorm:"column:delivery_number;type:int(10) unsigned;not null;default:0"`
	InvoiceDate           time.Time `gorm:"column:invoice_date;type:datetime;not null"`
	DeliveryDate          time.Time `gorm:"column:delivery_date;type:datetime;not null"`
	Valid                 uint32    `gorm:"column:valid;type:int(1) unsigned;not null;default:0"`
	DateAdd               time.Time `gorm:"index:date_add;column:date_add;type:datetime;not null"`
	DateUpd               time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type Pack struct {
	IDProductPack          uint32 `gorm:"primaryKey;column:id_product_pack;type:int(10) unsigned;not null"`
	IDProductItem          uint32 `gorm:"primaryKey;index:product_item;column:id_product_item;type:int(10) unsigned;not null"`
	IDProductAttributeItem uint32 `gorm:"primaryKey;index:product_item;column:id_product_attribute_item;type:int(10) unsigned;not null"`
	Quantity               uint32 `gorm:"column:quantity;type:int(10) unsigned;not null;default:1"`
}

type Page struct {
	IDPage     uint32 `gorm:"primaryKey;column:id_page;type:int(10) unsigned;not null"`
	IDPageType uint32 `gorm:"index:id_page_type;column:id_page_type;type:int(10) unsigned;not null"`
	IDObject   uint32 `gorm:"index:id_object;column:id_object;type:int(10) unsigned"`
}

type PageType struct {
	IDPageType uint32 `gorm:"primaryKey;column:id_page_type;type:int(10) unsigned;not null"`
	Name       string `gorm:"index:name;column:name;type:varchar(255);not null"`
}

type PageViewed struct {
	IDPage      uint32 `gorm:"primaryKey;column:id_page;type:int(10) unsigned;not null"`
	IDShopGroup uint32 `gorm:"column:id_shop_group;type:int(10) unsigned;not null;default:1"`
	IDShop      uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDDateRange uint32 `gorm:"primaryKey;column:id_date_range;type:int(10) unsigned;not null"`
	Counter     uint32 `gorm:"column:counter;type:int(10) unsigned;not null"`
}

type Pagenotfound struct {
	IDPagenotfound uint32    `gorm:"primaryKey;column:id_pagenotfound;type:int(10) unsigned;not null"`
	IDShop         uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDShopGroup    uint32    `gorm:"column:id_shop_group;type:int(10) unsigned;not null;default:1"`
	RequestURI     string    `gorm:"column:request_uri;type:varchar(256);not null"`
	HTTPReferer    string    `gorm:"column:http_referer;type:varchar(256);not null"`
	DateAdd        time.Time `gorm:"index:date_add;column:date_add;type:datetime;not null"`
}

type Product struct {
	IDProduct               uint32         `gorm:"primaryKey;index:product_manufacturer;column:id_product;type:int(10) unsigned;not null"`
	IDSupplier              uint32         `gorm:"index:product_supplier;column:id_supplier;type:int(10) unsigned"`
	IDManufacturer          uint32         `gorm:"index:product_manufacturer;column:id_manufacturer;type:int(10) unsigned"`
	IDCategoryDefault       uint32         `gorm:"index:id_category_default;column:id_category_default;type:int(10) unsigned"`
	IDShopDefault           uint32         `gorm:"column:id_shop_default;type:int(10) unsigned;not null;default:1"`
	IDTaxRulesGroup         uint32         `gorm:"column:id_tax_rules_group;type:int(11) unsigned;not null"`
	OnSale                  bool           `gorm:"column:on_sale;type:tinyint(1) unsigned;not null;default:0"`
	OnlineOnly              bool           `gorm:"column:online_only;type:tinyint(1) unsigned;not null;default:0"`
	Ean13                   string         `gorm:"column:ean13;type:varchar(13)"`
	Isbn                    string         `gorm:"column:isbn;type:varchar(32)"`
	Upc                     string         `gorm:"column:upc;type:varchar(12)"`
	Mpn                     string         `gorm:"column:mpn;type:varchar(40)"`
	Ecotax                  float64        `gorm:"column:ecotax;type:decimal(17,6);not null;default:0.000000"`
	Quantity                int            `gorm:"column:quantity;type:int(10);not null;default:0"`
	MinimalQuantity         uint32         `gorm:"column:minimal_quantity;type:int(10) unsigned;not null;default:1"`
	LowStockThreshold       int            `gorm:"column:low_stock_threshold;type:int(10)"`
	LowStockAlert           bool           `gorm:"column:low_stock_alert;type:tinyint(1);not null;default:0"`
	Price                   float64        `gorm:"column:price;type:decimal(20,6);not null;default:0.000000"`
	WholesalePrice          float64        `gorm:"column:wholesale_price;type:decimal(20,6);not null;default:0.000000"`
	Unity                   string         `gorm:"column:unity;type:varchar(255)"`
	UnitPriceRatio          float64        `gorm:"column:unit_price_ratio;type:decimal(20,6);not null;default:0.000000"`
	AdditionalShippingCost  float64        `gorm:"column:additional_shipping_cost;type:decimal(20,6);not null;default:0.000000"`
	Reference               string         `gorm:"index:reference_idx;column:reference;type:varchar(64)"`
	SupplierReference       string         `gorm:"index:supplier_reference_idx;column:supplier_reference;type:varchar(64)"`
	Location                string         `gorm:"column:location;type:varchar(64)"`
	Width                   float64        `gorm:"column:width;type:decimal(20,6);not null;default:0.000000"`
	Height                  float64        `gorm:"column:height;type:decimal(20,6);not null;default:0.000000"`
	Depth                   float64        `gorm:"column:depth;type:decimal(20,6);not null;default:0.000000"`
	Weight                  float64        `gorm:"column:weight;type:decimal(20,6);not null;default:0.000000"`
	OutOfStock              uint32         `gorm:"column:out_of_stock;type:int(10) unsigned;not null;default:2"`
	AdditionalDeliveryTimes bool           `gorm:"column:additional_delivery_times;type:tinyint(1) unsigned;not null;default:1"`
	QuantityDiscount        bool           `gorm:"column:quantity_discount;type:tinyint(1);default:0"`
	Customizable            int8           `gorm:"column:customizable;type:tinyint(2);not null;default:0"`
	UploadableFiles         int8           `gorm:"column:uploadable_files;type:tinyint(4);not null;default:0"`
	TextFields              int8           `gorm:"column:text_fields;type:tinyint(4);not null;default:0"`
	Active                  bool           `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	RedirectType            string         `gorm:"column:redirect_type;type:enum('','404','301-product','302-product','301-category','302-category');not null;default:''"`
	IDTypeRedirected        uint32         `gorm:"column:id_type_redirected;type:int(10) unsigned;not null;default:0"`
	AvailableForOrder       bool           `gorm:"column:available_for_order;type:tinyint(1);not null;default:1"`
	AvailableDate           datatypes.Date `gorm:"column:available_date;type:date"`
	ShowCondition           bool           `gorm:"column:show_condition;type:tinyint(1);not null;default:0"`
	Condition               string         `gorm:"column:condition;type:enum('new','used','refurbished');not null;default:new"`
	ShowPrice               bool           `gorm:"column:show_price;type:tinyint(1);not null;default:1"`
	Indexed                 bool           `gorm:"index:indexed;column:indexed;type:tinyint(1);not null;default:0"`
	Visibility              string         `gorm:"column:visibility;type:enum('both','catalog','search','none');not null;default:both"`
	CacheIsPack             bool           `gorm:"column:cache_is_pack;type:tinyint(1);not null;default:0"`
	CacheHasAttachments     bool           `gorm:"column:cache_has_attachments;type:tinyint(1);not null;default:0"`
	IsVirtual               bool           `gorm:"column:is_virtual;type:tinyint(1);not null;default:0"`
	CacheDefaultAttribute   uint32         `gorm:"column:cache_default_attribute;type:int(10) unsigned"`
	DateAdd                 time.Time      `gorm:"index:date_add;column:date_add;type:datetime;not null"`
	DateUpd                 time.Time      `gorm:"index:state;column:date_upd;type:datetime;not null"`
	AdvancedStockManagement bool           `gorm:"column:advanced_stock_management;type:tinyint(1);not null;default:0"`
	PackStockType           uint32         `gorm:"column:pack_stock_type;type:int(11) unsigned;not null;default:3"`
	State                   uint32         `gorm:"index:state;column:state;type:int(11) unsigned;not null;default:1"`
}

type ProductAttachment struct {
	IDProduct    uint32 `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDAttachment uint32 `gorm:"primaryKey;column:id_attachment;type:int(10) unsigned;not null"`
}

type ProductAttribute struct {
	IDProductAttribute uint32         `gorm:"primaryKey;index:id_product_id_product_attribute;column:id_product_attribute;type:int(10) unsigned;not null"`
	IDProduct          uint32         `gorm:"uniqueIndex:product_default;index:product_attribute_product;index:id_product_id_product_attribute;column:id_product;type:int(10) unsigned;not null"`
	Reference          string         `gorm:"index:reference;column:reference;type:varchar(64)"`
	SupplierReference  string         `gorm:"index:supplier_reference;column:supplier_reference;type:varchar(64)"`
	Location           string         `gorm:"column:location;type:varchar(64)"`
	Ean13              string         `gorm:"column:ean13;type:varchar(13)"`
	Isbn               string         `gorm:"column:isbn;type:varchar(32)"`
	Upc                string         `gorm:"column:upc;type:varchar(12)"`
	Mpn                string         `gorm:"column:mpn;type:varchar(40)"`
	WholesalePrice     float64        `gorm:"column:wholesale_price;type:decimal(20,6);not null;default:0.000000"`
	Price              float64        `gorm:"column:price;type:decimal(20,6);not null;default:0.000000"`
	Ecotax             float64        `gorm:"column:ecotax;type:decimal(17,6);not null;default:0.000000"`
	Quantity           int            `gorm:"column:quantity;type:int(10);not null;default:0"`
	Weight             float64        `gorm:"column:weight;type:decimal(20,6);not null;default:0.000000"`
	UnitPriceImpact    float64        `gorm:"column:unit_price_impact;type:decimal(20,6);not null;default:0.000000"`
	DefaultOn          bool           `gorm:"uniqueIndex:product_default;column:default_on;type:tinyint(1) unsigned"`
	MinimalQuantity    uint32         `gorm:"column:minimal_quantity;type:int(10) unsigned;not null;default:1"`
	LowStockThreshold  int            `gorm:"column:low_stock_threshold;type:int(10)"`
	LowStockAlert      bool           `gorm:"column:low_stock_alert;type:tinyint(1);not null;default:0"`
	AvailableDate      datatypes.Date `gorm:"column:available_date;type:date"`
}

type ProductAttributeCombination struct {
	IDAttribute        uint32 `gorm:"primaryKey;column:id_attribute;type:int(10) unsigned;not null"`
	IDProductAttribute uint32 `gorm:"primaryKey;index:id_product_attribute;column:id_product_attribute;type:int(10) unsigned;not null"`
}

type ProductAttributeImage struct {
	IDProductAttribute uint32 `gorm:"primaryKey;column:id_product_attribute;type:int(10) unsigned;not null"`
	IDImage            uint32 `gorm:"primaryKey;index:id_image;column:id_image;type:int(10) unsigned;not null"`
}

type ProductAttributeShop struct {
	IDProduct          uint32         `gorm:"uniqueIndex:id_product;column:id_product;type:int(10) unsigned;not null"`
	IDProductAttribute uint32         `gorm:"primaryKey;column:id_product_attribute;type:int(10) unsigned;not null"`
	IDShop             uint32         `gorm:"primaryKey;uniqueIndex:id_product;column:id_shop;type:int(10) unsigned;not null"`
	WholesalePrice     float64        `gorm:"column:wholesale_price;type:decimal(20,6);not null;default:0.000000"`
	Price              float64        `gorm:"column:price;type:decimal(20,6);not null;default:0.000000"`
	Ecotax             float64        `gorm:"column:ecotax;type:decimal(17,6);not null;default:0.000000"`
	Weight             float64        `gorm:"column:weight;type:decimal(20,6);not null;default:0.000000"`
	UnitPriceImpact    float64        `gorm:"column:unit_price_impact;type:decimal(20,6);not null;default:0.000000"`
	DefaultOn          bool           `gorm:"uniqueIndex:id_product;column:default_on;type:tinyint(1) unsigned"`
	MinimalQuantity    uint32         `gorm:"column:minimal_quantity;type:int(10) unsigned;not null;default:1"`
	LowStockThreshold  int            `gorm:"column:low_stock_threshold;type:int(10)"`
	LowStockAlert      bool           `gorm:"column:low_stock_alert;type:tinyint(1);not null;default:0"`
	AvailableDate      datatypes.Date `gorm:"column:available_date;type:date"`
}

type ProductCarrier struct {
	IDProduct          uint32 `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDCarrierReference uint32 `gorm:"primaryKey;column:id_carrier_reference;type:int(10) unsigned;not null"`
	IDShop             uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type ProductComment struct {
	IDProductComment int       `gorm:"primaryKey;column:id_product_comment;type:int(11);not null"`
	IDProduct        int       `gorm:"column:id_product;type:int(11);not null"`
	IDCustomer       int       `gorm:"column:id_customer;type:int(11);not null"`
	IDGuest          int       `gorm:"column:id_guest;type:int(11);not null"`
	CustomerName     string    `gorm:"column:customer_name;type:varchar(64);not null"`
	Title            string    `gorm:"column:title;type:varchar(64);not null"`
	Content          string    `gorm:"column:content;type:longtext;not null"`
	Grade            int       `gorm:"column:grade;type:int(11);not null"`
	Validate         bool      `gorm:"column:validate;type:tinyint(1);not null"`
	Deleted          bool      `gorm:"column:deleted;type:tinyint(1);not null"`
	DateAdd          time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type ProductCommentCriterion struct {
	IDProductCommentCriterion     int  `gorm:"primaryKey;column:id_product_comment_criterion;type:int(11);not null"`
	IDProductCommentCriterionType int  `gorm:"column:id_product_comment_criterion_type;type:int(11);not null"`
	Active                        bool `gorm:"column:active;type:tinyint(1);not null"`
}

type ProductCommentCriterionCategory struct {
	IDProductCommentCriterion uint32 `gorm:"primaryKey;column:id_product_comment_criterion;type:int(10) unsigned;not null"`
	IDCategory                uint32 `gorm:"primaryKey;index:id_category;column:id_category;type:int(10) unsigned;not null"`
}

type ProductCommentCriterionLang struct {
	IDProductCommentCriterion uint32 `gorm:"primaryKey;column:id_product_comment_criterion;type:int(11) unsigned;not null"`
	IDLang                    uint32 `gorm:"primaryKey;column:id_lang;type:int(11) unsigned;not null"`
	Name                      string `gorm:"column:name;type:varchar(64);not null"`
}

type ProductCommentCriterionProduct struct {
	IDProduct                 uint32 `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDProductCommentCriterion uint32 `gorm:"primaryKey;index:id_product_comment_criterion;column:id_product_comment_criterion;type:int(10) unsigned;not null"`
}

type ProductCommentGrade struct {
	IDProductComment          int `gorm:"primaryKey;index:IDX_367426EBACF38A54;column:id_product_comment;type:int(11);not null"`
	IDProductCommentCriterion int `gorm:"primaryKey;index:IDX_367426EB8375853C;column:id_product_comment_criterion;type:int(11);not null"`
	Grade                     int `gorm:"column:grade;type:int(11);not null"`
}

type ProductCommentReport struct {
	IDCustomer       int `gorm:"primaryKey;column:id_customer;type:int(11);not null"`
	IDProductComment int `gorm:"primaryKey;index:IDX_D22C9649ACF38A54;column:id_product_comment;type:int(11);not null"`
}

type ProductCommentUsefulness struct {
	IDCustomer       int  `gorm:"primaryKey;column:id_customer;type:int(11);not null"`
	IDProductComment int  `gorm:"primaryKey;index:IDX_E681E112ACF38A54;column:id_product_comment;type:int(11);not null"`
	Usefulness       bool `gorm:"column:usefulness;type:tinyint(1);not null"`
}

type ProductCountryTax struct {
	IDProduct int `gorm:"primaryKey;column:id_product;type:int(11);not null"`
	IDCountry int `gorm:"primaryKey;column:id_country;type:int(11);not null"`
	IDTax     int `gorm:"column:id_tax;type:int(11);not null"`
}

type ProductDownload struct {
	IDProductDownload uint32    `gorm:"primaryKey;column:id_product_download;type:int(10) unsigned;not null"`
	IDProduct         uint32    `gorm:"column:id_product;type:int(10) unsigned;not null"`
	DisplayFilename   string    `gorm:"column:display_filename;type:varchar(255)"`
	Filename          string    `gorm:"column:filename;type:varchar(255)"`
	DateAdd           time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateExpiration    time.Time `gorm:"column:date_expiration;type:datetime"`
	NbDaysAccessible  uint32    `gorm:"column:nb_days_accessible;type:int(10) unsigned"`
	NbDownloadable    uint32    `gorm:"column:nb_downloadable;type:int(10) unsigned;default:1"`
	Active            bool      `gorm:"column:active;type:tinyint(1) unsigned;not null;default:1"`
	IsShareable       bool      `gorm:"column:is_shareable;type:tinyint(1) unsigned;not null;default:0"`
}

type ProductGroupReductionCache struct {
	IDProduct uint32  `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDGroup   uint32  `gorm:"primaryKey;column:id_group;type:int(10) unsigned;not null"`
	Reduction float64 `gorm:"column:reduction;type:decimal(5,4);not null"`
}

type ProductLang struct {
	IDProduct        uint32 `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDShop           uint32 `gorm:"primaryKey;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang           uint32 `gorm:"primaryKey;index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	Description      string `gorm:"column:description;type:text"`
	DescriptionShort string `gorm:"column:description_short;type:text"`
	LinkRewrite      string `gorm:"column:link_rewrite;type:varchar(128);not null"`
	MetaDescription  string `gorm:"column:meta_description;type:varchar(512)"`
	MetaKeywords     string `gorm:"column:meta_keywords;type:varchar(255)"`
	MetaTitle        string `gorm:"column:meta_title;type:varchar(128)"`
	Name             string `gorm:"index:name;column:name;type:varchar(128);not null"`
	AvailableNow     string `gorm:"column:available_now;type:varchar(255)"`
	AvailableLater   string `gorm:"column:available_later;type:varchar(255)"`
	DeliveryInStock  string `gorm:"column:delivery_in_stock;type:varchar(255)"`
	DeliveryOutStock string `gorm:"column:delivery_out_stock;type:varchar(255)"`
}

type ProductSale struct {
	IDProduct uint32         `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	Quantity  uint32         `gorm:"index:quantity;column:quantity;type:int(10) unsigned;not null;default:0"`
	SaleNbr   uint32         `gorm:"column:sale_nbr;type:int(10) unsigned;not null;default:0"`
	DateUpd   datatypes.Date `gorm:"column:date_upd;type:date"`
}

type ProductShop struct {
	IDProduct               uint32         `gorm:"primaryKey;index:indexed;column:id_product;type:int(10) unsigned;not null"`
	IDShop                  uint32         `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
	IDCategoryDefault       uint32         `gorm:"index:id_category_default;column:id_category_default;type:int(10) unsigned"`
	IDTaxRulesGroup         uint32         `gorm:"column:id_tax_rules_group;type:int(11) unsigned;not null"`
	OnSale                  bool           `gorm:"column:on_sale;type:tinyint(1) unsigned;not null;default:0"`
	OnlineOnly              bool           `gorm:"column:online_only;type:tinyint(1) unsigned;not null;default:0"`
	Ecotax                  float64        `gorm:"column:ecotax;type:decimal(17,6);not null;default:0.000000"`
	MinimalQuantity         uint32         `gorm:"column:minimal_quantity;type:int(10) unsigned;not null;default:1"`
	LowStockThreshold       int            `gorm:"column:low_stock_threshold;type:int(10)"`
	LowStockAlert           bool           `gorm:"column:low_stock_alert;type:tinyint(1);not null;default:0"`
	Price                   float64        `gorm:"column:price;type:decimal(20,6);not null;default:0.000000"`
	WholesalePrice          float64        `gorm:"column:wholesale_price;type:decimal(20,6);not null;default:0.000000"`
	Unity                   string         `gorm:"column:unity;type:varchar(255)"`
	UnitPriceRatio          float64        `gorm:"column:unit_price_ratio;type:decimal(20,6);not null;default:0.000000"`
	AdditionalShippingCost  float64        `gorm:"column:additional_shipping_cost;type:decimal(20,6);not null;default:0.000000"`
	Customizable            int8           `gorm:"column:customizable;type:tinyint(2);not null;default:0"`
	UploadableFiles         int8           `gorm:"column:uploadable_files;type:tinyint(4);not null;default:0"`
	TextFields              int8           `gorm:"column:text_fields;type:tinyint(4);not null;default:0"`
	Active                  bool           `gorm:"index:date_add;index:indexed;column:active;type:tinyint(1) unsigned;not null;default:0"`
	RedirectType            string         `gorm:"column:redirect_type;type:enum('','404','301-product','302-product','301-category','302-category');not null;default:''"`
	IDTypeRedirected        uint32         `gorm:"column:id_type_redirected;type:int(10) unsigned;not null;default:0"`
	AvailableForOrder       bool           `gorm:"column:available_for_order;type:tinyint(1);not null;default:1"`
	AvailableDate           datatypes.Date `gorm:"column:available_date;type:date"`
	ShowCondition           bool           `gorm:"column:show_condition;type:tinyint(1);not null;default:1"`
	Condition               string         `gorm:"column:condition;type:enum('new','used','refurbished');not null;default:new"`
	ShowPrice               bool           `gorm:"column:show_price;type:tinyint(1);not null;default:1"`
	Indexed                 bool           `gorm:"index:indexed;column:indexed;type:tinyint(1);not null;default:0"`
	Visibility              string         `gorm:"index:date_add;column:visibility;type:enum('both','catalog','search','none');not null;default:both"`
	CacheDefaultAttribute   uint32         `gorm:"column:cache_default_attribute;type:int(10) unsigned"`
	AdvancedStockManagement bool           `gorm:"column:advanced_stock_management;type:tinyint(1);not null;default:0"`
	DateAdd                 time.Time      `gorm:"index:date_add;column:date_add;type:datetime;not null"`
	DateUpd                 time.Time      `gorm:"column:date_upd;type:datetime;not null"`
	PackStockType           uint32         `gorm:"column:pack_stock_type;type:int(11) unsigned;not null;default:3"`
}

type ProductSupplier struct {
	IDProductSupplier        uint32  `gorm:"primaryKey;column:id_product_supplier;type:int(11) unsigned;not null"`
	IDProduct                uint32  `gorm:"uniqueIndex:id_product;index:id_supplier;column:id_product;type:int(11) unsigned;not null"`
	IDProductAttribute       uint32  `gorm:"uniqueIndex:id_product;column:id_product_attribute;type:int(11) unsigned;not null;default:0"`
	IDSupplier               uint32  `gorm:"uniqueIndex:id_product;index:id_supplier;column:id_supplier;type:int(11) unsigned;not null"`
	ProductSupplierReference string  `gorm:"column:product_supplier_reference;type:varchar(64)"`
	ProductSupplierPriceTe   float64 `gorm:"column:product_supplier_price_te;type:decimal(20,6);not null;default:0.000000"`
	IDCurrency               uint32  `gorm:"column:id_currency;type:int(11) unsigned;not null"`
}

type ProductTag struct {
	IDProduct uint32 `gorm:"primaryKey;column:id_product;type:int(10) unsigned;not null"`
	IDTag     uint32 `gorm:"primaryKey;index:id_tag;index:id_lang;column:id_tag;type:int(10) unsigned;not null"`
	IDLang    uint32 `gorm:"index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
}

type Profile struct {
	IDProfile uint32 `gorm:"primaryKey;column:id_profile;type:int(10) unsigned;not null"`
}

type ProfileLang struct {
	IDLang    uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	IDProfile uint32 `gorm:"primaryKey;column:id_profile;type:int(10) unsigned;not null"`
	Name      string `gorm:"column:name;type:varchar(128);not null"`
}

type PscheckoutCart struct {
	IDPscheckoutCart          uint32    `gorm:"primaryKey;column:id_pscheckout_cart;type:int(10) unsigned;not null"`
	IDCart                    uint32    `gorm:"column:id_cart;type:int(10) unsigned;not null"`
	PaypalIntent              string    `gorm:"column:paypal_intent;type:varchar(20);default:CAPTURE"`
	PaypalOrder               string    `gorm:"column:paypal_order;type:varchar(20)"`
	PaypalStatus              string    `gorm:"column:paypal_status;type:varchar(20)"`
	PaypalFunding             string    `gorm:"column:paypal_funding;type:varchar(20)"`
	PaypalToken               string    `gorm:"column:paypal_token;type:varchar(1024)"`
	PaypalTokenExpire         time.Time `gorm:"column:paypal_token_expire;type:datetime"`
	PaypalAuthorizationExpire time.Time `gorm:"column:paypal_authorization_expire;type:datetime"`
	IsExpressCheckout         bool      `gorm:"column:isExpressCheckout;type:tinyint(1) unsigned;not null;default:0"`
	IsHostedFields            bool      `gorm:"column:isHostedFields;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd                   time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd                   time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type PscheckoutFundingSource struct {
	Name     string `gorm:"primaryKey;column:name;type:varchar(20);not null"`
	Active   bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	Position uint8  `gorm:"column:position;type:tinyint(2) unsigned;not null"`
	IDShop   uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(10) unsigned;not null"`
}

type PscheckoutOrderMatrice struct {
	IDOrderMatrice    uint32 `gorm:"primaryKey;column:id_order_matrice;type:int(10) unsigned;not null"`
	IDOrderPrestashop uint32 `gorm:"column:id_order_prestashop;type:int(10) unsigned;not null"`
	IDOrderPaypal     string `gorm:"column:id_order_paypal;type:varchar(20);not null"`
}

type PsgdprConsent struct {
	IDGdprConsent uint32    `gorm:"primaryKey;column:id_gdpr_consent;type:int(10) unsigned;not null"`
	IDModule      uint32    `gorm:"primaryKey;column:id_module;type:int(10) unsigned;not null"`
	Active        int       `gorm:"column:active;type:int(10);not null"`
	Error         int       `gorm:"column:error;type:int(10)"`
	ErrorMessage  string    `gorm:"column:error_message;type:text"`
	DateAdd       time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd       time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type PsgdprConsentLang struct {
	IDGdprConsent uint32 `gorm:"primaryKey;column:id_gdpr_consent;type:int(10) unsigned;not null"`
	IDLang        uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Message       string `gorm:"column:message;type:text"`
	IDShop        uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
}

type PsgdprLog struct {
	IDGdprLog   uint32    `gorm:"primaryKey;column:id_gdpr_log;type:int(10) unsigned;not null"`
	IDCustomer  uint32    `gorm:"column:id_customer;type:int(10) unsigned"`
	IDGuest     uint32    `gorm:"column:id_guest;type:int(10) unsigned"`
	ClientName  string    `gorm:"column:client_name;type:varchar(250)"`
	IDModule    uint32    `gorm:"column:id_module;type:int(10) unsigned;not null"`
	RequestType int       `gorm:"column:request_type;type:int(10);not null"`
	DateAdd     time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd     time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type Psreassurance struct {
	IDPsreassurance uint32    `gorm:"primaryKey;column:id_psreassurance;type:int(10) unsigned;not null"`
	Icon            string    `gorm:"column:icon;type:varchar(255)"`
	CustomIcon      string    `gorm:"column:custom_icon;type:varchar(255)"`
	Status          uint32    `gorm:"column:status;type:int(10) unsigned;not null"`
	Position        uint32    `gorm:"column:position;type:int(10) unsigned;not null"`
	IDShop          uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null"`
	TypeLink        uint32    `gorm:"column:type_link;type:int(10) unsigned"`
	IDCms           uint32    `gorm:"column:id_cms;type:int(10) unsigned"`
	DateAdd         time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd         time.Time `gorm:"column:date_upd;type:datetime"`
}

type PsreassuranceLang struct {
	IDPsreassurance uint32 `gorm:"primaryKey;column:id_psreassurance;type:int(10) unsigned;not null"`
	IDLang          uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null"`
	Title           string `gorm:"column:title;type:varchar(255);not null"`
	Description     string `gorm:"column:description;type:varchar(255);not null"`
	Link            string `gorm:"column:link;type:varchar(255);not null"`
}

type QuickAccess struct {
	IDQuickAccess uint32 `gorm:"primaryKey;column:id_quick_access;type:int(10) unsigned;not null"`
	NewWindow     bool   `gorm:"column:new_window;type:tinyint(1);not null;default:0"`
	Link          string `gorm:"column:link;type:varchar(255);not null"`
}

type QuickAccessLang struct {
	IDQuickAccess uint32 `gorm:"primaryKey;column:id_quick_access;type:int(10) unsigned;not null"`
	IDLang        uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name          string `gorm:"column:name;type:varchar(32);not null"`
}

type RangePrice struct {
	IDRangePrice uint32  `gorm:"primaryKey;column:id_range_price;type:int(10) unsigned;not null"`
	IDCarrier    uint32  `gorm:"uniqueIndex:id_carrier;column:id_carrier;type:int(10) unsigned;not null"`
	Delimiter1   float64 `gorm:"uniqueIndex:id_carrier;column:delimiter1;type:decimal(20,6);not null"`
	Delimiter2   float64 `gorm:"uniqueIndex:id_carrier;column:delimiter2;type:decimal(20,6);not null"`
}

type RangeWeight struct {
	IDRangeWeight uint32  `gorm:"primaryKey;column:id_range_weight;type:int(10) unsigned;not null"`
	IDCarrier     uint32  `gorm:"uniqueIndex:id_carrier;column:id_carrier;type:int(10) unsigned;not null"`
	Delimiter1    float64 `gorm:"uniqueIndex:id_carrier;column:delimiter1;type:decimal(20,6);not null"`
	Delimiter2    float64 `gorm:"uniqueIndex:id_carrier;column:delimiter2;type:decimal(20,6);not null"`
}

type Referrer struct {
	IDReferrer           uint32    `gorm:"primaryKey;column:id_referrer;type:int(10) unsigned;not null"`
	Name                 string    `gorm:"column:name;type:varchar(64);not null"`
	Passwd               string    `gorm:"column:passwd;type:varchar(255)"`
	HTTPRefererRegexp    string    `gorm:"column:http_referer_regexp;type:varchar(64)"`
	HTTPRefererLike      string    `gorm:"column:http_referer_like;type:varchar(64)"`
	RequestURIRegexp     string    `gorm:"column:request_uri_regexp;type:varchar(64)"`
	RequestURILike       string    `gorm:"column:request_uri_like;type:varchar(64)"`
	HTTPRefererRegexpNot string    `gorm:"column:http_referer_regexp_not;type:varchar(64)"`
	HTTPRefererLikeNot   string    `gorm:"column:http_referer_like_not;type:varchar(64)"`
	RequestURIRegexpNot  string    `gorm:"column:request_uri_regexp_not;type:varchar(64)"`
	RequestURILikeNot    string    `gorm:"column:request_uri_like_not;type:varchar(64)"`
	BaseFee              float64   `gorm:"column:base_fee;type:decimal(5,2);not null;default:0.00"`
	PercentFee           float64   `gorm:"column:percent_fee;type:decimal(5,2);not null;default:0.00"`
	ClickFee             float64   `gorm:"column:click_fee;type:decimal(5,2);not null;default:0.00"`
	DateAdd              time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type ReferrerCache struct {
	IDConnectionsSource uint32 `gorm:"primaryKey;column:id_connections_source;type:int(11) unsigned;not null"`
	IDReferrer          uint32 `gorm:"primaryKey;column:id_referrer;type:int(11) unsigned;not null"`
}

type ReferrerShop struct {
	IDReferrer         uint32  `gorm:"primaryKey;column:id_referrer;type:int(10) unsigned;not null"`
	IDShop             uint32  `gorm:"primaryKey;column:id_shop;type:int(10) unsigned;not null;default:1"`
	CacheVisitors      int     `gorm:"column:cache_visitors;type:int(11)"`
	CacheVisits        int     `gorm:"column:cache_visits;type:int(11)"`
	CachePages         int     `gorm:"column:cache_pages;type:int(11)"`
	CacheRegistrations int     `gorm:"column:cache_registrations;type:int(11)"`
	CacheOrders        int     `gorm:"column:cache_orders;type:int(11)"`
	CacheSales         float64 `gorm:"column:cache_sales;type:decimal(17,2)"`
	CacheRegRate       float64 `gorm:"column:cache_rps_rate;type:decimal(5,4)"`
	CacheOrderRate     float64 `gorm:"column:cache_order_rate;type:decimal(5,4)"`
}

type RequestSQL struct {
	IDRequestSQL int    `gorm:"primaryKey;column:id_request_sql;type:int(11);not null"`
	Name         string `gorm:"column:name;type:varchar(200);not null"`
	SQL          string `gorm:"column:sql;type:text;not null"`
}

type RequiredField struct {
	IDRequiredField int    `gorm:"primaryKey;column:id_required_field;type:int(11);not null"`
	ObjectName      string `gorm:"index:object_name;column:object_name;type:varchar(32);not null"`
	FieldName       string `gorm:"column:field_name;type:varchar(32);not null"`
}

type Risk struct {
	IDRisk  uint32 `gorm:"primaryKey;column:id_risk;type:int(11) unsigned;not null"`
	Percent int8   `gorm:"column:percent;type:tinyint(3);not null"`
	Color   string `gorm:"column:color;type:varchar(32)"`
}

type RiskLang struct {
	IDRisk uint32 `gorm:"primaryKey;index:id_risk;column:id_risk;type:int(10) unsigned;not null"`
	IDLang uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name   string `gorm:"column:name;type:varchar(20);not null"`
}

type SearchEngine struct {
	IDSearchEngine uint32 `gorm:"primaryKey;column:id_search_engine;type:int(10) unsigned;not null"`
	Server         string `gorm:"column:server;type:varchar(64);not null"`
	Getvar         string `gorm:"column:getvar;type:varchar(16);not null"`
}

type SearchIndex struct {
	IDProduct uint32 `gorm:"primaryKey;index:id_product;column:id_product;type:int(11) unsigned;not null"`
	IDWord    uint32 `gorm:"primaryKey;column:id_word;type:int(11) unsigned;not null"`
	Weight    uint16 `gorm:"index:id_product;column:weight;type:smallint(4) unsigned;not null;default:1"`
}

type SearchWord struct {
	IDWord uint32 `gorm:"primaryKey;column:id_word;type:int(10) unsigned;not null"`
	IDShop uint32 `gorm:"uniqueIndex:id_lang;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDLang uint32 `gorm:"uniqueIndex:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	Word   string `gorm:"uniqueIndex:id_lang;column:word;type:varchar(30);not null"`
}

type Sekeyword struct {
	IDSekeyword uint32    `gorm:"primaryKey;column:id_sekeyword;type:int(10) unsigned;not null"`
	IDShop      uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDShopGroup uint32    `gorm:"column:id_shop_group;type:int(10) unsigned;not null;default:1"`
	Keyword     string    `gorm:"column:keyword;type:varchar(256);not null"`
	DateAdd     time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type Shop struct {
	IDShop      int    `gorm:"primaryKey;column:id_shop;type:int(11);not null"`
	IDShopGroup int    `gorm:"index:IDX_667E487AF5C9E40;column:id_shop_group;type:int(11);not null"`
	Name        string `gorm:"column:name;type:varchar(64);not null"`
	IDCategory  int    `gorm:"column:id_category;type:int(11);not null"`
	ThemeName   string `gorm:"column:theme_name;type:varchar(255);not null"`
	Active      bool   `gorm:"column:active;type:tinyint(1);not null"`
	Deleted     bool   `gorm:"column:deleted;type:tinyint(1);not null"`
}

type ShopGroup struct {
	IDShopGroup   int    `gorm:"primaryKey;column:id_shop_group;type:int(11);not null"`
	Name          string `gorm:"column:name;type:varchar(64);not null"`
	ShareCustomer bool   `gorm:"column:share_customer;type:tinyint(1);not null"`
	ShareOrder    bool   `gorm:"column:share_order;type:tinyint(1);not null"`
	ShareStock    bool   `gorm:"column:share_stock;type:tinyint(1);not null"`
	Active        bool   `gorm:"column:active;type:tinyint(1);not null"`
	Deleted       bool   `gorm:"column:deleted;type:tinyint(1);not null"`
}

type ShopURL struct {
	IDShopURL   uint32 `gorm:"primaryKey;column:id_shop_url;type:int(11) unsigned;not null"`
	IDShop      uint32 `gorm:"index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	Domain      string `gorm:"uniqueIndex:full_shop_url;column:domain;type:varchar(150);not null"`
	DomainSsl   string `gorm:"uniqueIndex:full_shop_url_ssl;column:domain_ssl;type:varchar(150);not null"`
	PhysicalURI string `gorm:"uniqueIndex:full_shop_url;uniqueIndex:full_shop_url_ssl;column:physical_uri;type:varchar(64);not null"`
	VirtualURI  string `gorm:"uniqueIndex:full_shop_url;uniqueIndex:full_shop_url_ssl;column:virtual_uri;type:varchar(64);not null"`
	Main        bool   `gorm:"index:id_shop;column:main;type:tinyint(1);not null"`
	Active      bool   `gorm:"column:active;type:tinyint(1);not null"`
}

type SmartyCache struct {
	IDSmartyCache string    `gorm:"primaryKey;column:id_smarty_cache;type:char(40);not null"`
	Name          string    `gorm:"index:name;column:name;type:char(40);not null"`
	CacheID       string    `gorm:"index:cache_id;column:cache_id;type:varchar(254)"`
	Modified      time.Time `gorm:"index:modified;column:modified;type:timestamp;not null;default:current_timestamp()"`
	Content       string    `gorm:"column:content;type:longtext;not null"`
}

type SmartyLastFlush struct {
	Type      string    `gorm:"primaryKey;column:type;type:enum('compile','template');not null"`
	LastFlush time.Time `gorm:"column:last_flush;type:datetime;not null;default:0000-00-00 00:00:00"`
}

type SmartyLazyCache struct {
	TemplateHash string    `gorm:"primaryKey;column:template_hash;type:varchar(32);not null;default:''"`
	CacheID      string    `gorm:"primaryKey;column:cache_id;type:varchar(191);not null;default:''"`
	CompileID    string    `gorm:"primaryKey;column:compile_id;type:varchar(32);not null;default:''"`
	Filepath     string    `gorm:"column:filepath;type:varchar(255);not null;default:''"`
	LastUpdate   time.Time `gorm:"column:last_update;type:datetime;not null;default:0000-00-00 00:00:00"`
}

type SpecificPrice struct {
	IDSpecificPrice     uint32    `gorm:"primaryKey;column:id_specific_price;type:int(10) unsigned;not null"`
	IDSpecificPriceRule uint32    `gorm:"uniqueIndex:id_product_2;index:id_specific_price_rule;column:id_specific_price_rule;type:int(11) unsigned;not null"`
	IDCart              uint32    `gorm:"uniqueIndex:id_product_2;index:id_cart;column:id_cart;type:int(11) unsigned;not null"`
	IDProduct           uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;column:id_product;type:int(10) unsigned;not null"`
	IDShop              uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;index:id_shop;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDShopGroup         uint32    `gorm:"uniqueIndex:id_product_2;column:id_shop_group;type:int(11) unsigned;not null"`
	IDCurrency          uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;column:id_currency;type:int(10) unsigned;not null"`
	IDCountry           uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;column:id_country;type:int(10) unsigned;not null"`
	IDGroup             uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;column:id_group;type:int(10) unsigned;not null"`
	IDCustomer          uint32    `gorm:"uniqueIndex:id_product_2;index:id_product;index:id_customer;column:id_customer;type:int(10) unsigned;not null"`
	IDProductAttribute  uint32    `gorm:"uniqueIndex:id_product_2;index:id_product_attribute;column:id_product_attribute;type:int(10) unsigned;not null"`
	Price               float64   `gorm:"column:price;type:decimal(20,6);not null"`
	FromQuantity        string    `gorm:"uniqueIndex:id_product_2;index:id_product;index:from_quantity;column:from_quantity;type:mediumint(8) unsigned;not null"`
	Reduction           float64   `gorm:"column:reduction;type:decimal(20,6);not null"`
	ReductionTax        bool      `gorm:"column:reduction_tax;type:tinyint(1);not null;default:1"`
	ReductionType       string    `gorm:"column:reduction_type;type:enum('amount','percentage');not null"`
	From                time.Time `gorm:"uniqueIndex:id_product_2;index:id_product;index:from;column:from;type:datetime;not null"`
	To                  time.Time `gorm:"uniqueIndex:id_product_2;index:id_product;index:to;column:to;type:datetime;not null"`
}

type SpecificPricePriority struct {
	IDSpecificPricePriority int    `gorm:"primaryKey;column:id_specific_price_priority;type:int(11);not null"`
	IDProduct               int    `gorm:"primaryKey;unique;column:id_product;type:int(11);not null"`
	Priority                string `gorm:"column:priority;type:varchar(80);not null"`
}

type SpecificPriceRule struct {
	IDSpecificPriceRule uint32    `gorm:"primaryKey;column:id_specific_price_rule;type:int(10) unsigned;not null"`
	Name                string    `gorm:"column:name;type:varchar(255);not null"`
	IDShop              uint32    `gorm:"index:id_product;column:id_shop;type:int(11) unsigned;not null;default:1"`
	IDCurrency          uint32    `gorm:"index:id_product;column:id_currency;type:int(10) unsigned;not null"`
	IDCountry           uint32    `gorm:"index:id_product;column:id_country;type:int(10) unsigned;not null"`
	IDGroup             uint32    `gorm:"index:id_product;column:id_group;type:int(10) unsigned;not null"`
	FromQuantity        string    `gorm:"index:id_product;column:from_quantity;type:mediumint(8) unsigned;not null"`
	Price               float64   `gorm:"column:price;type:decimal(20,6)"`
	Reduction           float64   `gorm:"column:reduction;type:decimal(20,6);not null"`
	ReductionTax        bool      `gorm:"column:reduction_tax;type:tinyint(1);not null;default:1"`
	ReductionType       string    `gorm:"column:reduction_type;type:enum('amount','percentage');not null"`
	From                time.Time `gorm:"index:id_product;column:from;type:datetime;not null"`
	To                  time.Time `gorm:"index:id_product;column:to;type:datetime;not null"`
}

type SpecificPriceRuleCondition struct {
	IDSpecificPriceRuleCondition      uint32 `gorm:"primaryKey;column:id_specific_price_rule_condition;type:int(11) unsigned;not null"`
	IDSpecificPriceRuleConditionGroup uint32 `gorm:"index:id_specific_price_rule_condition_group;column:id_specific_price_rule_condition_group;type:int(11) unsigned;not null"`
	Type                              string `gorm:"column:type;type:varchar(255);not null"`
	Value                             string `gorm:"column:value;type:varchar(255);not null"`
}

type SpecificPriceRuleConditionGroup struct {
	IDSpecificPriceRuleConditionGroup uint32 `gorm:"primaryKey;column:id_specific_price_rule_condition_group;type:int(11) unsigned;not null"`
	IDSpecificPriceRule               uint32 `gorm:"primaryKey;column:id_specific_price_rule;type:int(11) unsigned;not null"`
}

type State struct {
	IDState     uint32 `gorm:"primaryKey;column:id_state;type:int(10) unsigned;not null"`
	IDCountry   uint32 `gorm:"index:id_country;column:id_country;type:int(11) unsigned;not null"`
	IDZone      uint32 `gorm:"index:id_zone;column:id_zone;type:int(11) unsigned;not null"`
	Name        string `gorm:"index:name;column:name;type:varchar(64);not null"`
	IsoCode     string `gorm:"column:iso_code;type:varchar(7);not null"`
	TaxBehavior int16  `gorm:"column:tax_behavior;type:smallint(1);not null;default:0"`
	Active      bool   `gorm:"column:active;type:tinyint(1);not null;default:0"`
}

type Statssearch struct {
	IDStatssearch uint32    `gorm:"primaryKey;column:id_statssearch;type:int(10) unsigned;not null"`
	IDShop        uint32    `gorm:"column:id_shop;type:int(10) unsigned;not null;default:1"`
	IDShopGroup   uint32    `gorm:"column:id_shop_group;type:int(10) unsigned;not null;default:1"`
	Keywords      string    `gorm:"column:keywords;type:varchar(255);not null"`
	Results       int       `gorm:"column:results;type:int(6);not null;default:0"`
	DateAdd       time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type Stock struct {
	IDStock            uint32  `gorm:"primaryKey;column:id_stock;type:int(11) unsigned;not null"`
	IDWarehouse        uint32  `gorm:"index:id_warehouse;column:id_warehouse;type:int(11) unsigned;not null"`
	IDProduct          uint32  `gorm:"index:id_product;column:id_product;type:int(11) unsigned;not null"`
	IDProductAttribute uint32  `gorm:"index:id_product_attribute;column:id_product_attribute;type:int(11) unsigned;not null"`
	Reference          string  `gorm:"column:reference;type:varchar(64);not null"`
	Ean13              string  `gorm:"column:ean13;type:varchar(13)"`
	Isbn               string  `gorm:"column:isbn;type:varchar(32)"`
	Upc                string  `gorm:"column:upc;type:varchar(12)"`
	Mpn                string  `gorm:"column:mpn;type:varchar(40)"`
	PhysicalQuantity   uint32  `gorm:"column:physical_quantity;type:int(11) unsigned;not null"`
	UsableQuantity     uint32  `gorm:"column:usable_quantity;type:int(11) unsigned;not null"`
	PriceTe            float64 `gorm:"column:price_te;type:decimal(20,6);default:0.000000"`
}

type StockAvailable struct {
	IDStockAvailable   uint32 `gorm:"primaryKey;column:id_stock_available;type:int(11) unsigned;not null"`
	IDProduct          uint32 `gorm:"uniqueIndex:product_sqlstock;index:id_product;column:id_product;type:int(11) unsigned;not null"`
	IDProductAttribute uint32 `gorm:"uniqueIndex:product_sqlstock;index:id_product_attribute;column:id_product_attribute;type:int(11) unsigned;not null"`
	IDShop             uint32 `gorm:"uniqueIndex:product_sqlstock;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	IDShopGroup        uint32 `gorm:"uniqueIndex:product_sqlstock;index:id_shop_group;column:id_shop_group;type:int(11) unsigned;not null"`
	Quantity           int    `gorm:"column:quantity;type:int(10);not null;default:0"`
	PhysicalQuantity   int    `gorm:"column:physical_quantity;type:int(11);not null;default:0"`
	ReservedQuantity   int    `gorm:"column:reserved_quantity;type:int(11);not null;default:0"`
	DependsOnStock     bool   `gorm:"column:depends_on_stock;type:tinyint(1) unsigned;not null;default:0"`
	OutOfStock         bool   `gorm:"column:out_of_stock;type:tinyint(1) unsigned;not null;default:0"`
	Location           string `gorm:"column:location;type:varchar(255);not null;default:''"`
}

type StockMvt struct {
	IDStockMvt        int64     `gorm:"primaryKey;column:id_stock_mvt;type:bigint(20);not null"`
	IDStock           int       `gorm:"index:id_stock;column:id_stock;type:int(11);not null"`
	IDOrder           int       `gorm:"column:id_order;type:int(11)"`
	IDSupplyOrder     int       `gorm:"column:id_supply_order;type:int(11)"`
	IDStockMvtReason  int       `gorm:"index:id_stock_mvt_reason;column:id_stock_mvt_reason;type:int(11);not null"`
	IDEmployee        int       `gorm:"column:id_employee;type:int(11);not null"`
	EmployeeLastname  string    `gorm:"column:employee_lastname;type:varchar(32)"`
	EmployeeFirstname string    `gorm:"column:employee_firstname;type:varchar(32)"`
	PhysicalQuantity  int       `gorm:"column:physical_quantity;type:int(11);not null"`
	DateAdd           time.Time `gorm:"column:date_add;type:datetime;not null"`
	Sign              int16     `gorm:"column:sign;type:smallint(6);not null;default:1"`
	PriceTe           float64   `gorm:"column:price_te;type:decimal(20,6);default:0.000000"`
	LastWa            float64   `gorm:"column:last_wa;type:decimal(20,6);default:0.000000"`
	CurrentWa         float64   `gorm:"column:current_wa;type:decimal(20,6);default:0.000000"`
	Referer           int64     `gorm:"column:referer;type:bigint(20)"`
}

type StockMvtReason struct {
	IDStockMvtReason uint32    `gorm:"primaryKey;column:id_stock_mvt_reason;type:int(11) unsigned;not null"`
	Sign             bool      `gorm:"column:sign;type:tinyint(1);not null;default:1"`
	DateAdd          time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd          time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Deleted          bool      `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type StockMvtReasonLang struct {
	IDStockMvtReason uint32 `gorm:"primaryKey;column:id_stock_mvt_reason;type:int(11) unsigned;not null"`
	IDLang           uint32 `gorm:"primaryKey;column:id_lang;type:int(11) unsigned;not null"`
	Name             string `gorm:"column:name;type:varchar(255);not null"`
}

type Store struct {
	IDStore   uint32    `gorm:"primaryKey;column:id_store;type:int(10) unsigned;not null"`
	IDCountry uint32    `gorm:"column:id_country;type:int(10) unsigned;not null"`
	IDState   uint32    `gorm:"column:id_state;type:int(10) unsigned"`
	City      string    `gorm:"column:city;type:varchar(64);not null"`
	Postcode  string    `gorm:"column:postcode;type:varchar(12);not null"`
	Latitude  float64   `gorm:"column:latitude;type:decimal(13,8)"`
	Longitude float64   `gorm:"column:longitude;type:decimal(13,8)"`
	Phone     string    `gorm:"column:phone;type:varchar(16)"`
	Fax       string    `gorm:"column:fax;type:varchar(16)"`
	Email     string    `gorm:"column:email;type:varchar(255)"`
	Active    bool      `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
	DateAdd   time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd   time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type StoreLang struct {
	IDStore  uint32 `gorm:"primaryKey;column:id_store;type:int(11) unsigned;not null"`
	IDLang   uint32 `gorm:"primaryKey;column:id_lang;type:int(11) unsigned;not null"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Address1 string `gorm:"column:address1;type:varchar(255);not null"`
	Address2 string `gorm:"column:address2;type:varchar(255)"`
	Hours    string `gorm:"column:hours;type:text"`
	Note     string `gorm:"column:note;type:text"`
}

type StoreShop struct {
	IDStore uint32 `gorm:"primaryKey;column:id_store;type:int(11) unsigned;not null"`
	IDShop  uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Supplier struct {
	IDSupplier uint32    `gorm:"primaryKey;column:id_supplier;type:int(10) unsigned;not null"`
	Name       string    `gorm:"column:name;type:varchar(64);not null"`
	DateAdd    time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd    time.Time `gorm:"column:date_upd;type:datetime;not null"`
	Active     bool      `gorm:"column:active;type:tinyint(1);not null;default:0"`
}

type SupplierLang struct {
	IDSupplier      uint32 `gorm:"primaryKey;column:id_supplier;type:int(10) unsigned;not null"`
	IDLang          uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Description     string `gorm:"column:description;type:text"`
	MetaTitle       string `gorm:"column:meta_title;type:varchar(255)"`
	MetaKeywords    string `gorm:"column:meta_keywords;type:varchar(255)"`
	MetaDescription string `gorm:"column:meta_description;type:varchar(512)"`
}

type SupplierShop struct {
	IDSupplier uint32 `gorm:"primaryKey;column:id_supplier;type:int(11) unsigned;not null"`
	IDShop     uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type SupplyOrder struct {
	IDSupplyOrder        uint32    `gorm:"primaryKey;column:id_supply_order;type:int(11) unsigned;not null"`
	IDSupplier           uint32    `gorm:"index:id_supplier;column:id_supplier;type:int(11) unsigned;not null"`
	SupplierName         string    `gorm:"column:supplier_name;type:varchar(64);not null"`
	IDLang               uint32    `gorm:"column:id_lang;type:int(11) unsigned;not null"`
	IDWarehouse          uint32    `gorm:"index:id_warehouse;column:id_warehouse;type:int(11) unsigned;not null"`
	IDSupplyOrderState   uint32    `gorm:"column:id_supply_order_state;type:int(11) unsigned;not null"`
	IDCurrency           uint32    `gorm:"column:id_currency;type:int(11) unsigned;not null"`
	IDRefCurrency        uint32    `gorm:"column:id_ref_currency;type:int(11) unsigned;not null"`
	Reference            string    `gorm:"index:reference;column:reference;type:varchar(64);not null"`
	DateAdd              time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd              time.Time `gorm:"column:date_upd;type:datetime;not null"`
	DateDeliveryExpected time.Time `gorm:"column:date_delivery_expected;type:datetime"`
	TotalTe              float64   `gorm:"column:total_te;type:decimal(20,6);default:0.000000"`
	TotalWithDiscountTe  float64   `gorm:"column:total_with_discount_te;type:decimal(20,6);default:0.000000"`
	TotalTax             float64   `gorm:"column:total_tax;type:decimal(20,6);default:0.000000"`
	TotalTi              float64   `gorm:"column:total_ti;type:decimal(20,6);default:0.000000"`
	DiscountRate         float64   `gorm:"column:discount_rate;type:decimal(20,6);default:0.000000"`
	DiscountValueTe      float64   `gorm:"column:discount_value_te;type:decimal(20,6);default:0.000000"`
	IsTemplate           bool      `gorm:"column:is_template;type:tinyint(1);default:0"`
}

type SupplyOrderDetail struct {
	IDSupplyOrderDetail       uint32  `gorm:"primaryKey;column:id_supply_order_detail;type:int(11) unsigned;not null"`
	IDSupplyOrder             uint32  `gorm:"index:id_supply_order;column:id_supply_order;type:int(11) unsigned;not null"`
	IDCurrency                uint32  `gorm:"column:id_currency;type:int(11) unsigned;not null"`
	IDProduct                 uint32  `gorm:"index:id_supply_order;index:id_product_product_attribute;column:id_product;type:int(11) unsigned;not null"`
	IDProductAttribute        uint32  `gorm:"index:id_product_attribute;index:id_product_product_attribute;column:id_product_attribute;type:int(11) unsigned;not null"`
	Reference                 string  `gorm:"column:reference;type:varchar(64);not null"`
	SupplierReference         string  `gorm:"column:supplier_reference;type:varchar(64);not null"`
	Name                      string  `gorm:"column:name;type:varchar(128);not null"`
	Ean13                     string  `gorm:"column:ean13;type:varchar(13)"`
	Isbn                      string  `gorm:"column:isbn;type:varchar(32)"`
	Upc                       string  `gorm:"column:upc;type:varchar(12)"`
	Mpn                       string  `gorm:"column:mpn;type:varchar(40)"`
	ExchangeRate              float64 `gorm:"column:exchange_rate;type:decimal(20,6);default:0.000000"`
	UnitPriceTe               float64 `gorm:"column:unit_price_te;type:decimal(20,6);default:0.000000"`
	QuantityExpected          uint32  `gorm:"column:quantity_expected;type:int(11) unsigned;not null"`
	QuantityReceived          uint32  `gorm:"column:quantity_received;type:int(11) unsigned;not null"`
	PriceTe                   float64 `gorm:"column:price_te;type:decimal(20,6);default:0.000000"`
	DiscountRate              float64 `gorm:"column:discount_rate;type:decimal(20,6);default:0.000000"`
	DiscountValueTe           float64 `gorm:"column:discount_value_te;type:decimal(20,6);default:0.000000"`
	PriceWithDiscountTe       float64 `gorm:"column:price_with_discount_te;type:decimal(20,6);default:0.000000"`
	TaxRate                   float64 `gorm:"column:tax_rate;type:decimal(20,6);default:0.000000"`
	TaxValue                  float64 `gorm:"column:tax_value;type:decimal(20,6);default:0.000000"`
	PriceTi                   float64 `gorm:"column:price_ti;type:decimal(20,6);default:0.000000"`
	TaxValueWithOrderDiscount float64 `gorm:"column:tax_value_with_order_discount;type:decimal(20,6);default:0.000000"`
	PriceWithOrderDiscountTe  float64 `gorm:"column:price_with_order_discount_te;type:decimal(20,6);default:0.000000"`
}

type SupplyOrderHistory struct {
	IDSupplyOrderHistory uint32    `gorm:"primaryKey;column:id_supply_order_history;type:int(11) unsigned;not null"`
	IDSupplyOrder        uint32    `gorm:"index:id_supply_order;column:id_supply_order;type:int(11) unsigned;not null"`
	IDEmployee           uint32    `gorm:"index:id_employee;column:id_employee;type:int(11) unsigned;not null"`
	EmployeeLastname     string    `gorm:"column:employee_lastname;type:varchar(255);default:''"`
	EmployeeFirstname    string    `gorm:"column:employee_firstname;type:varchar(255);default:''"`
	IDState              uint32    `gorm:"index:id_state;column:id_state;type:int(11) unsigned;not null"`
	DateAdd              time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type SupplyOrderReceiptHistory struct {
	IDSupplyOrderReceiptHistory uint32    `gorm:"primaryKey;column:id_supply_order_receipt_history;type:int(11) unsigned;not null"`
	IDSupplyOrderDetail         uint32    `gorm:"index:id_supply_order_detail;column:id_supply_order_detail;type:int(11) unsigned;not null"`
	IDEmployee                  uint32    `gorm:"column:id_employee;type:int(11) unsigned;not null"`
	EmployeeLastname            string    `gorm:"column:employee_lastname;type:varchar(255);default:''"`
	EmployeeFirstname           string    `gorm:"column:employee_firstname;type:varchar(255);default:''"`
	IDSupplyOrderState          uint32    `gorm:"index:id_supply_order_state;column:id_supply_order_state;type:int(11) unsigned;not null"`
	Quantity                    uint32    `gorm:"column:quantity;type:int(11) unsigned;not null"`
	DateAdd                     time.Time `gorm:"column:date_add;type:datetime;not null"`
}

type SupplyOrderState struct {
	IDSupplyOrderState uint32 `gorm:"primaryKey;column:id_supply_order_state;type:int(11) unsigned;not null"`
	DeliveryNote       bool   `gorm:"column:delivery_note;type:tinyint(1);not null;default:0"`
	Editable           bool   `gorm:"column:editable;type:tinyint(1);not null;default:0"`
	ReceiptState       bool   `gorm:"column:receipt_state;type:tinyint(1);not null;default:0"`
	PendingReceipt     bool   `gorm:"column:pending_receipt;type:tinyint(1);not null;default:0"`
	Enclosed           bool   `gorm:"column:enclosed;type:tinyint(1);not null;default:0"`
	Color              string `gorm:"column:color;type:varchar(32)"`
}

type SupplyOrderStateLang struct {
	IDSupplyOrderState uint32 `gorm:"primaryKey;column:id_supply_order_state;type:int(11) unsigned;not null"`
	IDLang             uint32 `gorm:"primaryKey;column:id_lang;type:int(11) unsigned;not null"`
	Name               string `gorm:"column:name;type:varchar(128)"`
}

type Tab struct {
	IDTab        int    `gorm:"primaryKey;column:id_tab;type:int(11);not null"`
	IDParent     int    `gorm:"column:id_parent;type:int(11);not null"`
	Position     int    `gorm:"column:position;type:int(11);not null"`
	Module       string `gorm:"column:module;type:varchar(64)"`
	ClassName    string `gorm:"column:class_name;type:varchar(64);not null"`
	RouteName    string `gorm:"column:route_name;type:varchar(256)"`
	Active       bool   `gorm:"column:active;type:tinyint(1);not null"`
	Enabled      bool   `gorm:"column:enabled;type:tinyint(1);not null"`
	HideHostMode bool   `gorm:"column:hide_host_mode;type:tinyint(1);not null"`
	Icon         string `gorm:"column:icon;type:varchar(32)"`
}

type TabAdvice struct {
	IDTab    int `gorm:"primaryKey;column:id_tab;type:int(11);not null"`
	IDAdvice int `gorm:"primaryKey;column:id_advice;type:int(11);not null"`
}

type TabLang struct {
	IDTab  int    `gorm:"primaryKey;index:IDX_3E3D6F36ED47AB56;column:id_tab;type:int(11);not null"`
	IDLang int    `gorm:"primaryKey;index:IDX_3E3D6F36BA299860;column:id_lang;type:int(11);not null"`
	Name   string `gorm:"column:name;type:varchar(128);not null"`
}

type TabModulePreference struct {
	IDTabModulePreference int    `gorm:"primaryKey;column:id_tab_module_preference;type:int(11);not null"`
	IDEmployee            int    `gorm:"uniqueIndex:employee_module;column:id_employee;type:int(11);not null"`
	IDTab                 int    `gorm:"uniqueIndex:employee_module;column:id_tab;type:int(11);not null"`
	Module                string `gorm:"uniqueIndex:employee_module;column:module;type:varchar(191);not null"`
}

type Tag struct {
	IDTag  uint32 `gorm:"primaryKey;column:id_tag;type:int(10) unsigned;not null"`
	IDLang uint32 `gorm:"index:id_lang;column:id_lang;type:int(10) unsigned;not null"`
	Name   string `gorm:"index:tag_name;column:name;type:varchar(32);not null"`
}

type TagCount struct {
	IDGroup uint32 `gorm:"primaryKey;index:id_group;column:id_group;type:int(10) unsigned;not null;default:0"`
	IDTag   uint32 `gorm:"primaryKey;column:id_tag;type:int(10) unsigned;not null;default:0"`
	IDLang  uint32 `gorm:"index:id_group;column:id_lang;type:int(10) unsigned;not null;default:0"`
	IDShop  uint32 `gorm:"index:id_group;column:id_shop;type:int(11) unsigned;not null;default:0"`
	Counter uint32 `gorm:"index:id_group;column:counter;type:int(10) unsigned;not null;default:0"`
}

type Tax struct {
	IDTax   uint32  `gorm:"primaryKey;column:id_tax;type:int(10) unsigned;not null"`
	Rate    float64 `gorm:"column:rate;type:decimal(10,3);not null"`
	Active  bool    `gorm:"column:active;type:tinyint(1) unsigned;not null;default:1"`
	Deleted bool    `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type TaxLang struct {
	IDTax  uint32 `gorm:"primaryKey;column:id_tax;type:int(10) unsigned;not null"`
	IDLang uint32 `gorm:"primaryKey;column:id_lang;type:int(10) unsigned;not null"`
	Name   string `gorm:"column:name;type:varchar(32);not null"`
}

type TaxRule struct {
	IDTaxRule       int    `gorm:"primaryKey;column:id_tax_rule;type:int(11);not null"`
	IDTaxRulesGroup int    `gorm:"index:id_tax_rules_group;index:category_getproducts;column:id_tax_rules_group;type:int(11);not null"`
	IDCountry       int    `gorm:"index:category_getproducts;column:id_country;type:int(11);not null"`
	IDState         int    `gorm:"index:category_getproducts;column:id_state;type:int(11);not null"`
	ZipcodeFrom     string `gorm:"index:category_getproducts;column:zipcode_from;type:varchar(12);not null"`
	ZipcodeTo       string `gorm:"column:zipcode_to;type:varchar(12);not null"`
	IDTax           int    `gorm:"index:id_tax;column:id_tax;type:int(11);not null"`
	Behavior        int    `gorm:"column:behavior;type:int(11);not null"`
	Description     string `gorm:"column:description;type:varchar(100);not null"`
}

type TaxRulesGroup struct {
	IDTaxRulesGroup int       `gorm:"primaryKey;column:id_tax_rules_group;type:int(11);not null"`
	Name            string    `gorm:"column:name;type:varchar(50);not null"`
	Active          int       `gorm:"column:active;type:int(11);not null"`
	Deleted         bool      `gorm:"column:deleted;type:tinyint(1) unsigned;not null"`
	DateAdd         time.Time `gorm:"column:date_add;type:datetime;not null"`
	DateUpd         time.Time `gorm:"column:date_upd;type:datetime;not null"`
}

type TaxRulesGroupShop struct {
	IDTaxRulesGroup uint32 `gorm:"primaryKey;column:id_tax_rules_group;type:int(11) unsigned;not null"`
	IDShop          uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type Timezone struct {
	IDTimezone uint32 `gorm:"primaryKey;column:id_timezone;type:int(10) unsigned;not null"`
	Name       string `gorm:"column:name;type:varchar(32);not null"`
}

type Translation struct {
	IDTranslation int    `gorm:"primaryKey;column:id_translation;type:int(11);not null"`
	IDLang        int    `gorm:"index:IDX_3B44757BBA299860;column:id_lang;type:int(11);not null"`
	Key           string `gorm:"column:key;type:text;not null"`
	Translation   string `gorm:"column:translation;type:text;not null"`
	Domain        string `gorm:"index:key;column:domain;type:varchar(80);not null"`
	Theme         string `gorm:"column:theme;type:varchar(32)"`
}

type Warehouse struct {
	IDWarehouse    uint32 `gorm:"primaryKey;column:id_warehouse;type:int(11) unsigned;not null"`
	IDCurrency     uint32 `gorm:"column:id_currency;type:int(11) unsigned;not null"`
	IDAddress      uint32 `gorm:"column:id_address;type:int(11) unsigned;not null"`
	IDEmployee     uint32 `gorm:"column:id_employee;type:int(11) unsigned;not null"`
	Reference      string `gorm:"column:reference;type:varchar(64)"`
	Name           string `gorm:"column:name;type:varchar(45);not null"`
	ManagementType string `gorm:"column:management_type;type:enum('WA','FIFO','LIFO');not null;default:WA"`
	Deleted        bool   `gorm:"column:deleted;type:tinyint(1) unsigned;not null;default:0"`
}

type WarehouseCarrier struct {
	IDCarrier   uint32 `gorm:"primaryKey;index:id_carrier;column:id_carrier;type:int(11) unsigned;not null"`
	IDWarehouse uint32 `gorm:"primaryKey;index:id_warehouse;column:id_warehouse;type:int(11) unsigned;not null"`
}

type WarehouseProductLocation struct {
	IDWarehouseProductLocation uint32 `gorm:"primaryKey;column:id_warehouse_product_location;type:int(11) unsigned;not null"`
	IDProduct                  uint32 `gorm:"uniqueIndex:id_product;column:id_product;type:int(11) unsigned;not null"`
	IDProductAttribute         uint32 `gorm:"uniqueIndex:id_product;column:id_product_attribute;type:int(11) unsigned;not null"`
	IDWarehouse                uint32 `gorm:"uniqueIndex:id_product;column:id_warehouse;type:int(11) unsigned;not null"`
	Location                   string `gorm:"column:location;type:varchar(64)"`
}

type WarehouseShop struct {
	IDShop      uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
	IDWarehouse uint32 `gorm:"primaryKey;index:id_warehouse;column:id_warehouse;type:int(11) unsigned;not null"`
}

type WebBrowser struct {
	IDWebBrowser uint32 `gorm:"primaryKey;column:id_web_browser;type:int(10) unsigned;not null"`
	Name         string `gorm:"column:name;type:varchar(64)"`
}

type WebserviceAccount struct {
	IDWebserviceAccount int    `gorm:"primaryKey;column:id_webservice_account;type:int(11);not null"`
	Key                 string `gorm:"index:key;column:key;type:varchar(32);not null"`
	Description         string `gorm:"column:description;type:text"`
	ClassName           string `gorm:"column:class_name;type:varchar(50);not null;default:WebserviceRequest"`
	IsModule            int8   `gorm:"column:is_module;type:tinyint(2);not null;default:0"`
	ModuleName          string `gorm:"column:module_name;type:varchar(50)"`
	Active              int8   `gorm:"column:active;type:tinyint(2);not null"`
}

type WebserviceAccountShop struct {
	IDWebserviceAccount uint32 `gorm:"primaryKey;column:id_webservice_account;type:int(11) unsigned;not null"`
	IDShop              uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}

type WebservicePermission struct {
	IDWebservicePermission int    `gorm:"primaryKey;column:id_webservice_permission;type:int(11);not null"`
	Resource               string `gorm:"uniqueIndex:resource_2;index:resource;column:resource;type:varchar(50);not null"`
	Method                 string `gorm:"uniqueIndex:resource_2;index:method;column:method;type:enum('GET','POST','PUT','DELETE','HEAD');not null"`
	IDWebserviceAccount    int    `gorm:"uniqueIndex:resource_2;index:id_webservice_account;column:id_webservice_account;type:int(11);not null"`
}

type Zone struct {
	IDZone uint32 `gorm:"primaryKey;column:id_zone;type:int(10) unsigned;not null"`
	Name   string `gorm:"column:name;type:varchar(64);not null"`
	Active bool   `gorm:"column:active;type:tinyint(1) unsigned;not null;default:0"`
}

type ZoneShop struct {
	IDZone uint32 `gorm:"primaryKey;column:id_zone;type:int(11) unsigned;not null"`
	IDShop uint32 `gorm:"primaryKey;index:id_shop;column:id_shop;type:int(11) unsigned;not null"`
}
