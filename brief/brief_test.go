package brief

import (
	"fmt"
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
			bbInput, err := os.ReadFile("tests/simple/" + testCase + ".input")
			if err != nil {
				t.Fatal(err)
			}
			bbOutput, err := os.ReadFile("tests/simple/" + testCase + ".output")
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

func TestOnOpenTofuDataset(t *testing.T) {
	inputs, err := os.ReadDir("tests/opentofu/inputs")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, input := range inputs {
		testCase := input.Name()

		t.Run(testCase, func(t *testing.T) {
			bbInput, err := os.ReadFile("tests/opentofu/inputs/" + testCase + "/plan")
			if err != nil {
				t.Fatal(err)
			}
			bbOutput, err := os.ReadFile("tests/opentofu/outputs/" + testCase + "/plan")
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
