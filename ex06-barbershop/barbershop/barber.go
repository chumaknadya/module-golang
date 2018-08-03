package barbershop

import (
	"fmt"
	"time"
)

type Barber struct {
	name   string
	WakeMe chan *Client
}

func (b *Barber) GetName() string {
	return b.name
}

func NewBarber(name string) *Barber {
	barber := new(Barber)
	barber.name = name
	barber.WakeMe = make(chan *Client)
	return barber
}

func (b *Barber) ManageShop(shop *Shop) {
	for {
		select {
		case client := <-shop.WaitingRoom:
			fmt.Printf("%s : 'Cuts hair %s'\n", b.GetName(), client.GetName())
			time.Sleep(time.Millisecond * 20)
			fmt.Printf("%s : 'I am finished %s'\n", b.GetName(), client.GetName())
		default:
			fmt.Printf("%s :'I waant to sleep ZZZzzzz'\n", b.GetName())
			client := <-b.WakeMe
			fmt.Printf("%s :'I am waked by %s'\n", b.GetName(), client.GetName())
		}
	}
}
