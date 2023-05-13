package repositories

import (
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"gorm.io/gorm"
)

type CategoriaRepositoryImpl struct {
	DB *gorm.DB
}

// CatergoriaRepository es la interfaz de CatergoriaRepository
type CategoriaRepository interface {
	GetCategorias() ([]models.Categoria, error)
}

/*
* Constructor de CatergoriaRepository
* @param DB: conexión a la base de datos
* @return instancia de CatergoriaRepository
 */
func NewCategoriaRepository(DB *gorm.DB) CategoriaRepository {
	return &CategoriaRepositoryImpl{DB: DB}
}

/*
* Método para obtener las categorias
* @return error si existe alguno o nil si no existe
 */
func (r *CategoriaRepositoryImpl) GetCategorias() ([]models.Categoria, error) {
	var categorias []models.Categoria
	if err := r.DB.Find(&categorias).Error; err != nil {
		return nil, err
	}
	return categorias, nil
}
