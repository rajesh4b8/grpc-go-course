### To generate code for Go:

```shell
protoc -Igreet/proto --go_opt=module=github.com/rajesh4b8/grpc-go-course --go_out=. --go-grpc_opt=module=github.com/rajesh4b8/grpc-go-course --go-grpc_out=. greet/proto/*.proto
```
