package teams

import (
	"time"
	"vollyemsk_tournament_gateway/models/users"
)

type Team struct {
	ID             int
	Name           string
	Description    string
	FoundationDate time.Time
	ContactUsers   []*users.User
	HomeGyms       []*TeamGym
}
