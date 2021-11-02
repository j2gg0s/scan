package example

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	scan "github.com/j2gg0s/scan/dialect/mysql"
)

// nolint: govet
func ExampleMySQL() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/example?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PWD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
		))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	ctx := context.Background()
	for _, dest := range []interface{}{&Project{}, &map[string]interface{}{}} {
		rows, err := db.QueryContext(
			ctx,
			`SELECT * FROM project WHERE org = ? LIMIT ?`,
			"uptrace", 1,
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := scan.Scan(rows, dest); err != nil {
			log.Fatal(err)
		}
		printProject(dest)
	}

	for _, dest := range []interface{}{&[]Project{}, &[]map[string]interface{}{}} {
		rows, err := db.QueryContext(
			ctx,
			`SELECT * FROM project WHERE org = ? ORDER BY project`,
			"kubernetes",
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := scan.Scan(rows, dest); err != nil {
			log.Fatal(err)
		}
		printProjects(dest)
	}

	// output:
	// uptrace bun 374 22
	// uptrace bun 374 22
	// kubernetes client-go 5300 2100
	// kubernetes example 4500 3400
	// kubernetes kubernetes 82200 30100
	// kubernetes client-go 5300 2100
	// kubernetes example 4500 3400
	// kubernetes kubernetes 82200 30100
}
