package main

import (
    "fmt"
)

type User struct {
    Id string
    AddressId string
}

func main() {
    s := NewServer()

    s.HandleFunc("GET", "/", func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "welcome!")
    })

    s.HandleFunc("GET", "/about", func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "about")
    })

    s.HandleFunc("GET", "/users/:user_id", func(c *Context) {
        u := User{Id: c.Params["user_id"].(string)}
        c.RenderXml(u)
    })

    s.HandleFunc("GET", "/users/:user_id/address/:address_id", func(c *Context) {
        u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
        c.RenderJson(u)
    })

    s.HandleFunc("POST", "/users", func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, c.Params)
    })

    s.HandleFunc("POST", "/users/:user_id/address", func(c *Context) {
        fmt.Fprintf(c.ResponseWriter, "create user %v's address\n",
            c.Params["user_id"])
    })

    s.Run(":8080")
}
