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

	//TODO: Authorization Check

	var requestBody models.UserDashboardRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Query all consents
	consentQueryResult := db.QueryByDataSubject(db.TableConsent, db.IndexDataSubject, requestBody.DataSubject)

	// Convert the query results to a slice of ConsentDAO
	var allConsents []models.ConsentDAO
	for _, item := range consentQueryResult.Items {
		var consent models.ConsentDAO
		err := utils.UnmarshalDynoNotation(item, &consent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allConsents = append(allConsents, consent)
	}

	//Query all sharing
	sharingQueryResult := db.QueryByDataSubject(db.TableDataSharing, db.IndexDataSubject, requestBody.DataSubject)

	// Convert the query results to a slice of ConsentDAO
	var allSharing []models.DataSharingDAO
	for _, item := range sharingQueryResult.Items {
		var sharing models.DataSharingDAO
		err := utils.UnmarshalDynoNotation(item, &sharing)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allSharing = append(allSharing, sharing)
	}

	//Query all usage
	usageQueryResult := db.QueryByDataSubject(db.TableDataUsage, db.IndexDataSubject, requestBody.DataSubject)

	// Convert the query results to a slice of ConsentDAO
	var allUsage []models.DataUsageDAO
	for _, item := range usageQueryResult.Items {
		var usage models.DataUsageDAO
		err := utils.UnmarshalDynoNotation(item, &usage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allUsage = append(allUsage, usage)
	}

	//Query all violations
	violationQueryResult := db.QueryByDataSubject(db.TableDataViolation, db.IndexDataSubject, requestBody.DataSubject)

	// Convert the query results to a slice of ConsentDAO
	var allViolations []models.ViolationDAO
	for _, item := range violationQueryResult.Items {
		var violation models.ViolationDAO
		err := utils.UnmarshalDynoNotation(item, &violation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allViolations = append(allViolations, violation)
	}

	//create user dashboard response
	response := utils.MapToUserDashboardResponse(allConsents, allSharing, allUsage, allViolations)

	c.JSON(http.StatusOK, response)
}
