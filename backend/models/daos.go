package models

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

// DynoNotation represents an object in dynamoDB.
// Used to represent key value data such as keys, table items...
type DynoNotation map[string]types.AttributeValue

type ConsentDAO struct {
	TraceID     string      `dynamodbav:"trace_id"`
	Timestamp   string      `dynamodbav:"timestamp"`
	DataSubject string      `dynamodbav:"data_subject"`
	Description string      `dynamodbav:"description"`
	Consents    []RecordDAO `dynamodbav:"consents"`
	ParentIDS   []string    `dynamodbav:"parent_ids"`
	TraceURI    string      `dynamodbav:"trace_uri"`
	TraceCERT   string      `dynamodbav:"trace_cert"`
}

type RecordDAO struct {
	Category string `dynamodbav:"category"`
	Uses     string `dynamodbav:"uses"`
	Subject  string `dynamodbav:"subject"`
}
