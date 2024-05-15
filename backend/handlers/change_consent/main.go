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
	r.POST("/change-consent", ChangeConsentHandler)
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(HandleRequest)
}

func ChangeConsentHandler(c *gin.Context) {

	//TODO: Authorization Check

	var requestBody models.ChangeConsentRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := db.GetItem(db.TableConsent, requestBody.TraceID)
	if err != nil {
		return
	}

	var consent models.ConsentDAO
	err = utils.UnmarshalDynoNotation(item, &consent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	consent.Consents = utils.MapDataRecords(requestBody.Consents)
	consent.Description = requestBody.Description

	consentDBItem, err := utils.MakeDynoNotation(consent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.PutItem(db.TableConsent, consentDBItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
