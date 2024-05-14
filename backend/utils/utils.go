package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"otrace_service/models"
	"reflect"
)

func MakeDynoNotation(obj interface{}) (models.DynoNotation, error) {
	avMap, err := attributevalue.MarshalMap(obj)
	if err != nil {
		return nil, err
	}

	dynoNotation := make(models.DynoNotation)
	for key, val := range avMap {
		dynoNotation[key] = val.(types.AttributeValue)
	}
	return dynoNotation, nil
}

// UnmarshalDynoNotation takes a DynoNotation and a pointer to the struct you want to fill.
func UnmarshalDynoNotation(dynoNotation models.DynoNotation, out interface{}) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("out argument must be a pointer to a struct")
	}

	err := attributevalue.UnmarshalMap(dynoNotation, out)
	if err != nil {
		return fmt.Errorf("error unmarshalling from DynoNotation: %v", err)
	}
	return nil
}
