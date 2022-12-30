package indodax

import (
	"fmt"
	"github.com/buger/jsonparser"
)

type PriceIncrements struct {
	Entries map[string]float64
}

func (is *PriceIncrements) UnmarshalJSON(b []byte) (err error) {
	incrementsBA, _, _, err := jsonparser.Get(b, "increments")
	if err != nil {
		return err
	}

	is.Entries = make(map[string]float64)

	return jsonparser.ObjectEach(
		incrementsBA,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			fv, fErr := jsonparser.GetFloat(value)
			pair := string(key)
			if fErr != nil {
				return fmt.Errorf("failed to parse [%s] price increment with value [%s]: %v", pair, string(value), fErr)
			}

			is.Entries[pair] = fv

			return nil
		},
	)
}
