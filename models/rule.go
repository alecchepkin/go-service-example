package models

import (
	"tracker-sms-svc/app"
)

type Rule struct {
	Id            int
	AdvertiserId  int
	PublisherId   int
	ShaveInterval int
	IsManual      bool
}

const (
	MaxInterval = 70
	DefInterval = 46
	MinInterval = 1
)

func (Rule) TableName() string {
	return "public.advertiser_to_publisher"
}

func FindRule(aid int, pid int) *Rule {
	db := app.GetPostgres()

	var rule Rule
	//db.Where("advertiser_id=?", aid).First(&rule)
	db.Where("advertiser_id=? and publisher_id=?", aid, pid).First(&rule)
	if rule.Id > 0 {
		return &rule
	}
	return nil

}

func (r Rule) SaveRule() {
	db := app.GetPostgres()
	sql := `insert into advertiser_to_publisher (advertiser_id, publisher_id, shave_interval)
values (?, ?, ?)
on conflict (
   advertiser_id,
   publisher_id)
   do update set shave_interval = EXCLUDED.shave_interval;`
	db.Exec(sql, r.AdvertiserId, r.PublisherId, r.ShaveInterval)
}

func (r *Rule) ChangeInterval(inc int) {
	i := r.ShaveInterval + inc
	if i > MaxInterval {
		i = MaxInterval
	}
	if i < MinInterval {
		i = MinInterval
	}
	r.ShaveInterval = i
}
