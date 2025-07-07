// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
    	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
    		url = "http://" + url
    	}
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
      
        defer resp.Body.Close()
        //打印状态码，不需要在异常流程里打印，因为err时resp == nil
        fmt.Fprintf(os.Stdout, "status code: %s\n\n", resp.Status)

        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: Copy faild %s: %v\n", url, err)
            os.Exit(1)
        }
        
    }
}
