package carta_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/seambiz/carta"
	td "github.com/seambiz/carta/testdata"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	// "github.com/golang/protobuf/proto"
)

var (
	update = false
	initDB = true
)

var ctx context.Context

// DBConnection is the connection string to use for testing
const dsnURI = "file:test1?mode=memory&cache=shared"

// NewMigratedDB returns a new connection to a migrated database
func NewMigratedDB(provider string, connection string, models ...interface{}) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Generate test data before running tests
// Start local server with bufconn
func setup() {
	ctx = context.Background()
}

func TestMain(m *testing.M) {
	updatePtr := flag.Bool("update", false, "update the golden file, results are always considered correct")
	initdbPtr := flag.Bool("initdb", false, "initialize and populate testing database")
	flag.Parse()
	update = *updatePtr
	initDB = *initdbPtr
	setup()
	code := m.Run()
	goldenFile := "testdata/mapper.golden"
	if update {
		// update golden file
		updateGoldenFile(goldenFile)
	} else {
		// compare existing results
		compareResults(goldenFile)
	}
	os.Exit(code)
}

func createDatabase(conn *sql.DB) {
	dump, err := os.ReadFile("testdata/sql/dump.sql")
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(string(dump))
	if err != nil {
		panic(err)
	}
}

func updateGoldenFile(goldenFile string) {
	jsonResult := generateResultBytes()
	if err := ioutil.WriteFile(goldenFile, jsonResult, 0o644); err != nil {
		log.Fatalln(err)
	}
}

func compareResults(goldenFile string) {
	goldenFileJson, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		log.Fatalln(err)
	}

	jsonResult := generateResultBytes()

	resultDiff := diff.New()
	d, err := resultDiff.Compare(goldenFileJson, jsonResult)
	if err != nil {
		log.Fatalln(err)
	}
	formatter := formatter.NewDeltaFormatter()
	diffString, err := formatter.Format(d)
	if diffString != "{}\n" {
		log.Println("Results Do Not Match Golden File, " +
			"if this is expecred result with go test with --update")
		log.Fatalln(diffString)
	}
}

func generateResultBytes() []byte {
	var jsonResult []byte
	if r, err := json.MarshalIndent(testResults, "", "    "); err != nil {
		log.Fatalln(err)
	} else {
		jsonResult = r
	}
	return jsonResult
}

func query(rawSql string) map[string]*sql.Rows {
	resp := map[string]*sql.Rows{}
	for dbName, db := range dbs {
		stmt, err := db.Prepare(rawSql)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		if rows, err := stmt.Query(); err != nil {
			log.Fatal(err)
		} else {
			resp[dbName] = rows
		}
	}
	return resp
}

func queryPG(rawSql string) (rows *sql.Rows) {
	var err error
	if rows, err = dbs[pg].Query(rawSql); err != nil {
		log.Fatal(err)
	}
	return
}

func queryMysql(rawSql string) (rows *sql.Rows) {
	var err error
	stmt, err := dbs[mysql].Prepare(rawSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if rows, err = stmt.Query(); err != nil {
		log.Fatal(err)
	}
	return
}

func TestBlog(m *testing.T) {
	ans := []byte{}
	for _, rows := range query(td.BlogQuery) {
		resp := []td.Blog{}
		if err := carta.Map(rows, &resp); err != nil {
			log.Fatal(err.Error())
		}
		e, _ := json.Marshal(resp)
		if len(ans) == 0 {
			ans = e
		} else if string(ans) != string(e) {
			log.Fatal(errors.New("Test Blog Produced Inconsistent Results"))
		}
		testResults["TestBlog"] = resp
	}
}

func TestNull(m *testing.T) {
	respPG := []td.NullTest{}
	if err := carta.Map(queryPG(td.NullQueryPG), &respPG); err != nil {
		log.Fatal(err.Error())
	}
	respMySQL := []td.NullTest{}
	if err := carta.Map(queryMysql(td.NullQueryMySql), &respMySQL); err != nil {
		log.Fatal(err.Error())
	}
	ansPG, _ := json.Marshal(respPG)
	ansMySQL, _ := json.Marshal(respMySQL)
	if string(ansPG) != string(ansMySQL) {
		log.Fatal(errors.New("Test Null Produced Inconsistent Results"))
	}
	testResults["TestNull"] = respPG
}

func TestNotNull(m *testing.T) {
	respPG := []td.NullTest{}
	if err := carta.Map(queryPG(td.NotNullQueryPG), &respPG); err != nil {
		log.Fatal(err.Error())
	}
	respMySQL := []td.NullTest{}
	if err := carta.Map(queryMysql(td.NotNullQueryMySQL), &respMySQL); err != nil {
		log.Fatal(err.Error())
	}
	ansPG, _ := json.Marshal(respPG)
	ansMySQL, _ := json.Marshal(respMySQL)
	if string(ansPG) != string(ansMySQL) {
		log.Println(string(ansMySQL))
		log.Fatal(errors.New("Test Not Null Produced Inconsistent Results"))
	}
	testResults["TestNotNull"] = respPG
}

func TestPGTypes(m *testing.T) {
	resp := []td.PGDTypes{}
	if err := carta.Map(queryPG(td.PGDTypesQuery), &resp); err != nil {
		log.Fatal(err.Error())
	}
	testResults["TestPGTypes"] = resp
}

func TestRelation(m *testing.T) {
	ans := []byte{}
	for _, rows := range query(td.RelationTestQuery) {
		resp := []td.RelationTest{}
		if err := carta.Map(rows, &resp); err != nil {
			log.Fatal(err.Error())
		}
		e, _ := json.Marshal(resp)
		if len(ans) == 0 {
			ans = e
		} else if string(ans) != string(e) {
			log.Fatal(errors.New("Test Blog Produced Inconsistent Results"))
		}
		testResults["TestRelation"] = resp
	}
}
