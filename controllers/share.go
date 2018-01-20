package controllers

import (
        "fmt"
        "net/http"
        "html/template"
        "models"
        "strconv"
        "regexp"
        )

type TagItem struct {
    Tag string
    Items []string
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
func unescaped (x string) interface{} {
    reg, _ := regexp.Compile(`(https?://[^\s]+)`)
    y := reg.ReplaceAll([]byte(x), []byte("<a href='${1}'>${1}</a>"))
    return template.HTML(y)
}

func Share(w http.ResponseWriter, r *http.Request) {
    action(w, r)

    tag := r.Form.Get("tag")
    res := models.Fetch(tag)
    t := template.New("share.html")
    t = t.Funcs(template.FuncMap{"unescaped": unescaped})
    t, _ = t.ParseFiles("views/share.html")
    ti := TagItem {Tag: tag, Items: res}
    t.Execute(w, ti)
}
func Copy(w http.ResponseWriter, r *http.Request) {
    action(w, r)
    
    tag := r.Form.Get("tag")
    goa := r.Form.Get("go")
    if r.Method == "GET" {
        w.Header().Set("Content-Type", "text/html")

        t, _ := template.New("copy.html").ParseFiles("views/copy.html")
        ti := TagItem {Tag: tag}
        t.Execute(w, ti)
    } else if r.Method == "POST" {
        n := models.Save(tag, r.Form.Get("content"))
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
