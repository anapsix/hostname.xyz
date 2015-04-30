package main

import (
  "flag"
  "fmt"
  "net"
  "net/http"
  "strconv"
  "strings"
  "os"
)

var listenPort int

func init() {
  flag.IntVar(&listenPort, "port", 8080, "listen port")
  flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    fmt.Fprintf(os.Stderr,"Client IP: %s\n", ip)
    host, _ := net.LookupAddr(ip)
    switch {
    case len(host) == 0:
      host = []string{"NXDOMAIN"}
    }
    fmt.Fprintf(os.Stderr,"Client Hostname: %s\n", strings.TrimRight(host[0],"."))
    fmt.Fprintf(w, "%s", strings.TrimRight(host[0],".") )
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Listening on port:", strconv.Itoa(listenPort))
    http.ListenAndServe(":" + strconv.Itoa(listenPort), nil)
}
