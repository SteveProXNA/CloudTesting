package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testwebapi/waf"
)

func AdminHelloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Healthy...!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("testHandler")
	web, err := url.Parse("https://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}
	log.Print(web)
	proxy := httputil.NewSingleHostReverseProxy(web)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = web.Path
	}
	proxy.ServeHTTP(w, req)
}

func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req.URL %s", r.URL)
		inter := waf.ProcessModSec(r.URL.String())
		if inter > 0 {
			log.Printf("==== Mod Security Blocked! ====")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
