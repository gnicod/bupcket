generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bupcket/bupcket.proto

serve:
	go run bupcket_server/main.go