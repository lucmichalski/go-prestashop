package cmd

import (
    "time"
)

type Product struct {
    DateAdd          time.Time
    DateUpd          time.Time
    IdProduct        uint
    Reference        string
    Price            float64
    IdShopDefault    uint
    IsVirtual        bool
    Name             string
    Description      string
    DescriptionShort string
    Active           bool
    ManufacturerId   uint
    ManufacturerName string
    ShopId           uint
    ShopName         string
    LinkRewrite      string
    IdImage          uint
    IdCategory       uint
    NameCategory     string
    PriceFinal       float64
    NbDownloadable   uint
    SavQuantity      uint
    BadgeDanger      bool
    Features         string
    FeatureValues    string
}

type Feature struct {
    IdProduct uint
    Value     string
    Name      string
}

type Category struct {
    ShopId       uint
    LangId       uint
    ParentId     uint
    CategoryLink string
    Category     string
    Description  string
    MetaTitle    string
}

type Customer struct {
    IdCustomer               uint
    Firstname                string
    Lastname                 string
    Fullname                 string
    Email                    string
    Note                     string
    Active                   uint
    Deleted                  uint
    Newsletter               uint
    Optin                    uint
    IpRegistrationNewsletter string
    Birthday                 time.Time
    DateAdd                  time.Time
    SocialTitle              string
    IdGroup                  uint
    GroupName                string
    ShopeName                string
    Company                  string
    TotalSpent               float64
    Connect                  time.Time
}

type Order struct {
    IdOrder               uint
    IdCurrency            uint
    IdPdf                 uint
    IdShop                uint
    IdLang                uint
    OrderPayment          string
    ModulePayment         string
    TotalPaid             float64
    TotalDiscounts        float64
    TotalDiscountsTaxIncl float64
    TotalDiscountsTaxExcl float64
    TotalPaidTaxIncl      float64
    TotalPaidTaxExcl      float64
    TotalPaidReal         float64
    TotalProducts         []uint8
    TotalProductsWt       []uint8
    TotalShipping         float64
    TotalShippingTaxIncl  float64
    TotalShippingTaxExcl  float64
    TotalWrapping         float64
    TotalWrappingTaxIncl  float64
    TotalWrappingTaxExcl  float64
    RoundMode             uint
    RoundType             uint
    InvoiceNumber         uint
    DeliveryNumber        string
    InvoiceDate           time.Time
    DeliveryDate          time.Time
    Valid                 uint
    DateAdd               time.Time
    DateUpd               time.Time
    CarrierTaxRate        float64
    IdAddressDelivery     uint
    CustomerName          string
    OrderStatusName       string
    OrderStatusColor      string
    OrderNew              uint
    IdCountry             uint
    CountryName           string
    Company               string
    VatNumber             string
    PhoneMobile           string
    Phone                 string
    DeliveryLastname      string
    DeliveryFirstname     string
    Address1              string
    Address2              string
    Postcode              string
    City                  string
    BadgeSuccess          uint
}
