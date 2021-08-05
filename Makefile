
proto:
	protoc \
		-I=../.. \
		-I=. \
		--go_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=api.yaml:. \
		--openapiv2_out=logtostderr=true,grpc_api_configuration=api.yaml:. \
		bootcamp.proto