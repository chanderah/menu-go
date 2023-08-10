package model

type PagingInfo struct {
	Filter string `json:"filter,omitempty"`
	FilterField
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
}

type FilterField struct {
	Column string `json:"column,omitempty"`
	Value  string `json:"value,omitempty"`
}
