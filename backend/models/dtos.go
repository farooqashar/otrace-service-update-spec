package models

type CreateConsentRequest struct {
	Timestamp                string       `json:"timestamp" binding:"required"`
	DataSubject              string       `json:"data_subject" binding:"required"`
	Description              string       `json:"description" binding:"required"`
	Consents                 []DataRecord `json:"consents" binding:"required"`
	ParentIDS                []string     `json:"parent_ids"`
	ProviderChallenge        string       `json:"provider_challenge" binding:"required"`
	ProviderChallengeMethod  string       `json:"provider_challenge_method" binding:"required"`
	RecipientChallenge       string       `json:"recipient_challenge" binding:"required"`
	RecipientChallengeMethod string       `json:"recipient_challenge_method" binding:"required"`
	TraceURI                 string       `json:"trace_uri" binding:"required"`
	TraceCERT                string       `json:"trace_cert" binding:"required"`
}

type CreateConsentResponse struct {
	TraceID string `json:"trace_id" binding:"required"`
}

type DeleteConsentRequest struct {
	TraceID string `json:"trace_id" binding:"required"`
}

type ChangeConsentRequest struct {
	TraceID     string       `json:"trace_id" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Consents    []DataRecord `json:"consents" binding:"required"`
}

type DataRecord struct {
	Category string `json:"category" binding:"required"`
	Uses     string `json:"uses" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
}

type ConsentRecord struct {
	TraceID     string       `json:"trace_id" binding:"required"`
	Timestamp   string       `json:"timestamp" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Consents    []DataRecord `json:"consents" binding:"required"`
}

type ShareDataRecord struct {
	TraceID     string       `json:"trace_id" binding:"required"`
	Timestamp   string       `json:"timestamp" binding:"required"`
	Description string       `json:"description" binding:"required"`
	DataShared  []DataRecord `json:"data_shared" binding:"required"`
}

type UseDataRecord struct {
	TraceID     string       `json:"trace_id" binding:"required"`
	Timestamp   string       `json:"timestamp" binding:"required"`
	Description string       `json:"description" binding:"required"`
	DataUsed    []DataRecord `json:"data_used" binding:"required"`
}

type ViolationRecord struct {
	TraceID     string       `json:"trace_id" binding:"required"`
	Timestamp   string       `json:"timestamp" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Violations  []DataRecord `json:"violations" binding:"required"`
}

type UserDashboardRequest struct {
	DataSubject string `json:"data_subject" binding:"required"`
}

type UserDashboardResponse struct {
	DataConsents   []ConsentRecord   `json:"data_consents" binding:"required"`
	DataSharing    []ShareDataRecord `json:"data_sharing"`
	DataUsage      []UseDataRecord   `json:"data_usage"`
	DataViolations []UseDataRecord   `json:"data_violations"`
}

type NotifyViolationRequest struct {
	DataViolations []UseDataRecord `json:"data_violations"  binding:"required"`
}
