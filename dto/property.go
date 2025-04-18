package dto

import "time"

type PropertyRequestDTO struct {
	Title        string  `json:"title" binding:"required"`
	Purpose      string  `json:"purpose" binding:"required,oneof=sale rent"`
	Price        float64 `json:"price" binding:"required,gt=0"`
	PropertyType string  `json:"property_type" binding:"required"`
	Bedrooms     int     `json:"bedrooms" binding:"required,gte=0"`
	Bathrooms    int     `json:"bathrooms" binding:"required,gte=0"`
	Size         float64 `json:"size" binding:"required,gt=0"`
	BuiltYear    int     `json:"built_year" binding:"omitempty,gt=0"`
	CountryID    uint32  `json:"country_id" binding:"required"`
	DivisionID   uint32  `json:"division_id" binding:"required"`
	DistrictID   uint32  `json:"district_id" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	Description  string  `json:"description" binding:"required"`
}

type PropertyResponseDTO struct {
	ID           uint32  `json:"id"`
	Title        string  `json:"title"`
	Purpose      string  `json:"purpose"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	PropertyType string  `json:"property_type"`
	BedRooms     int     `json:"bedrooms"`
	BathRooms    int     `json:"bathrooms"`
	Size         float64 `json:"size"`
	BuiltYear    int     `json:"built_year"`
	CountryID    uint32  `json:"country_id"`
	DivisionID   uint32  `json:"division_id"`
	DistrictID   uint32  `json:"district_id"`
	Address      string  `json:"address"`
	Description  string  `json:"description"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type PropertyListDTO struct {
	ID           uint32                     `json:"id"`
	Title        string                     `json:"title"`
	Purpose      string                     `json:"purpose"`
	Price        float64                    `json:"price"`
	PropertyType string                     `json:"property_type"`
	Country      CountryMinimalDTO          `json:"country"`
	Division     DivisionMinimal2DTO        `json:"division"`
	District     DistrictMinimalResponseDTO `json:"district"`
	Status       string                     `json:"status"`
	Address      string                     `json:"address"`
	Views        int                        `json:"views"`
	Inquiries    int                        `json:"inquiries"`
	CreatedAt    string                     `json:"created_at"`
}

type PropertyListResponseDTO struct {
	Properties []PropertyListDTO `json:"properties"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	PerPage    int               `json:"per_page"`
}

type PropertyFilterDTO struct {
	OwerID       uint    `form:"owner_id"`
	Purpose      string  `form:"purpose"`
	MinPrice     float64 `form:"min_price"`
	MaxPrice     float64 `form:"max_price"`
	PropertyType string  `form:"property_type"`
	BedRooms     int     `form:"bedrooms"`
	BathRooms    int     `form:"bathrooms"`
	MinSize      float64 `form:"min_size"`
	MaxSize      float64 `form:"max_size"`
	CountryID    uint32  `form:"country_id"`
	DivisionID   uint32  `form:"division_id"`
	DistrictID   uint32  `form:"district_id"`
	Page         int     `form:"page,default=1"`
	PerPage      int     `form:"per_page,default=10"`
	Status       string  `form:"status"`
}

type AmenitiesDTO struct {
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

type SecurityFeatureDTO struct {
	SecuritySystem bool `json:"securitySystem"`
	Doorman        bool `json:"doorman"`
	SecurityCamera bool `json:"securityCamera"`
	GatedCommunity bool `json:"gatedCommunity"`
	FireAlarm      bool `json:"fireAlarm"`
}

type TechnologyFeatureDTO struct {
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

type LuxuryFeatureDTO struct {
	Fireplace     bool `json:"fireplace"`
	Pool          bool `json:"pool"`
	Gym           bool `json:"gym"`
	WalkInClosets bool `json:"walkInClosets"`
	Jacuzzi       bool `json:"jacuzzi"`
	Sauna         bool `json:"sauna"`
}

type CommunityFeatureDTO struct {
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

type UtilsFeatureDTO struct {
	WaterIncluded        bool `json:"waterIncluded"`
	GasIncluded          bool `json:"gasIncluded"`
	ElectricityIncluded  bool `json:"electricityIncluded"`
	TrashRemovalIncluded bool `json:"trashRemovalIncluded"`
	InternetIncluded     bool `json:"internetIncluded"`
}

type EnergyFeatureDTO struct {
	SolarPanels               bool `json:"solarPanels"`
	EnergyEfficientAppliances bool `json:"energyEfficientAppliances"`
	GreenCertification        bool `json:"greenCertification"`
	EVCharging                bool `json:"evCharging"`
	RainwaterHarvesting       bool `json:"rainwaterHarvesting"`
	ProgrammableThermostat    bool `json:"programmableThermostat"`
}

type PropertyFeatureDTO struct {
	PropertyID uint     `json:"property_id"`
	Features   []string `json:"features"`

	Amenities         AmenitiesDTO         `json:"amenities"`
	SecurityFeature   SecurityFeatureDTO   `json:"securityFeature"`
	TechnologyFeature TechnologyFeatureDTO `json:"technologyFeature"`
	LuxuryFeature     LuxuryFeatureDTO     `json:"luxuryFeature"`
	CommunityFeature  CommunityFeatureDTO  `json:"communityFeature"`
	UtilsFeature      UtilsFeatureDTO      `json:"utilsFeature"`
	EnergyFeature     EnergyFeatureDTO     `json:"energyFeature"`
}

type PropertyFeatureDetailsDTO struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	PropertyID uint      `json:"property_id" binding:"required"`
	Features   []string  `json:"features"`

	Amenities         AmenitiesDTO         `json:"amenities"`
	SecurityFeature   SecurityFeatureDTO   `json:"securityFeature"`
	TechnologyFeature TechnologyFeatureDTO `json:"technologyFeature"`
	LuxuryFeature     LuxuryFeatureDTO     `json:"luxuryFeature"`
	CommunityFeature  CommunityFeatureDTO  `json:"communityFeature"`
	UtilsFeature      UtilsFeatureDTO      `json:"utilsFeature"`
	EnergyFeature     EnergyFeatureDTO     `json:"energyFeature"`
}
