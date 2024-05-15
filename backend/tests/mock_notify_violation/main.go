package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Service Starting")
	r := gin.Default()
	r.POST("/mock-data-recipient/notify-violation", NotifyViolationHandler)
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(HandleRequest)
}

func NotifyViolationHandler(c *gin.Context) {
	//TODO: ADD MOCK DATA RECIPIENT IMPLEMENTATION
	c.JSON(http.StatusNotImplemented, "API Not Implemented")
}
