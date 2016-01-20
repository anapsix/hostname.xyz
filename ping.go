package main

import (
  "flag"
  "fmt"
  "net"
  "net/http"
  "strconv"
  "os"
  "time"
  "github.com/tatsushid/go-fastping"
)

var listenPort int

func init() {
  flag.IntVar(&listenPort, "port", 8080, "listen port") 
  flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
    p := fastping.NewPinger()
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    fmt.Fprintf(os.Stderr,"Client IP: %s\n", ip)

    p.AddIP(ip)
    p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
//      fmt.Fprintf(w,"IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
      fmt.Fprintf(w, "%v", rtt)
    }

    err := p.Run()
    if err != nil {
      fmt.Println(err)
      fmt.Fprintf(w, "error")
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Listening on port:", strconv.Itoa(listenPort))
    http.ListenAndServe(":" + strconv.Itoa(listenPort), nil)
}
