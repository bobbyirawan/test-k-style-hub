package controller

import (
	"test-k-style-hub/dto"
	"test-k-style-hub/service"

	"github.com/labstack/echo/v4"
)

type MemberController struct {
	memberService service.MemberService
}

func NewMemberController(memberService service.MemberService) *MemberController {
	return &MemberController{
		memberService: memberService,
	}
}

func (m *MemberController) Create(ctx echo.Context) error {
	var (
		req = new(dto.CreateMemberRequest)
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := m.memberService.Create(req)
	if err != nil {
		return err
	}

	return ctx.JSON(201, res)

}

func (m *MemberController) Update(ctx echo.Context) error {
	var (
		req = new(dto.UpdateMemberRequest)
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := m.memberService.Update(req)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)

}

func (m *MemberController) Delete(ctx echo.Context) error {
	var (
		req = new(dto.DeleteMemberRequest)
	)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	res, err := m.memberService.Delete(req)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)

}

func (m *MemberController) FindAll(ctx echo.Context) error {
	res, err := m.memberService.FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)

}
