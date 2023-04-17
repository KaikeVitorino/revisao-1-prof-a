package q5

import (
	"github.com/revisao-1/utils"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	testCases := []struct {
		name      string
		temp      float64
		fromScale string
		toScale   string
		expected  float64
		expectErr bool
	}{
		{
			name:      "Test Celsius to Fahrenheit",
			temp:      25,
			fromScale: "C",
			toScale:   "F",
			expected:  77,
			expectErr: false,
		},
		{
			name:      "Test Celsius to Kelvin",
			temp:      30,
			fromScale: "C",
			toScale:   "K",
			expected:  303.15,
			expectErr: false,
		},
		{
			name:      "Test Fahrenheit to Celsius",
			temp:      50,
			fromScale: "F",
			toScale:   "C",
			expected:  10,
			expectErr: false,
		},
		{
			name:      "Test Fahrenheit to Kelvin",
			temp:      68,
			fromScale: "F",
			toScale:   "K",
			expected:  293.15,
			expectErr: false,
		},
		{
			name:      "Test Kelvin to Celsius",
			temp:      350,
			fromScale: "K",
			toScale:   "C",
			expected:  76.85,
			expectErr: false,
		},
		{
			name:      "Test Kelvin to Fahrenheit",
			temp:      400,
			fromScale: "K",
			toScale:   "F",
			expected:  260.33,
			expectErr: false,
		},
		{
			name:      "Test Invalid From Scale",
			temp:      25,
			fromScale: "X",
			toScale:   "F",
			expected:  0,
			expectErr: true,
		},
		{
			name:      "Test Invalid To Scale",
			temp:      25,
			fromScale: "C",
			toScale:   "Y",
			expected:  0,
			expectErr: true,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ConvertTemperature(tc.temp, tc.fromScale, tc.toScale)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !utils.AssertFloatWithPrecision(result, tc.expected, 1e-2) {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				}
			}
		})
	}
}
