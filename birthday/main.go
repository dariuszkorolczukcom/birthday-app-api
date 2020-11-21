package main

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var buf bytes.Buffer
var req BodyRequest
var b Birthday
var err error

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {

	log.Printf("Request body:\n %s", request.Body)
	// Unmarshal the json, return 404 if error
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 400}, nil
	}
	log.Printf("Request Object parsed:\n %v", req)

	err = b.setBirthday(req.BirthdayDate)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 400}, nil
	}
	log.Printf("Birthday Object with parsed time:\n %v", b)
	b.DecimalBirthday()
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
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

type BodyRequest struct {
	BirthdayDate string `json:"birthday"`
}

type Birthday struct {
	Born                        time.Time
	HoursRoundDecimalBirthday   []time.Time
	MinutesRoundDecimalBirthday []time.Time
	SecondsRoundDecimalBirthday []time.Time
}

func (b *Birthday) getBirthday() string {
	return b.Born.Format(time.RFC3339)
}

func (b *Birthday) setBirthday(date string) error {
	b.Born, err = time.Parse(
		time.RFC3339,
		date)
	return err
}

func (b *Birthday) DecimalBirthday() {
	for i := 1; i < 7; i++ {
		b.HoursRoundDecimalBirthday = append(b.HoursRoundDecimalBirthday, b.Born.Add(time.Hour*time.Duration(i*100000)))
	}
	for i := 1; i < 42; i++ {
		b.MinutesRoundDecimalBirthday = append(b.MinutesRoundDecimalBirthday, b.Born.Add(time.Minute*time.Duration(i*1000000)))

	}
	for i := 1; i < 27; i++ {
		b.SecondsRoundDecimalBirthday = append(b.SecondsRoundDecimalBirthday, b.Born.Add(time.Second*time.Duration(i*100000000)))
	}
}
