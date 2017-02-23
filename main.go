package main

import (
    "fmt"
    "net/http"
)

func main() {
    r := &router{make(map[string]map[string]HandlerFunc)}

    r.HandleFunc("GET", "/", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "welcome!")
    })))))

    r.HandleFunc("GET", "/about", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, "about")
    })))))

    r.HandleFunc("GET", "/users/:user_id", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        if c.Params["user_id"] == "0" {
            panic("id is zero")
        }
        fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n",
            c.Params["user_id"])
    })))))

    r.HandleFunc("GET", "/users/:user_id/address/:address_id", logHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        fmt.Fprintf(c.ResponseWriter, "retrieve user %v's address %v\n",
            c.Params["user_id"], c.Params["address_id"])
    }))))

    r.HandleFunc("POST", "/users", logHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        fmt.Fprintln(c.ResponseWriter, c.Params)
    }))))

    r.HandleFunc("POST", "/users/:user_id/address", logHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
        fmt.Fprintf(c.ResponseWriter, "create user %v's address\n",
            c.Params["user_id"])
    }))))

    http.ListenAndServe(":8080", r)
}
