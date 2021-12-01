package models

type (
	OrderId string

	// Order entity
	Order struct {
		id          OrderId `bson:"_id"`
		productCode string  `bson:"product_code"`
		quantity    int     `bson:"quantity"`
	}
)
