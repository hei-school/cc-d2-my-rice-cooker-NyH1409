package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	"strings"
)

type RiceCooker struct {
	isOn          bool
	isPluggedIn   bool
	containsWater bool
	cookedRice    bool
}

func (rc *RiceCooker) turnOnOrOff() {
	rc.isOn = !rc.isOn
}

func (rc *RiceCooker) plugInOrOff() {
	rc.isPluggedIn = !rc.isPluggedIn
}

func (rc *RiceCooker) addWaterOrEmpty() {
	rc.containsWater = !rc.containsWater
}

func (rc *RiceCooker) cookRice() {
	if rc.isOn && rc.isPluggedIn && rc.containsWater {
		fmt.Println("Cooking rice...")
		rc.cookedRice = true
	} else {
		fmt.Println("Cannot cook rice. Make sure the rice cooker is on, plugged in, and contains water.")
	}
}

func (rc *RiceCooker) displayStatus() {
	status := "OFF"
	if rc.isOn {
		status = "ON"
	}
	pluggedInStatus := "Not Plugged In"
	if rc.isPluggedIn {
		pluggedInStatus = "Plugged In"
	}
	waterStatus := "No Water"
	if rc.containsWater {
		waterStatus = "Water Present"
	}
	cookedRiceStatus := "No Cooked Rice"
	if rc.cookedRice {
		cookedRiceStatus = "Cooked Rice Ready"
	}

	fmt.Printf("Status: %s\nPower: %s\nWater: %s\nCooked Rice: %s\n", status, pluggedInStatus, waterStatus, cookedRiceStatus)
}

func main() {
	riceCooker := RiceCooker{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Rice Cooker Emulator:")
		fmt.Println("1. Turn On/Off")
		fmt.Println("2. Plug In/Off")
		fmt.Println("3. Add Water/Empty")
		fmt.Println("4. Cook Rice")
		fmt.Println("5. Display Status")
		fmt.Println("0. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			riceCooker.turnOnOrOff()
		case "2":
			riceCooker.plugInOrOff()
		case "3":
			riceCooker.addWaterOrEmpty()
		case "4":
			riceCooker.cookRice()
		case "5":
			riceCooker.displayStatus()
		case "0":
			fmt.Println("Exiting Rice Cooker Emulator.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
