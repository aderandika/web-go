package categorymodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * from categories`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	// Looping
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories

}

// Fungsi
func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at)
		VALUE (?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInserId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInserId > 0
}

// Model Data Detail
func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

// Model Data Update
func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

// Methode Delete
func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM  categories WHERE id = ?`, id)
	return err
}
