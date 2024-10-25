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
	testCases := []string{
		"basic_json_string_update",
		"basic_list",
		"basic_list_empty",
		"basic_list_null",
		"basic_map",
		"basic_map_empty",
		"basic_map_null",
		"basic_map_update",
		"basic_multiline_string_update",
		"basic_set",
		"basic_set_empty",
		"basic_set_null",
		"basic_set_update",
	}

	for _, testCase := range testCases {
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
