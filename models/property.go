package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
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
	PropertyID        uint           `json:"property_id"`
	Property          Property       `gorm:"foreignKey:PropertyID" json:"property"`
	Features          pq.StringArray `gorm:"type:text[]" json:"features"`
	Amenities         []byte         `gorm:"type:jsonb" json:"-"`
	SecurityFeature   []byte         `gorm:"type:jsonb" json:"-"`
	TechnologyFeature []byte         `gorm:"type:jsonb" json:"-"`
	LuxuryFeature     []byte         `gorm:"type:jsonb" json:"-"`
	CommunityFeature  []byte         `gorm:"type:jsonb" json:"-"`
	UtilsFeature      []byte         `gorm:"type:jsonb" json:"-"`
	EnergyFeature     []byte         `gorm:"type:jsonb" json:"-"`

	// Getter methods
	AmenitiesData         Amenities         `gorm:"-" json:"amenities"`
	SecurityFeatureData   SecurityFeature   `gorm:"-" json:"securityFeature"`
	TechnologyFeatureData TechnologyFeature `gorm:"-" json:"technologyFeature"`
	LuxuryFeatureData     LuxuryFeature     `gorm:"-" json:"luxuryFeature"`
	CommunityFeatureData  CommunityFeature  `gorm:"-" json:"communityFeature"`
	UtilsFeatureData      UtilsFeature      `gorm:"-" json:"utilsFeature"`
	EnergyFeatureData     EnergyFeature     `gorm:"-" json:"energyFeature"`
}

// AfterFind hook
func (p *PropertyFeature) AfterFind(tx *gorm.DB) error {

	if len(p.Amenities) > 0 {
		if err := json.Unmarshal(p.Amenities, &p.AmenitiesData); err != nil {
			return err
		}
	}

	if len(p.SecurityFeature) > 0 {
		if err := json.Unmarshal(p.SecurityFeature, &p.SecurityFeatureData); err != nil {
			return err
		}
	}

	if len(p.TechnologyFeature) > 0 {
		if err := json.Unmarshal(p.TechnologyFeature, &p.TechnologyFeatureData); err != nil {
			return err
		}
	}

	if len(p.LuxuryFeature) > 0 {
		if err := json.Unmarshal(p.LuxuryFeature, &p.LuxuryFeatureData); err != nil {
			return err
		}
	}

	if len(p.CommunityFeature) > 0 {
		if err := json.Unmarshal(p.CommunityFeature, &p.CommunityFeatureData); err != nil {
			return err
		}
	}

	if len(p.UtilsFeature) > 0 {
		if err := json.Unmarshal(p.UtilsFeature, &p.UtilsFeatureData); err != nil {
			return err
		}
	}

	if len(p.EnergyFeature) > 0 {
		if err := json.Unmarshal(p.EnergyFeature, &p.EnergyFeatureData); err != nil {
			return err
		}
	}

	return nil
}

func (p *PropertyFeature) BeforeSave(tx *gorm.DB) error {
	// Convert struct fields to JSON
	var err error

	p.Amenities, err = json.Marshal(p.AmenitiesData)
	if err != nil {
		return err
	}

	p.SecurityFeature, err = json.Marshal(p.SecurityFeatureData)
	if err != nil {
		return err
	}

	p.TechnologyFeature, err = json.Marshal(p.TechnologyFeatureData)
	if err != nil {
		return err
	}

	p.LuxuryFeature, err = json.Marshal(p.LuxuryFeatureData)
	if err != nil {
		return err
	}

	p.CommunityFeature, err = json.Marshal(p.CommunityFeatureData)
	if err != nil {
		return err
	}

	p.UtilsFeature, err = json.Marshal(p.UtilsFeatureData)
	if err != nil {
		return err
	}

	p.EnergyFeature, err = json.Marshal(p.EnergyFeatureData)
	if err != nil {
		return err
	}

	return nil
}
