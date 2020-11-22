package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/dariuszkorolczukcom/birthday-app-api/birthday/structs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {

	var buf bytes.Buffer
	var req structs.BodyRequest
	var b structs.Birthday
	var err error

	log.Printf("Request body:\n %s", request.Body)
	// Unmarshal the json, return 404 if error
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 400}, nil
	}
	log.Printf("Request Object parsed:\n %v", req)

	err = b.SetBirthday(req.BirthdayDate)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 400}, nil
	}
	log.Printf("Birthday Object with parsed time:\n %v", b)
	b.CountHoursRoundDecimalBirthday()
	b.CountMinutesRoundDecimalBirthday()
	b.CountSecondsRoundDecimalBirthday()
	body, err := json.Marshal(b)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "hello-handler",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
