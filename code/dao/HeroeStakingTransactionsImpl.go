package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/datasource"
	actions "spaceclan1/spaceclan-api/models/actions"
	"strings"
)

var (
	HeroeStakingTransactionsImpl = &heroestaking_transactionsimps{}
)

const (
	markers   = "(?,?,?,?,?,?,?,?,?,?)"
	insertSql = "INSERT IGNORE INTO `wax_data`.`heroestaking_transactions`(`trans_time`,`block_num`,`action`,`from`,`to`,`amount`,`symbol`,`memo`,`trx_id`,`action_ordinal`) VALUES %s"
)

type heroestaking_transactionsimps struct {
}

func (h heroestaking_transactionsimps) Create(a actions.Action) {
	stm, err := datasource.MainDb.Prepare(fmt.Sprintf(insertSql, markers))
	if err != nil {
		log.Fatal(err)
	}
	_, err = stm.Exec(a.Timestamp, a.BlockNum, a.Act.Name, a.Act.Data.From, a.Act.Data.To, a.Act.Data.Amount, a.Act.Data.Symbol, a.Act.Data.Memo, a.TrxID, a.ActionOrdinal)
	if err != nil {
		log.Fatal(err)

	}
	defer stm.Close()
}
func (h heroestaking_transactionsimps) CreateBulk(aa []actions.Action) {
	valueArgs := []interface{}{}
	chunks := h.chunks(aa, 100)
	tx, err := datasource.MainDb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(chunks); i++ {
		m := make([]string, 0)
		for _, a := range chunks[i] {
			m = append(m, markers)
			valueArgs = append(valueArgs, a.Timestamp, a.BlockNum, a.Act.Name, a.Act.Data.From, a.Act.Data.To, a.Act.Data.Amount, a.Act.Data.Symbol, a.Act.Data.Memo, a.TrxID, a.ActionOrdinal)
		}
		stm, err := tx.Prepare(fmt.Sprintf(insertSql, strings.Join(m, ",")))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		_, err = stm.Exec(valueArgs...)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func (h heroestaking_transactionsimps) chunks(xs []actions.Action, chunkSize int) [][]actions.Action {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]actions.Action, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided

}
