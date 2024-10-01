package db

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	// Mock the database
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer db.Close()

	// Mock the Ping function
	mock.ExpectPing().WillReturnError(nil)

	// Call the ConnectToDB function
	connStr := "postgres://tanryberdi:tanryberdi@localhost:5432/testdb?sslmode=disable"
	db, err = ConnectToDB(connStr)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestSaveRate(t *testing.T) {
	// Mock the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the Exec function
	mock.ExpectExec("INSERT INTO rates").WithArgs(75.0, 74.0, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the SaveRate function
	err = SaveRate(db, 75.0, 74.0, time.Now())

	// Assertions
	assert.NoError(t, err)
}
