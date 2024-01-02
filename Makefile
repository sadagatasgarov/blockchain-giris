build:
	@go build -o bin/blocker

run: build
	@./bin/blocker

test:
	@go test -v ./... -count=1 

git:
	@git add .
	@git commit -m "Public ve private keyler hazir"
	@git push -u origin main