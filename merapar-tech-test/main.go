package main

import (
	"html/template"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	random "github.com/xyproto/randomstring"
)

func main() {
	lambda.Start(handler)
}

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	t, err := template.New("index").Parse(`<h1>The Saved String Is {{ . }}!</h1>`)
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	err = t.Execute(&b, random.HumanFriendlyString(10))
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       b.String(),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil
}
