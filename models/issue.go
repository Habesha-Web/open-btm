package models

import (
	"gorm.io/gorm"
)

// Issue Database model info
// @Description App type information
type Issue struct {
	*gorm.Model
	ID               uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	IssueName        string `gorm:"not null; unique;" json:"issue_name,omitempty"`
	IssueStatus      string `gorm:"not null; unique;" json:"issue_status,omitempty"`
	IssueDescription string `gorm:"not null; unique;" json:"issue_description,omitempty"`
}

// IssuePost model info
// @Description IssuePost type information
type IssuePost struct {
	IssueName        string `gorm:"not null; unique;" json:"issue_name,omitempty"`
	IssueStatus      string `gorm:"not null; unique;" json:"issue_status,omitempty"`
	IssueDescription string `gorm:"not null; unique;" json:"issue_description,omitempty"`
}

// IssueGet model info
// @Description IssueGet type information
type IssueGet struct {
	ID               uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	IssueName        string `gorm:"not null; unique;" json:"issue_name,omitempty"`
	IssueStatus      string `gorm:"not null; unique;" json:"issue_status,omitempty"`
	IssueDescription string `gorm:"not null; unique;" json:"issue_description,omitempty"`
}

// IssuePut model info
// @Description IssuePut type information
type IssuePut struct {
	ID               uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	IssueName        string `gorm:"not null; unique;" json:"issue_name,omitempty"`
	IssueStatus      string `gorm:"not null; unique;" json:"issue_status,omitempty"`
	IssueDescription string `gorm:"not null; unique;" json:"issue_description,omitempty"`
}

// IssuePatch model info
// @Description IssuePatch type information
type IssuePatch struct {
	ID               uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	IssueName        string `gorm:"not null; unique;" json:"issue_name,omitempty"`
	IssueStatus      string `gorm:"not null; unique;" json:"issue_status,omitempty"`
	IssueDescription string `gorm:"not null; unique;" json:"issue_description,omitempty"`
}
