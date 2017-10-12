package simple_entity

import "testing"

func TestSetAndGetName(t *testing.T) {
	var u User;
	u.setName("Matthew");
	if u.getName() != "Matthew" {
		t.Error("Expected Matthew, got ", u.getName())
	}
}