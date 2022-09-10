package models

import "time"
import "github.com/ayimdomnic/go-web-server/types"

type Ticket struct {
	ID        uint              `json:"id" gorm:"primary_key"`
	PlayerId  string            `json:"player_id"`
	LottoId   string            `json:"lotto_id"`
	Status    types.LottoStatus `json:"status" sql:"type:status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt time.Time         `json:"deleted_at"`
}
