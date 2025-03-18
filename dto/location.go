package dto

type CountryRequestDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryResponseDTO struct {
	ID uint32 `json:"id"`
	CountryRequestDTO
}
