package dto

type (
	LikeReviewCreateRequest struct {
		MemberID int `json:"member_id"`
		ReviewID int `json:"review_id"`
	}

	LikeReviewCreateResponse struct {
		Message string `json:"message"`
	}
)

type (
	LikeReviewDeleteRequest struct {
		LikeReviewID int `param:"id"`
	}

	LikeReviewDeleteResponse struct {
		Message string `json:"message"`
	}
)
