package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	backend := flag.String("backend", "localhost:8080", "backend's host:port")
	https := flag.Bool("https", true, "terminate https?")
	crt := flag.String("crt", "crt", "certificate file path")
	key := flag.String("key", "key", "key file path")
    flag.Parse()

    localAddr := ":80"
    if *https {
        localAddr = ":443"
    }

    log.Println("Using following config:")
    log.Println("\thttps:", *https)
    log.Println("\tcrt:", *crt)
    log.Println("\tkey:", *key)
    log.Println("\tbackend:", *backend)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   *backend,
	})

    log.Println("Listening on:", localAddr)
	if *https {
		log.Fatal(http.ListenAndServeTLS(localAddr, *crt, *key, proxy))
	} else {
		log.Fatal(http.ListenAndServe(localAddr, proxy))
	}
}
