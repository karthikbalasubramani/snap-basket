package testing

type TestCaseResult struct {
	Name   string
	Status string
	Detail string
}

var TestResults []TestCaseResult

func RecordTestResult(name, status, detail string) {
	TestResults = append(TestResults, TestCaseResult{Name: name, Status: status, Detail: detail})
}
