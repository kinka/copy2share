package main

import (
	"fmt"
	"log"
	"net/http"
    ctrl "controllers"
)

func main() {
    http.HandleFunc("/", ctrl.Index)
	http.HandleFunc("/share", ctrl.Share)
	http.HandleFunc("/copy", ctrl.Copy)

    fmt.Println("start listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
