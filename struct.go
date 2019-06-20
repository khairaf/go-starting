package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	//atribut
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple/usixteenbitmax)
}

func main() {
	a_car := car {
							gas_pedal: 16535,
							brake_pedal: 0,
							steering_wheel: 12562,
							top_speed_kmh: 255.0,
	}
	//a_car := car {16535,0,12562,225.0}
	fmt.Println(a_car.gas_pedal, a_car.top_speed_kmh) //cara akses atributnya pake dot
	fmt.Println(a_car.kmh()) //cara panggil method
}

//Methods that just access values are called value receivers and methods that can modify information are pointer receivers.