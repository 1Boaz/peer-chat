package Init

import (
	"flag"
	"os"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	// Restores old Args when function is finished
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []string{"8080", "", "-5", "not_a_port", "1000000000"}

	for _, testCase := range cases {
		// Sets the Args to one of the array
		os.Args = []string{os.Args[0], "-port", testCase}

		// Rests CMD flags
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		// Calls the tested function
		port, err := parse()

		// When tests should fail verify correct error message
		switch testCase {
		case "not_a_port", "":
			// If case is invalid and parse did not return an error fail the test
			if err == nil {
				t.Errorf("Parse() should have returned an error for invalid port")
				// If it did return an error skip the comparison test for this case
			} else {
				continue
			}
		// Same for all the cases
		case "-5":
			if err == nil {
				t.Errorf("Parse() should have returned an error for too low port number")
			} else {
				continue
			}
		case "1000000000":
			if err == nil {
				t.Errorf("Parse() should have returned an error for too high port number")
			} else {
				continue
			}
		}

		// Converts the test case to integer, safe because parse should return error with non input integer
		testCaseInt, err := strconv.Atoi(testCase)
		if err != nil {
			t.Errorf("parse(%q) should have failed but did not", testCase)
		}

		// Compares parsed port to the inputted test case
		if port != testCaseInt {
			t.Errorf("Port should be %v got: %v", testCase, port)
		}
	}
}
