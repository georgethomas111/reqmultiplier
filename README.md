# README

Request multiplier multiplies the request that is sent to the server by 
acting as a reverse proxy which makes 2 requests to the http server.

The proxying is happening in the ip-layer.

# DESIGN
```
                    --> actual-service
req -> multiplexer -
                    --> actual-service
```

So the trick here is to make sure the HTTP request that is received 
by the TCP server should first write it into a buffer.
