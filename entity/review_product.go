package entity

type ReviewProducts struct {
	ReviewID   int `gorm:"index:idx_unique,unique"`
	ProductID  int
	MemberID   int
	ReviewDesc int `gorm:"size:500"`
	Entity
}

func (ReviewProducts) TableName() string {
	return "review_products"
}
