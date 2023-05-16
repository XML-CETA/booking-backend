#!/bin/bash

my_array=("accommodation_service" "reservation_service" "user_service" "auth_service")

for element in "${my_array[@]}"
do
    echo "$element"
	protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative ./proto/"$element"/"$element".proto
done

