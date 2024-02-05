package domain

type Product struct {
	Description           string      `json:"description"`
	IsPhysical            bool        `json:"is_physical"`
	ProductType           ProductType `json:"product_type"`
	RequiresFirstAidVideo bool        `json:"requires_first_aid_video"`
}
