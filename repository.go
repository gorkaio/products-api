package main

type ProductRepo interface {
	save(Product) (Product, error)
	load(ProductId) (Product, error)
}

type InMemoryRepo struct {
	items map[ProductId]Product
}

func GetProductRepository() ProductRepo {
	return &InMemoryRepo{
		items: make(map[ProductId]Product),
	}
}

func (r *InMemoryRepo) save(p Product) (Product, error) {
	r.items[p.Id] = p
	return p, nil
}

func (r *InMemoryRepo) load(id ProductId) (Product, error) {
	return r.items[id], nil
}
