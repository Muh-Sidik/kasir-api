package request

import "strconv"

type PaginateReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginateQuery struct {
	PaginateReq
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func Paginate(page, pageSize string) *PaginateQuery {
	pageInt, errPage := strconv.Atoi(page)
	pageSizeInt, errPageSize := strconv.Atoi(pageSize)

	if errPage != nil {
		pageInt = 1
	}
	if errPageSize != nil {
		pageSizeInt = 10
	}

	if pageInt < 1 {
		pageInt = 1
	}
	if pageSizeInt < 1 {
		pageSizeInt = 10
	}

	offset := (pageInt - 1) * pageSizeInt
	limit := pageInt

	res := &PaginateQuery{
		Limit:  limit,
		Offset: offset,
	}
	res.Page = pageInt
	res.PageSize = pageSizeInt
	return res
}
