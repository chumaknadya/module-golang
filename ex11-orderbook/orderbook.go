package orderbook

import "sort"

type Orderbook struct {
	Bids []*Order
	Asks []*Order
}

func New() *Orderbook {
	Bids := []*Order{}
	Asks := []*Order{}
	return &Orderbook{Bids, Asks}
}

func NewTrade(Bid *Order, Ask *Order, Volume uint64, Price uint64) *Trade {
	return &Trade{Bid, Ask, Volume, Price}
}

func (orderbook *Orderbook) AddAskOrder(order *Order) {
	//descending order
	if order.Volume > 0 {
		orderbook.Asks = append(orderbook.Asks, order)
		sort.Slice(orderbook.Asks, func(i, j int) bool {
			return orderbook.Asks[i].Price > orderbook.Asks[j].Price
		})
	}
}

func (orderbook *Orderbook) AddBidOrder(order *Order) {
	//ascending order
	if order.Volume > 0 {
		orderbook.Bids = append(orderbook.Bids, order)
		sort.Slice(orderbook.Bids, func(i, j int) bool {
			return orderbook.Bids[i].Price < orderbook.Bids[j].Price
		})
	}
}

func CheckConditions(trades []*Trade, order *Order) ([]*Trade, *Order) {
	if order.Volume > 0 && order.Price == 0 {
		return trades, order
	}
	return trades, nil
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side {
	case SideAsk:
		return orderbook.ProcessAsk(order)
	case SideBid:
		return orderbook.ProcessBid(order)
	}
	return nil, nil
}

func CheckIfVolumeSatisfiedOrderVolume(order *Order, trade *Trade, side *Order) bool {
	if side.Volume > order.Volume {
		trade.Volume = order.Volume
		side.Volume -= order.Volume
		order.Volume = 0
		return true
	}
	return false
}

func MakeOrderVolume(order *Order, trade *Trade, side *Order) {
	trade.Volume = side.Volume
	order.Volume -= side.Volume
	side.Volume = 0
}

func (orderbook *Orderbook) ProcessAsk(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	for i := 0; i < len(orderbook.Asks); i++ {
		ask := orderbook.Asks[i]
		if ask.Price >= order.Price || order.Price == 0 {
			trade := NewTrade(order, ask, 0, ask.Price)
			if !CheckIfVolumeSatisfiedOrderVolume(order, trade, ask) {
				MakeOrderVolume(order, trade, ask)
				orderbook.Asks = append(orderbook.Asks[:i], orderbook.Asks[i+1:]...)
				i--
			}
			trades = append(trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}
	orderbook.AddBidOrder(order)
	return CheckConditions(trades, order)
}

func (orderbook *Orderbook) ProcessBid(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	for i := 0; i < len(orderbook.Bids); i++ {
		bid := orderbook.Bids[i]
		if order.Price == 0 || bid.Price <= order.Price {
			trade := NewTrade(bid, order, 0, bid.Price)
			if !CheckIfVolumeSatisfiedOrderVolume(order, trade, bid) {
				MakeOrderVolume(order, trade, bid)
				orderbook.Bids = append(orderbook.Bids[:i], orderbook.Bids[i+1:]...)
				i--
			}
			trades = append(trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}

	orderbook.AddAskOrder(order)
	return CheckConditions(trades, order)
}
