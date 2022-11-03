package factoryMethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	ak47, err := CreateGun(AK_47)
	if err != nil {
		t.Errorf("CreateGun: %s", err)
	}
	if ak47.GetName() != "Ak-47" || ak47.GetPower() != 3 {
		t.Errorf("Wrong weapon name or weapon power")
	}
	if !ak47.Fire() {
		t.Errorf("Gun cannot fire")
	}

	ak74, err := CreateGun(AK_74)
	if err != nil {
		t.Errorf("CreateGun: %s", err)
	}
	if ak74.GetName() != "Ak-74" || ak74.GetPower() != 7 {
		t.Errorf("Wrong weapon name or weapon power")
	}
	if !ak74.Fire() {
		t.Errorf("Gun cannot fire")
	}

	ak103, err := CreateGun(AK_103)
	if err != nil {
		t.Errorf("CreateGun: %s", err)
	}
	if ak103.GetName() != "Ak-103" || ak103.GetPower() != 5 {
		t.Errorf("Wrong weapon name or weapon power")
	}
	if !ak103.Fire() {
		t.Errorf("Gun cannot fire")
	}
}
