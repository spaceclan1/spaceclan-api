package dao

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/config"
	"spaceclan1/spaceclan-api/datasource"
	"spaceclan1/spaceclan-api/models"
	"time"
)

var (
	Cacher = &cacher{}
	ctx    = context.Background()
)

const (
	aggSqlDays     = "SELECT h.`action`, h.`from`, h.`to`, h.`symbol`,  h.`memo`, SUM(h.`amount`) amount FROM `heroestaking_transactions` h WHERE h.`trans_time`>=? AND h.`trans_time`<? GROUP BY h.`action`, h.`from`, h.`to`,h.`symbol`, h.`memo`"
	aggSqlMonth    = "SELECT h.`action`, h.`from`, h.`to`, h.`symbol`,  h.`memo`, SUM(h.`amount`) amount FROM `heroestaking_transactions_agg_day` h WHERE h.`date`>=? AND h.`date`<? GROUP BY h.`action`, h.`from`, h.`to`,h.`symbol`, h.`memo`"
	aggInsertDay   = "INSERT INTO `heroestaking_transactions_agg_day` (`date`,`action`,`from`,`to`,`symbol`,`memo`,`amount`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE amount=?"
	aggInsertMonth = "INSERT INTO `heroestaking_transactions_agg_month` (`date`,`action`,`from`,`to`,`symbol`,`memo`,`amount`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE amount=?"
)

type cacher struct {
}

func (c cacher) GetAggregation(from time.Time, to time.Time) []models.Heroestaking_transactions_agg {
	rows, err := datasource.MainDb.Query(aggSqlDays, from, to)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	l := make([]models.Heroestaking_transactions_agg, 0)
	for rows.Next() {
		o := models.Heroestaking_transactions_agg{}
		rows.Scan(&o.Action, &o.From, &o.To, &o.Symbol, &o.Memo, &o.Amount)
		o.Date = from.Format(config.SQL_DATETIME_FORMAT)
		l = append(l, o)
	}
	return l
}

func (c cacher) SaveAndCacheAggregatedDay(data []models.Heroestaking_transactions_agg) {
	if len(data) == 0 {
		return
	}
	tx, err := datasource.MainDb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stm, err2 := tx.Prepare(aggInsertDay)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, d := range data {
		_, err = stm.Exec(d.Date, d.Action, d.From, d.To, d.Symbol, d.Memo, d.Amount, d.Amount)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		//to avoid second loop data will be send to redis here
		k := "DAY:" + d.To + ":" + d.Date
		k2 := "DAY_USER" + d.To
		j, _ := json.Marshal(d)
		datasource.Rdb.Set(ctx, k, j, config.REDISTTL)
		datasource.Rdb.HSet(ctx, k2, d.Date)
	}
	tx.Commit()
}

func (c cacher) GetAggregationMonth(from time.Time, to time.Time) []models.Heroestaking_transactions_agg {
	rows, err := datasource.MainDb.Query(aggSqlMonth, from, to)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	l := make([]models.Heroestaking_transactions_agg, 0)
	for rows.Next() {
		o := models.Heroestaking_transactions_agg{}
		rows.Scan(&o.Action, &o.From, &o.To, &o.Symbol, &o.Memo, &o.Amount)
		o.Date = from.Format(config.SQL_DATE_FORMAT)
		l = append(l, o)
	}
	return l
}

func (c cacher) SaveAndCacheAggregatedMonth(data []models.Heroestaking_transactions_agg) {
	tx, err := datasource.MainDb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stm, err2 := tx.Prepare(aggInsertMonth)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, d := range data {
		_, err = stm.Exec(d.Date, d.Action, d.From, d.To, d.Symbol, d.Memo, d.Amount, d.Amount)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		j, _ := json.Marshal(d)
		k := "MONTH:" + d.Date + ":" + d.Action + ":" + d.From + ":" + d.To + ":" + d.Symbol + ":" + d.Memo + ":"
		datasource.Rdb.Set(ctx, k, j, config.REDISTTL)
		datasource.Rdb.HSet(ctx, "MONTH:"+d.To, k)
		datasource.Rdb.Set(ctx, "MONTH:"+d.To+":"+d.From, j, config.REDISTTL)
	}
	tx.Commit()
}

func (c cacher) GetMonthlyRewardFromCache(wallet string) {
	k := "MONTH_USER:" + wallet
	datasource.Rdb.HGetAll(ctx, k)
}

func (c cacher) GetDailyRewardFromCache() {

}
