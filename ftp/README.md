# FTP State Machine

A finite state machine that validates a flow of FTP commands and it's replies.

## Build

To build -

```
go build ./...
```

## Test

To test -

```
go test ./...
```

## Description

The library exposes APIs to invoke the state machine based on the request and response line messages that are generated when the FTP client and server communicate. The image capture from Wireshark below highlights the request and response messages for a login sequence -

![FTP Wireshark Capture](https://github.com/r00tu53r/protos/raw/main/ftp/img/ftp_packets.png)

