package services

import (
	"testing"
	mocks "testing_demo/mocks/services"
)

/* type MockMessageService struct {
	mock.Mock
}

func (s MockMessageService) Send(msg string) bool {
	fmt.Println("Mock Send() method invoked")
	args := s.Mock.Called(msg)
	return args.Bool(0)
}

func TestMessageProcessor(t *testing.T) {
	mockMessageService := MockMessageService{}
	mockMessageService.Mock.On("Send", "Hello").Return(true)

	processor := MessageProcessor{mockMessageService}
	processor.Process("Hello")

	mockMessageService.AssertExpectations(t)
}
*/

func TestMessageProcessor(t *testing.T) {
	mockMessageService := &mocks.MessageService{}
	mockMessageService.Mock.On("Send", "Hello").Return(true)

	processor := MessageProcessor{mockMessageService}
	processor.Process("Hello")

	mockMessageService.AssertExpectations(t)
}
