package dto

type UserFilterDTO struct {
	Role          *string `form:"role"`
	Status        *string `form:"status"`
	EmailVerified *bool   `form:"email_verified"`
	Search        *string `form:"search"`
	SortBy        *string `form:"sort_by"`
	SortOrder     *string `form:"sort_order"`
}

// NewUserFilterDTO creates a new filter with default values
func NewUserFilterDTO() UserFilterDTO {
	return UserFilterDTO{}
}

// GetSortField returns the field to sort by with a default value
func (f UserFilterDTO) GetSortField() string {
	if f.SortBy == nil {
		return "first_name"
	}

	// Validate sort field to prevent SQL injection
	validFields := map[string]bool{
		"id": true, "first_name": true, "last_name": true,
		"email": true, "joined_at": true, "status": true,
	}

	if _, ok := validFields[*f.SortBy]; ok {
		return *f.SortBy
	}

	return "first_name" // Default
}

// GetSortOrder returns the sort order with a default value
func (f UserFilterDTO) GetSortOrder() string {
	if f.SortOrder == nil {
		return "ASC"
	}

	order := *f.SortOrder
	if order == "desc" || order == "DESC" {
		return "DESC"
	}

	return "ASC" // Default
}
