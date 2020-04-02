package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    domain, ok := os.LookupEnv("DOMAIN")
    if ok {
      fmt.Println("Listening on http://0.0.0.0:80")
      fmt.Printf("Redirecting all requests to https://%s:443.\n", domain)
    } else {
      fmt.Println("USAGE: Set $DOMAIN environment variable, then run http2https.")
      os.Exit(1)
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        url := "https://" + domain + r.URL.Path
        http.Redirect(w, r, url, 308)
    })

    err := http.ListenAndServe(":80", nil)
    if err != nil {
      panic(err)
    }
}
