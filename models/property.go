package models

import (
	"time"

	"gorm.io/gorm"
)

type PropertyString string

const (
	PurposeSale PropertyString = "sale"
	PurposeRent PropertyString = "rent"
)

type PropertyStatus string

const (
	StatusActive  PropertyStatus = "active"
	StatusDraft   PropertyStatus = "draft"
	StatusPending PropertyStatus = "pending"
)

type Property struct {
	gorm.Model
	OwnerID      uint           `json:"owner_id"`
	Owner        User           `gorm:"foreignKey:OwnerID" json:"owner"`
	Title        string         `gorm:"type:varchar(255);not null" json:"title"`
	Purpose      PropertyString `gorm:"type:varchar(20);not null" json:"purpose"`
	Price        float64        `gorm:"not null" json:"price"`
	Status       PropertyStatus `gorm:"type:varchar(20);default:draft" json:"status"`
	PropertyType string         `gorm:"type:varchar(50);not null" json:"property_type"`
	Bedrooms     int            `gorm:"not null" json:"bedrooms"`
	Bathrooms    int            `gorm:"not null" json:"bathrooms"`
	Size         float64        `gorm:"not null" json:"size"`
	BuiltYear    int            `json:"built_year"`

	CountryID  uint32   `json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryID" json:"country"`
	DivisionID uint32   `json:"division_id"`
	Division   Division `gorm:"foreignKey:DivisionID" json:"division"`
	DistrictID uint32   `json:"district_id"`
	District   District `gorm:"foreignKey:DistrictID" json:"district"`

	Address string `gorm:"type:varchar(255);not null" json:"address"`

	Description  string     `gorm:"type:text;not null" json:"description"`
	ApprovedAt   *time.Time `json:"approved_at"`
	ApprovedByID *uint      `json:"approved_by_id"`
	ApprovedBy   *User      `gorm:"foreignKey:ApprovedByID" json:"approved_by"`
}
