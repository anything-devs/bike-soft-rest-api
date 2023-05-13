package repositories

import (
	"fmt"

	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"gorm.io/gorm"
)

// ProductoRepositoryImpl es la implementación de ProductoRepository
type ProductoRepositoryImpl struct {
	DB *gorm.DB
}

// ProductoRepository es la interfaz de ProductoRepository
type ProductoRepository interface {
	CrearProducto(producto *models.Producto) error
	EliminarProducto(id string, producto *models.Producto) error
}

/*
* Constructor de ProductoRepository
* @param DB: conexión a la base de datos
* @return instancia de ProductoRepository
 */
func NewProductoRepository(DB *gorm.DB) ProductoRepository {
	return &ProductoRepositoryImpl{DB: DB}
}

/*
* Método para crear los productos
* @param producto: producto a crear
* @return error si existe alguno o nil si no existe
 */
func (r *ProductoRepositoryImpl) CrearProducto(producto *models.Producto) error {
	if err := r.DB.Where("codigo= ?", producto.Codigo).First(&producto).Error; err != nil {
		return err
	}
	if err := r.DB.Create(producto).Error; err != nil {
		return err
	}
	return nil
}

/*
* Método para eliminar los productos
* @param id: id del producto a eliminar
* @param producto: producto a eliminar
* @return error si existe alguno o nil si no existe
 */
func (r *ProductoRepositoryImpl) EliminarProducto(id string, producto *models.Producto) error {
	if err := r.DB.Where("id = ?", id).First(&producto).Error; err != nil {
		return err
	}
	if err := r.DB.Delete(&producto).Error; err != nil {
		return err
	}
	return nil
}

// MockProductoRepository es la implementación de ProductoRepository para pruebas
type MockProductoRepository struct {
	Productos                map[uint]*models.Producto
	ProductoCreado           *models.Producto
	ProductoEliminado        *models.Producto
	Err                      error
	llamadoProductoCreado    bool
	llamadoProductoEliminado bool
}

/*
* Constructor de ProductoRepository para pruebas
* @return instancia de ProductoRepository para pruebas
 */
func NewMockProductoRepository() *MockProductoRepository {
	return &MockProductoRepository{
		Productos: map[uint]*models.Producto{},
	}
}

/*
* Método para crear los productos para pruebas
* @param producto: producto a crear para pruebas
* @return error si existe alguno o nil si no existe para pruebas
 */
func (m *MockProductoRepository) CrearProducto(producto *models.Producto) error {
	if m.Err != nil {
		return m.Err
	}
	m.llamadoProductoCreado = true
	m.ProductoCreado = producto
	m.Productos[producto.ID] = producto
	return nil
}

/*
* Método para eliminar los productos para pruebas
* @param id: id del producto a eliminar para pruebas
* @param producto: producto a eliminar para pruebas
* @return error si existe alguno o nil si no existe para pruebas
 */
func (m *MockProductoRepository) EliminarProducto(id string, producto *models.Producto) error {
	if m.Err != nil {
		return m.Err
	}
	if _, ok := m.Productos[producto.ID]; !ok {
		return fmt.Errorf("producto no encontrado")
	}
	m.llamadoProductoEliminado = true
	m.ProductoEliminado = producto
	delete(m.Productos, producto.ID)
	return nil
}
