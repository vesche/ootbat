package main

import (
	"fmt"
	"time"
	"io/ioutil"
    "os/exec"
    "strconv"
    "strings"
)

func getBatteryPercentage() int {
    data, _ := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
    percentString := strings.TrimSuffix(string(data), "\n")
    percentInt, _ := strconv.Atoi(percentString)
    return percentInt
}

func main() {
    for {
        var batteryPercentage int
        batteryPercentage = getBatteryPercentage()

        if batteryPercentage < 5 {
            cmd := exec.Command("/usr/bin/aplay", "/opt/ootbat/OOT_LowHealth.wav")
            err := cmd.Run()
            if err != nil {
                fmt.Println(err)
            }
        }

        time.Sleep(10 * time.Second)
    }
}