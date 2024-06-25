package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
//	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/t|ools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func ListAddressQuery(User string) ([]entities.Address, error) {
	fmt.Println("Comienza ListAddressQuery")
	fmt.Println("User: ")
	fmt.Println(User)

	//Creamos la variable que almacenara cada registro devuelto de cada categoria de la base de datos.
	var Addr []entities.Address

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return Addr, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses"
	sentencia +=" WHERE Add_Id = '" + User + "'"

	var rows *sql.Rows

	rows, err = secundary.Db.Query(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return Addr, err
	}

	defer rows.Close()

	for rows.Next() {
		var a entities.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err:= rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)

		if err != nil {
			fmt.Println(err.Error())
			return Addr, err
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String

		Addr = append(Addr, a)
	}

	fmt.Println("Select Addresses > Ejecución Exitosa")
	fmt.Println("Cambiar texto")

	return Addr, nil
}