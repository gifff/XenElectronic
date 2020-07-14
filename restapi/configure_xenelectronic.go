// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gifff/xenelectronic/models"
	"github.com/gifff/xenelectronic/restapi/operations"
	"github.com/gifff/xenelectronic/restapi/operations/carts"
	"github.com/gifff/xenelectronic/restapi/operations/categories"
	"github.com/gifff/xenelectronic/restapi/operations/orders"
)

//go:generate swagger generate server --target ../../xenelectronic --name Xenelectronic --spec ../swagger.yml

func configureFlags(api *operations.XenelectronicAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Server Options",
			LongDescription:  "Server Options",
			Options:          &appConfig,
		},
	}
}

func configureAPI(api *operations.XenelectronicAPI) http.Handler {
	configureDependencies()

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CartsAddOneProductIntoCartHandler = carts.AddOneProductIntoCartHandlerFunc(func(params carts.AddOneProductIntoCartParams) middleware.Responder {
		cartItem, err := CartService.AddProductIntoCart(params.CartID.String(), *params.Body.ProductID)
		if err != nil {
			return carts.NewAddOneProductIntoCartDefault(500).WithPayload(formatError(1000, err))
		}

		payload := &models.CartItem{
			ID: cartItem.ID,
			Product: &models.Product{
				ID:          cartItem.Product.ID,
				CategoryID:  &cartItem.Product.CategoryID,
				Name:        &cartItem.Product.Name,
				Description: &cartItem.Product.Description,
				Photo:       cartItem.Product.Photo,
				Price:       &cartItem.Product.Price,
			},
			ProductID: &cartItem.Product.ID,
		}
		return carts.NewAddOneProductIntoCartOK().WithPayload(payload)
	})

	api.OrdersCheckoutHandler = orders.CheckoutHandlerFunc(func(params orders.CheckoutParams) middleware.Responder {
		order, err := OrderRepository.CheckoutFromCart(
			params.Body.CartID.String(),
			*params.Body.CustomerName,
			params.Body.CustomerEmail.String(),
			*params.Body.CustomerAddress,
		)
		if err != nil {
			return orders.NewCheckoutDefault(500).WithPayload(formatError(1000, err))
		}

		custEmail := strfmt.Email(order.CustomerEmail)
		payload := &models.Order{
			ID:                   strfmt.UUID(order.ID),
			CustomerName:         &order.CustomerName,
			CustomerEmail:        &custEmail,
			CustomerAddress:      &order.CustomerAddress,
			PaymentAmount:        order.PaymentAmount,
			PaymentMethod:        order.PaymentMethod,
			PaymentAccountNumber: order.PaymentAccountNumber,
			CartItems:            make(models.CartItems, len(order.CartItems)),
		}
		for i := range order.CartItems {
			p := &models.Product{
				ID:          order.CartItems[i].Product.ID,
				CategoryID:  &order.CartItems[i].Product.CategoryID,
				Name:        &order.CartItems[i].Product.Name,
				Description: &order.CartItems[i].Product.Description,
				Photo:       order.CartItems[i].Product.Photo,
				Price:       &order.CartItems[i].Product.Price,
			}

			payload.CartItems[i] = &models.CartItem{
				ID:        order.CartItems[i].ID,
				Product:   p,
				ProductID: &order.CartItems[i].Product.ID,
			}
		}
		return orders.NewCheckoutCreated().WithPayload(payload)
	})

	api.CartsCreateCartHandler = carts.CreateCartHandlerFunc(func(params carts.CreateCartParams) middleware.Responder {
		cartID, err := CartService.CreateCart()
		if err != nil {
			return carts.NewCreateCartDefault(500).WithPayload(formatError(1000, err))
		}

		cartIDUUID := strfmt.UUID(cartID)
		return carts.NewCreateCartOK().WithPayload(&carts.CreateCartOKBody{
			CartID: &cartIDUUID,
		})
	})

	api.CategoriesListCategoriesHandler = categories.ListCategoriesHandlerFunc(func(params categories.ListCategoriesParams) middleware.Responder {
		allCategories, err := CategoryService.ListAllCategories()
		if err != nil {
			return categories.NewListCategoriesDefault(500).WithPayload(formatError(1000, err))
		}

		payload := make([]*models.Category, len(allCategories))
		for i := range allCategories {
			payload[i] = &models.Category{
				ID:   allCategories[i].ID,
				Name: &allCategories[i].Name,
			}
		}

		return categories.NewListCategoriesOK().WithPayload(payload)
	})

	api.CartsListProductsInCartHandler = carts.ListProductsInCartHandlerFunc(func(params carts.ListProductsInCartParams) middleware.Responder {
		cartItems, err := CartService.ListProductsInCart(params.CartID.String())
		if err != nil {
			return carts.NewListProductsInCartDefault(500).WithPayload(formatError(1000, err))
		}

		payload := make([]*models.CartItem, len(cartItems))
		for i := range cartItems {
			p := &models.Product{
				ID:          cartItems[i].Product.ID,
				CategoryID:  &cartItems[i].Product.CategoryID,
				Name:        &cartItems[i].Product.Name,
				Description: &cartItems[i].Product.Description,
				Photo:       cartItems[i].Product.Photo,
				Price:       &cartItems[i].Product.Price,
			}

			payload[i] = &models.CartItem{
				ID:        cartItems[i].ID,
				Product:   p,
				ProductID: &cartItems[i].Product.ID,
			}
		}

		return carts.NewListProductsInCartOK().WithPayload(payload)
	})

	api.CategoriesListProductsOfCategoryHandler = categories.ListProductsOfCategoryHandlerFunc(func(params categories.ListProductsOfCategoryParams) middleware.Responder {

		products, err := CategoryService.ListProductsByCategoryID(params.CategoryID, *params.Since, *params.Limit)
		if err != nil {
			return categories.NewListProductsOfCategoryDefault(500).WithPayload(formatError(1000, err))
		}

		payload := make([]*models.Product, len(products))
		for i := range products {
			payload[i] = &models.Product{
				ID:          products[i].ID,
				CategoryID:  &products[i].CategoryID,
				Name:        &products[i].Name,
				Description: &products[i].Description,
				Photo:       products[i].Photo,
				Price:       &products[i].Price,
			}
		}

		return categories.NewListProductsOfCategoryOK().WithPayload(payload)
	})

	api.CartsRemoveOneProductFromCartHandler = carts.RemoveOneProductFromCartHandlerFunc(func(params carts.RemoveOneProductFromCartParams) middleware.Responder {
		err := CartService.RemoveProductFromCart(params.CartID.String(), params.ProductID)
		if err != nil {
			return carts.NewRemoveOneProductFromCartDefault(500).WithPayload(formatError(1000, err))
		}

		return carts.NewRemoveOneProductFromCartNoContent()
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
