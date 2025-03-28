package dto

type CountryRequestDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryUpdateRequestDTO struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status bool   `json:"status"`
}

type CountryResponseDTO struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    bool   `json:"status"`
	Divisions int64  `json:"divisions"`
}

type CountryMinimalDTO struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type DivisionRequestDTO struct {
	Name      string `json:"name"`
	CountryId uint32 `json:"country_id"`
}

type DivisionUpdateRequestDTO struct {
	CountryID uint32 `json:"country_id"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
}

type DivisionResponseDTO struct {
	ID        uint32            `json:"id"`
	Name      string            `json:"name"`
	Country   CountryMinimalDTO `json:"country"`
	Status    bool              `json:"status"`
	Districts int64             `json:"districts"`
}

type DivisionMinimalDTO struct {
	ID      uint32            `json:"id"`
	Name    string            `json:"name"`
	Country CountryMinimalDTO `json:"country"`
}

type DistrictRequestDTO struct {
	Name       string `json:"name"`
	DivisionId uint32 `json:"division_id"`
}

type DistrictResponseDTO struct {
	ID       uint32             `json:"id"`
	Name     string             `json:"name"`
	Division DivisionMinimalDTO `json:"division"`
	Status   bool               `json:"status"`
}

type DistrictUpdateRequestDTO struct {
	Name       string `json:"name"`
	DivisionId uint32 `json:"division_id"`
	Status     bool   `json:"status"`
}

// public response DTO

// Public DTOs for API responses

type PublicCountryDTO struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type PublicDivisionDTO struct {
	ID      uint32           `json:"id"`
	Name    string           `json:"name"`
	Country PublicCountryDTO `json:"country"`
}

type PublicDistrictDTO struct {
	ID       uint32            `json:"id"`
	Name     string            `json:"name"`
	Country  PublicCountryDTO  `json:"country"`
	Division PublicDivisionDTO `json:"division"`
}

// Response wrapper with pagination info
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}
