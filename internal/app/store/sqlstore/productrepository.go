package sqlstore

import (
	"api-online-store/internal/app/filter"
	"api-online-store/internal/app/model"
	"api-online-store/internal/app/store"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	ctx := context.Background()
	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = tx.QueryRow(
		"INSERT INTO "+p.GetTableName()+" (title, description, price ) VALUES ($1, $2, $3) RETURNING id",
		p.Title,
		p.Description,
		p.Price,
	).Scan(&p.ID)

	if err != nil {
		tx.Rollback()
		return err
	}
	if len(p.Tags) > 0 {
		//FIXME:: Требуется удалять дублирующиеся id
		for _, s := range p.Tags {
			pt := model.ProductTag{}
			pt.ProductId = p.ID
			pt.TagId = s.ID
			ptr := ProductTagRepository{store: r.store, tx: tx}
			err = ptr.Add(&pt)
			if err != nil {
				tx.Rollback()
				return err
			}

		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
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
		p.ID,
	).Scan(&p.ID)
}

//List ...
func (r *ProductRepository) List(filter *filter.Product) ([]model.Product, error) {
	p := []model.Product{}
	var productModel *model.Product
	filterSql := filter.Apply(filter)
	pt := productModel.GetTableName()
	vp := productModel.GetViewTags()
	jsonFunc := "json_agg(json_build_object('id', " + vp + ".tag_id, 'title', " + vp + ".tag_title )) as tags"
	sql := "SELECT " + pt + ".*, " + jsonFunc + " FROM " +
		pt + " LEFT JOIN " + vp + " ON " + pt + ".id = " + vp + ".id " + filterSql
	fmt.Println(sql)
	rows, err := r.store.db.Query(sql)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		var pr model.Product
		var tagsData string
		var arrTags []model.Tag
		err := rows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Price, &tagsData)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(tagsData), &arrTags)
		if err != nil {
			log.Fatal(err)
		}
		pr.Tags = arrTags
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
