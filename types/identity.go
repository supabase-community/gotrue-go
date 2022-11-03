package types

import (
	"time"

	"github.com/google/uuid"
)

type Identity struct {
	ID           string                 `json:"id"`
	UserID       uuid.UUID              `json:"user_id"`
	IdentityData map[string]interface{} `json:"identity_data,omitempty"`
	Provider     string                 `json:"provider"`
	LastSignInAt *time.Time             `json:"last_sign_in_at,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}
