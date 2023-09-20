package utils

type Page int
type PageSize int
type SortBy string
type Direction bool

type PageRequest struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
	SortBy    string `form:"sortBy"`
	Direction string `form:"direction"`
}

func NewPageRequest(
	page int,
	pageSize int,
	sortBy string,
	direction string,
) PageRequest {
	return PageRequest{
		Page:      page,
		PageSize:  pageSize,
		SortBy:    sortBy,
		Direction: direction,
	}
}
