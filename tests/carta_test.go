package tests

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/seambiz/carta"
	carta_mysql "github.com/seambiz/carta/tests/carta_msyql"
	"github.com/seambiz/seambiz/sdb"
	"github.com/stretchr/testify/require"
)

func queryMysql(rawSql string) (rows *sql.Rows) {
	var err error
	stmt, err := conn.Prepare(rawSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if rows, err = stmt.Query(); err != nil {
		log.Fatal(err)
	}
	return
}

func TestSelect(t *testing.T) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actors.actor_id`,")
	stmt.AppendStr("    a.first_name AS `actors.first_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY f.film_id ")
	stmt.AppendStr("LIMIT 100")

	type Dest struct {
		Film []struct {
			carta_mysql.Film
			Actors []carta_mysql.Actor `db:"actors"`
		} `db:"film"`
	}

	dest := Dest{}

	rows := queryMysql(stmt.Query())
	err := carta.Map(rows, &dest)
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, dest)
}

func TestSelectActor(t *testing.T) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actor_id`,")
	stmt.AppendStr("    a.first_name AS `first_name`, ")
	stmt.AppendStr("    a.last_name AS `last_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY a.actor_id ")
	stmt.AppendStr("LIMIT 100")

	type Dest struct {
		ActorID   uint   `json:"actor_id" db:"actor_id"`
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name" db:"last_name"`

		Film []carta_mysql.Film
	}

	dest := []Dest{}

	rows := queryMysql(stmt.Query())
	err := carta.Map(rows, &dest)
	if err != nil {
		t.Fatal(err)
	}

	snaps.MatchSnapshot(t, dest)
}

func TestJetSelectActor(t *testing.T) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actor.actor_id`,")
	stmt.AppendStr("    a.first_name AS `actor.first_name`, ")
	stmt.AppendStr("    a.last_name AS `actor.last_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY a.actor_id ")
	stmt.AppendStr("LIMIT 100")

	var dest []struct {
		carta_mysql.Actor

		Film []carta_mysql.Film
	}

	sql := mysql.RawStatement(stmt.Query())

	err := sql.Query(conn, &dest)
	require.NoError(t, err)

	snaps.MatchSnapshot(t, dest)
}

func BenchmarkJetSelectActor(t *testing.B) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actor.actor_id`,")
	stmt.AppendStr("    a.first_name AS `actor.first_name`, ")
	stmt.AppendStr("    a.last_name AS `actor.last_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY a.actor_id ")
	stmt.AppendStr("LIMIT 100")

	for i := 0; i < t.N; i++ {
		var dest []struct {
			carta_mysql.Actor

			Film []carta_mysql.Film
		}

		sql := mysql.RawStatement(stmt.String())

		err := sql.Query(conn, &dest)
		require.NoError(t, err)
	}
}

func BenchmarkSelectActor(b *testing.B) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actor_id`,")
	stmt.AppendStr("    a.first_name AS `first_name`, ")
	stmt.AppendStr("    a.last_name AS `last_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY a.actor_id ")
	stmt.AppendStr("LIMIT 100")

	for i := 0; i < b.N; i++ {
		var dest []struct {
			ActorID   uint   `json:"actor_id" db:"actor_id"`
			FirstName string `json:"first_name" db:"first_name"`
			LastName  string `json:"last_name" db:"last_name"`

			Film []carta_mysql.Film
		}

		rows := queryMysql(stmt.String())
		err := carta.Map(rows, &dest)
		require.NoError(b, err)
	}
}

func BenchmarkSelectActorOptimal(b *testing.B) {
	stmt := sdb.NewSQLStatement()
	stmt.AppendStr("SELECT ")
	stmt.AppendStr("    f.film_id AS `film.film_id`,")
	stmt.AppendStr("    f.title AS `film.title`,")
	stmt.AppendStr("    f.description AS `film.description`,")
	stmt.AppendStr("    f.release_year AS `film.release_year`,")
	stmt.AppendStr("    a.actor_id AS `actor_id`,")
	stmt.AppendStr("    a.first_name AS `first_name`, ")
	stmt.AppendStr("    a.last_name AS `last_name` ")
	stmt.AppendStr("FROM sakila.film f ")
	stmt.AppendStr("INNER JOIN sakila.film_actor fa ON f.film_id = fa.film_id ")
	stmt.AppendStr("INNER JOIN sakila.actor a ON fa.actor_id = a.actor_id ")
	stmt.AppendStr("ORDER BY a.actor_id ")
	stmt.AppendStr("LIMIT 100")

	var dest []struct {
		ActorID   uint   `json:"actor_id" db:"actor_id"`
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name" db:"last_name"`

		Film []carta_mysql.Film
	}

	actors := make(map[uint]bool)

	for i := 0; i < b.N; i++ {
		rows := queryMysql(stmt.String())
		for rows.Next() {
			var film carta_mysql.Film

			var actorID uint
			var firstName string
			var lastName string

			err := rows.Scan(&film.FilmID, &film.Title, &film.Description, &film.ReleaseYear, &actorID, &firstName, &lastName)
			require.NoError(b, err)

			if _, ok := actors[actorID]; !ok {
				actors[actorID] = true

				dest = append(dest, struct {
					ActorID   uint   `json:"actor_id" db:"actor_id"`
					FirstName string `json:"first_name" db:"first_name"`
					LastName  string `json:"last_name" db:"last_name"`

					Film []carta_mysql.Film
				}{
					ActorID:   actorID,
					FirstName: firstName,
					LastName:  lastName,
				})
			}
			dest[len(dest)-1].Film = append(dest[len(dest)-1].Film, film)
		}
	}
}
