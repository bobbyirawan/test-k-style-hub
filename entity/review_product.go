package entity

type ReviewProducts struct {
	ReviewID       int `gorm:"index:idx_unique,unique"`
	ProductID      int
	MemberID       int
	ReviewDesc     string        `gorm:"size:500"`
	ListLikeReview []LikeReviews `json:"ListLikeReview,omitempty" gorm:"Foreignkey:ReviewID;references:ReviewID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entity
}

func (ReviewProducts) TableName() string {
	return "review_products"
}
