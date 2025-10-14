package birdeye

import (
	"reflect"
	"strconv"
	"strings"
)

// DefaultTag is the struct tag key for default values
const DefaultTag = "default"

// ApplyDefaults applies default values to a struct based on "default" tags
// This allows us to use non-pointer fields in Opts structs and still have optional parameters
func ApplyDefaults(opts any) error {
	if opts == nil {
		return nil
	}

	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		// Get the default tag value
		defaultValue := fieldType.Tag.Get(DefaultTag)
		if defaultValue == "" {
			continue
		}

		// Skip if field already has a value (not zero value)
		if !field.IsZero() {
			continue
		}

		// Apply default value based on field type
		if err := setFieldValue(field, defaultValue); err != nil {
			return err
		}
	}

	return nil
}

// setFieldValue sets a field value from a string representation
func setFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intVal)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(uintVal)

	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatVal)

	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)

	default:
		// For custom types, try to find a string representation
		if field.Type().Implements(reflect.TypeOf((*interface{ String() string })(nil)).Elem()) {
			// This is a custom type that implements String() method
			// We'll need to handle this case by case
			return nil
		}
	}

	return nil
}

// ParseDefaultTag parses a default tag value that might contain multiple options
// Format: "value" or "value|description"
func ParseDefaultTag(tagValue string) (value, description string) {
	parts := strings.SplitN(tagValue, "|", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

// ApplyDefaultsAndBuildParams applies default values and builds API parameters map
// It excludes fields that are not API parameters (like Chains, OnLimitExceeded)
func ApplyDefaultsAndBuildParams(opts any) (map[string]any, error) {
	if opts == nil {
		return make(map[string]any), nil
	}

	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return make(map[string]any), nil
	}

	params := make(map[string]any)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		fieldName := fieldType.Name

		// Skip fields that are not API parameters
		if isNonAPIField(fieldName) {
			continue
		}

		// Get the default tag value
		defaultValue := fieldType.Tag.Get(DefaultTag)

		// Apply default value if field is zero and has default tag
		if field.IsZero() && defaultValue != "" {
			if err := setFieldValue(field, defaultValue); err != nil {
				return nil, err
			}
		}

		// Add field to params if it has a non-zero value
		if !field.IsZero() {
			paramName := toSnakeCase(fieldName)

			// Handle special cases for API parameters
			switch field.Kind() {
			case reflect.Bool:
				// Convert bool to string for API
				if field.Bool() {
					params[paramName] = "true"
				}
			default:
				params[paramName] = field.Interface()
			}
		}
	}

	return params, nil
}

// isNonAPIField checks if a field should be excluded from API parameters
func isNonAPIField(fieldName string) bool {
	nonAPIFields := []string{
		"Chains",
		"OnLimitExceeded",
		"Chain", // for single chain fields
	}

	for _, field := range nonAPIFields {
		if fieldName == field {
			return true
		}
	}
	return false
}

// toSnakeCase converts CamelCase to snake_case
func toSnakeCase(s string) string {
	// Handle special cases first
	specialCases := map[string]string{
		"UIAmountMode":        "ui_amount_mode",
		"CheckLiquidity":      "check_liquidity",
		"IncludeLiquidity":    "include_liquidity",
		"AfterTime":           "after_time",
		"BeforeTime":          "before_time",
		"BeforeBlockNumber":   "before_block_number",
		"AfterBlockNumber":    "after_block_number",
		"TimeFrom":            "time_from",
		"TimeTo":              "time_to",
		"SortType":            "sort_type",
		"SortBy":              "sort_by",
		"TxType":              "tx_type",
		"TimeFrame":           "time_frame",
		"SearchMode":          "search_mode",
		"SearchBy":            "search_by",
		"MinLiquidity":        "min_liquidity",
		"MaxLiquidity":        "max_liquidity",
		"MinMarketCap":        "min_market_cap",
		"MaxMarketCap":        "max_market_cap",
		"MinFDV":              "min_fdv",
		"MaxFDV":              "max_fdv",
		"FilterValue":         "filter_value",
		"VerifyToken":         "verify_token",
		"PlatformID":          "platform_id",
		"ScrollID":            "scroll_id",
		"ChangeType":          "change_type",
		"CountLimit":          "count_limit",
		"MemePlatformEnabled": "meme_platform_enabled",
	}

	if result, exists := specialCases[s]; exists {
		return result
	}

	// Default conversion for other fields
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
