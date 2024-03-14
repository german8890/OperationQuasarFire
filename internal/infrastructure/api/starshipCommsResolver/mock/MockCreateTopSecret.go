package mock

import (
	"github.com/stretchr/testify/mock"
)

type CommunicationServicesMock struct {
	mock.Mock
}

func (m *CommunicationServicesMock) GetLocation(distances ...float32) (float32, float32) {
	args := m.Called(distances)
	return args.Get(0).(float32), args.Get(1).(float32)
}

func (m *CommunicationServicesMock) GetMessage(messages ...[]string) string {
	args := m.Called(messages)
	return args.Get(0).(string)
}
