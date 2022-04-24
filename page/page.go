package page

type Pagination struct {
	CurrPage     int
	PageSize     int
	MaxCount     int
	MaxPageCount int
}

func New(currPage int, pageSize int, maxCount int) (res *Pagination) {
	res = &Pagination{}
	res.set(currPage, pageSize, maxCount)
	return
}

func (p *Pagination) set(currPage int, pageSize int, maxCount int) {
	p.CurrPage = currPage
	p.PageSize = pageSize
	p.MaxCount = maxCount

	if p.PageSize <= 0 {
		p.PageSize = 15
	}
	if p.MaxCount <= 0 {
		p.MaxCount = 0
	}
	if p.CurrPage <= 0 {
		p.CurrPage = 1
	}
	p.MaxPageCount = p.MaxCount / p.PageSize
	if p.MaxCount%p.PageSize > 0 {
		p.MaxPageCount += 1
	}
	if p.MaxPageCount <= 0 {
		p.MaxPageCount = 1
	}
}

func (p *Pagination) Offset() (offset int) {
	if p.CurrPage > p.MaxPageCount {
		return p.MaxCount
	}
	return (p.CurrPage - 1) * p.PageSize
}
