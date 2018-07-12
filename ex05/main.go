package main

import (
	"./barbershop"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello) It is my barbershop!!!  Have a good time")
	barber := barbershop.NewBarber("Nadinka")
	shop := barbershop.NewShop(barber, 3)

	go barber.ManageShop(shop)
	time.Sleep(time.Second * 3)
	clients := []string{"Kolya", "Vasya", "Maryna", "Katuha"}
	for _, c := range clients {
		client := barbershop.NewClient(c)
		go client.EnterShop(shop)
	}

	fmt.Scanln()
}
