package dto

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
	ID           uint32              `json:"id"`
	Title        string              `json:"title"`
	Purpose      string              `json:"purpose"`
	Price        float64             `json:"price"`
	Status       string              `json:"status"`
	PropertyType string              `json:"property_type"`
	BedRooms     int                 `json:"bedrooms"`
	BathRooms    int                 `json:"bathrooms"`
	Size         float64             `json:"size"`
	BuiltYear    int                 `json:"built_year"`
	Country      CountryResponseDTO  `json:"country"`
	Division     DivisionResponseDTO `json:"division"`
	District     DistrictResponseDTO `json:"district"`
	Address      string              `json:"address"`
	Description  string              `json:"description"`
	CreatedAt    string              `json:"created_at"`
	UpdatedAt    string              `json:"updated_at"`
}

type PropertyListResponseDTO struct {
	Properties []PropertyResponseDTO `json:"properties"`
	Total      int64                 `json:"total"`
	Page       int                   `json:"page"`
	PerPage    int                   `json:"per_page"`
}

type PropertyFilterDTO struct {
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
}
