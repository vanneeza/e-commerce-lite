package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	logincontroller "github.com/vanneeza/e-commerce-lite/api/controller/login"
	productcontroller "github.com/vanneeza/e-commerce-lite/api/controller/product"
	stockcontroller "github.com/vanneeza/e-commerce-lite/api/controller/stock"
	storecontroller "github.com/vanneeza/e-commerce-lite/api/controller/store"
	transactioncontroller "github.com/vanneeza/e-commerce-lite/api/controller/transaction"
	"github.com/vanneeza/e-commerce-lite/internal/domain"
	productrepository "github.com/vanneeza/e-commerce-lite/internal/repository/product"
	stockrepository "github.com/vanneeza/e-commerce-lite/internal/repository/stock"
	storerepository "github.com/vanneeza/e-commerce-lite/internal/repository/store"
	transactionrepository "github.com/vanneeza/e-commerce-lite/internal/repository/transaction"
	productservice "github.com/vanneeza/e-commerce-lite/internal/service/product"
	stockservice "github.com/vanneeza/e-commerce-lite/internal/service/stock"
	storeservice "github.com/vanneeza/e-commerce-lite/internal/service/store"
	transactionservice "github.com/vanneeza/e-commerce-lite/internal/service/transaction"
	"github.com/vanneeza/e-commerce-lite/utils/middleware"
	"github.com/xendit/xendit-go"

	customercontroller "github.com/vanneeza/e-commerce-lite/api/controller/customer"
	customerrepository "github.com/vanneeza/e-commerce-lite/internal/repository/customer"
	customerservice "github.com/vanneeza/e-commerce-lite/internal/service/customer"
)

func Run(db *sql.DB, jwtKey, xenditKey string) *gin.Engine {
	xendit.Opt.SecretKey = xenditKey

	r := gin.Default()

	storeRepository := storerepository.NewStoreRepository()
	storeService := storeservice.NewStoreService(db, storeRepository)
	storeController := storecontroller.NewStoreController(storeService)

	customerRepository := customerrepository.NewCustomerRepository()
	customerService := customerservice.NewCustomerService(db, customerRepository)
	customerController := customercontroller.NewCustomerController(customerService)

	productRepository := productrepository.NewProductRepository()
	productDomain := domain.ProductDomain{
		ProductRepository: productRepository,
		StoreRepository:   storeRepository,
	}

	productService := productservice.NewProductService(db, productDomain)
	productController := productcontroller.NewProductController(productService)

	stockRepository := stockrepository.NewStockRepository()
	stockDomain := domain.StockDomain{
		StockRepository:   stockRepository,
		ProductRepository: productRepository,
		StoreRepository:   storeRepository,
	}

	stockService := stockservice.NewStockService(db, stockDomain)
	stockController := stockcontroller.NewStockController(stockService)

	transactionRepository := transactionrepository.NewTransactionRepository()
	txRepository := domain.TxRepository{
		TxRepository:       transactionRepository,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
		StoreRepository:    storeRepository,
	}

	transactionService := transactionservice.NewTransactionService(db, txRepository)
	transactionController := transactioncontroller.NewTransactionController(transactionService)

	loginContoller := logincontroller.NewLoginController(customerService)
	eCommerce := r.Group("e-commerce/v1/")
	eCommerce.POST("xendit-callback", transactionController.XenditCallback)
	eCommerce.POST("login", loginContoller.Login)

	store := eCommerce.Group("stores")
	{
		store.POST("/", storeController.Register)
		store.GET("/", storeController.GetAll)
		store.GET("/:id", storeController.GetById)
		store.PUT("/:id", storeController.Edit)
		store.DELETE("/:id", storeController.Unreg)
	}

	product := eCommerce.Group("products")
	{
		product.POST("/", productController.Register)
		product.GET("/", productController.GetAll)
		product.GET("/:id", productController.GetById)
		product.PUT("/:id", productController.Edit)
		product.DELETE("/:id", productController.Unreg)
	}

	customer := eCommerce.Group("customers")
	customer.Use(middleware.AuthMiddleware(jwtKey))
	{

		customer.POST("/", customerController.Register)
		customer.GET("/", customerController.GetAll)
		customer.GET("/:id", customerController.GetById)
		customer.PUT("/:id", customerController.Edit)
		customer.DELETE("/:id", customerController.Unreg)

		customer.POST("/transaction", transactionController.MakeOrder)

	}

	stock := eCommerce.Group("stocks")
	// stock.Use(middleware.AuthMiddleware(jwtKey))
	{

		stock.POST("/", stockController.Register)
		stock.GET("/", stockController.GetAll)
		stock.GET("/:id", stockController.GetById)
		stock.PUT("/:id", stockController.Edit)
		stock.DELETE("/:id", stockController.Unreg)

	}

	return r
}
