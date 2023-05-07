# To generate grpc communication use the following command
> protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative ./proto/greeter/greeter-service.proto

# Build and run the project using
> docker-compose build
> docker-compose up

# If there are errors in the files use this
> go get
