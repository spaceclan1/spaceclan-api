package dao

import (
	"context"
	"encoding/json"
	"spaceclan1/spaceclan-api/config"
	"spaceclan1/spaceclan-api/datasource"
	"spaceclan1/spaceclan-api/models"
)

var (
	Cacher = &cacher{}
	ctx    = context.Background()
)

type cacher struct {
}

func (c cacher) GetToFrom(to string, from string) ([]models.Heroestaking_transactions_agg, error) {
	cm := datasource.Rdb.Get(ctx, "DAY:"+to)
	r, err := cm.Result()
	if err != nil {
		return nil, err
	}
	d := make([]models.Heroestaking_transactions_agg, 0)
	ret := make([]models.Heroestaking_transactions_agg, 0)
	json.Unmarshal([]byte(r), &d)
	for _, s2 := range d {
		if s2.From == from {
			ret = append(ret, s2)
		}
	}
	return ret, nil
}

func (c cacher) CacheDay(a []models.Heroestaking_transactions_agg) {
	t := make(map[string][]models.Heroestaking_transactions_agg)
	rewards_to := make(map[string]map[string]float32)
	if len(a) > 0 {
		for _, agg := range a {
			t[agg.To] = append(t[agg.To], agg)
			if "VIP2 Payout" == agg.Memo {
				if _, ok := rewards_to["TOTAL"]; !ok {
					rewards_to["TOTAL"] = make(map[string]float32, 0)
				}
				if _, ok := rewards_to["TOTAL"][agg.Date]; !ok {
					rewards_to["TOTAL"][agg.Date] = 0
				}
				rewards_to["TOTAL"][agg.Date] += agg.Amount

				if _, ok := rewards_to[agg.To]; !ok {
					rewards_to[agg.To] = make(map[string]float32, 0)
				}
				if _, ok := rewards_to[agg.To][agg.Date]; !ok {
					rewards_to[agg.To][agg.Date] = 0
				}
				rewards_to[agg.To][agg.Date] += agg.Amount
			}
		}
		for to, aggs := range t {
			j, _ := json.Marshal(aggs)
			datasource.Rdb.Set(ctx, "DAY:"+to, j, config.REDISTTL)
		}
		for to, agg := range rewards_to {
			for date, amount := range agg {
				datasource.Rdb.HSet(ctx, "REWARD:"+to, date, amount)
			}
		}
	}
}

func (c cacher) GetRewards(s string) (map[string]string, error) {
	cm := datasource.Rdb.HGetAll(ctx, "REWARD:"+s)
	r, err := cm.Result()
	if err != nil {
		return nil, err
	}
	return r, nil
}
