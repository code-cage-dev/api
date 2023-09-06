package solution

import (
	"github.com/code-cage-dev/api/pkg/entity"
	"github.com/google/uuid"
)

type Entity struct {
	entity.Base
	UserID      uuid.UUID `json:"user_id"`
	ChallengeID uuid.UUID `json:"challenge_id"`
	Language    string    `json:"language"`
	Code        string    `json:"code"`
}
