package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (events.APIGatewayProxyResponse, error){
	
	body, _ := json.Marshal(map[string]string{
		"message": "Hello from Go Lambda!",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(body),
	}, nil

}	

func main(){
	lambda.Start(handler)
}
