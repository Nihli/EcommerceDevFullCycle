package database

import (
	"database/sql"
	"ecommerceDevFullCycle/internal/entity"
)

// conexão com o banco
type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

// retorna ou a lista de categoria ou um erro. Se o erro tiver vazio, func rodou corretamente.
func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	//faz um select no db
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil { //nil == vazio
		return nil, err
	}
	//ultima instrução a ser rodada nessa func
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		//se tiver erro (erro n for vazio), retorna erro. Senão rows.Scan pega item a item e joga na var category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	//retorna categorias com o erro em branco
	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	//QueryRow consulta apenas uma linha
	//busca no banco e ja seta valor na variavel
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
