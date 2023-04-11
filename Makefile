PROTO_DIR=pb

stubs:
	protoc -I $(PROTO_DIR)\
	       	--go_out=$(PROTO_DIR)\
	       	--go_opt=paths=source_relative\
	       	--go-grpc_out=$(PROTO_DIR)\
	       	--go-grpc_opt=paths=source_relative\
		$(PROTO_DIR)/*.proto

.PHONY: stubs