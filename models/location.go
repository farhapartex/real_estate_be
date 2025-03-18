package models

import "time"

type Base struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Country struct {
	Base
	Name string `gorm:"size:150;not null" json:"name"`
	Code string `gorm:"size:10;not null" json:"code"`
}

type Division struct {
	Base
	Name      string  `gorm:"size:150;not null" json:"name"`
	CountryId uint32  `gorm:"index" json:"country_id"`
	Country   Country `gorm:"foreignKey:CountryId" json:"country"`
}

type District struct {
	Base
	Name       string   `gorm:"size:150;not null" json:"name"`
	CountryId  uint32   `gorm:"index" json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryId" json:"country"`
	DivisionId uint32   `gorm:"index" json:"division_id"`
	Division   Division `gorm:"foreignKey:DivisionId" json:"division"`
}
