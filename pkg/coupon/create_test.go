package coupon

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldCreateCoupon(t *testing.T) {
	test := Coupon{
		ID:        "free food",
		Brand:     "Mac",
		Value:     20.2,
		CreatedAt: "2019-10-13",
		Expiry:    "2019-10-14",
	}

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare mock database with SQL command executed in function
	query := "INSERT INTO coupons \\(name, brand, value, created, expiry\\) VALUES \\( \\?, \\?, \\?, \\?, \\? \\)"
	mock.ExpectPrepare(query).ExpectExec().WithArgs(test.Name, test.Brand, test.Value, test.CreatedAt, test.Expiry).WillReturnResult(sqlmock.NewResult(1, 1))

	// Test actual function with mock database
	if _, err := Create(db, test.Name, test.Brand, test.Value, test.CreatedAt, test.Expiry); err != nil {
		t.Errorf("Expected err = nil Got %v", err)
	}

	// Make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
