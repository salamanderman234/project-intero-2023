package config

import (
	"github.com/asaskevich/govalidator"
)

func RegisterCustomValidationRules() {
	govalidator.CustomTypeTagMap.Set("mustUint", func(i, o interface{}) bool {
		// _, ok := o.(uint)
		return false
	})
}