package models

import (
	"tracker-sms-svc/app"
)

type Adv struct {
	Id   int
	Name string
	Cost float64
}

func (Adv) TableName() string {
	return "public.advertiser"
}

func FindAdv(aid int) *Adv {
	db := app.GetPostgres()

	var adv Adv
	db.First(&adv, aid)
	if adv.Id > 0 {
		return &adv
	}
	return nil

}
