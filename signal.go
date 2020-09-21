package wireless

import (
	"strconv"
	"strings"
)

// Signal represents the current signal state
type Signal struct {
	RSSI            int    `json:"rssi"`
	LinkSpeed       string `json:"link_speed"`
	Noise           string `json:"noise"`
	Frequency       string `json:"frequency"`
	Width           string `json:"width"`
	CenterFrequency string `json:"center_frequency"`
}

// NewSignal will return the state of the signal when given the raw output
func NewSignal(data string) Signal {
	s := Signal{}
	for _, l := range strings.Split(data, "\n") {
		bits := strings.Split(strings.TrimSpace(l), "=")
		if len(bits) < 2 {
			continue
		}

		switch bits[0] {
		case "RSSI":
			rssi, _ := strconv.Atoi(bits[1])
			s.RSSI = rssi
		case "LINKSPEED":
			s.LinkSpeed = bits[1]
		case "NOISE":
			s.Noise = bits[1]
		case "FREQUENCY":
			s.Frequency = bits[1]
		case "WIDTH":
			s.Width = bits[1]
		case "CETNER_FRQ1":
			s.CenterFrequency = bits[1]
		}
	}

	return s
}
