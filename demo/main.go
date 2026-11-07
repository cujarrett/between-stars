// Package main charts a course between two stars.
package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const warpFactor = 9.2

var ErrNoFuel = errors.New("not enough fuel to jump")

// Star is a point of light, Distance light years from Sol.
type Star struct {
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}

type Ship struct {
	Fuel int
}

// Jump returns how long the ship takes to reach the next star.
func (s *Ship) Jump(from, to Star) (time.Duration, error) {
	gap := math.Abs(to.Distance - from.Distance)
	if s.Fuel < int(gap) {
		return 0, fmt.Errorf("%s -> %s: %w", from.Name, to.Name, ErrNoFuel)
	}
	s.Fuel -= int(gap)
	years := gap / math.Pow(warpFactor, 3)
	return time.Duration(years * 365 * 24 * float64(time.Hour)), nil
}

func main() {
	ship := &Ship{Fuel: 10}
	route := []Star{{"Sol", 0}, {"Proxima Centauri", 4.24}, {"Tau Ceti", 11.9}}

	for i := 1; i < len(route); i++ {
		eta, err := ship.Jump(route[i-1], route[i])
		if err != nil {
			fmt.Println("adrift:", err)
			return
		}
		fmt.Printf("arrived at %s in %v\n", route[i].Name, eta.Round(time.Hour))
	}
}
