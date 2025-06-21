package utils

import (
	"time"
)

// Verifca se o valor date foi inserido no valor correto
func IsValidDate(dateStr string) bool {

	t, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return false
	}
	_ = t
	return true
}
