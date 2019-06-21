# gomavlink
MAVLink Encoding and Decoding in Go 

[![GoDoc](https://godoc.org/github.com/daedaleanai/gomavlink?status.svg)](https://godoc.org/github.com/daedaleanai/gomavlink)

MAVLink is documented on https://mavlink.io/en/

The top level package defines Encoder and Decoder types that can handle the dialects generated from the XML files in the subpackages.

The dialect subpackages are entirely auto-generated by the [![gomavgen](https://github.com/daedaleanai/gomavgen)] tool.

usable for dumping ardupilot tlogs, see tlog/dump.go

NOTE ON THE LICENSE: 
The generated files here are derived from [![Original XML code](https://github.com/mavlink/mavlink/tree/master/message_definitions/v1.0)] under the LGPL.
The Go code is more liberal i really don't care, but i don't want to become a license manager, so i'm slapping the same on this.
Ping me if you want a more liberal one.


WIP! a little patience please.

TODO:
- imprint git version of xml files used to generate
- V2 signatures
- nice api for commands/parameters
- decent test suite
- test V2 decoding/encoding
- test V1 encoding
- String() method for generated message types that prints [n]byte as string 
- figure out if we can decide between V1 and V2 purely on message id.  i.e. are mid<255 always sent as V1 with no signature?  
if so, we could only generate V1 or V2 encode/decode based on message id. 
but then why have extended fields if you dont have a way to send messages as v1?  

