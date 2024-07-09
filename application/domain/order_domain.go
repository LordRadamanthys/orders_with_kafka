package domain

type OrderEvent struct {
	OrderId  int    `json:"orderId" bson:"orderId,omitempty" copy:"CodeOrder"`
	ClientId int    `json:"clientId" bson:"clientId,omitempty" copy:"CodeClient"`
	Items    []Item `json:"itens" bson:"itens,omitempty" copy:"Items"`
}

type Item struct {
	Product string  `json:"product" bson:"product,omitempty" copy:"Product"`
	Quanty  int     `json:"quanty" bson:"quanty,omitempty" copy:"Quanty"`
	Price   float64 `json:"price" bson:"price" copy:"Price"`
}
