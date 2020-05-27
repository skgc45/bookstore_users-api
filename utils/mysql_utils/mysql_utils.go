package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/skgc45/bookstore_users_api/utils/errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

// SQLError Handling
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewBadRequestError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062: // Duplicate entry
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
