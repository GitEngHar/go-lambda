GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o bootstrap main.go && zip function.zip bootstrap
