package server

type Result struct {
	Error   bool        `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message,omitempty"`
	Item    interface{} `json:"item,omitempty"`
}

// Wrap creates a Wrapper instance and adds the initial namespace and data to be returned.
func Wrap(data interface{}) Result {
	return Result{
		Item: data,
	}
}

func (r Result) AddMessage(message string) Result {
	r.Message = message
	return r
}

func (r Result) AddError(err string, details interface{}) Result {
	r.Message = err
	r.Details = details
	r.Error = true
	return r
}
