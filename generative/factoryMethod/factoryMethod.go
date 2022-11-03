package factoryMethod

import (
	"Design-Patterns/pkg/logger"
	"fmt"
)

type gunType string

const (
	AK_47  gunType = "Ak-47"
	AK_74  gunType = "Ak-74"
	AK_103 gunType = "Ak-103"
)

type IGun interface {
	Fire() bool
	PrintDetails()
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) Fire() bool {
	logger.Trace("%s shoots", g.name)
	return true
}

func (g *Gun) PrintDetails() {
	logger.Trace("\nGun:\nName - %s\nPower - %d", g.name, g.power)
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "Ak-47",
			power: 3,
		},
	}
}

type Ak74 struct {
	Gun
}

func newAk74() IGun {
	return &Ak74{
		Gun{
			name:  "Ak-74",
			power: 7,
		},
	}
}

type Ak103 struct {
	Gun
}

func newAk103() IGun {
	return &Ak103{
		Gun{
			name:  "Ak-103",
			power: 5,
		},
	}
}

func CreateGun(gunType gunType) (IGun, error) {
	if gunType == "Ak-47" {
		return newAk47(), nil
	}
	if gunType == "Ak-74" {
		return newAk74(), nil
	}
	if gunType == "Ak-103" {
		return newAk103(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
