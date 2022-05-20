package appModel

type OrderModel interface {
	GetAllOrder() ([]Order, error)
	AddOrder(Order) (Order, error)
}
