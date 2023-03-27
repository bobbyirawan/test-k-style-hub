package entity

type LikeReviews struct {
	LikeReviewID int `gorm:"primaryKey;autoIncrement:true;index:idx_unique,unique;"`
	ReviewID     int
	MemberID     int
}

func (LikeReviews) TableName() string {
	return "like_reviews"
}
