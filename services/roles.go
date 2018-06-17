package services

import (
	"database/sql"
)

// RoleService provides a service to retrieve roles
type RoleService struct {
	db *sql.DB
}

// Role represents the Hero's Role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetRoleService instantiates and returns a new RoleService
func GetRoleService(db *sql.DB) *RoleService {
	return &RoleService{db: db}
}

// GetRoles returns all the roles
func (rs *RoleService) GetRoles() (*[]Role, error) {
	getRolesStmt, err := rs.db.Prepare(`SELECT * FROM roles`)

	if err != nil {
		return nil, err
	}

	defer getRolesStmt.Close()

	rows, err := getRolesStmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []Role{}

	for rows.Next() {
		var r Role
		err := rows.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return &roles, nil
}
