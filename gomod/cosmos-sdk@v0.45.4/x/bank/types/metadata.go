package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)








func (m Metadata) Validate() error {
	if strings.TrimSpace(m.Name) == "" {
		return errors.New("name field cannot be blank")
	}

	if strings.TrimSpace(m.Symbol) == "" {
		return errors.New("symbol field cannot be blank")
	}

	if err := sdk.ValidateDenom(m.Base); err != nil {
		return fmt.Errorf("invalid metadata base denom: %w", err)
	}

	if err := sdk.ValidateDenom(m.Display); err != nil {
		return fmt.Errorf("invalid metadata display denom: %w", err)
	}

	var (
		hasDisplay      bool
		currentExponent uint32
	)

	seenUnits := make(map[string]bool)

	for i, denomUnit := range m.DenomUnits {

		if i == 0 {

			if denomUnit.Denom != m.Base {
				return fmt.Errorf("metadata's first denomination unit must be the one with base denom '%s'", m.Base)
			}
			if denomUnit.Exponent != 0 {
				return fmt.Errorf("the exponent for base denomination unit %s must be 0", m.Base)
			}
		} else if currentExponent >= denomUnit.Exponent {
			return errors.New("denom units should be sorted asc by exponent")
		}

		currentExponent = denomUnit.Exponent

		if seenUnits[denomUnit.Denom] {
			return fmt.Errorf("duplicate denomination unit %s", denomUnit.Denom)
		}

		if denomUnit.Denom == m.Display {
			hasDisplay = true
		}

		if err := denomUnit.Validate(); err != nil {
			return err
		}

		seenUnits[denomUnit.Denom] = true
	}

	if !hasDisplay {
		return fmt.Errorf("metadata must contain a denomination unit with display denom '%s'", m.Display)
	}

	return nil
}


func (du DenomUnit) Validate() error {
	if err := sdk.ValidateDenom(du.Denom); err != nil {
		return fmt.Errorf("invalid denom unit: %w", err)
	}

	seenAliases := make(map[string]bool)
	for _, alias := range du.Aliases {
		if seenAliases[alias] {
			return fmt.Errorf("duplicate denomination unit alias %s", alias)
		}

		if strings.TrimSpace(alias) == "" {
			return fmt.Errorf("alias for denom unit %s cannot be blank", du.Denom)
		}

		seenAliases[alias] = true
	}

	return nil
}
