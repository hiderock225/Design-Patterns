package factoryMethod

import (
	"Design-Patterns/pkg/logger"
	"fmt"
)

type GunType string

const (
	Ak47  GunType = "Ak47"
	Ak74  GunType = "Ak74"
	Ak103 GunType = "Ak103"
)

type IGun interface {
	Fire()
	PrintDetails()
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) Fire() {
	logger.Trace("%s shoots", g.name)
}

func (g *Gun) PrintDetails() {
	logger.Trace("\nGun:\nName - %s\nPower - %d", g.name, g.power)
}

func CreateGun(gunType GunType) (IGun, error) {
	switch gunType {
	case Ak47:
		return &Gun{
			name:  "Ak-47",
			power: 3,
		}, nil
	case Ak74:
		return &Gun{
			name:  "Ak-74",
			power: 7,
		}, nil
	case Ak103:
		return &Gun{
			name:  "Ak-103",
			power: 5,
		}, nil
	default:
		return nil, fmt.Errorf("unknown gun")
	}
}
