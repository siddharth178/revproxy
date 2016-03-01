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
	localPort := flag.String("listenOn", "443", "port to listen on")
    flag.Parse()

    if *https {
        *localPort = "443"
    } else {
        *localPort = "80"
    }

    log.Println("Using following config:")
    log.Println("\thttps:", *https)
    log.Println("\tlocalPort:", *localPort)
    log.Println("\tcrt:", *crt)
    log.Println("\tkey:", *key)
    log.Println("\tbackend:", *backend)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   *backend,
	})

	if *https {
		log.Fatal(http.ListenAndServeTLS(":"+*localPort, *crt, *key, proxy))
	} else {
		log.Fatal(http.ListenAndServe(":"+*localPort, proxy))
	}
}
