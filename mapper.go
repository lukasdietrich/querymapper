package querymapper

import (
	"fmt"
	"net/url"
	r "reflect"
	"strconv"
	"strings"
)

func getValue(v url.Values, key string) (string, bool) {
	if s, ok := v[key]; !ok || len(s) == 0 {
		return "", false
	}

	return s[len(s)-1], true
}

// MapQuery sets the values of m to
// the respective parameters in v
func MapQuery(v url.Values, m interface{}) error {
	s := r.ValueOf(m).Elem()
	t := s.Type()

	for i := s.NumField() - 1; i >= 0; i-- {
		fv, ft := s.Field(i), t.Field(i)
		key := ft.Tag.Get("query")

		if key == "" {
			key = strings.ToLower(ft.Name)
		}

		sval, ok := getValue(v, key)
		if !ok {
			return fmt.Errorf("querymapper: missing query value for '%s'", key)
		}

		switch fv.Kind() {
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64:
			if val, err := strconv.ParseUint(sval, 10, 64); err != nil {
				return err
			}

			fv.SetUint(val)

		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			if val, err := strconv.ParseInt(sval, 10, 64); err != nil {
				return err
			}

			fv.SetInt(val)

		case r.Float32, r.Float64:
			if val, err := strconv.ParseFloat(sval, 64); err != nil {
				return err
			}

			fv.SetFloat(val)

		case r.Bool:
			if val, err := strconv.ParseBool(sval); err != nil {
				return err
			}

			fv.SetBool(val)

		case r.String:
			fv.SetString(sval)

		default:
			return fmt.Errorf("querymapper: unsupported type '%s'",
				fv.Type().Name())
		}
	}

	return nil
}
