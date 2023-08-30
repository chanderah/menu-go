package model

type PagingInfo struct {
	Field     Field  `json:"field,omitempty"`
	Filter    string `json:"filter,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
	// Page      int    `json:"page,omitempty"`
}

type Field struct {
	Column string      `json:"column,omitempty"`
	Value  interface{} `json:"value,omitempty"`
}
