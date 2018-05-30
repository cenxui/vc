build:
	go get github.com/aws/aws-sdk-go
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/aws/aws-lambda-go/events
	env GOOS=linux go build -ldflags="-s -w" -o bin/cameraGet camera/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/cameraPost camera/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/monitorGet monitor/get/main.go