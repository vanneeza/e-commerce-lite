package storerepository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

var DummyStore = []entity.Store{
	{
		Id:      "qwerty123456",
		Name:    "Vanneeza Store",
		Email:   "Vnza@gmail.com",
		NoHp:    "089453234534",
		Address: "Jln Raya Bogor No 10",
		Balance: 1000,
	},
	{
		Id:      "123456qwerty",
		Name:    "Chauzar Store",
		Email:   "Cha@gmail.com",
		NoHp:    "08423534234",
		Address: "Jln Raya Bogor No 15",
		Balance: 5000,
	},
}

func TestStoreRepository_Create_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repository := &StoreRepositoryImpl{}

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)
	defer tx.Rollback()

	mock.ExpectExec("INSERT INTO tbl_store").
		WithArgs(DummyStore[0].Name, DummyStore[0].Email, DummyStore[0].NoHp, DummyStore[0].Address, DummyStore[0].Balance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdStore, err := repository.Create(context.Background(), tx, DummyStore[0])

	assert.NoError(t, err)
	assert.Equal(t, DummyStore[0], createdStore)
	assert.NoError(t, mock.ExpectationsWereMet())
}
