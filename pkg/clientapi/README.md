
## Dev

Note: Below instructions are for Ubuntu 16.04. Please modify for other platforms accordingly.

If you make changes to ``api.proto``, please regenerate ``api.pb.go`` by following below steps.

1. Download ``protoc`` version 3.0 from [here](https://github.com/google/protobuf/releases).

2. Place ``protoc`` binary in your ``PATH``.

3. Install ``protoc-gen-go``.

```
go get github.com/golang/protobuf/protoc-gen-go
```

This would place ``protoc-gen-go`` binary in your ``$GOROOT/bin``. Verify ``$GOROOT/bin`` is in your ``$PATH``.

```
# env | grep PATH
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin:/root/gostuff/bin
GOPATH=/root/gostuff
```

4. Run ``protoc`` to generate ``clientapi.pb.go``.

```
protoc -I=$PWD --go_out=plugins=grpc:.  $PWD/clientapi.proto
```
