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

type DivisionRequestDTO struct {
	Name      string `json:"name"`
	CountryId uint32 `json:"country_id"`
}

type DivisionUpdateRequestDTO struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type DivisionResponseDTO struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	CountryID uint32 `json:"country_id"`
	Status    bool   `json:"status"`
	Districts int64  `json:"districts"`
}
