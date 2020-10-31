package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getBatteryPercentage() int {
	data, _ := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
	percentString := strings.TrimSuffix(string(data), "\n")
	percentInt, _ := strconv.Atoi(percentString)
	return percentInt
}

func getBatteryStatus() string {
	data, _ := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")
	statusString := strings.TrimSuffix(string(data), "\n")
	return statusString
}

func main() {
	for {
		var batteryPercentage int
		batteryPercentage = getBatteryPercentage()

		var batteryStatus string
		batteryStatus = getBatteryStatus()

		if batteryPercentage < 5 && batteryStatus != "Charging" {
			cmd := exec.Command("/usr/bin/aplay", "/opt/ootbat/OOT_LowHealth.wav")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}
