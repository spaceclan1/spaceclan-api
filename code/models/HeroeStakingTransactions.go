package models

import "time"

type Heroestaking_transactions struct {
	Id            int
	TransTime     time.Time
	BlockNum      int64
	Action        string
	From          string
	To            string
	amount        float32
	Symbol        string
	Memo          string
	TrxId         string
	ActionOrdinal int
	Timestamp     time.Time
}
