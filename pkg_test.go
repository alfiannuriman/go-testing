package main

import "testing"

func TestReturnValidatePasswordLength_WithLengthLess(t *testing.T) {
	expectation := false
	actual := validatePasswordLength("12345")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnValidatePasswordLength_WithLengthMore(t *testing.T) {
	expectation := true
	actual := validatePasswordLength("1234567")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnValidatePasswordLength_WithLengthEqual(t *testing.T) {
	expectation := true
	actual := validatePasswordLength("123456")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnValidatePasswordCharacter(t *testing.T) {
	type testCase struct {
		description string
		input       string
		expected    bool
	}

	testCases := []testCase{
		{
			description: "With no number",
			input:       "mypassword",
			expected:    false,
		},
		{
			description: "With no uppercase",
			input:       "mypassword",
			expected:    false,
		},
		{
			description: "With no lowercase",
			input:       "MYPASSWORD",
			expected:    false,
		},
		{
			description: "With no special character",
			input:       "myPassword123",
			expected:    false,
		},
		{
			description: "Expected positive return",
			input:       "myPassword123$",
			expected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := validatePasswordCharacter(tc.input)

			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestReturnvalidateInputedPassword_WithLengthNegativeReturn(t *testing.T) {
	expectation := false
	actual := validateInputedPassword("P123$")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnvalidateInputedPassword_WithLengthPositiveReturn(t *testing.T) {
	expectation := false
	actual := validateInputedPassword("MypasswordWithEnoughLength")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnvalidateInputedPassword_WithCharacterNegativeReturn(t *testing.T) {
	expectation := false
	actual := validateInputedPassword("MypasswordWithEnoughLength123")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnvalidateInputedPassword_WithCharacterPositiveReturn(t *testing.T) {
	expectation := false
	actual := validateInputedPassword("Pw1$")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestReturnvalidateInputedPassword_ExpectedPositiveReturn(t *testing.T) {
	expectation := true
	actual := validateInputedPassword("MyPassword123$$")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
