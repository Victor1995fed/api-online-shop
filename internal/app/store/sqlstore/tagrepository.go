package sqlstore

import (
	"api-online-store/internal/app/model"
	"api-online-store/internal/app/store"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

//TagRepository ...
type TagRepository struct {
	store *Store
}

// Create ...
func (r *TagRepository) Create(m *model.Tag) error {
	if err := m.Validate(m.GetSupportedScenarioValidation()["CREATE"]); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO "+m.GetTableName()+" (title) VALUES ($1) RETURNING id",
		m.Title,
	).Scan(&m.ID)
}

// Find ...
func (r *TagRepository) Find(id int) (*model.Tag, error) {
	m := &model.Tag{}
	if err := r.store.db.QueryRow(
		"SELECT id, title FROM "+m.GetTableName()+" WHERE id= $1",
		id,
	).Scan(
		&m.ID,
		&m.Title,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return m, nil
}

// Update ...
func (r *TagRepository) Update(m *model.Tag) error {
	if err := m.Validate(m.GetSupportedScenarioValidation()["UPDATE"]); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"UPDATE "+m.GetTableName()+" SET title=$1 WHERE id=$2 RETURNING id",
		m.Title,
		m.ID,
	).Scan(&m.ID)
}

//List ...
func (r *TagRepository) List(m map[string]string) ([]model.Tag, error) {
	s := []model.Tag{}
	c, err := strconv.Atoi(m["count"])
	if err != nil {
		return s, err
	}
	if c > 100 {
		c = 100
	}
	var tagModel *model.Tag

	rows, err := r.store.db.Query("SELECT id, title  FROM "+tagModel.GetTableName()+" LIMIT  $1", c)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		var pr model.Tag
		err := rows.Scan(&pr.ID, &pr.Title)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pr)
		s = append(s, pr)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return s, nil
}

// Delete ...
func (r *TagRepository) Delete(m *model.Tag) error {
	if err := m.Validate(m.GetSupportedScenarioValidation()["DELETE"]); err != nil {
		return err
	}
	_, err := r.store.db.Exec(
		"DELETE FROM "+m.GetTableName()+" WHERE id=$1",
		m.ID,
	)
	return err
}
