// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateIssueInput struct {
	IssueName        string `json:"issue_name"`
	IssueStatus      string `json:"issue_status"`
	IssueDescription string `json:"issue_description"`
}

type CreateRequirementInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTestInput struct {
	Name           string `json:"name"`
	Steps          string `json:"steps"`
	Expectedresult string `json:"expectedresult"`
}

type CreateTestTestsetInput struct {
	RunStatus string `json:"run_status"`
	Run       string `json:"run"`
	Sevierity string `json:"sevierity"`
}

type CreateTestsetInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Issue struct {
	ID               uint   `json:"id"`
	IssueName        string `json:"issue_name"`
	IssueStatus      string `json:"issue_status"`
	IssueDescription string `json:"issue_description"`
}

type Mutation struct {
}

type Query struct {
}

type Requirement struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Tests       []*Test `json:"tests"`
}

type Test struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Steps          string `json:"steps"`
	ExpectedResult string `json:"expected_result"`
	RequirementID  string `json:"requirement_id"`
}

type TestTestset struct {
	ID        uint      `json:"id"`
	TestID    string    `json:"test_id"`
	TestsetID string    `json:"testset_id"`
	RunStatus RunStatus `json:"run_status"`
	Run       string    `json:"run"`
	Sevierity Severity  `json:"sevierity"`
	Issues    []*Issue  `json:"issues"`
}

type Testset struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Tests       []*Test `json:"tests"`
}

type UpdateIssueInput struct {
	ID               uint   `json:"id"`
	IssueName        string `json:"issue_name"`
	IssueStatus      string `json:"issue_status"`
	IssueDescription string `json:"issue_description"`
}

type UpdateRequirementInput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTestInput struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Steps          string `json:"steps"`
	ExpectedResult string `json:"expected_result"`
}

type UpdateTestTestsetInput struct {
	ID        uint   `json:"id"`
	RunStatus string `json:"run_status"`
	Run       string `json:"run"`
	Sevierity string `json:"sevierity"`
}

type UpdateTestsetInput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RunStatus string

const (
	RunStatusNa         RunStatus = "NA"
	RunStatusNr         RunStatus = "NR"
	RunStatusBlocked    RunStatus = "Blocked"
	RunStatusPassed     RunStatus = "Passed"
	RunStatusFailed     RunStatus = "Failed"
	RunStatusInProgress RunStatus = "InProgress"
)

var AllRunStatus = []RunStatus{
	RunStatusNa,
	RunStatusNr,
	RunStatusBlocked,
	RunStatusPassed,
	RunStatusFailed,
	RunStatusInProgress,
}

func (e RunStatus) IsValid() bool {
	switch e {
	case RunStatusNa, RunStatusNr, RunStatusBlocked, RunStatusPassed, RunStatusFailed, RunStatusInProgress:
		return true
	}
	return false
}

func (e RunStatus) String() string {
	return string(e)
}

func (e *RunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RunStatus", str)
	}
	return nil
}

func (e RunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Severity string

const (
	SeverityCritical Severity = "Critical"
	SeverityHigh     Severity = "High"
	SeverityMedium   Severity = "Medium"
	SeverityLow      Severity = "Low"
)

var AllSeverity = []Severity{
	SeverityCritical,
	SeverityHigh,
	SeverityMedium,
	SeverityLow,
}

func (e Severity) IsValid() bool {
	switch e {
	case SeverityCritical, SeverityHigh, SeverityMedium, SeverityLow:
		return true
	}
	return false
}

func (e Severity) String() string {
	return string(e)
}

func (e *Severity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Severity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Severity", str)
	}
	return nil
}

func (e Severity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
