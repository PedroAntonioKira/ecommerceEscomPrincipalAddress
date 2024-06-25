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
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/utils"
)

func AddAddressDatabase(addr entities.Address, User string)  error {
	fmt.Println("Comienza Registro de Address")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()
	fmt.Println("Despues de la conexión")
	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name )"
	sentencia += " VALUES ('" + User + "','" + addr.AddAddress + "','" + addr.AddCity + "','" + addr.AddState + "','"
	sentencia += addr.AddPostalCode + "','" + addr.AddPhone + "','" + addr.AddTitle + "','" + addr.AddName +  "')" 

	fmt.Println(sentencia)

	//Ejecutamos la sentencia SQL
	_, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Insert Address > Ejecución Exitosa")

	return nil

}



