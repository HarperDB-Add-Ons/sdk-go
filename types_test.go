package harperdb

import "testing"

type Player struct {
	Record        // embed the Record type in your custom types
	Name   string `json:"name"`
}

func ExampleRecord() Player {
	return Player{
		Name: "Jane",
		Record: Record{
			UpdatedTime: 123123.23324,
		},
	}
}

func TestRecordFloatTime(t *testing.T) {
	str := ExampleRecord()

	// Will panic if there's an issue.
	str.UpdatedTime.ToTime()
	str.CreatedTime.ToTime()
}
