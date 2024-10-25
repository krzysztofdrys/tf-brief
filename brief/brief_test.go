package brief

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
)

func TestSimpleScenarios(t *testing.T) {
	testCases := []string{
		"simple",
		"nested",
		"sensitive_value",
	}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			bbInput, err := os.ReadFile("tests/" + testCase + ".input")
			if err != nil {
				t.Fatal(err)
			}
			bbOutput, err := os.ReadFile("tests/" + testCase + ".output")
			if err != nil {
				t.Fatal(err)
			}

			input := string(bbInput)
			output := strings.Split(string(bbOutput), "\n")

			result := Plan(strings.Split(input, "\n"))

			if diff := cmp.Diff(output, result); diff != "" {
				t.Errorf("Plan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
