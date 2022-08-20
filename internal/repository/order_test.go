package repository

import (
	"github.com/magiconair/properties/assert"
	"github.com/onemgvv/wb-l0/internal/domain"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestOrderRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewOrderRepository(db)

	type args struct {
		id   string
		data string
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				id:   "ab1",
				data: `{"id":"ab1","name":"test",count:11}`,
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(args.id)
				mock.ExpectQuery("INSERT INTO").WithArgs(args.id, args.data).WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			got, err := r.Create(testCase.args.id, testCase.args.data)
			if testCase.wantError {
				assert.Equal(t, "", err.Error())
			} else {
				assert.Equal(t, testCase.args.id, got)
			}
		})
	}
}

func TestOrderRepository_GetOrder(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewOrderRepository(db)

	order := domain.Order{
		UID:  "f4bde549-ef88-4cff-a974-fd1abee0e598",
		Data: `{"orderID": "542", "sum": "1.300", "curr": "RUB"}`,
	}
	type mockBehavior func(id string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		id           string
		order        domain.Order
		wantError    bool
	}{
		{
			name: "OK",
			id:   "f4bde549-ef88-4cff-a974-fd1abee0e598",
			mockBehavior: func(id string) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "data"}).AddRow(order.UID, order.Data)
				mock.ExpectQuery("SELECT * FROM orders where id = 'f4bde549-ef88-4cff-a974-fd1abee0e598'").WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty ID",
			id:   "",
			mockBehavior: func(id string) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "data"}).AddRow(order.UID, order.Data)
				mock.ExpectQuery("SELECT * FROM orders where id = ''").WillReturnRows(rows)

				mock.ExpectRollback()
			},
		},
		{
			name: "Not Found",
			id:   "43123",
			mockBehavior: func(id string) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "data"}).AddRow(order.UID, order.Data)
				mock.ExpectQuery("SELECT * FROM orders where id = '43123'").WillReturnRows(rows)

				mock.ExpectRollback()
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.id)

			got, err := r.GetOrder(testCase.id)
			if testCase.wantError {
				t.Errorf("[ERROR Get Ordder Test: %s]: %s", testCase.name, err)
			} else {
				assert.Equal(t, testCase.order, got)
			}
		})
	}
}
