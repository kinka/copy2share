package main

import (
	"fmt"
	"log"
	"net/http"
    "strings"
    ctrl "controllers"
)

func dispatch(w http.ResponseWriter, r *http.Request) {
    p := r.URL.String()
    
    if strings.HasPrefix(p, "/share?") {
        ctrl.Share(w, r)
    } else if strings.HasPrefix(p, "/copy?") {
        ctrl.Copy(w, r)
    } else {
        ctrl.Index(w, r)        
    }
}

func main() {
    http.HandleFunc("/", dispatch)
	//http.HandleFunc("/share", ctrl.Share)
	//http.HandleFunc("/copy", ctrl.Copy)

    fmt.Println("start listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
