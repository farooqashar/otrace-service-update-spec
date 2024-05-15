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
	r.POST("/use-data", ShareDataHandler)
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(HandleRequest)
}

func ShareDataHandler(c *gin.Context) {

	//TODO: Authorization Check

	var requestBody models.UseDataRecord
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Get Consent by TraceId
	item, err := db.GetItem(db.TableConsent, requestBody.TraceID)
	if err != nil {
		return
	}

	var consentDAO models.ConsentDAO
	err = utils.UnmarshalDynoNotation(item, &consentDAO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//check if data usage activity in the scope of the consent

	useDataDAO := utils.MapDataRecords(requestBody.DataUsed)

	violationRecordDAOs := utils.CheckActivitiesUnderConsents(consentDAO.Consents, useDataDAO)
	if len(violationRecordDAOs) != 0 {
		log.Printf("violation detected, no consents exists for this data activity")

		violationDAO := models.ViolationDAO{
			TraceID:      requestBody.TraceID,
			Timestamp:    requestBody.Timestamp,
			DataSubject:  consentDAO.DataSubject,
			Description:  requestBody.Description,
			DataViolated: violationRecordDAOs,
		}
		//Create Violation Record
		dataViolationDBItem, err := utils.MakeDynoNotation(violationDAO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//Store Violation Record to DB
		err = db.PutItem(db.TableDataViolation, dataViolationDBItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//Create Usage Record
	dataUsageDAO := utils.MapToDataUsageDAO(requestBody, consentDAO.DataSubject)
	dataUsageDBItem, err := utils.MakeDynoNotation(dataUsageDAO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Store Usage Record to DB
	err = db.PutItem(db.TableDataUsage, dataUsageDBItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Created")
}
