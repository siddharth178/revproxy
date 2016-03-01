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
```
#### example
```
sudo revproxy -backend=demo.io:80 -crt=/tmp/crt -key=/tmp/key
```
This will start revproxy process that listens on local port 443 and takes the traffic to backend running on host 'demo.io' on port 80. It will use certificate and key from location specified in arguments.

Output
```
sudo revproxy -backend=demo.io:80
2016/03/01 06:30:42 Using following config:
2016/03/01 06:30:42     https: true
2016/03/01 06:30:42     crt: crt
2016/03/01 06:30:42     key: key
2016/03/01 06:30:42     backend: demo.io:80
2016/03/01 06:30:42 Listening on: :443
```
