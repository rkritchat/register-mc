package abs

import "database/sql"

type ServiceInterface interface {
	ValidateRequestMsg() error
	ValidateBusinessRule(db *sql.DB) error
	Execute(db *sql.DB) (string, error)
}

type Db interface {
	save()
	update()
	delete()
}