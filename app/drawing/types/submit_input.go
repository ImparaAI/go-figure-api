package types

type SubmitInput struct {
	Points []OriginalPoint `json:"points"`
	Image string `json:"image"`
}