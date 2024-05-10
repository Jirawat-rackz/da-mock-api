package model

// Pagination is a struct for paginate
type Pagination struct {
	Page  int `form:"page,default=1"`
	Limit int `form:"limit,default=15"`
}

func (p Pagination) IsValid() bool {
	if p.Page <= 0 || p.Limit <= 0 {
		return false
	}
	return true
}
