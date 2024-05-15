package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"otrace_service/models"
	"reflect"
)

// MakeDynoNotation takes a struct and returns a pointer to DynoNotation
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

// UnmarshalDynoNotation takes a DynoNotation and a pointer to the struct.
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

func MapConsentRequestToConsentDAO(request models.CreateConsentRequest, traceId string) models.ConsentDAO {
	return models.ConsentDAO{
		TraceID:     traceId,
		Timestamp:   request.Timestamp,
		DataSubject: request.DataSubject,
		Description: request.Description,
		Consents:    MapDataRecords(request.Consents),
		ParentIDS:   request.ParentIDS,
		TraceURI:    request.TraceURI,
		TraceCERT:   request.TraceCERT,
	}
}

func MapToDataSharingDAO(record models.ShareDataRecord, dataSubject string) models.DataSharingDAO {
	return models.DataSharingDAO{
		TraceID:     record.TraceID,
		Timestamp:   record.Timestamp,
		DataSubject: dataSubject,
		Description: record.Description,
		DataShared:  MapDataRecords(record.DataShared),
	}
}

func MapToDataUsageDAO(record models.UseDataRecord, dataSubject string) models.DataUsageDAO {
	return models.DataUsageDAO{
		TraceID:     record.TraceID,
		Timestamp:   record.Timestamp,
		DataSubject: dataSubject,
		Description: record.Description,
		DataUsed:    MapDataRecords(record.DataUsed),
	}
}

func MapDataRecords(records []models.DataRecord) []models.RecordDAO {
	result := make([]models.RecordDAO, len(records))
	for i, record := range records {
		result[i] = models.RecordDAO{
			Category: record.Category,
			Uses:     record.Uses,
			Subject:  record.Subject,
		}
	}
	return result
}

func MapToCreateConsentResponse(traceId string) models.CreateConsentResponse {
	return models.CreateConsentResponse{
		TraceID: traceId,
	}
}

func MapToUserDashboardResponse(allConsents []models.ConsentDAO, allSharing []models.DataSharingDAO, allUsage []models.DataUsageDAO, allViolations []models.ViolationDAO) models.UserDashboardResponse {
	return models.UserDashboardResponse{
		DataConsents:   convertConsentDAOsToConsentRecords(allConsents),
		DataSharing:    convertDataSharingDAOsToShareDataRecords(allSharing),
		DataUsage:      convertDataUsageDAOsToUseDataRecords(allUsage),
		DataViolations: convertDataViolationsDAOsToUseDataRecords(allViolations),
	}
}

func convertConsentDAOsToConsentRecords(daos []models.ConsentDAO) []models.ConsentRecord {
	var records []models.ConsentRecord

	for _, dao := range daos {
		var consentsData []models.DataRecord
		for _, record := range dao.Consents {
			consentsData = append(consentsData, models.DataRecord{
				Category: record.Category,
				Uses:     record.Uses,
				Subject:  record.Subject,
			})
		}
		records = append(records, models.ConsentRecord{
			TraceID:     dao.TraceID,
			Timestamp:   dao.Timestamp,
			Description: dao.Description,
			Consents:    consentsData,
		})
	}

	return records
}

// convertDataSharingDAOsToShareDataRecords converts a list of DataSharingDAO to a list of ShareDataRecord
func convertDataSharingDAOsToShareDataRecords(daos []models.DataSharingDAO) []models.ShareDataRecord {
	var records []models.ShareDataRecord

	for _, dao := range daos {
		var sharedData []models.DataRecord
		for _, record := range dao.DataShared {
			sharedData = append(sharedData, models.DataRecord{
				Category: record.Category,
				Uses:     record.Uses,
				Subject:  record.Subject,
			})
		}
		records = append(records, models.ShareDataRecord{
			TraceID:     dao.TraceID,
			Timestamp:   dao.Timestamp,
			Description: dao.Description,
			DataShared:  sharedData,
		})
	}

	return records
}

// convertDataUsageDAOsToUseDataRecords converts a list of DataUsageDAO to a list of UseDataRecord
func convertDataUsageDAOsToUseDataRecords(daos []models.DataUsageDAO) []models.UseDataRecord {
	var records []models.UseDataRecord

	for _, dao := range daos {
		var usedData []models.DataRecord
		for _, record := range dao.DataUsed {
			usedData = append(usedData, models.DataRecord{
				Category: record.Category,
				Uses:     record.Uses,
				Subject:  record.Subject,
			})
		}
		records = append(records, models.UseDataRecord{
			TraceID:     dao.TraceID,
			Timestamp:   dao.Timestamp,
			Description: dao.Description,
			DataUsed:    usedData,
		})
	}

	return records
}

// convertDataUsageDAOsToUseDataRecords converts a list of DataUsageDAO to a list of UseDataRecord
func convertDataViolationsDAOsToUseDataRecords(daos []models.ViolationDAO) []models.ViolationRecord {
	var records []models.ViolationRecord

	for _, dao := range daos {
		var violations []models.DataRecord
		for _, record := range dao.DataViolated {
			violations = append(violations, models.DataRecord{
				Category: record.Category,
				Uses:     record.Uses,
				Subject:  record.Subject,
			})
		}
		records = append(records, models.ViolationRecord{
			TraceID:     dao.TraceID,
			Timestamp:   dao.Timestamp,
			Description: dao.Description,
			Violations:  violations,
		})
	}

	return records
}

// recordMatches checks if two DataRecords match based on Category, Uses
func recordMatches(consentRecord, activityRecord models.RecordDAO) bool {
	return consentRecord.Category == activityRecord.Category &&
		consentRecord.Uses == activityRecord.Uses
}

// CheckActivitiesUnderConsents checks if all data activity are consented, returns a list of unauthorized data activities
func CheckActivitiesUnderConsents(consents []models.RecordDAO, dataActivity []models.RecordDAO) []models.RecordDAO {
	var allViolationRecords []models.RecordDAO
	for _, share := range dataActivity {
		isAllowed := false
		for _, consent := range consents {
			if recordMatches(consent, share) {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			allViolationRecords = append(allViolationRecords, share)
		}
	}
	return allViolationRecords
}
