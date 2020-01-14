package timeutil

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		format   string
		value    string
		expected time.Time
	}{
		{"dd/MM/YYYY HH:mm:ss.SSSSSSSSS P", "03/02/2007 23:10:05.000000004 P",
			time.Date(2007, time.February, 3, 23, 10, 5, 4, time.UTC)},

		{"dd/MMMM/yyyy:HH:mm:ss Z", "30/August/2015:21:44:25 -0500",
			time.Date(2015, time.August, 30, 21, 44, 25, 0, time.FixedZone("", -5*60*60))},

		{"dd/MMMM/yyyy:hh:m:s a Z P", "30/August/2015:03:4:5 PM -0500 P",
			time.Date(2015, time.August, 30, 15, 4, 5, 0, time.FixedZone("", -5*60*60))},

		{"dd/MMMM/yyyy:hh:m:s a Z", "30/August/2015:03:4:25 PM -0500",
			time.Date(2015, time.August, 30, 15, 4, 25, 0, time.FixedZone("", -5*60*60))},

		{"YYYY-MM-dd HH:mm:ss.SSS", "2012-12-22 12:53:30.000",
			time.Date(2012, time.December, 22, 12, 53, 30, 0, time.UTC)},
		{"E d-MMMM-YY HH:mm:ss.SSS", "Mon 1-may-17 12:53:30.000",
			time.Date(2017, time.May, 1, 12, 53, 30, 0, time.UTC)},
		{"[EEE MMM dd HH:mm:ss y]", "[Sun Jan 11 10:43:35 2015]",
			time.Date(2015, time.January, 11, 10, 43, 35, 0, time.UTC)},
		{"[EEEE MMM dd HH:mm:ss y]", "[Sunday Jan 11 10:43:35 2015]",
			time.Date(2015, time.January, 11, 10, 43, 35, 0, time.UTC)},
		{"[EEEE M dd h:mm:ss y]", "[Sunday 1 11 9:43:35 2015]",
			time.Date(2015, time.January, 11, 9, 43, 35, 0, time.UTC)},

		{"dd/MMMM/yyyy:hh:m:s a ZZ", "30/August/2015:03:4:25 PM -05:00",
			time.Date(2015, time.August, 30, 15, 4, 25, 0, time.FixedZone("", -5*60*60))},

		{"YYYY-MM-dd''HH:mm:ss", "2017-02-18'16:33:21",
			time.Date(2017, time.February, 18, 16, 33, 21, 0, time.UTC)},

		{"YYYY-MM-dd'T'HH:mm:ss", "2017-02-18T16:33:21",
			time.Date(2017, time.February, 18, 16, 33, 21, 0, time.UTC)},

		{"YYYY-MM-dd'T'HH:mm:ss'Z'", "2017-02-18T16:33:21Z",
			time.Date(2017, time.February, 18, 16, 33, 21, 0, time.UTC)},

		{"YYYY-MM-dd HH:mm:ss.SSS", "2012-12-22 12:53:30.123",
			time.Date(2012, time.December, 22, 12, 53, 30, 123*1000000, time.UTC)},
		{"YYYY-MM-dd HH:mm:ss.SS", "2012-12-22 12:53:30.12",
			time.Date(2012, time.December, 22, 12, 53, 30, 12*10000000, time.UTC)},
		{"YYYY-MM-dd HH:mm:ss.S", "2012-12-22 12:53:30.1",
			time.Date(2012, time.December, 22, 12, 53, 30, 1*100000000, time.UTC)},
	}

	for _, test := range tests {
		rTime, err := ParseTime(test.format, test.value)
		if err != nil {
			t.Error(err)
		}
		if !test.expected.Equal(rTime) {
			t.Errorf("expected: %q, got : %q", test.expected, rTime)
		}
	}
}
