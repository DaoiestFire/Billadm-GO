package dto

type TrQueryResult struct {
	Items        []*TransactionRecordDto `json:"items"`
	Total        int64                   `json:"total"`
	TrStatistics map[string]float64      `json:"trStatistics"`
}
