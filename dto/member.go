package dto

type (
	CreateMemberRequest struct {
		Username  string `json:"username"`
		Gender    string `json:"Gender"`
		SkinType  string `json:"skin_type"`
		SkinColor string `json:"skin_color"`
	}

	CreateMemberResponse struct {
		Message string `json:"message"`
	}
)

type (
	UpdateMemberRequest struct {
		MemberID  int    `param:"id"`
		Username  string `json:"username"`
		Gender    string `json:"Gender"`
		SkinType  string `json:"skin_type"`
		SkinColor string `json:"skin_color"`
	}

	UpdateMemberResponse struct {
		Message string `json:"message"`
	}
)

type (
	DeleteMemberRequest struct {
		MemberID int `param:"id"`
	}

	DeleteMemberResponse struct {
		Message string `json:"message"`
	}
)

type FindAllMemberResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
