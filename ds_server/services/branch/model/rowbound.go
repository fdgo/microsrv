package model

type RowBound struct {
	Offset int
	Limit  int
}

func NewRowBound(pageNum int, pageSize int) RowBound {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return RowBound{
		Limit:  pageSize,
		Offset: (pageNum - 1) * pageSize,
	}
}
