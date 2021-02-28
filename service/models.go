package service

// TargetConfig is the info pair of the target url to be tested.
type TargetConfig struct {
	URL                string
	ExpectedStatusCode int
	IgnoreAnalysis     bool
}

// Result is the return msg combination struct for the tester.
type Result struct {
	URL     string
	Msg     string
	Succeed bool
}
