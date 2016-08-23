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

2016/08/21 11:52:18 ACK Interrupt: 55550000000000000000ff00ff000000ff00ff00
2016/08/21 11:52:18 ACK Frame: 00ff00ff
2016/08/21 11:52:18 ACK Response: 00ff01ff31cf
```

```
$ ./bin/gateway-id /dev/cu.usbmodem1411
2016/08/23 15:23:14 Sending REQUEST_GATEWAY_ID - 00ff02fed4a686
2016/08/23 15:23:14 Reading Response: 00ff00ff
2016/08/23 15:23:14 Reading Response: 00ff12eed5a74230434539383436314230303037303015
2016/08/23 15:23:14 Response: 00ff12eed5a74230434539383436314230303037303015
2016/08/23 15:23:14 Received - 00ff12eed5a74230434539383436314230303037303015
2016/08/23 15:23:14 Gateway ID: B0CE98461B000700
```
