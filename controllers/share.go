package controllers

import (
        "fmt"
        "net/http"
        "strings"
        "html/template"
        "models"
        "strconv"
        )

type Copyform struct {
    Tag string
}

func action(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
//	fmt.Fprintln(w, r.Form)
//	fmt.Fprintln(w, "from: ", r.URL)
//
//	for k, v := range r.Form {
//		fmt.Printf("key: %v, val: %v\n", k, strings.Join(v, ""))
//	}
}

func Share(w http.ResponseWriter, r *http.Request) {
    action(w, r)

    tag := r.Form["tag"][0]
    content := models.Fetch(tag)
    t := template.New("share board")
    t, _ = t.Parse(`{{.}}`)
    t.Execute(w, content)
}
func Copy(w http.ResponseWriter, r *http.Request) {
    action(w, r)
    
    tag := r.Form.Get("tag")
    goa := r.Form.Get("go")
    if r.Method == "GET" {
        w.Header().Set("Content-Type", "text/html")

        t, _ := template.New("copy.html").ParseFiles("views/copy.html")
        cf := Copyform {Tag: tag}
        t.Execute(w, cf)
    } else if r.Method == "POST" {
        n := models.Save(r.Form["tag"][0], strings.Join(r.Form["content"], ""))
        if goa == "share" {
            http.Redirect(w, r, "/goweb/share?tag=" + tag, http.StatusFound)
        } else {
            fmt.Fprintf(w, strconv.Itoa(n))
        }
    }
}
func Index(w http.ResponseWriter, r *http.Request) {
    action(w, r)
}
