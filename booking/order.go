package booking

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

var (
	ErrTicketIsNotAvailable = errors.New("ticket is not available")
)

type OrderStatus string

const (
	Created  OrderStatus = "created"
	Reserved OrderStatus = "reserved"
	Rejected OrderStatus = "rejected"
)

type Order struct {
	ID        xid.ID      `json:"id,omitempty"`
	Email     string      `json:"email,omitempty"`
	ConcertID xid.ID      `json:"concert_id,omitempty"`
	Status    OrderStatus `json:"status,omitempty"`
	Date      time.Time   `json:"date"`
}

type ReservationRequest struct {
	ConcertID xid.ID `json:"concert_id,omitempty"`
	OrderID   xid.ID `json:"order_id,omitempty"`
}

type CreateOrderRequest struct {
	Email     string `json:"email,omitempty"`
	ConcertID string `json:"concert_id,omitempty"`
}
