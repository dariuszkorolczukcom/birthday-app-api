package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(birthdayFunction())
	if err != nil {
		return Response{StatusCode: 404}, err
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

func birthdayFunction() (b Birthday) {
	p := fmt.Println
	f := fmt.Printf
	now := time.Now()
	b.Born = time.Date(
		1986, 8, 22, 16, 30, 00, 0, time.UTC)
	p("\nbirthday:")
	p(b.getBirthday())
	p("\nnow:")
	p(now.Format(time.RFC3339))
	diff := now.Sub(b.Born)
	p("\ndifference:")
	p(diff)
	p("\ndifference hours:")
	p(diff.Hours())
	p("\ndifference minutes:")
	f("%.0f\n", diff.Minutes())
	p("\ndifference seconds:")
	f("%.0f\n", diff.Seconds())
	for i := 1; i < 7; i++ {
		f("\nbirthday plus %d00k hours:", i)
		b.HoursRoundDecimalBirthday = append(b.HoursRoundDecimalBirthday, b.Born.Add(time.Hour*time.Duration(i*100000)))
	}
	for i := 1; i < 42; i++ {
		f("\nbirthday plus %dM minutes:", i)
		b.MinutesRoundDecimalBirthday = append(b.MinutesRoundDecimalBirthday, b.Born.Add(time.Minute*time.Duration(i*1000000)))

	}
	for i := 1; i < 27; i++ {
		f("\nbirthday plus %d00M seconds:", i)
		b.SecondsRoundDecimalBirthday = append(b.SecondsRoundDecimalBirthday, b.Born.Add(time.Second*time.Duration(i*100000000)))

	}
	return
}

type Birthday struct {
	Born                        time.Time
	HoursRoundDecimalBirthday   []time.Time
	MinutesRoundDecimalBirthday []time.Time
	SecondsRoundDecimalBirthday []time.Time
}

func (b Birthday) getBirthday() string {
	return b.Born.Format(time.RFC3339)
}
