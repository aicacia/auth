package repository

import (
	"time"

	"github.com/lib/pq"
)

type ResourceRowST struct {
	Id            int32          `db:"id"`
	ApplicationId int32          `db:"application_id"`
	Description   string         `db:"description"`
	URI           string         `db:"uri"`
	Actions       pq.StringArray `db:"actions"`
	UpdatedAt     time.Time      `db:"updated_at"`
	CreatedAt     time.Time      `db:"created_at"`
}

func GetResources(applicationId int32) ([]ResourceRowST, error) {
	return All[ResourceRowST](`SELECT r.*
		FROM resources r
		WHERE r.application_id = $1 
		ORDER BY r.updated_at DESC;`, applicationId)
}

func GetResourceById(id int32) ([]ResourceRowST, error) {
	return All[ResourceRowST](`SELECT r.*
		FROM resources r
		WHERE r.id = $1 
		LIMIT 1;`, id)
}

type CreateResourceST struct {
	ApplicationId int32          `db:"application_id"`
	Description   string         `db:"description"`
	URI           string         `db:"uri"`
	Actions       pq.StringArray `db:"actions"`
}

func CreateResource(params CreateResourceST) (ResourceRowST, error) {
	return NamedGet[ResourceRowST](`INSERT INTO resources (application_id, description, uri, actions)
		VALUES (:application_id, :description, :uri, :actions)
		RETURNING *;`, params)
}

type UpdateResourceST struct {
	Id          int32     `db:"id"`
	Description *string   `db:"description"`
	URI         *string   `db:"uri"`
	Actions     *[]string `db:"actions"`
}

func UpdateResource(params UpdateResourceST) (*ResourceRowST, error) {
	return NamedGetOptional[ResourceRowST](`UPDATE resources SET
		description=COALESCE(:description, description)
		uri=COALESCE(:uri, uri)
		actions=COALESCE(:id, actions)
	WHERE id=:id
	RETURNING *;`, params)
}
