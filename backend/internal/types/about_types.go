package types

// ApiSummary
//
// @public
type ApiSummary struct {
	Healthy  bool     `json:"health"`
	Versions []string `json:"versions"`
	Title    string   `json:"title"`
	Message  string   `json:"message"`
}
