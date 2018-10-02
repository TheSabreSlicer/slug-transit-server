package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "math/big"
    "time"
    "io/ioutil"
)

type ActiveBus struct {
    ID string
    Lat big.Rat
    Lon big.Rat
    Type string
}

type ActiveBusList struct {
    Collection []ActiveBus
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/active", GetActive).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

var client = &http.Client{Timeout: 2 * time.Second}

func GetActive(w http.ResponseWriter, r *http.Request) {
    resp, err := client.Get("http://bts.ucsc.edu:8081/location/get")
    if err != nil {
        return
    }
    defer resp.Body.Close()

    bodyBytes, err2 := ioutil.ReadAll(resp.Body)

    if err2 != nil {
        return
    }
    
    busses := make([]ActiveBus,0)
    json.Unmarshal(bodyBytes, &busses)
    enc := json.NewEncoder(w)
    enc.Encode(busses)
    return
}
