package models

import "gorm.io/gorm"

// Document Database model info
// @Description App type information
type Document struct {
	*gorm.Model
	ID            uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name          string `gorm:"not null; " json:"issue_name,omitempty"`
	File_URL      string `gorm:"not null; " json:"issue_status,omitempty"`
	RequirementID uint   `gorm:"foreignkey:RequirementID OnDelete:SET NULL" json:"requirement_id,omitempty" swaggertype:"number"`
	SprintID      uint   `gorm:"foreignkey:SprintID OnDelete:SET NULL" json:"sprint_id,omitempty" swaggertype:"number"`
	TestRunID     uint   `gorm:"foreignkey:TestRunID OnDelete:SET NULL" json:"test_run_id,omitempty" swaggertype:"number"`
}

// DocumentPost model info
// @Description DocumentPost type information
type DocumentPost struct {
	Name          string `gorm:"not null; " json:"issue_name,omitempty"`
	File_URL      string `gorm:"not null; " json:"issue_status,omitempty"`
	RequirementID uint   `gorm:"foreignkey:RequirementID OnDelete:SET NULL" json:"requirement_id,omitempty" swaggertype:"number"`
	SprintID      uint   `gorm:"foreignkey:SprintID OnDelete:SET NULL" json:"sprint_id,omitempty" swaggertype:"number"`
	TestRunID     uint   `gorm:"foreignkey:TestRunID OnDelete:SET NULL" json:"test_run_id,omitempty" swaggertype:"number"`
}

// DocumentGet model info
// @Description DocumentGet type information
type DocumentGet struct {
	ID            uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name          string `gorm:"not null; " json:"issue_name,omitempty"`
	File_URL      string `gorm:"not null; " json:"issue_status,omitempty"`
	RequirementID uint   `gorm:"foreignkey:RequirementID OnDelete:SET NULL" json:"requirement_id,omitempty" swaggertype:"number"`
	SprintID      uint   `gorm:"foreignkey:SprintID OnDelete:SET NULL" json:"sprint_id,omitempty" swaggertype:"number"`
	TestRunID     uint   `gorm:"foreignkey:TestRunID OnDelete:SET NULL" json:"test_run_id,omitempty" swaggertype:"number"`
}

// DocumentPut model info
// @Description DocumentPut type information
type DocumentPut struct {
	ID       uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name     string `gorm:"not null; " json:"issue_name,omitempty"`
	File_URL string `gorm:"not null; " json:"issue_status,omitempty"`
}

// DocumentPatch model info
// @Description DocumentPatch type information
type DocumentPatch struct {
	ID       uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name     string `gorm:"not null; " json:"issue_name,omitempty"`
	File_URL string `gorm:"not null; " json:"issue_status,omitempty"`
}
