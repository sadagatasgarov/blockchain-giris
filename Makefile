build:
	@go build -o bin/blocker

run: build
	@./bin/blocker

test:
	@go test -v ./... -count=1 

git:
	@git add .
	@git commit -m "adrese seedler de yazildi.test de yazildi"
	@git push -u origin main