package coupon

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldDeleteCoupon(t *testing.T) {
	test := Coupon{
		Name:      "free food",
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

	// Prepare mock database with SQL command to be executed
	queryDelete := "delete from coupons where name=?"
	mock.ExpectPrepare(queryDelete).
		ExpectExec().WithArgs(test.Name).
		WillReturnResult(sqlmock.NewResult(0, 1)) // no inserted id, 1 affected row

	// Test actual function with mock database
	if err = Delete(db, "name", test.Name); err != nil {
		t.Errorf("Expected err = nil Got %v", err)
	}

	// Make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteCouponByValue(t *testing.T) {
	test := Coupon{
		Name:      "free food",
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

	// Prepare mock database with SQL command to be executed
	queryDelete := "delete from coupons where value=?"
	mock.ExpectPrepare(queryDelete).
		ExpectExec().WithArgs(test.Value).
		WillReturnResult(sqlmock.NewResult(0, 1)) // no inserted id, 1 affected row

	// Test actual function with mock database
	if err = DeleteByValue(db, "value", test.Value); err != nil {
		t.Errorf("Expected err = nil Got %v", err)
	}

	// Make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
