package sqlstore

import (
	"api-online-store/internal/app/model"
	"database/sql"
)

//ProductRepository ...
type ProductTagRepository struct {
	store *Store
	tx    *sql.Tx
}

// Add ...
func (r *ProductTagRepository) Add(p *model.ProductTag) error {
	return r.tx.QueryRow(
		"INSERT INTO "+p.GetTableName()+" (product_id, tag_id ) VALUES ($1, $2) RETURNING id",
		p.ProductId,
		p.TagId,
	).Scan(&p.ID)

}

// Delete ...
func (r *ProductTagRepository) UnlinkProduct(p *model.ProductTag) error {
	_, err := r.tx.Exec(
		"DELETE FROM "+p.GetTableName()+" WHERE product_id = $1",
		p.ProductId,
	)
	return err
}
