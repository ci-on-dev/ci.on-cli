package assert

type Result struct {
	Suite      string
	Test       string
	Passed     bool
	Assertions []AssertionResult
}

type Results struct {
	Passed  bool
	Results []Result
}
