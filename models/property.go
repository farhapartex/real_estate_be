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

type Amenities struct {
	AirConditioning  bool    `json:"airConditioning"`
	Heating          bool    `json:"heating"`
	Furnished        bool    `json:"furnished"`
	PetsAllowed      bool    `json:"petsAllowed"`
	WasherDryer      bool    `json:"washerDryer"`
	InUnitLaundry    bool    `json:"inUnitLaundry"`
	Elevator         bool    `json:"elevator"`
	OutdoorSpace     bool    `json:"outdoorSpace"`
	Balcony          bool    `json:"balcony"`
	DisabilityAccess bool    `json:"disabilityAccess"`
	HardwoodFloors   bool    `json:"hardwoodFloors"`
	Fireplace        bool    `json:"fireplace"`
	Pool             bool    `json:"pool"`
	Gym              bool    `json:"gym"`
	Parking          int     `json:"parking"`
	Garages          int     `json:"garages"`
	CeilingHeight    float64 `json:"ceilingHeight"`
	LotSize          float64 `json:"lotSize"`
}

type SecurityFeature struct {
	SecuritySystem bool `json:"securitySystem"`
	Doorman        bool `json:"doorman"`
	SecurityCamera bool `json:"securityCamera"`
	GatedCommunity bool `json:"gatedCommunity"`
	FireAlarm      bool `json:"fireAlarm"`
}

type TechnologyFeature struct {
	InternetWifi    bool `json:"internetWifi"`
	SmartHome       bool `json:"smartHome"`
	Dishwasher      bool `json:"dishwasher"`
	GarbageDisposal bool `json:"garbageDisposal"`
	CableTV         bool `json:"cableTV"`
	Refrigerator    bool `json:"refrigerator"`
	Microwave       bool `json:"microwave"`
	StoveOven       bool `json:"stoveOven"`
	CeilingFans     bool `json:"ceilingFans"`
}

type LuxuryFeature struct {
	Fireplace     bool `json:"fireplace"`
	Pool          bool `json:"pool"`
	Gym           bool `json:"gym"`
	WalkInClosets bool `json:"walkInClosets"`
	Jacuzzi       bool `json:"jacuzzi"`
	Sauna         bool `json:"sauna"`
}

type CommunityFeature struct {
	Concierge       bool `json:"concierge"`
	BusinessCenter  bool `json:"businessCenter"`
	ConferenceRoom  bool `json:"conferenceRoom"`
	GuestParking    bool `json:"guestParking"`
	Playground      bool `json:"playground"`
	BBQArea         bool `json:"bbqArea"`
	CommunityGarden bool `json:"communityGarden"`
	TennisCourt     bool `json:"tennisCourt"`
	BasketballCourt bool `json:"basketballCourt"`
}

type UtilsFeature struct {
	WaterIncluded        bool `json:"waterIncluded"`
	GasIncluded          bool `json:"gasIncluded"`
	ElectricityIncluded  bool `json:"electricityIncluded"`
	TrashRemovalIncluded bool `json:"trashRemovalIncluded"`
	InternetIncluded     bool `json:"internetIncluded"`
}

type EnergyFeature struct {
	SolarPanels               bool `json:"solarPanels"`
	EnergyEfficientAppliances bool `json:"energyEfficientAppliances"`
	GreenCertification        bool `json:"greenCertification"`
	EVCharging                bool `json:"evCharging"`
	RainwaterHarvesting       bool `json:"rainwaterHarvesting"`
	ProgrammableThermostat    bool `json:"programmableThermostat"`
}

type PropertyFeature struct {
	gorm.Model
	PropertyID        uint              `json:"property_id"`
	Property          Property          `gorm:"foreignKey:PropertyID" json:"property"`
	Features          []string          `gorm:"type:text[]" json:"features"`
	Amenities         Amenities         `gorm:"type:jsonb" json:"amenities"`
	SecurityFeature   SecurityFeature   `gorm:"type:jsonb" json:"securityFeature"`
	TechnologyFeature TechnologyFeature `gorm:"type:jsonb" json:"technologyFeature"`
	LuxuryFeature     LuxuryFeature     `gorm:"type:jsonb" json:"luxuryFeature"`
	CommunityFeature  CommunityFeature  `gorm:"type:jsonb" json:"communityFeature"`
	UtilsFeature      UtilsFeature      `gorm:"type:jsonb" json:"utilsFeature"`
	EnergyFeature     EnergyFeature     `gorm:"type:jsonb" json:"energyFeature"`
}
