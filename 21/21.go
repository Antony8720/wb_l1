// Реализовать паттерн «адаптер» на любом примере.

// Адаптер - это структурный паттерн проектирования, котоырй позволяет объектам с несовместимыми интерфейсами работать вместе
// Преимущества: отделяет и скрывает от клиента подробности проеобразования различных интерфейсов
// Недостатки: усложняет код программы из-за введения дополнительных классов
package main

import "fmt"

type Client struct{
}

func (c *Client) InsertRuChargerIntoSocket(ch Charger){
	fmt.Println("Client insert Russian charger into socket")
	ch.InsertIntoRuSocket()
}

type Charger interface{
	InsertIntoRuSocket()
}

type RuCharger struct{
}

func(rsck *RuCharger) InsertIntoRuSocket(){
	fmt.Println("Russian charger is plugged into Russian socket")
}

type UsCharger struct{
}

func(ussck *UsCharger)InsertIntoUsSocket() {
	fmt.Println("Us charger is plugged into Us socket")
}

type UsAdapter struct{
	UsaCharger *UsCharger
}

func (usad *UsAdapter) InsertIntoRuSocket() {
	fmt.Println("Adapter converts Ru charger format into Us charger format")
	usad.UsaCharger.InsertIntoUsSocket()
}

func main() {
	client := &Client{}
	ruCharger := &RuCharger{}

	client.InsertRuChargerIntoSocket(ruCharger)

	usCharger := &UsCharger{}
	usAdapter := &UsAdapter{
		UsaCharger: usCharger,
	}

	client.InsertRuChargerIntoSocket(usAdapter)
}