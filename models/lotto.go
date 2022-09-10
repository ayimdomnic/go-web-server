package models

import "time"
import "github.com/ayimdomnic/go-web-server/types"

// const (
// 	admin  UserRoles = "admin"
// 	player UserRoles = "player"
// 	staff  UserRoles = "staff"
// )

// const (
// 	active   TicketStatus = "active"
// 	inactive TicketStatus = "inactive"
// 	won      TicketStatus = "won"
// )

// const (
// 	live      LottoStatus = "live"
// 	drawing   LottoStatus = "drawing"
// 	finished  LottoStatus = "finished"
// 	cancelled LottoStatus = "cancelled"
// )

type Lotto struct {
	ID            uint              `json:"id" gorm:"primary_key"`
	Name          string            `json:"name"`
	ImageUrl      string            `json:"image_url"`
	StartDate     time.Time         `json:"start_date"`
	EndDate       time.Time         `json:"end_date"`
	Status        types.LottoStatus `json:"status" sql:"type:status"`
	WinningNumber string            `json:"winnig_number"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     time.Time         `json:"deleted_at"`
	CreatedBy     string            `json:"created_by"`
	Tickets []Ticket
}
