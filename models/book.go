package models

import "gorm.io/gorm"

type Book struct{
    gorm.Model
    ID      int64 `gorm:"primaryKey" json:"id"`
    Title    string `json:"title"`
    Publisher   string `json:"publisher"`
    Genre   string `json:"genre"`
    Stock   int `json:"stock"`
}
