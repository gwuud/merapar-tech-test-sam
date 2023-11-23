package main

import (
	"html/template"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xyproto/randomstring"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	t, err := template.New("index").Parse(`<h1>The Saved String Is {{ . }}!</h1>`)
	if err != nil {
		log.Fatal(err)
	}

	var rendered strings.Builder
	err = t.Execute(&rendered, randomstring.HumanFriendlyString(10))
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body:       rendered.String(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
