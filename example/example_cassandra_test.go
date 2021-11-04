package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"

	"github.com/j2gg0s/scan/dialect/cassandra"
)

type Project struct {
	Project string
	Org     string
	Star    int
	Fork    int
}

func printProject(p interface{}) {
	if p, ok := p.(*Project); ok {
		fmt.Println(p.Org, p.Project, p.Star, p.Fork)
	}
	if p, ok := p.(*map[string]interface{}); ok {
		m := *p
		fmt.Println(m["org"], m["project"], m["star"], m["fork"])
	}
}

func printProjects(projects interface{}) {
	if projects, ok := projects.(*[]Project); ok {
		for _, p := range *projects {
			printProject(&p)
		}
	}

	if projects, ok := projects.(*[]map[string]interface{}); ok {
		for _, p := range *projects {
			printProject(&p)
		}
	}

	if values, ok := projects.([]interface{}); ok {
		orgs := values[0].(*[]string)
		projects := values[1].(*[]string)
		stars := values[2].(*[]int)
		forks := values[3].(*[]int)
		for i := 0; i < len(*orgs); i++ {
			fmt.Println((*orgs)[i], (*projects)[i], (*stars)[i], (*forks)[i])
		}
	}
}

// nolint: govet
func ExampleCassandra() {
	cluster := gocql.NewCluster(
		fmt.Sprintf("%s:%s", os.Getenv("CQLSH_HOST"), os.Getenv("CQLSH_PORT")))
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.One

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	ctx := context.Background()
	fmt.Println()

	for _, dest := range []interface{}{
		&Project{},
		&map[string]interface{}{},
	} {
		iter := session.Query(
			`SELECT * FROM project_by_org WHERE org = ? LIMIT ?`,
			"uptrace", 1,
		).WithContext(ctx).Iter()

		if err := cassandra.Scan(iter, dest); err != nil {
			log.Fatal(err)
		}
		printProject(dest)
	}

	for _, dest := range []interface{}{
		&[]Project{},
		&[]map[string]interface{}{},
	} {
		iter := session.Query(
			`SELECT * FROM project_by_org WHERE org = ?`,
			"kubernetes",
		).WithContext(ctx).Iter()

		if err := cassandra.Scan(iter, dest); err != nil {
			log.Fatal(err)
		}
		printProjects(dest)
	}

	for _, dest := range [][]interface{}{
		[]interface{}{&[]string{}, &[]string{}, &[]int{}, &[]int{}},
	} {
		iter := session.Query(
			`SELECT org, project, star, fork FROM project_by_org WHERE org = ?`,
			"kubernetes",
		).WithContext(ctx).Iter()

		if err := cassandra.Scan(iter, dest...); err != nil {
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
	// kubernetes client-go 5300 2100
	// kubernetes example 4500 3400
	// kubernetes kubernetes 82200 30100
}
