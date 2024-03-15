package service

import (
	"strings"

	"github.com/OperationQuasarFire/internal/pkg/ports"
)

type serviceSplit struct {
	ports ports.CommunicationServices
}

func NewServiceSplit(port ports.CommunicationServices) *serviceSplit {
	return &serviceSplit{}
}

func (s *serviceSplit) GetLocation(distances ...float32) (x, y float32) {

	kenobiCoords := [2]float32{-500.0, -200.0}
	skywalkerCoords := [2]float32{100.0, -100.0}
	satoCoords := [2]float32{500.0, 100.0}

	distanceKenobi := float64(distances[0])
	distanceSkywalker := float64(distances[1])
	distanceSato := float64(distances[2])

	emisorX := float32(0) // Inicializar con un valor temporal
	emisorY := float32(0) // Inicializar con un valor temporal

	dx1 := kenobiCoords[0] - emisorX
	dy1 := kenobiCoords[1] - emisorY
	dx2 := skywalkerCoords[0] - emisorX
	dy2 := skywalkerCoords[1] - emisorY
	dx3 := satoCoords[0] - emisorX
	dy3 := satoCoords[1] - emisorY

	A := float64(2 * dx2)
	B := float64(2 * dy2)
	C := float64(distanceKenobi*distanceKenobi - distanceSkywalker*distanceSkywalker - float64(dx1*dx1) - float64(dy1*dy1) + float64(dx2*dx2) + float64(dy2*dy2))
	D := float64(2 * dx3)
	E := float64(2 * dy3)
	F := float64(distanceKenobi*distanceKenobi - distanceSato*distanceSato - float64(dx1*dx1) - float64(dy1*dy1) + float64(dx3*dx3) + float64(dy3*dy3))

	denominator := B*D - A*E
	if denominator == 0 {
		return 0, 0
	}

	emisorX = float32((C*E - F*B) / denominator)
	emisorY = float32((C*D - A*F) / denominator)

	return emisorX, emisorY
}

func (s *serviceSplit) GetMessage(messages ...[]string) string {

	adjustedMessage := make([]string, 5)

	for _, m := range messages {
		for i, world := range m {
			if adjustedMessage[i] == "" {
				adjustedMessage[i] = world
			}
		}
	}

	finalMessage := strings.Join(adjustedMessage, " ")
	return finalMessage
}
