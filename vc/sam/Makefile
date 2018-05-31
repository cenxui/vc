build:
	go get github.com/aws/aws-sdk-go
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/aws/aws-lambda-go/events
	env GOOS=linux go build -ldflags="-s -w" -o bin/camera/get camera/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/camera/getAll camera/getAll/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/camera/post camera/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/monitor/get monitor/get/main.go