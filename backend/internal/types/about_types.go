package types

// ApiSummary
//
// @public
type ApiSummary struct {
	Healthy  bool     `json:"health"`
	Versions []string `json:"versions"`
	Title    string   `json:"title"`
	Message  string   `json:"message"`
	Build    Build
}

type Build struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"buildTime"`
}
