# Go GRPC + SQL Demo ( Future Workshop ? )

We are gonna create an API simulating a car service shop. This workshop can do multiple services, like exchanging tires, do technical diagnostics, change parts and wash the car. All cars coming here will be washed as a last step and for changing parts on the car, it needs to go though diagnostics first.

Given all the above, we should expose some endpoints:

- Register car owner
- Register a car
- Register a work order
- Start a given service
- Finish a given service
- Finish work order

### Running locally

To run on dev with hot reload, install [air](https://github.com/cosmtrek/air) and run:

```
air -c air.toml
```

The project also uses a PostgreSQL database and the database name needs to be created upfront. The db name can be informed on the `.env` file. By default it expects a database called `carservice-dev`.

### Generating code from our .proto file

We are gonna use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to also expose our gRPC service as a Rest API to be used by the Frontend apps. To install the extra deps, use this commmand.

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grp
```

To generate the files from our .proto run:

```
make protoc-generate
```

### gRPC

Our protobuf

```
service BackOfficeService {
  rpc RegisterOwner(Owner) returns (google.protobuf.Empty)
  rpc RegisterCar(Car) returns (google.protobuf.Empty)
}

service WorkOrderService {
  rpc RegisterWorkOrder(WorkOrderRequest) returns (google.protobuf.Empty)
  rpc GetRunningWorkOrders(RunningWorkOrdersQuery) returns (RunningWorkOrdersResponse)
  rpc StartWorkOrderService(StartWorkOrderServiceRequest) returns (google.protobuf.Empty)
  rpc FinishWorkOrderService(FinishWorkOrderServiceRequest) returns (google.protobuf.Empty)
  rpc EndWorkOrder(EndWorkOrderRequest) returns (google.protobuf.Empty)
}
```

Our service will expose a gRPC API based on the above protobuf definition.

### Rest API and Swagger

The same .proto file has annotations to generate a gateway/bridge between Rest and GRPC. Those annotation are provided by the `grpc-gateway` project mentioned before.

### TODO

- Validate/enforce order of services.
- Add Auth with Auth0
