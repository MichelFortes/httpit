package runner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MichelFortes/httpit/pkg/model"
)

type Runner struct {
	client *http.Client
}

func NewRunner() *Runner {
	return &Runner{
		client: &http.Client{},
	}
}

func (r *Runner) Run(scheme *model.TestScheme) error {
	baseURL := fmt.Sprintf("%s://%s:%d", scheme.Protocol, scheme.Host, scheme.Port)

	for _, test := range scheme.Tests {
		fmt.Printf("Running test: %s\n", test.Description)

		url := baseURL + test.Path

		var body *bytes.Buffer
		if test.Payload != "" {
			body = bytes.NewBufferString(test.Payload)
		} else {
			body = bytes.NewBuffer(nil)
		}

		req, err := http.NewRequest(test.Method, url, body)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		for k, v := range test.Headers {
			for _, val := range v {
				req.Header.Add(k, val)
			}
		}

		resp, err := r.client.Do(req)
		if err != nil {
			fmt.Printf("  [FAIL] Request failed: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		// Validation
		passed := true

		// 1. Status Code
		if uint16(resp.StatusCode) != test.ExpectedResult.StatusCode {
			fmt.Printf("  [FAIL] Status Code: got %d, expected %d\n", resp.StatusCode, test.ExpectedResult.StatusCode)
			passed = false
		} else {
			fmt.Printf("  [PASS] Status Code: %d\n", resp.StatusCode)
		}

		// 2. Headers
		for k, expectedVals := range test.ExpectedResult.Headers {
			actualVals, ok := resp.Header[k]
			if !ok {
				fmt.Printf("  [FAIL] Header %s missing\n", k)
				passed = false
				continue
			}
			// Check if at least one value matches (or exact match? keeping simple for now)
			found := false
			for _, ev := range expectedVals {
				for _, av := range actualVals {
					if ev == av {
						found = true
						break
					}
				}
			}
			if !found {
				fmt.Printf("  [FAIL] Header %s: expected one of %v, got %v\n", k, expectedVals, actualVals)
				passed = false
			} else {
				fmt.Printf("  [PASS] Header %s\n", k)
			}
		}

		// 3. Payload
		if test.ExpectedResult.Payload != "" {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("  [FAIL] Failed to read body: %v\n", err)
				passed = false
			} else {
				bodyStr := string(bodyBytes)
				if bodyStr != test.ExpectedResult.Payload {
					fmt.Printf("  [FAIL] Payload mismatch\n")
					// fmt.Printf("   Expected: %s\n", test.ExpectedResult.Payload)
					// fmt.Printf("   Got:      %s\n", bodyStr)
					passed = false
				} else {
					fmt.Printf("  [PASS] Payload\n")
				}
			}
		}

		if !passed {
			// return fmt.Errorf("test failed") // Optional: stop on failure?
		}
		fmt.Println(strings.Repeat("-", 20))
	}
	return nil
}
