package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    domain, ok := os.LookupEnv("DOMAIN")
    if ok {
      fmt.Printf("Listening on http://%s:80.\n", domain)
    } else {
      fmt.Print("USAGE: Set $DOMAIN environment variable, then run http2https.\n")
      os.Exit(1)
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        url := "https://" + domain + r.URL.Path
        http.Redirect(w, r, url, 307)
    })

    err := http.ListenAndServe(":80", nil)
    if err != nil {
      panic(err)
    }
}
