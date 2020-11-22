package structs

import (
	"testing"
	"time"
)

var EwciaBday string = "1991-02-11T00:00:00Z"
var DarouBday string = "1986-08-22T16:30:00Z"

func TestParseDate(t *testing.T) {

	var b Birthday

	b.SetBirthday(EwciaBday)

	parsedDate, _ := time.Parse(
		time.RFC3339,
		EwciaBday)

	if !b.Born.Equal(parsedDate) {
		t.Errorf("Date was incorrectly parsed: %v, want: %v.", b.GetBirthday(), parsedDate)
	}
}
