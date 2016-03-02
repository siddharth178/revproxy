# revproxy
a simple reverse proxy

it is nothing but go-lang's inbuilt reverse proxy functionality exposed as handy utility.


## how to use revproxy
```
Usage of $GOROOT/bin/revproxy:
  -backend string
        backend's host:port (default "localhost:8080")
  -crt string
        certificate file path (default "crt")
  -https
        terminate https? (default true)
  -key string
        key file path (default "key")
  -listenOn string
        port to listen on (default "443")
```
#### example
```
sudo ~/go/bin/revproxy -backend=demo.io:80 -crt=/tmp/crt -key=/tmp/key
```
This will start revproxy process that listens on local port 443 and takes the traffic to backend running on host 'demo.io' on port 80. It will use certificate and key from location specified in arguments.
