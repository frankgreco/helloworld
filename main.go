package main

import (
    "net/http"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
)

type Response struct {
    Msg string `json:"msg"`
}

func main() {
    r := httprouter.New()
    r.GET("/foo", Handler)
    http.ListenAndServe(":8080", r)
}

func Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    res := Response{
        Msg: "bar",
    }
    message, _ := json.Marshal(res)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(message)
}
