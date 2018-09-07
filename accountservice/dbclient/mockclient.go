package dbclient

import (
	"github.com/rabih/go-blog/accountservice/model"
	"github.com/stretchr/testify/mock"
)

type MockBoltClient struct {
	mock.Mock
}

// From here, we'll declare three functions that makes our MockBoltClient fulfill the interface IBoltClient that we declared in part 3.
func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}
func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}
