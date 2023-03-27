package repository

import (
	"test-k-style-hub/entity"

	"gorm.io/gorm"
)

type adaptorMemberRepository struct {
}

func NewMemberRepository() MemberRepository {
	return &adaptorMemberRepository{}
}

type MemberRepository interface {
	FindAll(db *gorm.DB) (*[]entity.Members, error)
	Create(db *gorm.DB, member *entity.Members) error
	Update(db *gorm.DB, member *entity.Members) error
	HardDelete(db *gorm.DB, member *entity.Members) error
}

func (a *adaptorMemberRepository) FindAll(db *gorm.DB) (*[]entity.Members, error) {
	member := new([]entity.Members)
	if err := db.Preload("ListReview").Find(member).Error; err != nil {
		return nil, err
	}
	return member, nil
}

func (a *adaptorMemberRepository) Create(db *gorm.DB, member *entity.Members) error {
	if err := db.Create(member).Error; err != nil {
		return err
	}

	return nil
}

func (a *adaptorMemberRepository) Update(db *gorm.DB, member *entity.Members) error {
	if err := db.Save([]entity.Members{*member}).Error; err != nil {
		return err
	}

	return nil
}

func (a *adaptorMemberRepository) HardDelete(db *gorm.DB, member *entity.Members) error {
	if err := db.Unscoped().Delete(member).Error; err != nil {
		return err
	}

	return nil
}
