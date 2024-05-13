package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"traceability_service/models"
)

var ginLambda *ginadapter.GinLambda
var dynamodbService *dynamodb.Client

func init() {
	log.Printf("Service Starting")
	r := gin.Default()
	r.POST("/create-consent", CreateConsentHandler)
	ginLambda = ginadapter.New(r)

	// build the required scan params
	fmt.Println("DB Connect:")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// using the config value, create the dynamodb client
	dynamodbService = dynamodb.NewFromConfig(cfg)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(HandleRequest)
}

func CreateConsentHandler(c *gin.Context) {
	var requestBody models.CreateConsentRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	traceID := uuid.NewString()

	//put into db
	_, err := dynamodbService.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("tracingTable"),
		Item: map[string]types.AttributeValue{
			"trace_id": &types.AttributeValueMemberS{Value: traceID},
		},
	})
	if err != nil {
		log.Printf("could not put the dyanmodb table! error: %v", err)
	} else {
		log.Printf("Item added successfully with trace_id: %v", traceID)
	}

	// Define the input parameters for the Scan operation
	input := &dynamodb.ScanInput{
		TableName: aws.String("tracingTable"),
	}

	// Execute the Scan operation
	result, err := dynamodbService.Scan(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to scan table, %v", err)
	}

	c.JSON(http.StatusCreated, result.Items)
}
