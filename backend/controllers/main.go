package controllers

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"page_info,omitempty"`
	Results  any    `json:"results,omitempty"`
}
