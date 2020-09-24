package core

import (
	"net/http"
	"sort"
)

type GBonFilter struct {
	Priority int
	Handler  http.HandlerFunc
}

type ByPriority []*GBonFilter

func (b ByPriority) Len() int           { return len(b) }
func (b ByPriority) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByPriority) Less(i, j int) bool { return b[i].Priority < b[j].Priority }

func (filters ByPriority) AddFilter(filter *GBonFilter) ByPriority {
	filters = append(filters, filter)
	sort.Sort(ByPriority(filters))

	return filters
}
