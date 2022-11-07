generate:
	protoc -I ./common/model \
	  --go_out ./common/model --go_opt paths=source_relative \
	  --go-grpc_out ./common/model --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out ./common/model --grpc-gateway_opt paths=source_relative \
	  --openapiv2_out docs/openapiv2 \
	  ./common/model/*.proto