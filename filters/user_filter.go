// filters/user_filter.go
package filters

import (
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
	"gorm.io/gorm"
)

type UserFilterManager struct {
	Filter  *BaseFilter
	SortBy  string
	SortDir string
}

func NewUserFilterManager(filterDTO dto.UserFilterDTO) *UserFilterManager {
	filter := NewFilter(&models.User{})

	// Define valid sort fields for users
	validSortFields := map[string]bool{
		"id":         true,
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"joined_at":  true,
		"status":     true,
	}

	sortField := "first_name" // Default
	if filterDTO.SortBy != nil {
		if _, ok := validSortFields[*filterDTO.SortBy]; ok {
			sortField = *filterDTO.SortBy
		}
	}

	sortDir := "ASC" // Default
	if filterDTO.SortOrder != nil {
		order := *filterDTO.SortOrder
		if order == "desc" || order == "DESC" {
			sortDir = "DESC"
		}
	}

	manager := &UserFilterManager{
		Filter:  filter,
		SortBy:  sortField,
		SortDir: sortDir,
	}

	manager.AddStandardFilters(filterDTO)

	return manager
}

// AddStandardFilters adds standard user filters based on the DTO
func (uf *UserFilterManager) AddStandardFilters(filterDTO dto.UserFilterDTO) {
	// Role filter
	if filterDTO.Role != nil && *filterDTO.Role != "" {
		uf.Filter.AddFilter(FilterOption{
			Field:     "role",
			Operator:  "=",
			Value:     filterDTO.Role,
			IsPointer: true,
		})
	}

	// Status filter
	if filterDTO.Status != nil && *filterDTO.Status != "" {
		uf.Filter.AddFilter(FilterOption{
			Field:     "status",
			Operator:  "=",
			Value:     filterDTO.Status,
			IsPointer: true,
		})
	}

	// Email verification filter
	if filterDTO.EmailVerified != nil {
		uf.Filter.AddFilter(FilterOption{
			Field:     "email_verified",
			Operator:  "=",
			Value:     filterDTO.EmailVerified,
			IsPointer: true,
		})
	}

	// Search filter (as OR conditions)
	if filterDTO.Search != nil && *filterDTO.Search != "" {
		// This will be handled separately during Apply() because it's an OR condition
	}
}

// Apply applies all filters and sorting to a query
func (uf *UserFilterManager) Apply(db *gorm.DB, filterDTO dto.UserFilterDTO) *gorm.DB {
	// Apply base filters
	query := uf.Filter.Build(db)

	// Handle search as a special case (OR condition)
	if filterDTO.Search != nil && *filterDTO.Search != "" {
		searchTerm := *filterDTO.Search
		query = OrCondition(query,
			func(db *gorm.DB) *gorm.DB {
				return db.Where("first_name LIKE ?", "%"+searchTerm+"%")
			},
			func(db *gorm.DB) *gorm.DB {
				return db.Where("last_name LIKE ?", "%"+searchTerm+"%")
			},
			func(db *gorm.DB) *gorm.DB {
				return db.Where("email LIKE ?", "%"+searchTerm+"%")
			},
		)
	}

	// Apply sorting
	validFields := map[string]bool{
		"id":         true,
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"joined_at":  true,
		"status":     true,
	}

	query = ApplySorting(query, uf.SortBy, uf.SortDir, validFields)

	return query
}

// AddCustomFilter adds a custom filter option
func (uf *UserFilterManager) AddCustomFilter(field, operator string, value interface{}, isPointer bool) *UserFilterManager {
	uf.Filter.AddFilter(FilterOption{
		Field:     field,
		Operator:  operator,
		Value:     value,
		IsPointer: isPointer,
	})
	return uf
}
