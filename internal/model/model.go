package model

import "fmt"

// Query 数据库查询参数
type Query[T any] struct {
	Page       int64
	PageSize   int64
	TID        int64
	Conditions T
}

func (q Query[T]) OffsetLimit() string {
	if q.Page > 0 && q.PageSize > 0 {
		offset := (q.Page - 1) * q.PageSize
		limit := q.PageSize
		return fmt.Sprintf(" LIMIT %d,%d ", offset, limit)
	}
	return ""
}

// Result 数据查询结果
type Result[T any] struct {
	Total int64
	Data  T
}
