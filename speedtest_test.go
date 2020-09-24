package speedtest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTest(t *testing.T) {
	var (
		tests = []struct {
			name      string
			forceFail bool
			err       bool
			result    *Results
		}{
			{
				name:   "Valid Test",
				result: &validTestResult,
			},
			{
				name:   "Not found",
				result: &invalidTypeResult,
				err:    true,
			},
			{
				name:      "Force Fail",
				forceFail: true,
				err:       true,
			},
		}
	)

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if err := setupTest(test.forceFail, test.result); err != nil {
				t.Fatalf("Failed to setup path: %v\n", err)
			}

			result, err := NewTest()
			if (err != nil) != test.err {
				t.Fatalf("Expected err to be %v\nGot: %v\n", test.err, err)
			}

			// Since we recieved an error, we don't need to check the results
			// that we got back as they're supposed to be nil and we can
			// pretend like they are.
			if err != nil {
				return
			}

			if test.result != nil {
				if !cmp.Equal(*test.result, result, cmpopts.IgnoreUnexported(Results{})) {
					t.Fatalf("Results are not equal.\nWanted: %#v\nRecieved: %#v\n", test.result, result)
				}
			}
		})
	}
}

func setupTest(forceError bool, result *Results) error {
	// Get our current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get working directory: %w", err)
	}

	// Get the current path structure incase we need other things.
	currentPath := os.Getenv("PATH")

	// Prefix our path so we know our speedtest is being called.
	if err = os.Setenv("PATH", currentDir+":"+currentPath); err != nil {
		return fmt.Errorf("Failed to set env: %w", err)
	}

	// Tells our script to return exit(1).
	if forceError {
		if err = os.Setenv("SPEEDTEST_SCRIPT_ERROR", "true"); err != nil {
			return fmt.Errorf("Failed to set env: %w", err)
		}
	}

	if result != nil {
		byts, err := json.Marshal(result)
		if err != nil {
			return err
		}
		if err = os.Setenv("SPEEDTEST_SCRIPT_VALUE", string(byts)); err != nil {
			return fmt.Errorf("Failed to set env: %w", err)
		}
	}

	return nil
}
