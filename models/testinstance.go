package models

import "gorm.io/gorm"

// TestInstance Database model info
// @Description App type information
type TestInstance struct {
	*gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	TestID    uint      `gorm:"foreignkey:TestID uniqueIndex:idx_test_instance_test;  OnDelete:SET NULL" json:"test_id,omitempty" swaggertype:"number"`
	TestsetID uint      `gorm:"foreignkey:TestsetID uniqueIndex:idx_test_instance_test; OnDelete:SET NULL" json:"testset_id,omitempty" swaggertype:"number"`
	Severity  string    `gorm:"not null; " json:"severity,omitempty"`
	Issues    []Issue   `gorm:"association_foreignkey:TestInstanceID; constraint:OnUpdate:CASCADE; OnDelete:SET NULL;" json:"issues,omitempty"`
	TestRuns  []TestRun `gorm:"association_foreignkey:TestInstanceID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"test_runs,omitempty"`
}

// @Description App type information
type InstanceTest struct {
	ID             uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name           string `gorm:"not null; " json:"name,omitempty"`
	Description    string `gorm:"not null; " json:"description,omitempty"`
	Steps          string `gorm:"not null; " json:"steps,omitempty"`
	ExpectedResult string `gorm:"not null; " json:"expected_result,omitempty"`
	Severity       string `gorm:"not null; " json:"severity,omitempty"`
}

// TestInstancePost model info
// @Description TestInstancePost type information
type TestInstancePost struct {
	TestID    uint   `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_id,omitempty" swaggertype:"number"`
	TestsetID uint   `gorm:"foreignkey:TestsetID OnDelete:SET NULL" json:"testset_id,omitempty" swaggertype:"number"`
	Severity  string `gorm:"not null; " json:"severity,omitempty"`
}

// TestInstanceGet model info
// @Description TestInstanceGet type information
type TestInstanceGet struct {
	ID       uint    `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Severity string  `gorm:"not null; " json:"severity,omitempty"`
	Issues   []Issue `gorm:"many2many:test_instance_issues; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"issues,omitempty"`
}

// TestInstancePut model info
// @Description TestInstancePut type information
type TestInstancePut struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	TestID    uint `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_id,omitempty" swaggertype:"number"`
	TestsetID uint `gorm:"foreignkey:TestsetID OnDelete:SET NULL" json:"testset_id,omitempty" swaggertype:"number"`
}

// TestInstancePatch model info
// @Description TestInstancePatch type information
type TestInstancePatch struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	TestID    uint `gorm:"foreignkey:TestID OnDelete:SET NULL" json:"test_id,omitempty" swaggertype:"number"`
	TestsetID uint `gorm:"foreignkey:TestsetID OnDelete:SET NULL" json:"testset_id,omitempty" swaggertype:"number"`
}
