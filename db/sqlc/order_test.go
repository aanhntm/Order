package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		ID:          2,
		UserName:    sql.NullString{String: "My Anh", Valid: true},
		ProductName: sql.NullString{String: "Banana", Valid: true},
		Amount:      3,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	fmt.Println(order.Amount)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.ID, order.ID)
	require.NotZero(t, order.ID)
	require.NotZero(t, order.Amount)

	return order
}

func TestCreateOrder(t *testing.T) {
	CreateOrder(t)
}

/*
func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.getOneOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserName, order2.UserName)
}

func TestGetManyOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}
	arg := getManyOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.getManyOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}*/
