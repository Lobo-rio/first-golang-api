package models

type Password struct {
	NewPassword     string `json:"newpassword,omitempty"`
	CurrentPassword string `json:"currentpassword,omitempty"'
}
