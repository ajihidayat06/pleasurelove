package models

type GetListStruct struct {
	Filters map[string][2]interface{}
	Page    int
	Limit   int
	OrderBy string
	SortBy  string
}
