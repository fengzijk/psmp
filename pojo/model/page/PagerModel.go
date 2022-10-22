package page

import "math"

// 分页对象
type PagerModel struct {
	Page      int         `form:"pageNum"  json:"pageNum"`     //当前页
	PageSize  int         `form:"pageSize"  json:"pageSize"`   //每页条数
	Total     int         `form:"total"  json:"total"`         //总条数
	PageCount int         `form:"pageCount"  json:"pageCount"` //总页数
	Nums      []int       `form:"nums"  json:"nums"`           //分页序数
	NumsCount int         `form:"numsCount"  json:"numsCount"` //总页序数
	List      interface{} `json:"list"`
}

func CreatePager(page, pageSize, total int, list interface{}) *PagerModel {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	pageCount := math.Ceil(float64(total) / float64(pageSize))

	pager := new(PagerModel)
	pager.Page = page
	pager.PageSize = pageSize
	pager.Total = total
	pager.PageCount = int(pageCount)
	pager.NumsCount = 7
	pager.setNums()
	pager.List = list
	return pager
}

func (this *PagerModel) setNums() {
	this.Nums = []int{}
	if this.PageCount == 0 {
		return
	}

	half := math.Floor(float64(this.NumsCount) / float64(2))
	begin := this.Page - int(half)
	if begin < 1 {
		begin = 1
	}

	end := begin + this.NumsCount - 1
	if end >= this.PageCount {
		begin = this.PageCount - this.NumsCount + 1
		if begin < 1 {
			begin = 1
		}
		end = this.PageCount
	}

	for i := begin; i <= end; i++ {
		this.Nums = append(this.Nums, i)
	}
}
