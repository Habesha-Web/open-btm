package models

import (
	"gorm.io/gorm"
)

// TestTestset Database model info
// @Description App type information
type TestTestset struct {
	*gorm.Model
	ID        uint    `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	TestID    string  `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_id,omitempty" swaggertype:"number"`
	TestsetID string  `gorm:"foreignkey:TestsetID OnDelete:SET NULL" json:"testset_id,omitempty" swaggertype:"number"`
	RunStatus string  `gorm:"not null; unique;" json:"run_status,omitempty"`
	Run       string  `gorm:"not null; unique;" json:"run,omitempty"`
	Sevierity string  `gorm:"not null; unique;" json:"sevierity,omitempty"`
	Issues    []Issue `gorm:"many2many:test_testset_issues; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"issues,omitempty"`
}

// TestTestsetPost model info
// @Description TestTestsetPost type information
type TestTestsetPost struct {
	RunStatus string `gorm:"not null; unique;" json:"run_status,omitempty"`
	Run       string `gorm:"not null; unique;" json:"run,omitempty"`
	Sevierity string `gorm:"not null; unique;" json:"sevierity,omitempty"`
}

// TestTestsetGet model info
// @Description TestTestsetGet type information
type TestTestsetGet struct {
	ID uint `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`

	RunStatus string  `gorm:"not null; unique;" json:"run_status,omitempty"`
	Run       string  `gorm:"not null; unique;" json:"run,omitempty"`
	Sevierity string  `gorm:"not null; unique;" json:"sevierity,omitempty"`
	Issues    []Issue `gorm:"many2many:test_testset_issues; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"issues,omitempty"`
}

// TestTestsetPut model info
// @Description TestTestsetPut type information
type TestTestsetPut struct {
	ID uint `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`

	RunStatus string `gorm:"not null; unique;" json:"run_status,omitempty"`
	Run       string `gorm:"not null; unique;" json:"run,omitempty"`
	Sevierity string `gorm:"not null; unique;" json:"sevierity,omitempty"`
}

// TestTestsetPatch model info
// @Description TestTestsetPatch type information
type TestTestsetPatch struct {
	ID uint `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`

	RunStatus string `gorm:"not null; unique;" json:"run_status,omitempty"`
	Run       string `gorm:"not null; unique;" json:"run,omitempty"`
	Sevierity string `gorm:"not null; unique;" json:"sevierity,omitempty"`
}
