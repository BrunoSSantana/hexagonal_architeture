package application_test

import (
	"testing"

	"github.com/BrunoSSanta/hexagonal-go/application"
	mock_application "github.com/BrunoSSanta/hexagonal-go/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistance := mock_application.NewMockProductPersistence(ctrl)
	persistance.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistance,
	}

	result, err := service.Get("any_id")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistance := mock_application.NewMockProductPersistence(ctrl)
	persistance.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistance}

	result, err := service.Create("Product 01", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
