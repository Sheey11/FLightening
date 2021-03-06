package sqlconn

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func FindAirlineByOriginAndDest(origin, dest, page int) (*sql.Rows, error) {
	_sql, _, _ := dialect.Select(
		goqu.I("airline.id").As("id"),
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"c1.id":             origin,
		"c2.id":             dest,
		"p1.city":           goqu.I("c1.id"),
		"p2.city":           goqu.I("c2.id"),
		"p1.id":             goqu.I("airline.origin"),
		"p2.id":             goqu.I("airline.destination"),
		"airline.affiliate": goqu.I("ac.id"),
		"airline.model":     goqu.I("am.id"),
	}).Limit(10).Offset(uint(page-1) * 10).Distinct().ToSQL()

	rows, err := db.Query(_sql)
	return rows, err
}

func FindAirlineByOrigin(origin, page int) (*sql.Rows, error) {
	_sql, _, _ := dialect.Select(
		goqu.I("airline.id").As("id"),
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"c1.id":             origin,
		"p1.city":           goqu.I("c1.id"),
		"p2.city":           goqu.I("c2.id"),
		"p1.id":             goqu.I("airline.origin"),
		"p2.id":             goqu.I("airline.destination"),
		"airline.affiliate": goqu.I("ac.id"),
		"airline.model":     goqu.I("am.id"),
	}).Limit(10).Offset(uint(page-1) * 10).Distinct().ToSQL()

	rows, err := db.Query(_sql)
	return rows, err
}

func FindAirlineByDest(dest, page int) (*sql.Rows, error) {
	_sql, _, _ := dialect.Select(
		goqu.I("airline.id").As("id"),
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"c2.id":             dest,
		"p1.city":           goqu.I("c1.id"),
		"p2.city":           goqu.I("c2.id"),
		"p1.id":             goqu.I("airline.origin"),
		"p2.id":             goqu.I("airline.destination"),
		"airline.affiliate": goqu.I("ac.id"),
		"airline.model":     goqu.I("am.id"),
	}).Limit(10).Offset(uint(page-1) * 10).Distinct().ToSQL()

	rows, err := db.Query(_sql)
	return rows, err
}

func FindAirlineById(id int) *sql.Row {
	_sql, _, _ := dialect.Select(
		goqu.I("airline.id").As("id"),
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"airline.id":        id,
		"p1.city":           goqu.I("c1.id"),
		"p2.city":           goqu.I("c2.id"),
		"p1.id":             goqu.I("airline.origin"),
		"p2.id":             goqu.I("airline.destination"),
		"airline.affiliate": goqu.I("ac.id"),
		"airline.model":     goqu.I("am.id"),
	}).ToSQL()

	row := db.QueryRow(_sql)
	return row
}

func FetchAllAirlines(page uint) (*sql.Rows, error) {
	_sql, _, _ := dialect.Select(
		goqu.I("airline.id").As("id"),
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"p1.city":           goqu.I("c1.id"),
		"p2.city":           goqu.I("c2.id"),
		"p1.id":             goqu.I("airline.origin"),
		"p2.id":             goqu.I("airline.destination"),
		"airline.affiliate": goqu.I("ac.id"),
		"airline.model":     goqu.I("am.id"),
	}).ToSQL()

	return db.Query(_sql)
}

func AddAirline(name string, affiliate, model, origin, destination int) error {
	_sql, _, _ := dialect.From("airlines").Insert().Rows(
		goqu.Record{
			"name":        name,
			"affiliate":   affiliate,
			"model":       model,
			"origin":      origin,
			"destination": destination,
		},
	).ToSQL()
	_, err := db.Exec(_sql)
	return err
}

func UpadteAirline(id int, name string, affiliate, model, origin, destination int) error {
	_sql, _, _ := dialect.From("airlines").Update().Set(
		goqu.Record{
			"name":        name,
			"affiliate":   affiliate,
			"model":       model,
			"origin":      origin,
			"destination": destination,
		},
	).ToSQL()
	_, err := db.Exec(_sql)
	return err
}
