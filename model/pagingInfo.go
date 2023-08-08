package model

type PagingInfo struct {
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Filter    string `json:"filter,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SortOrder int    `json:"sortOrder,omitempty"`
	RowCount  int    `json:"rowCount,omitempty"`
}
