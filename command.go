package main

import "fmt"

type command interface {
	execute()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type device interface {
	on()
	off()
}

type onCommand struct {
	device device
}

func (o *onCommand) execute() {
	o.device.on()
}

type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	o.device.off()
}

type tv struct{}

func (t *tv) on() {
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	fmt.Println("Turning tv off")
}

type airConditioner struct{}

func (a *airConditioner) on() {
	fmt.Println("Turning airConditioner on")
}

func (a *airConditioner) off() {
	fmt.Println("Turning airConditioner off")
}

func Tv() {
	tv := &tv{}
	onTvCommand := &onCommand{device: tv}

	offTvCommand := &offCommand{device: tv}

	onTvButton := &button{command: onTvCommand}
	onTvButton.press()

	offTvButton := &button{command: offTvCommand}
	offTvButton.press()
}

func AirConditioner() {
	airConditioner := &airConditioner{}
	onAirConditionerCommand := &onCommand{device: airConditioner}

	offAirConditionerCommand := &offCommand{device: airConditioner}

	onAirConditionerButton := &button{command: onAirConditionerCommand}
	onAirConditionerButton.press()

	offAirConditionerButton := &button{command: offAirConditionerCommand}
	offAirConditionerButton.press()

}

func ExampleCommand() {
	fmt.Println("----------命令模式 Command Start----------")
	Tv()

	AirConditioner()

	fmt.Println("----------命令模式 Command End----------")
}
