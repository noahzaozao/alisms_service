echo 'protoc'
protoc -I/usr/local/include -I. \
    --proto_path=$GOPATH/src:. \
    --micro_out=./proto/alisms/ \
    --go_out=./proto/alisms/ \
    -I./proto/alisms/ alisms.proto &&
go build
echo 'success'
