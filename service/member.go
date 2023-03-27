package service

import (
	"fmt"
	"test-k-style-hub/dto"
	"test-k-style-hub/entity"
	"test-k-style-hub/repository"

	"gorm.io/gorm"
)

type adaptorMemberService struct {
	memberRepo repository.MemberRepository
	db         *gorm.DB
}

func NewMemberService(
	db *gorm.DB,
	memberRepo repository.MemberRepository,
) MemberService {
	return &adaptorMemberService{
		db:         db,
		memberRepo: memberRepo,
	}
}

type MemberService interface {
	Create(req *dto.CreateMemberRequest) (*dto.CreateMemberResponse, error)
	Update(req *dto.UpdateMemberRequest) (*dto.UpdateMemberResponse, error)
	Delete(req *dto.DeleteMemberRequest) (*dto.DeleteMemberResponse, error)
	FindAll() (*dto.FindAllMemberResponse, error)
}

func (a *adaptorMemberService) Create(req *dto.CreateMemberRequest) (*dto.CreateMemberResponse, error) {
	var (
		member = new(entity.Members)
		res    = new(dto.CreateMemberResponse)
	)

	member.Username = req.Username
	member.Gender = req.Gender
	member.SkinType = req.SkinType
	member.SkinColor = req.SkinColor

	if err := a.memberRepo.Create(a.db, member); err != nil {
		return nil, err
	}

	res.Message = "Member Created!!"

	return res, nil
}

func (a *adaptorMemberService) Update(req *dto.UpdateMemberRequest) (*dto.UpdateMemberResponse, error) {
	var (
		member = new(entity.Members)
		res    = new(dto.UpdateMemberResponse)
	)

	member.MemberID = req.MemberID
	member.Username = req.Username
	member.Gender = req.Gender
	member.SkinType = req.SkinType
	member.SkinColor = req.SkinColor

	if err := a.memberRepo.Update(a.db, member); err != nil {
		return nil, err
	}

	res.Message = fmt.Sprintf("Member ID : %d Updated!!", req.MemberID)

	return res, nil
}

func (a *adaptorMemberService) Delete(req *dto.DeleteMemberRequest) (*dto.DeleteMemberResponse, error) {
	var (
		member = new(entity.Members)
		res    = new(dto.DeleteMemberResponse)
	)

	member.MemberID = req.MemberID

	if err := a.memberRepo.HardDelete(a.db, member); err != nil {
		return nil, err
	}

	res.Message = fmt.Sprintf("Member ID : %d Deleted!!", req.MemberID)

	return res, nil
}

func (a *adaptorMemberService) FindAll() (*dto.FindAllMemberResponse, error) {
	var (
		listMember = new([]entity.Members)
		err        error
		// member     = new(entity.Members)
		res = new(dto.FindAllMemberResponse)
	)

	if listMember, err = a.memberRepo.FindAll(a.db); err != nil {
		return nil, err
	}

	res.Data = listMember
	res.Message = "request was successfully"

	return res, nil
}
