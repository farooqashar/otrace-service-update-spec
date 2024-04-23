package models

type ConsentRequest struct {
	TraceID                  string    `json:"trace_id"`
	Time                     string    `json:"time"`
	DataSubject              string    `json:"data_subject"`
	Description              string    `json:"description"`
	Consents                 []Consent `json:"consents"`
	ParentIDS                []string  `json:"parent_ids"`
	ProviderChallenge        string    `json:"provider_challenge"`
	ProviderChallengeMethod  string    `json:"provider_challenge_method"`
	RecipientChallenge       string    `json:"recipient_challenge"`
	RecipientChallengeMethod string    `json:"recipient_challenge_method"`
	TraceURI                 string    `json:"trace_uri"`
	TraceCERT                string    `json:"trace_cert"`
}

type Consent struct {
	Category string `json:"category"`
	Uses     string `json:"uses"`
	Subject  string `json:"subject"`
}

type ShareRequest struct {
	TraceID     string       `json:"trace_id"`
	Time        string       `json:"time"`
	Description string       `json:"description"`
	DataShared  []DataShared `json:"data_shared"`
}

type DataShared struct {
	Category string `json:"category"`
	Uses     string `json:"uses"`
	Subject  string `json:"subject"`
}
