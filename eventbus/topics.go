package eventbus

type Topic string

const (
	InvokeFunction Topic = "invoke-function"
	Audit          Topic = "audit"
	MessageToQueue Topic = "message-to-queue"
)
