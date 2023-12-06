package domain

import "math"

type Pagination struct {
	CurrentPage  uint 			`json:"current_page"`
	NextPage     uint 			`json:"next_page"`
	PreviousPage uint 			`json:"previous_page"`
	MaxPage      uint 			`json:"max_page"`
	Queries 	 map[string]any `json:"queries"`
}

func CreatePagination(currentPage uint, maxPage uint, q map[string]any) Pagination {
	next := uint(math.Min(float64(currentPage) + 1, float64(maxPage)))
	previous := uint(math.Max(float64(currentPage) - 1, float64(1)))
	return Pagination{
		CurrentPage: currentPage,
		NextPage: next,
		PreviousPage: previous,
		MaxPage: maxPage,
		Queries: q,
	}
}