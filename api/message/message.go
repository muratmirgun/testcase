package message

import "encoding/json"

// Message represents data about a message.
type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func (m Message) Validate() bool {
	if m.Sender == "" {
		return false
	}

	if m.Receiver == "" {
		return false
	}

	if m.Message == "" {
		return false
	}

	return true
}

func (m Message) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, &m)
}
