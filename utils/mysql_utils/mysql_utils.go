package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/skgc45/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

// SQLError Handling
func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewBadRequestError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062: // Duplicate entry
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request")
}
