package v1

// Status is a struct that contains Service Name and whether it is up.
type Status struct {
	IsUp string `json:"isUp"`
}

type Message struct {
	Author string
	Text   string
	Time   string
}

type Messages struct {
	Messages []Message
}
