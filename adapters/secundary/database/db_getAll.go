package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
//	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func ListCategoryQuery(CategId int, Slug string) ([]entities.Product, error) {
	fmt.Println("Comienza SelectCategories")

	//Creamos la variable que almacenara cada registro devuelto de cada categoria de la base de datos.
	var Categ []entities.Product



	fmt.Println("Cambiar texto")
	fmt.Println("Select Categories > Ejecución Exitosa : Optimizado01")
	fmt.Println("Cambiar texto")

	return Categ, nil
}