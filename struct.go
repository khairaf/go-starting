package main

import "fmt"

const usixteenbitmax float64 = 2
const kmh_multiple float64 = 0.5

type car struct {
	//atribut
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

//gets a copy, receiver type
func (c car) kmh() float64 {
	//c.top_speed_kmh = 30
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple/usixteenbitmax)
}

//modify the struct itself via pointer type
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car {
							gas_pedal: 15,
							brake_pedal: 0,
							steering_wheel: 12562,
							top_speed_kmh: 10.0,
	}
	//a_car := car {16535,0,12562,225.0}
	fmt.Println(a_car.gas_pedal, a_car.top_speed_kmh) //15, 10 cara akses atributnya pake dot
	fmt.Println(a_car.kmh()) //150 cara panggil method

	//pointer
	a_car.new_top_speed(20)
	fmt.Println(a_car.kmh()) //300
	fmt.Println(a_car.top_speed_kmh) //20
}

//Methods that just access values are called value receivers and methods that can modify information are pointer receivers.