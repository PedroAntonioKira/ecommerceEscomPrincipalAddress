package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/repositories"
)

func AddAddressUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Use Case"

	status, response = repositories.AddAddressRepositories(body, user)

	return status, response
}