package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Palabra struct {
    Uuid string `json:"uuid"`
    Palabra string  `json:"palabra"`
    Acepciones []string  `json:"acepciones"`
}

func handlePalabra(responseWriter  http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    reqPalabra := vars["palabra"]
    fmt.Println(reqPalabra)
    var resPalabra Palabra
    resPalabra = Palabra{Uuid: "151651", Palabra: reqPalabra, Acepciones: nil}
    responseWriter.Header().Set("Content-Type", "application-json")

    responseWriter.WriteHeader(http.StatusOK)
    fmt.Println(resPalabra)
    json.NewEncoder(responseWriter).Encode(resPalabra)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
    fmt.Println("home")
}

func handle() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/palabra/{palabra}", handlePalabra)
    myRouter.HandleFunc("/", handleHome)
    http.ListenAndServe(":8080", myRouter)
}

func main() {
    handle()
}
