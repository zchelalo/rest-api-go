package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

func New(page, perPage, total int) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
	}, nil
}

func (meta *Meta) Offset() int {
	return (meta.Page - 1) * meta.PerPage
}

func (meta *Meta) Limit() int {
	return meta.PerPage
}
