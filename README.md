# Msgpack forward example for Fluent Bit

Simple example showing how to use the Forward protocol to send data to Fluent Bit from Golang.

Using the 2nd definition of Forward here: <https://github.com/fluent/fluent-bit/blob/1fa0e94a09e4155f8a6d8a0efe36a5668cdc074e/plugins/in_forward/fw_prot.c#L417>

```c
* Forward format 2 (message mode) : [tag, time, map, ...]
```

To test first run up the receiver, e.g.:

```shell
$ docker run --rm -it -p 24224:24224 fluent/fluent-bit:1.9.7 -i forward -o stdout -m '*'
Fluent Bit v1.9.7
* Copyright (C) 2015-2022 The Fluent Bit Authors
* Fluent Bit is a CNCF sub-project under the umbrella of Fluentd
* https://fluentbit.io

[2022/08/20 14:34:10] [ info] [fluent bit] version=1.9.7, commit=265783ebe9, pid=1
[2022/08/20 14:34:10] [ info] [storage] version=1.2.0, type=memory-only, sync=normal, checksum=disabled, max_chunks_up=128
[2022/08/20 14:34:10] [ info] [cmetrics] version=0.3.5
[2022/08/20 14:34:10] [ info] [input:forward:forward.0] listening on 0.0.0.0:24224
[2022/08/20 14:34:10] [ info] [sp] stream processor started
[2022/08/20 14:34:10] [ info] [output:stdout:stdout.0] worker #0 started
```

Now build and send the data to it:

```shell
$ go mod download
$ go build ./...
$ ./msgpack-demo-fluent-bit
25 bytes sent on iteration 0
25 bytes sent on iteration 1
25 bytes sent on iteration 2
25 bytes sent on iteration 3
```

You should start seeing the following appearing in the receiver:

```shell
[2022/08/20 14:34:10] [ info] [output:stdout:stdout.0] worker #0 started
[0] iamatag: [1661006053.000000000, {"key"=>"value"}]
[0] iamatag: [1661006053.000000000, {"key"=>"value"}]
[0] iamatag: [1661006053.000000000, {"key"=>"value"}]
[0] iamatag: [1661006053.000000000, {"key"=>"value"}]
[0] iamatag: [1661006053.000000000, {"key"=>"value"}]
```
