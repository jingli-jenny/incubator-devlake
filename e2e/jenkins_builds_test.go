package e2e

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test should only run once main_test is complete and ready

type JenkinsBuild struct {
	Id int
}

func TestJenkinsBuilds(t *testing.T) {
	var jenkinsBuilds []JenkinsBuild
	db, err := InitializeDb()
	assert.Nil(t, err)
	if err != nil {
		log.Fatal(err)
	}
	sqlCommand := "SELECT id FROM jenkins_builds;"
	rows, err := db.Query(sqlCommand)
	if err != nil {
		fmt.Println("KEVIN >>> err", err)
	}
	assert.Nil(t, err)
	defer rows.Close()
	for rows.Next() {
		var jenkinsBuild JenkinsBuild
		if err := rows.Scan(&jenkinsBuild.Id); err != nil {
			panic(err)
		}
		jenkinsBuilds = append(jenkinsBuilds, jenkinsBuild)
	}
	assert.Equal(t, len(jenkinsBuilds), 2732)
}
