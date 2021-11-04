package dao

import (
	"fmt"
	"log"
	"spaceclan1/spaceclan-api/datasource"
	actions "spaceclan1/spaceclan-api/models/actions"
	"strings"
)

var (
	HeroeStakingTransactionsImpl = &heroestaking_transactions{}
)

const (
	markers   = "(?,?,?,?,?,?,?,?,?,?)"
	insertSql = "INSERT IGNORE INTO `wax_data`.`heroestaking_transactions`(`trans_time`,`block_num`,`action`,`from`,`to`,`amount`,`symbol`,`memo`,`trx_id`,`action_ordinal`) VALUES %s"
)

type heroestaking_transactions struct {
}

func (h heroestaking_transactions) Create(a actions.Action) {
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
func (h heroestaking_transactions) CreateBulk(aa []actions.Action) {
	valueArgs := []interface{}{}
	m := make([]string, 0)
	for _, a := range aa {
		m = append(m, markers)
		valueArgs = append(valueArgs, a.Timestamp, a.BlockNum, a.Act.Name, a.Act.Data.From, a.Act.Data.To, a.Act.Data.Amount, a.Act.Data.Symbol, a.Act.Data.Memo, a.TrxID, a.ActionOrdinal)
	}
	stm, err := datasource.MainDb.Prepare(fmt.Sprintf(insertSql, strings.Join(m, ",")))
	if err != nil {
		log.Fatal(err)
	}
	_, err = stm.Exec(valueArgs...)
	if err != nil {
		log.Fatal(err)
	}
}
