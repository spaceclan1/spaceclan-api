package models

type ActionRes struct {
	QueryTimeMs float64 `json:"query_time_ms"`
	Cached      bool    `json:"cached"`
	Lib         int     `json:"lib"`
	Total       struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	Actions []Action `json:"actions"`
}
