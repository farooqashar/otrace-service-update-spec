package db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"otrace_service/models"
)

var dbClient *dynamodb.Client

func init() {
	fmt.Println("DB Connect:")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(Region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	dbClient = dynamodb.NewFromConfig(cfg)
}

// PutItem inserts an item (key + attributes) in to a dynamodb table.
func PutItem(tableName string, item models.DynoNotation) (err error) {
	_, err = dbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName), Item: item,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetItem returns an item if found based on the key provided.
// the key could be either a primary or composite key and values map.
func GetItem(tableName string, key models.DynoNotation) (item models.DynoNotation, err error) {
	resp, err := dbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{Key: key, TableName: aws.String(tableName)})
	if err != nil {
		return nil, err
	}
	return resp.Item, nil
}

func QueryByDataSubject(tableName string, indexName string, dataSubject string) *dynamodb.QueryOutput {
	// Query input using the GSI for data_subject
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("data_subject = :dataSubject"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":dataSubject": &types.AttributeValueMemberS{Value: dataSubject},
		},
	}

	// Execute the query
	result, err := dbClient.Query(context.TODO(), input)
	if err != nil {
		log.Fatalf("Query failed to execute, %v", err)
	}
	return result
}
