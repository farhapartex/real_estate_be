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
