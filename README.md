# README

Request multiplier multiplies the requests by making multiple requests 
to the destination server. 

This can be thought of like a curl command which can be used to send 
multiple requests. 


# DESIGN
```
                    --> actual-service
req -> multiplexer -
                    --> actual-service
```

# INSTALL
```
(xenial)george@localhost:~/workspace/src/github.com/georgethomas111/reqmultiplier$ go install github.com/georgethomas111/reqmultiplier
(xenial)george@localhost:~/workspace/src/github.com/georgethomas111/reqmultiplier$ reqmultiplier -help
Usage of reqmultiplier:
  -destURL string
        URL to hit. (default "http://google.com")
  -multiply int
        Number of times to multiply request. (default 10)
  -proto string
        Protocol to use. (default "GET")
  -reqBody string
        Request body to send.
  -timeout int
        timeout between requests.
(xenial)george@localhost:~/workspace/src/github.com/georgethomas111/reqmultiplier$ 
```
