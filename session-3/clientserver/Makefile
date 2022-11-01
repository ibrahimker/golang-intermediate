.PHONY: generate-proto
generate-proto:
	protoc --proto_path=common/model --go_out=plugins=grpc:common/model/. --go_opt=paths=source_relative common/model/*.proto