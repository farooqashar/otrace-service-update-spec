package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	db "otrace_service/config"
	"otrace_service/models"
	"otrace_service/utils"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Service Starting")
	r := gin.Default()
	r.POST("/user-dashboard", UserDashboardHandler)
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(HandleRequest)
}

func UserDashboardHandler(c *gin.Context) {
	var requestBody models.UserDashboardRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryResult := db.QueryByDataSubject(db.TableConsent, db.IndexDataSubject, requestBody.DataSubject)

	// Convert the query results to a slice of ConsentDAO
	var allConsents []models.ConsentDAO
	for _, item := range queryResult.Items {
		var consent models.ConsentDAO
		err := utils.UnmarshalDynoNotation(item, &consent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allConsents = append(allConsents, consent)
	}

	c.JSON(http.StatusOK, allConsents)
}
