package models

type ConsentRequest struct {
	Timestamp                string        `json:"timestamp"`
	DataSubject              string        `json:"data_subject"`
	Description              string        `json:"description"`
	Consents                 []TraceRecord `json:"consents"`
	ParentIDS                []string      `json:"parent_ids"`
	ProviderChallenge        string        `json:"provider_challenge"`
	ProviderChallengeMethod  string        `json:"provider_challenge_method"`
	RecipientChallenge       string        `json:"recipient_challenge"`
	RecipientChallengeMethod string        `json:"recipient_challenge_method"`
	TraceURI                 string        `json:"trace_uri"`
	TraceCERT                string        `json:"trace_cert"`
}

type ConsentResponse struct {
	TraceID string `json:"trace_id"`
}

type TraceRecord struct {
	Category string `json:"category"`
	Uses     string `json:"uses"`
	Subject  string `json:"subject"`
}

type ShareRequest struct {
	TraceID     string        `json:"trace_id"`
	Time        string        `json:"time"`
	Description string        `json:"description"`
	DataShared  []TraceRecord `json:"data_shared"`
}
