package models

/*
LoginResponse contain token to return in login
*/
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
