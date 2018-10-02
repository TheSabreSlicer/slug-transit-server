package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/shopspring/decimal"
)

type ActiveBus struct {
    ID string
    Lat Decimal
    Lon Decimal
    Type string
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/active", GetActive).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

var myClient = &http.Client{Timeout: 2 * time.Second}

func getActive(target interface{}) error {
    r, err := myClient.Get("http://bts.ucsc.edu:8081/location/get")
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}
