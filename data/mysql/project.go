package mysql

import (
	"boyter/portfold/data"
	"database/sql"
)

type ProjectModel struct {
	DB *sql.DB
}

func (m *ProjectModel) Insert(project data.Project) (*data.Project, error) {
	return nil, nil
}

func (m *ProjectModel) Get(id int) (*data.Project, error) {
	return nil, nil
}
