package filters

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

// FilterOption represents a single filter condition
type FilterOption struct {
	Field      string      // Database field name
	Operator   string      // SQL operator (e.g., "=", "LIKE", ">", etc.)
	Value      interface{} // Value to compare against
	IsPointer  bool        // Whether the value is a pointer
	JoinFields []string    // For filtering on related models
}

type FilterBuilder interface {
	AddFilter(option FilterOption) FilterBuilder
	Build(db *gorm.DB) *gorm.DB
}

type BaseFilter struct {
	Options []FilterOption
	Model   interface{}
}

func NewFilter(model interface{}) *BaseFilter {
	return &BaseFilter{
		Options: []FilterOption{},
		Model:   model,
	}
}

func (f *BaseFilter) AddFilter(option FilterOption) FilterBuilder {
	f.Options = append(f.Options, option)
	return f
}

func (f *BaseFilter) Build(db *gorm.DB) *gorm.DB {
	query := db

	for _, option := range f.Options {
		if option.IsPointer {
			// Handle pointer values (check if nil or empty)
			v := reflect.ValueOf(option.Value)
			if v.IsNil() {
				continue
			}

			// Dereference the pointer
			value := reflect.Indirect(v).Interface()

			// Handle specific operators
			switch option.Operator {
			case "=":
				query = query.Where(option.Field+" = ?", value)
			case "LIKE":
				if strValue, ok := value.(string); ok && strValue != "" {
					likeValue := "%" + strValue + "%"
					query = query.Where(option.Field+" LIKE ?", likeValue)
				}
			case "IN":
				query = query.Where(option.Field+" IN ?", value)
			default:
				query = query.Where(option.Field+" "+option.Operator+" ?", value)
			}
		} else {
			// Handle non-pointer values
			switch option.Operator {
			case "=":
				query = query.Where(option.Field+" = ?", option.Value)
			case "LIKE":
				if strValue, ok := option.Value.(string); ok && strValue != "" {
					likeValue := "%" + strValue + "%"
					query = query.Where(option.Field+" LIKE ?", likeValue)
				}
			case "IN":
				query = query.Where(option.Field+" IN ?", option.Value)
			default:
				query = query.Where(option.Field+" "+option.Operator+" ?", option.Value)
			}
		}
	}

	return query
}

// OrCondition applies multiple conditions with OR logic
func OrCondition(db *gorm.DB, conditions ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	if len(conditions) == 0 {
		return db
	}

	return db.Where(func(db *gorm.DB) {
		subQuery := db
		for i, condition := range conditions {
			if i == 0 {
				subQuery = condition(subQuery)
			} else {
				subQuery = subQuery.Or(condition(db.Session(&gorm.Session{})))
			}
		}
	})
}

// SortOption defines sorting parameters
type SortOption struct {
	Field     string
	Direction string // "ASC" or "DESC"
}

// ApplySorting applies sort options to a query
func ApplySorting(db *gorm.DB, sortField, sortOrder string, validFields map[string]bool) *gorm.DB {
	// Validate sort field
	if _, ok := validFields[sortField]; !ok {
		sortField = "id" // Default
	}

	// Validate sort order
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "ASC" // Default
	}

	return db.Order(sortField + " " + sortOrder)
}
