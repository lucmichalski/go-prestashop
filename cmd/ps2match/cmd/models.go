package cmd

import (
    "gorm.io/gorm"
)

type Match struct {
    gorm.Model
    IdProduct   uint
    IdImage     uint
    ShopId      uint
    Fingerprint string
    Filepath    string
    Mimetype    string
}

type Response struct {
    Status string   `json:"status,omitempty"`
    Error  []string `json:"error,omitempty"`
    Method string   `json:"method,omitempty"`
    Result []string `json:"result,omitempty"`
}

type Params struct {
    Url      string      `url:"url,omitempty"`
    Filepath string      `url:"filepath,omitempty"`
    Metadata interface{} `url:"metadata,omitempty"`
}

type Image struct {
    LangId      uint   `json:"id_lang,omitempty"`
    ShopId      uint   `json:"id_shop,omitempty"`
    IdProduct   uint   `json:"id_product,omitempty"`
    IdImage     uint   `json:"id_image,omitempty"`
    LinkRewrite string `json:"link_rewrite,omitempty"`
    Legend      string `json:"legend,omitempty"`
}
