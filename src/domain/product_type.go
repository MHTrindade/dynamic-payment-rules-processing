package domain

type ProductType string

const (
	CommonProductType       ProductType = "comum"
	PhysicalProductType     ProductType = "físico"
	BookProductType         ProductType = "livro"
	SubscriptionProductType ProductType = "assinatura"
	VideoProductType        ProductType = "video"
)
