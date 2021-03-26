package sqlstore

import (
	"api-online-store/internal/app/filter"
	"api-online-store/internal/app/model"
	"api-online-store/internal/app/store"
	"api-online-store/tools/helpers"
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
	ptr := ProductTagRepository{store: r.store, tx: tx}
	err = r.AddTags(p, &ptr, tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err != nil {
		tx.Rollback()
		return err
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
	pt := p.GetTableName()
	vp := p.GetViewTags()
	var tagsData string
	var arrTags []model.Tag
	jsonFunc := "json_agg(json_build_object('id', " + vp + ".tag_id, 'title', " + vp + ".tag_title )) as tags"
	if err := r.store.db.QueryRow(
		"SELECT "+pt+".*, "+jsonFunc+" FROM "+
			pt+" LEFT JOIN "+vp+" ON "+pt+".id = "+vp+".id WHERE "+pt+".id = $1  GROUP by "+pt+".id",
		id,
	).Scan(
		&p.ID,
		&p.Title,
		&p.Description,
		&p.Price,
		&tagsData,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	err := json.Unmarshal([]byte(tagsData), &arrTags)
	if err != nil {
		log.Fatal(err)
	}
	p.Tags = arrTags
	return p, nil
}

// Update ...
func (r *ProductRepository) Update(p *model.Product) error {
	if err := p.Validate(p.GetSupportedScenarioValidation()["UPDATE"]); err != nil {
		return err
	}

	ctx := context.Background()
	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = tx.QueryRow(
		"UPDATE "+p.GetTableName()+" SET title=$1, description=$2, price=$3  WHERE id=$4 RETURNING id",
		p.Title,
		p.Description,
		p.Price,
		p.ID,
	).Scan(&p.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	//Remove all tags
	pt := model.ProductTag{}
	ptr := ProductTagRepository{store: r.store, tx: tx}
	pt.ProductId = p.ID
	err = ptr.UnlinkProduct(&pt)
	if err != nil {
		tx.Rollback()
		return err
	}

	//Add new tags
	err = r.AddTags(p, &ptr, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
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

func (*ProductRepository) AddTags(p *model.Product, ptr *ProductTagRepository, tx *sql.Tx) error {
	if len(p.Tags) > 0 {
		ah := helpers.ArrayHelper{}
		var TagsId []int
		for _, s := range p.Tags {
			pt := model.ProductTag{}
			_, found := ah.Find(TagsId, s.ID)
			if found {
				continue
			}
			TagsId = append(TagsId, s.ID)
			pt.ProductId = p.ID
			pt.TagId = s.ID
			err := ptr.Add(&pt)
			if err != nil {
				//tx.Rollback()
				return err
			}

		}
	}
	return nil
}
