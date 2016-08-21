# go-pen-gateway

Get chatty with the Smithsonian Gateway Board

This is a work in progress.

```
$ ./bin/list-ports

2016/08/21 00:40:53    Found 2 ports. Here are the Smithsonian Gateways
2016/08/21 00:40:53    Name: /dev/cu.usbmodem1421
2016/08/21 00:40:53    Description: Smithsonian Gateway
2016/08/21 00:40:53    Transport: 0
```

```
$ ./bin/ack /dev/cu.usbmodem1421

ACK Interrupt: 55550000000000000000ff00ff000000ff00ff00
ACK Frame: 00ff00ff
ACK Response: 00ff01ff31cf
```
