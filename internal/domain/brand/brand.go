package brand

type Brand struct {
	id   int    // brand id
	name string // brand name
}

type BrandStorage interface {
	Create(b *Brand) *Brand
	Read(b *Brand) *Brand
	Update(b *Brand) *Brand
	Delete(id int) error
}
