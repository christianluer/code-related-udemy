package exercises

import (
	"slices"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("James", "Bond", "vanilla", "mint")
	if p.First != "James" || p.Last != "Bond" || !slices.Equal(p.Flavors, []string{"vanilla", "mint"}) {
		t.Errorf("NewPerson = %+v", p)
	}
}

func TestByLastName(t *testing.T) {
	people := []Person{
		{First: "James", Last: "Bond"},
		{First: "Jenny", Last: "Moneypenny", Flavors: []string{"lemon"}},
	}
	m := ByLastName(people)
	if len(m) != 2 || m["Bond"].First != "James" || m["Moneypenny"].Flavors[0] != "lemon" {
		t.Errorf("ByLastName = %+v", m)
	}
}

func TestNewVehicle(t *testing.T) {
	v := NewVehicle("Aston Martin", "DB5", 2, "silver", false)
	want := Vehicle{Make: "Aston Martin", Model: "DB5", Doors: 2, Color: "silver"}
	if v.Make != want.Make || v.Model != want.Model || v.Doors != want.Doors || v.Color != want.Color || v.Engine.Electric {
		t.Errorf("NewVehicle = %+v, want %+v", v, want)
	}
	if e := NewVehicle("Tesla", "S", 4, "red", true); !e.Engine.Electric {
		t.Errorf("electric vehicle has Engine.Electric = false")
	}
}

func TestIsElectric(t *testing.T) {
	if !IsElectric(Vehicle{Engine: Engine{Electric: true}}) {
		t.Error("IsElectric(electric) = false")
	}
	if IsElectric(Vehicle{}) {
		t.Error("IsElectric(gas) = true")
	}
}

func TestNewProfile(t *testing.T) {
	p := NewProfile("James", "martini", "water")
	if p.First != "James" || !slices.Equal(p.FavDrinks, []string{"martini", "water"}) {
		t.Errorf("NewProfile = %+v", p)
	}
	if p.Friends == nil {
		t.Fatal("Friends is a nil map; writing to it would panic")
	}
	p.Friends["Q"] = 7
}

func TestFullNames(t *testing.T) {
	people := []Person{{First: "James", Last: "Bond"}, {First: "Ian", Last: "Fleming"}}
	if got, want := FullNames(people), []string{"James Bond", "Ian Fleming"}; !slices.Equal(got, want) {
		t.Errorf("FullNames = %v, want %v", got, want)
	}
}
