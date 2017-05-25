package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type IP struct {
	Ip      string    `json:"ip"`
	Updated time.Time `json:"updated"`
}

type IPs []IP

var ips IPs

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/add", Add)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8888",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Last 5 ips:\n")
	for _, ip := range ips {
		var str = fmt.Sprintf("ip: %v time: %v", ip.Ip, ip.Updated)
		fmt.Fprintln(w, str)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	if len(ips) == 5 {
		// truncate if gt 5
		_, ips = ips[0], ips[1:]
	}

	ip := IP{r.RemoteAddr, time.Now()}
	ips = append(ips, ip)
}
