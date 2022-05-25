package kafkamodels

type UserCreatedMessage struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
