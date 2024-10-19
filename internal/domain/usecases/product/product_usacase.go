package product
import(

    entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
    repository "github.com/StuardAP/clean_rest_golang/internal/domain/repositories"
)
type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) *ProductUseCase {
    return &ProductUseCase{productRepository: productRepository}
}

func (u *ProductUseCase) GetAllProducts() ([]entity.Product, error) {
	return u.productRepository.GetAllProducts()
}

func (u *ProductUseCase) GetProductByID(id string) (*entity.Product, error) {
	return u.productRepository.GetProductByID(id)
}

func (u *ProductUseCase) CreateProduct(name string, price float64) (*entity.Product, error) {
	product := &entity.Product{
		Name: name,
		Price: price,
	}
	if err := u.productRepository.CreateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}
