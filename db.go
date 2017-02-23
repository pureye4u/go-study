package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
    db, _ := gorm.Open("postgres", "host=localhost user=pureye4u dbname=pureye4u sslmode=disable password=730359")
    // defer db.Close()

    db.AutoMigrate(&User{})

    return db
}
