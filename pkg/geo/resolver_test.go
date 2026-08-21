package geo

import (
	"errors"
	"testing"
)

// --- Mocks ---

type MockProvider struct {
	result *GeolocationResult
	err    error
}

func (p *MockProvider) Resolve() (*GeolocationResult, error) {
	return p.result, p.err
}

type MockValidator struct {
	shouldPass bool
}

func (v *MockValidator) Evaluate(payload string) bool {
	return v.shouldPass
}

type MockRule struct {
	priorities map[string]int
}

func (r *MockRule) Evaluate(result *GeolocationResult) int {
	if p, ok := r.priorities[result.Source]; ok {
		return p
	}
	return 0
}

// --- Tests ---

func TestResolveBestVector(t *testing.T) {
	tests := []struct {
		name       string
		providers  []GeolocationProvider
		validator  TrustValidator
		rules      []PriorityRule
		wantSource string
		wantErr    bool
	}{
		{
			name: "Happy path - Highest priority wins",
			providers: []GeolocationProvider{
				&MockProvider{result: &GeolocationResult{Source: "P1"}, err: nil},
				&MockProvider{result: &GeolocationResult{Source: "P2"}, err: nil},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{"P1": 1, "P2": 2}},
			},
			wantSource: "P2",
			wantErr:    false,
		},
		{
			name: "Fallback - First provider fails, second succeeds",
			providers: []GeolocationProvider{
				&MockProvider{result: nil, err: errors.New("fail")},
				&MockProvider{result: &GeolocationResult{Source: "P2"}, err: nil},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{"P2": 1}},
			},
			wantSource: "P2",
			wantErr:    false,
		},
		{
			name: "All fail - Should return error",
			providers: []GeolocationProvider{
				&MockProvider{result: nil, err: errors.New("fail1")},
				&MockProvider{result: nil, err: errors.New("fail2")},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{}},
			},
			wantSource: "",
			wantErr:    true,
		},
		{
			name: "Auto mode preference - Higher priority rule for auto",
			providers: []GeolocationProvider{
				&MockProvider{result: &GeolocationResult{Source: "Manual", IsManual: true}, err: nil},
				&MockProvider{result: &GeolocationResult{Source: "Auto", IsAuto: true}, err: nil},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{"Manual": 1, "Auto": 10}},
			},
			wantSource: "Auto",
			wantErr:    false,
		},
		{
			name: "Manual mode preference - Higher priority rule for manual",
			providers: []GeolocationProvider{
				&MockProvider{result: &GeolocationResult{Source: "Manual", IsManual: true}, err: nil},
				&MockProvider{result: &GeolocationResult{Source: "Auto", IsAuto: true}, err: nil},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{"Manual": 10, "Auto": 1}},
			},
			wantSource: "Manual",
			wantErr:    false,
		},
		{
			name: "Edge Case - Provider returns nil",
			providers: []GeolocationProvider{
				&MockProvider{result: nil, err: nil},
			},
			validator: &MockValidator{shouldPass: true},
			rules: []PriorityRule{
				&MockRule{priorities: map[string]int{}},
			},
			wantSource: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := NewResolver(tt.providers, tt.validator, tt.rules)
			got, err := resolver.ResolveBestVector()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveBestVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Source != tt.wantSource {
				t.Errorf("ResolveBestVector() gotSource = %v, want %v", got.Source, tt.wantSource)
			}
		})
	}
}
