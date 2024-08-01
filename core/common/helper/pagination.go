package helper

type PaginationMeta struct {
	CurrentPage int64 `json:"current_page"`
	From        int64 `json:"from"`
	LastPage    int64 `json:"last_page"`
	PerPage     int64 `json:"per_page"`
	To          int64 `json:"to"`
	Total       int64 `json:"total"`
}

type Pagination struct {
	Data any             `json:"data"`
	Meta *PaginationMeta `json:"meta"`
}

func NewPaginationMeta(currentPage, perPage, total int64) *PaginationMeta {
	lastPage := (total / perPage) + 1
	from := (currentPage - 1) * perPage
	to := currentPage * perPage
	return &PaginationMeta{
		CurrentPage: currentPage,
		From:        from,
		LastPage:    lastPage,
		PerPage:     perPage,
		To:          to,
		Total:       total,
	}
}

func NewPagination(data any, meta *PaginationMeta) *Pagination {
	return &Pagination{
		Data: data,
		Meta: meta,
	}
}
