package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	serviceModel "github.com/woka20/sirclo-ecommerce-task/order/src/services/model"
	"github.com/woka20/sirclo-ecommerce-task/order/src/usecase"
	"github.com/woka20/sirclo-ecommerce-task/order/utils"

	"github.com/gorilla/mux"
)

// HttpOrderHandler model
type HttpOrderHandler struct {
	orderUseCase usecase.OrderUseCase
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpOrderHandler(orderUseCase usecase.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: orderUseCase}
}

// Me http handler function, for get Member by its ID from Authorization
func (h *HttpOrderHandler) Me() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("CustomerId")

		memberResult := <-h.orderUseCase.FindMemberByID(memberID)

		if memberResult.Error != nil {
			log.Printf("Error get Member = %s", memberResult.Error.Error())
			utils.JsonResponse(res, "Member not found", http.StatusInternalServerError)
			return
		}

		member, ok := memberResult.Result.(serviceModel.Customer)

		if !ok {
			utils.JsonResponse(res, "Result is not member", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, member, http.StatusOK)

	})
}

// GetProduct http handler function, for get product by ID
func (h *HttpOrderHandler) GetProduct() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("CustomerId")

		fmt.Println(memberID)

		paths := mux.Vars(req)
		productIDStr := paths["id"]
		productID, _ := strconv.Atoi(productIDStr)

		productResult := <-h.orderUseCase.FindProductByID(productID)

		if productResult.Error != nil {
			log.Printf("Error get Product = %s", productResult.Error.Error())
			utils.JsonResponse(res, "Cannot Get Product", http.StatusInternalServerError)
			return
		}

		product, ok := productResult.Result.(serviceModel.Product)

		if !ok {
			utils.JsonResponse(res, "Result is not product", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, product, http.StatusOK)

	})
}

// GetProducts http handler function, for get all products
func (h *HttpOrderHandler) GetProducts() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		memberID := req.Header.Get("CustomerId")

		fmt.Println(memberID)

		productResult := <-h.orderUseCase.FindProductAll()

		if productResult.Error != nil {
			log.Printf("Error get Product = %s", productResult.Error.Error())
			utils.JsonResponse(res, "Cannot Get Products", http.StatusInternalServerError)
			return
		}

		products, ok := productResult.Result.(serviceModel.Products)

		if !ok {
			utils.JsonResponse(res, "Result is not products", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, products, http.StatusOK)

	})
}
