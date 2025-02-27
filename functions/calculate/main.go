package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// var greetingService greeting.Servicer

// func init() {
// 	greetingService = greeting.NewService()
// }

type Request struct {
	OrderedItems      uint   `json:"ordered_items"`
	AvailablePkgSizes []uint `json:"available_package_sizes"`
}

type Response struct {
	Result []int `json:"result"`
}

func createErrorResponse(message string) *events.APIGatewayV2HTTPResponse {
	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 500,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: fmt.Sprintf(`{"error":"%s"}`, message),
	}
}

func handle(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	log.Printf("request %+v", request)

	req := Request{}
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return createErrorResponse("error parsing request"), nil
	}

	log.Printf("parsed %+v", req)

	packageSizes := make([]int, len(req.AvailablePkgSizes))
	for i, v := range req.AvailablePkgSizes {
		packageSizes[i] = int(v)
	}
	result := calculate(int(req.OrderedItems), packageSizes)
	resp := Response{Result: result}

	respRaw, err := json.Marshal(resp)
	if err != nil {
		return createErrorResponse("error marshaling response"), nil
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(respRaw),
	}, nil
}

func main() {
	lambda.Start(handle)
}
