package models

import (
	"fmt"
	"github.com/prometheus/common/log"
	"tracker-sms-svc/app"
	"time"
)

type Stat struct {
	AdvertiserId int
	PublisherId  int
	Payout       float64
	Revenue      float64
	Cost         float64
}

const MaxRevenue = 10

func FindStat() []Stat {
	sql := fmt.Sprintf(`
select advertiser_id,
       publisher_id,
       floor(sum(payout_not_shaved), 2) payout,
       floor(sum(revenue), 2)           revenue,
       floor(payout / revenue, 2)       cost
from (select advertiser_id,
             publisher_id,
             payout,
             is_shaved,
             payout * (1 - is_shaved) payout_not_shaved,
             revenue
      from tracker.conversions_local
      where date >= '%s' and lower(dictGetString('adv_info','name',advertiser_id)) like '%%feed%%' and goal_id='conversion'
       )
group by advertiser_id, publisher_id
having revenue > %d
order by advertiser_id, publisher_id;
`, time.Now().AddDate(0, 0, -2).Format("2006-01-02"), MaxRevenue)

	rows, err := app.GetClickhouse().Query(sql)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error(err)
		}

	}()

	var stats []Stat
	for rows.Next() {
		var (
			advertiserId int
			publisherId  int
			payout       float64
			revenue      float64
			cost         float64
		)
		if err := rows.Scan(&advertiserId, &publisherId, &payout, &revenue, &cost); err != nil {
			panic(err)
		}
		stat := Stat{advertiserId, publisherId, payout, revenue, cost}
		stats = append(stats, stat)

	}
	return stats

}
