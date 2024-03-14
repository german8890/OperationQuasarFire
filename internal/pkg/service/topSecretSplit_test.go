package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceSplit_GetLocation(t *testing.T) {
	service := NewServiceSplit(nil)

	x, y := service.GetLocation(100, 142.14, 125)
	fmt.Printf("Calculated: x = %f, y = %f\n", x, y)
}

func TestServiceSplit_GetMessage(t *testing.T) {
	service := NewServiceSplit(nil)

	message := service.GetMessage([]string{"este", "", "", "mensaje", ""}, []string{"", "es", "", "", "secreto"})
	assert.Equal(t, "este es  mensaje secreto", message)

	message = service.GetMessage([]string{"", "", "", ""}, []string{"", "", "", ""})
	assert.Equal(t, "    ", message)
}
