package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// クエリパラメータから name を取り出し
	name := req.QueryStringParameters["name"]
	if name == "" {
		name = "World"
	}

	// レスポンス本文
	body := fmt.Sprintf("Hello, %s!", name)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
