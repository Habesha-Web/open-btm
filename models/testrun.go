package models

import (
	"time"

	"gorm.io/gorm"
)

// TestRun Database model info
// @Description App type information
type TestRun struct {
	*gorm.Model
	ID             uint       `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	TestInstanceID uint       `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_instance_id,omitempty" swaggertype:"number"`
	RunStatus      string     `gorm:"not null; " json:"run_status,omitempty"`
	Result         string     `gorm:"not null; " json:"run,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty,string"`
	CreatedAt      time.Time  `json:"created_at,omitempty,string"`
	Documents      []Document `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// TestRunPost model info
// @Description TestRunPost type information
type TestRunPost struct {
	TestInstanceID uint   `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_instance_id,omitempty" swaggertype:"number"`
	RunStatus      string `gorm:"not null; " json:"run_status,omitempty"`
	Result         string `gorm:"not null; " json:"run,omitempty"`
}

// TestRunGet model info
// @Description TestRunGet type information
type TestRunGet struct {
	ID        uint       `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	RunStatus string     `gorm:"not null; " json:"run_status,omitempty"`
	Result    string     `gorm:"not null; " json:"run,omitempty"`
	Documents []Document `gorm:"association_foreignkey:RequirementID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"documents,omitempty"`
}

// TestRunPut model info
// @Description TestRunPut type information
type TestRunPut struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	RunStatus string `gorm:"not null; " json:"run_status,omitempty"`
	Result    string `gorm:"not null; " json:"run,omitempty"`
}

// TestRunPatch model info
// @Description TestRunPatch type information
type TestRunPatch struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	RunStatus string `gorm:"not null; " json:"run_status,omitempty"`
	Result    string `gorm:"not null; " json:"run,omitempty"`
}
