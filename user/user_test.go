package user_test

import (
	"errors"
	"testing"

	"github.com/an1l4/testing-gomock-mockgen/mocks"
	"github.com/an1l4/testing-gomock-mockgen/user"
	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mockDoer := mocks.NewMockDoer(mockctrl)
	testUser := &user.User{Doer: mockDoer}

	mockDoer.EXPECT().DoSometing(123, "Hello mock").Return(nil).Times(1)

	testUser.Use()

}

func TestUseReturnErrorFromDo(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	dummyError := errors.New("dummy error")

	mockDoer := mocks.NewMockDoer(mockctrl)
	testUser := &user.User{Doer: mockDoer}

	mockDoer.EXPECT().DoSometing(123, "Hello mock").Return(dummyError).Times(1)

	err := testUser.Use()

	if err != dummyError {
		t.Fail()
	}
}
