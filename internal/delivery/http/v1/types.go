package v1

type PairsRequest struct {
	Pairs []string `json:"pairs,omitempty"`
}

type ErrResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
