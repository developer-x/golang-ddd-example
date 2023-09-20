package utils

type PageResponse[T any] struct {
	Content   []T
	Page      int
	PageSize  int
	PageCount int
}

func NewPageResponse[T any](
	content []T,
	page int,
	pageSize int,
	pageCount int,
) PageResponse[T] {
	return PageResponse[T]{
		Content:   content,
		Page:      page,
		PageSize:  pageSize,
		PageCount: pageCount,
	}
}
