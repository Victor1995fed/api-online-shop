package sqlstore

import (
	"api-online-store/internal/app/model"
	"api-online-store/internal/app/store"
	"database/sql"
)

//ProductRepository ...
type ProductRepository struct {
	store *Store
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) error {
	if err := p.Validate(); err != nil {
		return err
	}
	// if err := u.BeforeCreate(); err != nil {
	// 	return err
	// }
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
		"SELECT id, email, encrypted_password FROM users WHERE id= $1",
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
func (r *ProductRepository) Update(id int) (*model.Product, error) {
	p := &model.Product{}
	//...
	return p, nil
}

// Delete ...
func (r *ProductRepository) Delete(id int) (bool, error) {
	// p := &model.Product{}
	//...
	return true, nil
}
