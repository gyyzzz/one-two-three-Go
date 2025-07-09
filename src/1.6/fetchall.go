// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
    "strings"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    filename := sanitizeFilename(url) + ".html"
    file, err := os.Create(filename)
    if err != nil {
        ch <- fmt.Sprintf("create file error for %s: %v", url, err)
        return
    }
    defer file.Close()

    // 将 body 内容写入文件，同时统计大小
    nbytes, err := io.Copy(file, resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("write error for %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s  -> saved to %s", secs, nbytes, url, filename)
}

// sanitizeFilename 将 URL 简化为合法文件名
func sanitizeFilename(url string) string {
    url = strings.TrimPrefix(url, "http://")
    url = strings.TrimPrefix(url, "https://")
    url = strings.ReplaceAll(url, "/", "_")
    return url
}
