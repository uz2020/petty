module t

go 1.16

require (
	go.etcd.io/etcd/client/v3 v3.5.1
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace t => ../
