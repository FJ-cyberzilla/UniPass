package geo

// GeolocationResult holds the resolved vector data and its source tier
type GeolocationResult struct {
	Latitude  float64
	Longitude float64
	Source    string
	IsAuto    bool
	IsManual  bool
	Error     error
}

// GeolocationProvider defines the interface for all geolocation sources
type GeolocationProvider interface {
	Resolve() (*GeolocationResult, error)
}

// TrustValidator defines the interface for validating geolocation data
type TrustValidator interface {
	Evaluate(locationPayload string) bool
}

// PriorityRule defines the interface for scoring geolocation results
type PriorityRule interface {
	Evaluate(result *GeolocationResult) int
}
