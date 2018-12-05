package data

import "database/sql"

type ProjectModel struct {
	DB *sql.DB
}

func (m *ProjectModel) Insert(project Project) (*Project, error) {
	return nil, nil
}

func (m *ProjectModel) Get(id int) (*Project, error) {
	return nil, nil
}
