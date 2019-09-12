package main

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/prometheus/common/log"
	"math"
	. "tracker-sms-svc/models"
)

func main() {

	stats := FindStat()

	l := len(stats)
	created := 0
	changed := 0

	for i, stat := range stats {
		log.Infof("(%d/%d)", i+1, l)
		pretty.Log("stat", stat)
		aid := stat.AdvertiserId
		pid := stat.PublisherId
		cost := stat.Cost
		rule := FindRule(aid, pid)
		if rule == nil {
			rule = &Rule{AdvertiserId: aid, PublisherId: pid, ShaveInterval: DefInterval}
			pretty.Log("created rule", rule)
		} else {
			pretty.Log("exists rule", rule)
		}

		if rule.IsManual {
			log.Infoln("skip rule, it is manual")
			continue
		}
		adv := FindAdv(aid)
		if adv == nil {
			log.Infoln("skip. Adv not found", aid)
			continue
		}

		advCost := adv.Cost

		if adv.Cost == 0 {
			log.Infoln("skip. AdvCost=0")
			continue
		}

		inc := intervalCalc(cost, advCost)

		rule.ChangeInterval(inc)
		log.Infoln("ShaveInterval:", rule.ShaveInterval)

		if rule.Id > 0 {
			changed += 1
		} else {
			created += 1
		}
		rule.SaveRule()

	}

	log.Infoln("Total Stats:", l)
	log.Infoln("Rules Created:", created)
	log.Infoln("Rules Changed:", changed)

}
func intervalCalc(cost float64, advCost float64) int {

	diff := int(math.Abs(math.Round((cost - advCost) * 100)))

	r := func(msg string, inc int) int {
		msg = fmt.Sprintf("[%s]{cost:%.2f, advCost:%.2f, diff:%d} [inc:%d]", msg, cost, advCost, diff, inc)
		log.Infoln(msg)
		return inc
	}
	if cost >= 1.1 {
		return r("cost >= 1.1", -20)
	}

	if cost >= 1.0 {
		return r("cost >= 1.0", -10)
	}

	msg := "cost < 1.0 "
	if cost >= advCost {
		msg += "cost >= advCost "
		if diff >= 5 {
			return r(msg+"diff >= 5", -5)
		}
		return r(msg+"diff < 5", -diff)

	}

	if cost < advCost {
		msg += "cost < advCost"
		if diff >= 10 {
			return r(msg+"diff >= 10", 10)
		}
		return r(msg+"diff < 10", diff)

	}

	return r("default", 0)
}

/*func dd(args ...interface{}) {
	pretty.Println(args)
	os.Exit(20)
}
*/
