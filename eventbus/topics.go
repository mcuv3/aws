package eventbus

type Topic struct {
	Name    string
	Message interface{}
}

type AuditTopic struct {
	Name    string
	Message AuditMessage
}
type AuditMessage struct {
	FullMethod string `json:"full_method"`
	AccountId  string `json:"account_id"`
}
