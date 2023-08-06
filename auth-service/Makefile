gen-api:
	protoc  --proto_path=./api \
            --proto_path=./third_party \
			--go_out=paths=source_relative:./api \
         	--go-grpc_out=paths=source_relative:./api \
           --grpc-gateway_out=paths=source_relative:./api \
       		 api/v1/proto/auth-service.proto

build:
	docker compose up --build

run:
	docker compose up