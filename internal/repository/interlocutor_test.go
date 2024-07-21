package repository_test

import (
	"AnonimousChat/internal/data"
	"AnonimousChat/internal/repository"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var columnsTable = []string{
	"tag",
	"source",
	"id",
	"count_connections",
	"sum_donation",
}

var (
	testError = errors.New("test error")

	emptyResult = sqlmock.NewResult(0, 0)
	emptyRows   = sqlmock.NewRows(columnsTable)
)

var testCases = []data.Interlocutor{
	{Tag: "example1", Source: "source1", ID: 1001, CountConnections: 5, SumDonation: 50},
	{Tag: "example2", Source: "source2", ID: 1002, CountConnections: 15, SumDonation: 100},
	{Tag: "example3", Source: "source3", ID: 1003, CountConnections: 20, SumDonation: 200},
	{Tag: "example4", Source: "source4", ID: 1004, CountConnections: 10, SumDonation: 150},
	{Tag: "example5", Source: "source5", ID: 1005, CountConnections: 25, SumDonation: 300},
	{Tag: "example6", Source: "source6", ID: 1006, CountConnections: 30, SumDonation: 400},
	{Tag: "example7", Source: "source7", ID: 1007, CountConnections: 35, SumDonation: 500},
	{Tag: "example8", Source: "source8", ID: 1008, CountConnections: 40, SumDonation: 600},
	{Tag: "example9", Source: "source9", ID: 1009, CountConnections: 45, SumDonation: 700},
	{Tag: "example10", Source: "source10", ID: 1010, CountConnections: 50, SumDonation: 800},
}

func TestErrorMigrationInterlocutorTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS interlocutor").
		WillReturnResult(emptyResult).
		WillReturnError(testError)

	repo := repository.NewInterlocutorRepository(db)

	err = repo.MigrationInterlocutorTable()
	if err != testError {
		t.Fatalf("Expected test error, got: %v", err)
	}
}

func TestSuccessfulMigrationInterlocutorTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS interlocutor").
		WillReturnResult(emptyResult).
		WillReturnError(nil)

	repo := repository.NewInterlocutorRepository(db)

	err = repo.MigrationInterlocutorTable()
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}
}

func TestSuccessfulRegistration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec("INSERT INTO interlocutor").
			WithArgs(
				test.Tag,
				test.Source,
				test.ID,
				test.CountConnections,
				test.SumDonation,
			).
			WillReturnResult(emptyResult).
			WillReturnError(nil)

		err = repo.Registration(test)
		if err != nil {
			t.Fatalf("Expected nil, got error: %v\ntest data: %v", err, test)
		}
	}
}

func TestErrorRegistration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec("INSERT INTO interlocutor").
			WithArgs(
				test.Tag,
				test.Source,
				test.ID,
				test.CountConnections,
				test.SumDonation,
			).
			WillReturnResult(emptyResult).
			WillReturnError(testError)

		err = repo.Registration(test)
		if err != testError {
			t.Fatalf("Expected test err, got: %v\ntest data: %v", err, test)
		}
	}
}

func TestErrorGetInterlocutor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectQuery("SELECT *").
			WithArgs(
				test.Tag,
			).
			WillReturnRows(emptyRows).
			WillReturnError(testError)

		_, err := repo.GetInterlocutor(test.Tag)
		if err != testError {
			t.Fatalf("Expected test error, got: %v\ntest data: %v", err, test)
		}
	}
}

func TestGetInterlocutor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	var rows *sqlmock.Rows
	for _, test := range testCases {
		rows = sqlmock.NewRows(columnsTable)
		rows.AddRow(
			test.Tag,
			test.Source,
			test.ID,
			test.CountConnections,
			test.SumDonation,
		)

		mock.ExpectQuery("SELECT *").
			WithArgs(
				test.Tag,
			).
			WillReturnRows(rows).
			WillReturnError(nil)

		user, err := repo.GetInterlocutor(test.Tag)
		if err != nil {
			t.Errorf("Expected nil, got error: %v\n test data: %v\n", err, test)
			continue
		}

		if test != user {
			t.Errorf("%v != %v\n", test, user)
		}
	}
}

func TestErrorAddCountConnections(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec(`UPDATE interlocutor SET count_connections`).
			WithArgs(
				test.Tag,
			).
			WillReturnResult(emptyResult).
			WillReturnError(testError)

		err = repo.AddCountConnections(test.Tag)
		if err != testError {
			t.Errorf("Expected test error, got: %v\ntest data: %v\n", err, test)
		}
	}
}

func TestSuccessfulAddCountConnections(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec(`UPDATE interlocutor SET count_connections`).
			WithArgs(
				test.Tag,
			).
			WillReturnResult(emptyResult).
			WillReturnError(nil)

		err = repo.AddCountConnections(test.Tag)
		if err != nil {
			t.Errorf("Expected nil, got: %v\ntest data: %v\n", err, test)
		}
	}
}

func TestErrorAddSumDonation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec(`UPDATE interlocutor SET sum_donation`).
			WithArgs(
				test.Tag,
				test.SumDonation,
			).
			WillReturnResult(emptyResult).
			WillReturnError(testError)

		err = repo.AddSumDonation(test.Tag, test.SumDonation)
		if err != testError {
			t.Errorf("Expected test error, got: %v\ntest data: %v\n", err, test)
		}
	}
}

func TestSuccessfulAddSumDonation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mock, error: %v", err)
	}

	repo := repository.NewInterlocutorRepository(db)

	for _, test := range testCases {
		mock.ExpectExec(`UPDATE interlocutor SET sum_donation`).
			WithArgs(
				test.Tag,
				test.SumDonation,
			).
			WillReturnResult(emptyResult).
			WillReturnError(nil)

		err = repo.AddSumDonation(test.Tag, test.SumDonation)
		if err != nil {
			t.Errorf("Expected nil, got: %v\ntest data: %v\n", err, test)
		}
	}
}
