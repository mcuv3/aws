package eventbus

type Topic string

func (t Topic) String() string {
	return string(t)
}

const (
	InvokeFunction Topic = "invoke-function"
	Audit          Topic = "audit"
	MessageToQueue Topic = "message-to-queue"
)
