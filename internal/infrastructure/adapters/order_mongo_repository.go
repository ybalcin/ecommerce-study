package adapters

import "github.com/ybalcin/ecommerce-study/pkg/mgo"

// OrderRepository implements repositories.OrderRepository
type OrderRepository struct {
	orders mgo.Collection
}
