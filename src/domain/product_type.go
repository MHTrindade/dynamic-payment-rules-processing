package domain

type ProductType string

const (
	CommonProductType   ProductType = "comum"
	PhysicalProductType ProductType = "físico"
	BookProductType     ProductType = "livro"
	UnknownProductType  ProductType = "desconhecido"
)
