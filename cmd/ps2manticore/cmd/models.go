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
    NameCategory     string
    PriceFinal       float64
    NbDownloadable   uint
    SavQuantity      uint
    BadgeDanger      bool
    Features         string
    FeatureValues    string
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
