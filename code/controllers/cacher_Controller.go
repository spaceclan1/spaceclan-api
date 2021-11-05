package controllers

import (
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/config"
	"spaceclan1/spaceclan-api/dao"
	"time"
)

var (
	CacherController = &cacher_controller{}
	oneDay           = time.Hour * 24
)

type cacher_controller struct {
	main_controller
}

func (c cacher_controller) CacheAndAggregate() {
	log.Info("started agg")
	c.aggregate()
}

func (c cacher_controller) aggregate() {
	t := time.Now()
	t = t.UTC()

	//get last aggregation date
	o := dao.OptionsImpl.Get("day_agg")
	ld, err := time.Parse(config.SQL_DATE_FORMAT, o.Value)
	if err != nil {
		log.Fatal(err)
	}
	d := t.Sub(ld)
	log.WithFields(log.Fields{
		"d":       d,
		"48hour":  time.Hour * 48,
		"compare": d > time.Hour*48,
	}).Info()
	if d > time.Hour*48 {
		c.aggregateDay(ld)
	} else {
		// aggregate current day
		c.aggregateDay(t)
		dao.OptionsImpl.Set("day_agg", t.Truncate(time.Duration(oneDay)).Format(config.SQL_DATE_FORMAT))
		// aggregate yesterday (in case we missed some during day change midnight)
		yesterday := t.Add(time.Duration(-oneDay))
		c.aggregateDay(yesterday)

		// aggregate current Month
		c.aggregateMonth(t)

		// aggregate yesterday (in case we missed some during day change midnight)
		prevMonth := t.AddDate(0, -1, 0)
		c.aggregateMonth(prevMonth)
	}
}

func (c cacher_controller) aggregateDay(t time.Time) {
	start := t.Truncate(time.Duration(oneDay))
	end := start.Add(time.Duration(oneDay))
	aggData := dao.Cacher.GetAggregation(start, end)
	log.WithFields(log.Fields{
		"t":     t,
		"start": start,
		"end":   end,
	}).Info()
	dao.Cacher.SaveAndCacheAggregatedDay(aggData)
}

func (c cacher_controller) aggregateMonth(t time.Time) {
	year, month, _ := t.Date()
	currentLocation := t.Location()

	start := time.Date(year, month, 1, 0, 0, 0, 0, currentLocation)
	end := time.Date(year, month+1, 0, 23, 59, 59, 999, currentLocation)

	log.WithFields(log.Fields{
		"start": start,
		"end":   end,
	}).Info()
	aggData := dao.Cacher.GetAggregationMonth(start, end)
	dao.Cacher.SaveAndCacheAggregatedMonth(aggData)
}
