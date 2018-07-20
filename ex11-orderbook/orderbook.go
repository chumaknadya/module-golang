package orderbook

import "fmt"

type Orderbook struct {
	// TODO
	Bids       []*Order
	Ask        []*Order
	OpenOrders map[uint]*Order
}

func New() *Orderbook {
	// TODO
	Bids := make([]*Order, 0)
	Ask := make([]*Order, 0)
	OpenOrders := make(map[uint]*Order)
	return &Orderbook{Bids, Ask, OpenOrders}
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	// TODO
	switch order.Kind.String() {
	case "MARKET":
		fmt.Println("market")

	case "LIMIT":
		fmt.Println("limit")
		if order.Side.String() == "BID" {
			fmt.Println("BID")
			if len(orderbook.Ask) == 0 {
				orderbook.Bids = append(orderbook.Bids, order)
			} else {
				for _, ask := range orderbook.Ask {
					if order.Price == ask.Price {
						fmt.Println("order", order.Price)
						fmt.Println("sell price", ask.Price)
						fmt.Println("Exchange done BUY")
					}
				}
			}
		}

		if order.Side.String() == "ASK" {
			fmt.Println("ASK")
			if len(orderbook.Bids) == 0 {
				orderbook.Ask = append(orderbook.Ask, order)
			} else {
				for _, bid := range orderbook.Bids {
					if order.Price >= bid.Price {
						fmt.Println("Exchange done SELL")
					}
				}
			}
		}
	}
	return nil, nil
}
