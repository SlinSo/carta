package tests

import (
	"testing"
	"time"

	"github.com/blockloop/scan"
	"github.com/gkampitakis/go-snaps/snaps"
	carta_mysql "github.com/seambiz/carta/tests/carta_msyql"
	"github.com/seambiz/seambiz/sdb"
	"github.com/stretchr/testify/require"
)

func TestBockLoopSelect(t *testing.T) {
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
		ActorID    uint      `json:"actor_id" db:"actor.actor_id" sql:"primary_key"`
		FirstName  string    `json:"first_name" db:"actor.first_name"`
		LastName   string    `json:"last_name" db:"actor.last_name"`
		LastUpdate time.Time `json:"last_update" db:"actor.last_update"`

		Film []carta_mysql.Film `db:"film"`
	}

	rows := queryMysql(stmt.Query())
	err := scan.RowsStrict(&dest, rows)
	require.NoError(t, err)

	snaps.MatchSnapshot(t, dest)
}
