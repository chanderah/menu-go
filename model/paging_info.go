package model

type PagingInfo struct {
	Field     `json:"field,omitempty"`
	Filter    string `json:"filter,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
	CountRows int64  `json:"countRows,omitempty"`
}

type Field struct {
	Column string `json:"column,omitempty"`
	Value  string `json:"value,omitempty"`
}
