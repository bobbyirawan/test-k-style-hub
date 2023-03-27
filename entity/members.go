package entity

type Members struct {
	MemberID       int              `gorm:"primaryKey;autoIncrement:true;index:idx_unique,unique;"`
	Username       string           `gorm:"index:idx_unique,unique;size:100"`
	Gender         string           `gorm:"size:255"`
	SkinType       string           `gorm:"size:50"`
	SkinColor      string           `gorm:"size:50"`
	ListReview     []ReviewProducts `json:"ListReview,omitempty" gorm:"Foreignkey:MemberID;references:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ListLikeReview []LikeReviews    `json:"ListLikeReview,omitempty" gorm:"Foreignkey:MemberID;references:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entity
}

func (Members) TableName() string {
	return "members"
}
