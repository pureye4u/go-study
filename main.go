package main

import (
    "fmt"
)

type User struct {
    Id string
    AddressId string
}

func main() {
    db := InitDB()

    s := NewServer()

    s.HandleFunc("GET", "/", func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "welcome!")
    })

    s.HandleFunc("GET", "/about", func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "about")
    })

    s.HandleFunc("GET", "/users/:user_id", func(c *Context) {
        u := User{Id: c.Params["user_id"].(string)}
        c.RenderJson(u)
    })

    s.HandleFunc("GET", "/users/:user_id/address/:address_id", func(c *Context) {
        u := User{Id: c.Params["user_id"].(string), AddressId: c.Params["address_id"].(string)}
        c.RenderJson(u)
    })

    s.HandleFunc("POST", "/users", func(c *Context) {
        u := User{Id: c.Params["user_id"].(string)}
        db.Create(&u)
        c.RenderJson(u)
    })

    s.HandleFunc("POST", "/users/:user_id/address", func(c *Context) {
        u := User{}
        db.Where("Id = ?", c.Params["user_id"].(string)).First(&u)
        u.AddressId = c.Params["address_id"].(string)
        db.Save(&u)
        c.RenderJson(u)
    })

    s.Run(":8080")
}
