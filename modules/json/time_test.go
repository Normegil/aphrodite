package json_test

import (
	"testing"
	"time"
	"github.com/normegil/aphrodite/modules/json"
	"github.com/normegil/aphrodite/modules/test"
)

func TestTimeToString(t *testing.T) {
	cases := []time.Time{
		time.Now(),
		time.Now().Add(48 * time.Hour),
	}
	for _, testTime := range cases {
		value := json.JSONTime(testTime).String()
		expected := testTime.Format(time.RFC3339)
		if expected != value {
			t.Error(test.Format("JSONTime.String()", "Returned message don't correspond to expectd output", expected, value))
		}
	}
}

func TestTimeMarshallJSON(t *testing.T) {
	cases := []struct {
		testName string
		input    time.Time
	}{
		{"JSON - Empty field", time.Now(),},
		{"JSON - Classic case", "Test", "\"Test\""},
	}
	for _, params := range cases {
		bytes, err := ErrorJSON{errors.New(params.input)}.MarshalJSON()
		if nil != err {
			t.Fatal(err.Error())
		}
		if params.output != string(bytes) {
			t.Error(test.Format(params.testName, "Malformed JSON", params.output, params.input))
		}
	}
}
