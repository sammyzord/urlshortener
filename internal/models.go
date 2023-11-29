package internal

import (
	"github.com/sqids/sqids-go"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID   uint64
	Code *string `gorm:"index:code_unique,unique"`
	Url  string
}

func (u *URL) AfterCreate(tx *gorm.DB) (err error) {

	s, _ := sqids.New()

	id, _ := s.Encode([]uint64{u.ID})

	tx.Model(u).Update("code", id)

	return
}
