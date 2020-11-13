package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
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
	f, err := os.Open("/opt/ootbat/OOT_LowHealth.wav")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	for {
		var batteryPercentage int
		batteryPercentage = getBatteryPercentage()

		var batteryStatus string
		batteryStatus = getBatteryStatus()

		if batteryPercentage <= 5 && batteryStatus != "Charging" {
			streamer.Seek(0)
			speaker.Play(streamer)
		}

		time.Sleep(10 * time.Second)
	}
}
