package client

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

var environments = map[string]string{
	"test": "https://api.test.fincode.jp",
	"live": "https://api.fincode.jp",
}

type Spec struct {
	APIKey      string `json:"api_key"`
	Environment string `json:"environment"`
}

func (s *Spec) SetDefaults() {
	if s.Environment == "" {
		s.Environment = "test"
	}
}

func (s *Spec) Validate() error {
	errors := make([]string, 0)

	if s.APIKey == "" {
		errors = append(errors, `"api_key" is required`)
	}

	if _, ok := environments[s.Environment]; !ok {
		validValues := maps.Keys(environments)
		sort.Strings(validValues)
		errors = append(errors, fmt.Sprintf(`invalid "environment". Expected one of %q, got %q`, validValues, s.Environment))
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid plugin spec: %s", strings.Join(errors, ". "))
	}

	return nil
}
