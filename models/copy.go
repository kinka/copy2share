package models

import (
       "fmt" 
       "os"
       "bufio"
       "strings"
        )

var gap = "==========hehe=========="

func Save(tag string, content string) int {
    f := openTag(tag)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    text := content + gap
    for scanner.Scan() {
        text += scanner.Text()
    }

    f.Seek(0, 0)
    writer := bufio.NewWriter(f)
    n, _ := writer.WriteString(text)
    writer.Flush()
    fmt.Println(n, tag, text)

    return n
}

func Fetch(tag string) string {
    f := openTag(tag)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var content string
    for scanner.Scan() {
        content += scanner.Text()
    }
    content = strings.Replace(content, gap, "\r\n\r\n", -1)
    fmt.Println(content)
    return content
}

func replace(content string) {

}
func openTag(tag string) *os.File {
    tagsDir := "/home/kinka/copy2share/src/tags/"
    os.MkdirAll(tagsDir, 0777)
    f, _ := os.OpenFile(tagsDir + tag, os.O_RDWR|os.O_CREATE, 0666)

    return f
}
