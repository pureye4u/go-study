package main

import (
    "net/http"
    "encoding/json"
    "encoding/xml"
)

type Context struct {
    Params map[string]interface{}
    ResponseWriter http.ResponseWriter
    Request *http.Request
}

type HandlerFunc func(*Context)

func (c *Context) RenderJson(v interface{}) {
    c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Method", "GET, OPTIONS")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    c.ResponseWriter.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
        c.RenderErr(http.StatusInternalServerError, err)
    }
}

func (c *Context) RenderXml(v interface{}) {
    c.ResponseWriter.Header().Set("Content-Type", "application/xml; charset=utf-8")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Method", "GET, OPTIONS")
    c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    c.ResponseWriter.WriteHeader(http.StatusOK)

    if err := xml.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
        c.RenderErr(http.StatusInternalServerError, err)
    }
}

func (c *Context) RenderErr(code int, err error) {
    if err != nil {
        if code > 0 {
            http.Error(c.ResponseWriter, http.StatusText(code), code)
        } else {
            defaultErr := http.StatusInternalServerError
            http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
        }
    }
}
