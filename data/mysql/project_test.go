// +build integration

package mysql

import (
	"boyter/portfold/data"
	"testing"
)

func TestProjectGet(t *testing.T) {
	project := ProjectModel{DB: connect(t)}
	_, err := project.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}

func TestProjectInsertGet(t *testing.T) {
	projectModel := ProjectModel{DB: connect(t)}

	zeproject, _ := projectModel.Insert(data.Project{
		Name: "sample name",
	})

	project, err := projectModel.Get(zeproject.Id)

	if err != nil {
		t.Error("Not expecting error")
	}

	if project.Id != zeproject.Id {
		t.Error("Expected id to match")
	}

	projectModel.Delete(*project)
}
