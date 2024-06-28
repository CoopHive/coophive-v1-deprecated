package module

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubst: [subt] can correctly substitute json encoded values into the template string.
func TestSubst(t *testing.T) {
	format := "Hello, %s!"
	inputs := []string{"hiro"}
	expectedOutput := "Hello, hiro!"

	jsonEncodedInputs := make([]any, 0, len(inputs))

	for _, input := range inputs {
		inputJ, err := json.Marshal(input)
		if err != nil {
			t.Fatalf("json marshall failed %v", err)
		}
		jsonEncodedInputs = append(jsonEncodedInputs, JSONEncodedInput(inputJ))
	}
	t.Logf("jsonEncodedInputs -%v %d", jsonEncodedInputs, len(jsonEncodedInputs))

	actualOutput := subt(format, jsonEncodedInputs...)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}
func TestOrGet(t *testing.T) {
	tests := []struct {
		name     string
		inputA   JSONEncodedInput
		inputB   JSONEncodedInput
		expected string
	}{
		{
			name:     "Both inputs empty",
			inputA:   "",
			inputB:   "",
			expected: "",
		},
		{
			name:     "First input empty, second non-empty",
			inputA:   "",
			inputB:   `"inputB"`,
			expected: "inputB",
		},
		{
			name:     "First input non-empty, second empty",
			inputA:   `"inputA"`,
			inputB:   "",
			expected: "inputA",
		},
		{
			name:     "Both inputs non-empty",
			inputA:   `"inputA"`,
			inputB:   `"inputB"`,
			expected: `inputA`,
		},
		{
			name:     "JSON: First input empty, second non-empty",
			inputA:   "",
			inputB:   encodeJson("inputB"),
			expected: "inputB",
		},
		{
			name:     "JSON:First input non-empty, second empty",
			inputA:   encodeJson("inputA"),
			inputB:   "",
			expected: "inputA",
		},
		{
			name:     "JSON: Both inputs non-empty",
			inputA:   encodeJson("inputA"),
			inputB:   encodeJson("inputB"),
			expected: "inputA",
		},
	}

	funcAliases := []func(a, b any) string{
		or,
		get,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// actual := or(tt.inputA, tt.inputB)
			// assert.Equal(t, tt.expected, actual)

			for _, f1 := range funcAliases {
				actual := f1(tt.inputA, tt.inputB)
				assert.Equal(t, tt.expected, actual)
			}

		})
	}
}
