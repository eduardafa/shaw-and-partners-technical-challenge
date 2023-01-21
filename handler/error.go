package handler

type ApiError struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Debug  *Debug `json:"debug,omitempty"`
}

type Debug struct {
	File      string   `json:"file,omitempty"`
	Line      int      `json:"line,omitempty"`
	RootCause string   `json:"root_cause,omitempty"`
	Stack     ApiStack `json:"stack,omitempty"`
}

type ApiStack struct {
	error
}

func (r *ApiError) Error() string {
	return r.Detail
}
