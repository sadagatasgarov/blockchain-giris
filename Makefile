build:
	@go build -o bin/blocker

run: build
	@./bin/blocker

test:
	@go test -v ./... -count=1

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/*.proto

git:
	@git add .
	@git commit -m "Peer discovery hele nodleri refector edirik23.59"
	@git push -u origin main

.PHONY: proto