package barbershop

import (
	"fmt"
	"time"
)

type Client struct {
	name string
}

func (client *Client) GetName() string {
	return client.name
}

func NewClient(name string) *Client {
	client := new(Client)
	client.name = name
	return client
}

func (client *Client) EnterShop(shop *Shop) {
	for i := 0; i < 3; i++ {
		if i > 0 {
			fmt.Printf("%s: 'I am waiting'\n", client.GetName())
		}
		select {
		case shop.WaitingRoom <- client:
			fmt.Printf("%s: 'I found a seat'\n", client.GetName())
			select {
			case shop.Barber.WakeMe <- client:
				fmt.Printf("%s: 'Wake up , barber'\n", client.GetName())
			default:
			}
			return
		default:
			fmt.Printf("%s: 'I am go out'\n", client.GetName())
			time.Sleep(time.Millisecond * 100)
		}
	}
	fmt.Printf("%s: 'I didn`t found a seat'\n", client.GetName())
}
