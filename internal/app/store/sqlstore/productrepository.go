package sqlstore

import (
	"api-online-store/internal/app/model"
	"api-online-store/internal/app/store"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

//ProductRepository ...
type ProductRepository struct {
	store *Store
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) error {
	if err := p.Validate(p.GetSupportedScenarioValidation()["CREATE"]); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO "+p.GetTableName()+" (title, description, price, image_url ) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
	).Scan(&p.ID)
}

// Find ...
func (r *ProductRepository) Find(id int) (*model.Product, error) {
	p := &model.Product{}
	if err := r.store.db.QueryRow(
		"SELECT id, title, description, price  FROM "+p.GetTableName()+" WHERE id= $1",
		id,
	).Scan(
		&p.ID,
		&p.Title,
		&p.Description,
		&p.Price,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return p, nil
}

// Update ...
func (r *ProductRepository) Update(p *model.Product) error {
	if err := p.Validate(p.GetSupportedScenarioValidation()["UPDATE"]); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"UPDATE "+p.GetTableName()+" SET title=$1, description=$2, price=$3, image_url=$4  WHERE id=$5 RETURNING id",
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.ID,
	).Scan(&p.ID)
}

//List ...
func (r *ProductRepository) List(m map[string]string) ([]model.Product, error) {
	p := []model.Product{}
	c, err := strconv.Atoi(m["count"])
	if err != nil {
		return p, err
	}
	if c > 100 {
		c = 100
	}
	var productModel *model.Product

	rows, err := r.store.db.Query("SELECT id, title, description, price  FROM "+productModel.GetTableName()+" LIMIT  $1", c)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var pr model.Product
		err := rows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Price)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pr)
		p = append(p, pr)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return p, nil
}

// Delete ...
func (r *ProductRepository) Delete(p *model.Product) error {
	if err := p.Validate(p.GetSupportedScenarioValidation()["DELETE"]); err != nil {
		return err
	}
	_, err := r.store.db.Exec(
		"DELETE FROM "+p.GetTableName()+" WHERE id=$1",
		p.ID,
	)
	return err
}
