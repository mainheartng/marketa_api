PROTO_DIR=protos

generate:
	protoc --go_out=. --go_opt=module=github.com/mainheartng/marketa_api --go-grpc_out=protos --proto_path=.  --go-grpc_opt=module=github.com/mainheartng/marketa_api/protos $(PROTO_DIR)/*.proto
#  --go_out=. --go_opt=module=github.com/inidaname/mosque_location --go-grpc_out=protos --go-grpc_opt=module=github.com/inidaname/mosque_location/protos --proto_path=protos