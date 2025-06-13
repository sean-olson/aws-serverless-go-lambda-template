build:
	GOOS=linux GOARCH=amd64 go build -o bin/hello main.go
	zip -j bin/hello.zip bin/hello

deploy: build
	serverless deploy

clean:
	rm -f bin/hello bin/hello.zip