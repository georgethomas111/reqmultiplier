# README

Request multiplier multiplies the request that is sent to the server by 
as a reverse proxy which makes 2 requests.

The proxying is happening in the ip-layer.

# DESIGN

                    --> actual-service
req -> multiplexer -
                    --> actual-service