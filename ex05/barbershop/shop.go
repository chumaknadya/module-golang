package barbershop

type Shop struct {
	Barber      *Barber
	WaitingRoom chan *Client
}

func NewShop(barber *Barber, NumberOfSeats int) *Shop {
	shop := new(Shop)
	shop.Barber = barber
	shop.WaitingRoom = make(chan *Client, NumberOfSeats)
	return shop
}
