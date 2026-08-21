package geo

import (
	"errors"
	"fmt"
)

// Resolver executes the cascade based on configured providers and validator
type Resolver struct {
	providers []GeolocationProvider
	validator TrustValidator
	rules     []PriorityRule
}

// NewResolver creates a new Resolver with injected dependencies
func NewResolver(providers []GeolocationProvider, validator TrustValidator, rules []PriorityRule) *Resolver {
	return &Resolver{
		providers: providers,
		validator: validator,
		rules:     rules,
	}
}

// ResolveBestVector executes the cascade based on configured priorities and settings
func (r *Resolver) ResolveBestVector() (*GeolocationResult, error) {
	var bestResult *GeolocationResult
	highestPriority := -1

	for _, provider := range r.providers {
		result, err := provider.Resolve()
		if err != nil || result == nil {
			continue
		}

		// If it's an IP source, validate it
		if result.Source == "ip" {
			ipStr := fmt.Sprintf("%f,%f", result.Latitude, result.Longitude)
			if !r.validator.Evaluate(ipStr) {
				continue
			}
		}

		// Evaluate rules
		currentPriority := -1
		for _, rule := range r.rules {
			if p := rule.Evaluate(result); p > currentPriority {
				currentPriority = p
			}
		}

		if currentPriority > highestPriority {
			highestPriority = currentPriority
			bestResult = result
		}
	}

	if bestResult != nil {
		return bestResult, nil
	}

	return nil, errors.New("resolver: all priority cascade tiers exhausted or disabled")
}
