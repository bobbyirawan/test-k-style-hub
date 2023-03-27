package entity

type Products struct {
	ProductID   int              `gorm:"primaryKey;autoIncrement:true;index:idx_unique,unique;"`
	ProductName string           `gorm:"index:idx_unique,unique;size:100"`
	Price       float64          `gorm:"size:500"`
	ListReview  []ReviewProducts `json:"ListReview" gorm:"Foreignkey:ProductID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entity
}

func (Products) TableName() string {
	return "products"
}
