package models

import (
	"gorm.io/gorm"
)

// Test Database model info
// @Description App type information
type Test struct {
	*gorm.Model
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; unique;" json:"name,omitempty"`
	Steps          string `gorm:"not null; unique;" json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; unique;" json:"expected_result,omitempty"`
	RequirementID  string `gorm:"foreignkey:RequirementID OnDelete:SET NULL" json:"requirement_id,omitempty" swaggertype:"number"`
}

// TestPost model info
// @Description TestPost type information
type TestPost struct {
	Name           string `gorm:"not null; unique;" json:"name,omitempty"`
	Steps          string `gorm:"not null; unique;" json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; unique;" json:"expected_result,omitempty"`
}

// TestGet model info
// @Description TestGet type information
type TestGet struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; unique;" json:"name,omitempty"`
	Steps          string `gorm:"not null; unique;" json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; unique;" json:"expected_result,omitempty"`
}

// TestPut model info
// @Description TestPut type information
type TestPut struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; unique;" json:"name,omitempty"`
	Steps          string `gorm:"not null; unique;" json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; unique;" json:"expected_result,omitempty"`
}

// TestPatch model info
// @Description TestPatch type information
type TestPatch struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; unique;" json:"name,omitempty"`
	Steps          string `gorm:"not null; unique;" json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; unique;" json:"expected_result,omitempty"`
}
