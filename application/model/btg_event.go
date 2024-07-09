package model

type OrderEvent struct {
	CodeOrder  int    `json:"codigoPedido"`
	CodeClient int    `json:"codigoCliente"`
	Items      []Item `json:"itens"`
}

type Item struct {
	Product string  `json:"produto"`
	Quanty  int     `json:"quantidade"`
	Price   float64 `json:"preco"`
}
