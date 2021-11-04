package models

type Action struct {
	Timestamp string `json:"timestamp"`
	BlockNum  int64  `json:"block_num"`
	TrxID     string `json:"trx_id"`
	Act       struct {
		Account       string `json:"account"`
		Name          string `json:"name"`
		Authorization []struct {
			Actor      string `json:"actor"`
			Permission string `json:"permission"`
		} `json:"authorization"`
		Data struct {
			From     string  `json:"from"`
			To       string  `json:"to"`
			Amount   float64 `json:"amount"`
			Symbol   string  `json:"symbol"`
			Memo     string  `json:"memo"`
			Quantity string  `json:"quantity"`
		} `json:"data"`
	} `json:"act"`
	Notified             []string `json:"notified"`
	CPUUsageUs           int      `json:"cpu_usage_us"`
	NetUsageWords        int      `json:"net_usage_words"`
	GlobalSequence       int64    `json:"global_sequence"`
	Producer             string   `json:"producer"`
	ActionOrdinal        int      `json:"action_ordinal"`
	CreatorActionOrdinal int      `json:"creator_action_ordinal"`
}
