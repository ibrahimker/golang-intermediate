.PHONY: generate-proto
generate-proto:
	protoc -I common/model/google/api --proto_path=common/model --go_out=plugins=grpc:common/model/. --go_opt=paths=source_relative common/model/*.proto