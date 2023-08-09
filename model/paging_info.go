package model

type PagingInfo struct {
	Filter    string `json:"filter,omitempty"`
	Limit     int    `json:"limit,"`
	Offset    int    `json:"offset"`
	SortField string `json:"sortField,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
	RowCount  int    `json:"rowCount,omitempty"`
}
