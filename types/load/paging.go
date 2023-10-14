package load

import "fmt"

const DefaultPagingSize = 100

const MaxPagingSize = 2000

const FirstPage = 1

type Page struct {
	Size    int64 `json:"page_size"` // Page Size, default 100
	Current int64 `json:"page_no"`   // From One
}

func PageALL() *Page {
	return &Page{
		Size:    MaxPagingSize,
		Current: FirstPage,
	}
}

type Paging struct {
	Size      int64 `json:"page_size"`  // Page Size, default 100
	Current   int64 `json:"page_no"`    // From One
	Total     int64 `json:"page_total"` // The Page Count
	ItemTotal int64 `json:"item_total"` // The Item Count
}

func (paging *Paging) ToString() string {
	return fmt.Sprintf("size: %d|current: %d|total: %d|item_total: %d",
		paging.Size, paging.Current, paging.Total, paging.ItemTotal)
}

func (paging *Paging) WithItemTotal(itemTotal int64) {
	if paging.Size == 0 {
		paging.Size = DefaultPagingSize
	}
	if paging.Current < FirstPage {
		paging.Current = FirstPage
	}

	if itemTotal == 0 {
		paging.Total = 0
		paging.Current = FirstPage
		return
	}
	paging.ItemTotal = itemTotal
	paging.Total = paging.ItemTotal / paging.Size
	if paging.ItemTotal%paging.Size != 0 {
		paging.Total += 1
	}

	if paging.Current > paging.Total {
		paging.Current = paging.Total
	}

	if paging.Current < 1 {
		paging.Current = 1
	}
}

func (paging *Paging) Skip() int64 {
	return (paging.Current - 1) * paging.Size
}

func (paging *Paging) Limit() int64 {
	return paging.Size
}

func PagingOfPage(page *Page) Paging {
	if page == nil {
		return PagingOf(DefaultPagingSize, FirstPage)
	}
	return PagingOf(page.Size, page.Current)
}

func PagingOf(size int64, current int64) Paging {
	if size <= 0 {
		size = DefaultPagingSize
	}
	if size > MaxPagingSize {
		size = MaxPagingSize
	}
	if current < FirstPage {
		current = FirstPage
	}
	return Paging{
		Size:      size,
		Current:   current,
		Total:     0,
		ItemTotal: 0,
	}
}

func PagingALL() Paging {
	return PagingOf(MaxPagingSize, FirstPage)
}

type Pagination struct {
	Paging Paging `bson:"paging" json:"paging"`
	Items  []any  `bson:"items" json:"items"`
}

func PagingWrap[T any](paging Paging, arr []*T, wrap func(m *T) any) Pagination {
	if len(arr) == 0 {
		return Pagination{
			Paging: paging,
			Items:  []any{},
		}
	}
	items := make([]any, len(arr))
	for i, m := range arr {
		items[i] = wrap(m)
	}
	return Pagination{
		Paging: paging,
		Items:  items,
	}
}
