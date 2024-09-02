package wire

type Transaction struct {
	ID       string  `json:"id"`
	SourceID string  `json:"sourceId"`
	TargetID string  `json:"targetId"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
}
