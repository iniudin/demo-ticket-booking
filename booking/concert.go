package booking

import (
	"github.com/rs/xid"
	"time"
)

type Concert struct {
	ID        xid.ID    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Location  string    `json:"location,omitempty"`
	Date      time.Time `json:"date"`
	Remaining int       `json:"remaining"`
}

type CreateConcertRequest struct {
	Name      string    `json:"name,omitempty"`
	Location  string    `json:"location,omitempty"`
	Date      time.Time `json:"date"`
	Remaining int       `json:"remaining"`
}
