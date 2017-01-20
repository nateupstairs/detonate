package detonate

import (
	"encoding/json"
	"testing"

	"github.com/nateupstairs/detonate"
)

func TestCreation(t *testing.T) {
	d := detonate.Create(500, "Test Error", "I'm testing!")

	if d.Code != 500 || d.Error != "Test Error" || d.Message != "I'm testing!" {
		t.Fail()
	}
}

func TestToJSON(t *testing.T) {
	var dd = new(Detonation)

	d := detonate.Create(500, "Test Error", "I'm testing!")
	j, err := d.ToJSON()
	if err != nil {
		t.Fail()
	}
	err = json.Unmarshal(j, &dd)
	if err != nil {
		t.Fail()
	}
	if d.Code != 500 || d.Error != "Test Error" || d.Message != "I'm testing!" {
		t.Fail()
	}
}
