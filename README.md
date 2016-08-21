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

2016/08/21 10:23:10 "\x00\xff\x01\xff4\xcc"
```
