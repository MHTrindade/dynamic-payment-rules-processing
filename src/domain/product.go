package domain

type Product struct {
	Description           string      `json:"description"`
	IsPhysical            bool        `json:"is_physical"`  // Indica se o produto é físico
	ProductType           ProductType `json:"product_type"` // Tipo de produto (livro, físico, associação, vídeo, etc.)
	RequiresFirstAidVideo bool        `json:"requires_first_aid_video"`
}
