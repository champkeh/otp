package hotp

import "testing"

func TestGenerateHOTPValue(t *testing.T) {
	cases := []struct {
		counter  uint64
		expected string
	}{
		{counter: 0, expected: "755224"},
		{counter: 1, expected: "287082"},
		{counter: 2, expected: "359152"},
		{counter: 3, expected: "969429"},
		{counter: 4, expected: "338314"},
		{counter: 5, expected: "254676"},
		{counter: 6, expected: "287922"},
		{counter: 7, expected: "162583"},
		{counter: 8, expected: "399871"},
		{counter: 9, expected: "520489"},
	}

	for _, item := range cases {
		value, err := GenerateHOTPValue([]byte("12345678901234567890"), item.counter, Param{Digits: 6})
		if err != nil {
			t.Error(err)
		}
		if value != item.expected {
			t.Fail()
		}
	}
}
