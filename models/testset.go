package models

import "gorm.io/gorm"

// Testset Database model info
// @Description App type information
type Testset struct {
	*gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; " json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
	Tests       []Test `gorm:"many2many:test_testsets; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"tests,omitempty"`
}

// TestsetPost model info
// @Description TestsetPost type information
type TestsetPost struct {
	Name        string `gorm:"not null; " json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}

// TestsetGet model info
// @Description TestsetGet type information
type TestsetGet struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; " json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
	Tests       []Test `gorm:"many2many:test_testsets; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"tests,omitempty"`
}

// TestsetPut model info
// @Description TestsetPut type information
type TestsetPut struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; " json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}

// TestsetPatch model info
// @Description TestsetPatch type information
type TestsetPatch struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `gorm:"not null; " json:"name,omitempty"`
	Description string `gorm:"not null; " json:"description,omitempty"`
}
