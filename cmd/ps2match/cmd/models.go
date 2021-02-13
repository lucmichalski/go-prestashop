package cmd

import (
    "gorm.io/gorm"
)

type Match struct {
    gorm.Model
    IdProduct   uint
    IdImage     uint
    IdShop      uint
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

type MatchSearch struct {
    Error  []interface{}  `json:"error"`
    Method string         `json:"method"`
    Result []ResultSearch `json:"result"`
    Status string         `json:"status"`
}

type ResultSearch struct {
    Filepath string  `json:"filepath"`
    Score    float64 `json:"score"`
}

type Add struct {
    Url      string      `url:"url,omitempty"`
    Filepath string      `url:"filepath,omitempty"`
    Metadata interface{} `url:"metadata,omitempty"`
}

type Search struct {
    Url             string `url:"url,omitempty"`
    Image           string `url:"image,omitempty"`
    AllOrientations bool   `url:"all_orientations,omitempty"`
}

type Image struct {
    IdLang      uint   `json:"id_lang,omitempty"`
    IdShop      uint   `json:"id_shop,omitempty"`
    IdProduct   uint   `json:"id_product,omitempty"`
    IdImage     uint   `json:"id_image,omitempty"`
    LinkRewrite string `json:"link_rewrite,omitempty"`
    Legend      string `json:"legend,omitempty"`
}
