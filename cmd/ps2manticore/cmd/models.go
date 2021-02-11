package cmd

import (
    "time"
)

type Product struct {
    DateAdd           time.Time
    DateUpd           time.Time
    ShopId            string
    LangId            string
    CategoryId        uint
    CategoryLink      string
    ManufacturerId    uint
    Afo               string
    AvailableForOrder string
    Price             int
    Product           string
    Category          string
    Manufacturer      string
    IdProduct         uint
    LinkRewrite       string
    Name              string
    Description       string
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
