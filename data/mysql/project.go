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

func (m *ProjectModel) GetPaged(userId int, offset int, perPage int) ([]*data.Project, error) {
	stmt := `
		SELECT	id,
             	name,
				created,
				updated
		FROM	project
		ORDER BY updated, id DESC
		LIMIT 10
`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []*data.Project{}

	for rows.Next() {
		project := &data.Project{}

		err := rows.Scan(&project.Id, &project.Name, &project.Created, &project.Updated)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (m *ProjectModel) Get(id int) (*data.Project, error) {
	stmt := `
		SELECT	id,
             	name,
				created,
				updated
		FROM	project
		WHERE	id = ?
`
	row := m.DB.QueryRow(stmt, id)
	project := &data.Project{}

	err := row.Scan(&project.Id, &project.Name, &project.Created, &project.Updated)
	if err == sql.ErrNoRows {
		return nil, data.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return project, nil
}
