package dto

type (
	ProductFindByIDRequest struct {
		ProductID int `query:"product_id"`
	}

	ProductFindByIDResponse struct {
		Message string                        `json:"message"`
		Data    []ProductFindByIDResponseData `json:"data"`
	}

	ProductFindByIDResponseData struct {
		Username         string `json:"username"`
		Gender           string `json:"gender"`
		SkinType         string `json:"skin_type"`
		SkinColor        string `json:"skin_color"`
		DescReview       string `json:"desc_review"`
		JumlahLikeReview int    `json:"jumlah_like_review"`
	}
)
