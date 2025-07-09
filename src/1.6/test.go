package main


import (
	"fmt"
	"net/http"
	"time"
	"os"
	"io"
	"strings"

)



func main() {

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	filename := sanitizeFilename(url) + ".html"
	file, err := os.Create(filename)
	if err != nil {
        ch <- fmt.Sprintf("create file error for %s: %v", url, err)
        return
    }
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("write error for %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s  -> saved to %s", secs, nbytes, url, filename)

}

func sanitizeFilename(url string) string {
    url = strings.TrimPrefix(url, "http://")
    url = strings.TrimPrefix(url, "https://")
    url = strings.ReplaceAll(url, "/", "_")
    return url
}

