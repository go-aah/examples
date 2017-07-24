package models

import "time"

// UserInfo struct to hold user data.
type UserInfo struct {
	PrincipalValue string    `json:"principal_value"`
	AccessTime     time.Time `json:"access_time"`
	Message        string    `json:"message,omitempty"`
}
