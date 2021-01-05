# Go Grpc Workshop

We are gonna create an API simulating a car service shop. This workshop can do multiple services, like exchanging tires, do technical diagnostics, change parts and wash the car. All cars coming here will be washed as a last step and for changing parts on the car, it needs to go though diagnostics first.

Given all the above, we should expose some endpoints:

- Register car owner
- Register a car
- Register a work order
- Start a given service
- Finish a given service
- Finish work order

### gRPC

Our protobuf

```
service CarService {
  rpc RegisterOwner(Owner) returns (google.protobuf.Empty)
  rpc RegisterCar(Car) returns (google.protobuf.Empty)
  rpc RegisterWorkOrder(WorkOrderRequest) returns (google.protobuf.Empty)
  rpc StartWorkOrderService(StartWorkOrderServiceRequest) returns (google.protobuf.Empty)
  rpc FinishWorkOrderService(FinishWorkOrderServiceRequest) returns (google.protobuf.Empty)
  rpc EndWorkOrder(EndWorkOrderRequest)
}
```

Our service will expose a gRPC API based on the above protobuf definition.

### Generating code from our .proto file

We are gonna use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to also expose our gRPC service as a Rest API to be used by the Frontend apps. To install the extra deps, use this commmand.

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grp

```

To generate the files from our .proto, run:

```
protoc -I . \
  --go_out . --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  carshop.proto
```
