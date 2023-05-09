package main

import "errors"

type ProductRepo interface {
	save(Product) (Product, error)
	retrieve(ProductId) (Product, error)
	retrieveAll() (map[ProductId]Product, error)
	delete(ProductId) (ProductId, error)
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

func (r *InMemoryRepo) retrieve(id ProductId) (Product, error) {
	if item, ok := r.items[id]; ok {
		return item, nil
	}

	return Product{}, errors.New("item not found")
}

func (r *InMemoryRepo) retrieveAll() (map[ProductId]Product, error) {
	return r.items, nil
}

func (r *InMemoryRepo) delete(id ProductId) (ProductId, error) {
	delete(r.items, id)
	return id, nil
}
