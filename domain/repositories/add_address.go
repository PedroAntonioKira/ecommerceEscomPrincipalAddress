package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/queries_category"
)

func AddAddressRepositories(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Repositories"

	status, response = queries_category.AddAddressQuery(body, user)

	return status, response
}