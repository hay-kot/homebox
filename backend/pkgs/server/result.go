package server

type Result struct {
	Error   bool        `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message,omitempty"`
	Item    interface{} `json:"item,omitempty"`
}

type Results struct {
	Items any `json:"items"`
}

// Wrap creates a Wrapper instance and adds the initial namespace and data to be returned.
func Wrap(data interface{}) Result {
	return Result{
		Item: data,
	}
}
