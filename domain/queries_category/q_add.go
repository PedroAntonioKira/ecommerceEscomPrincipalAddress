package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"
//	"fmt"

	//importaciones externas (descargadas)
	//"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func AddAddressQuery(body string, User string) (int, string) {
	// Creamos la variable que tiene la estructura de todo lo que conforma a un producto.
	var t entities.Address

	// Decodificamos lo que nos viene en el endpoint (json) para guardarlo en la estructura.
	err := json.Unmarshal([]byte(body), &t)

	//Validamos que nos traigan cada uno de los parametros
	if err != nil{
		return 400 , "Error en los datos recibidos " +  err.Error()
	}

	if t.AddAddress == ""{
		return 400, "Debe de especificar el Address"
	}

	if t.AddName == ""{
		return 400, "Debe de especificar el Name"
	}

	if t.AddTitle == ""{
		return 400, "Debe de especificar el Title"
	}

	if t.AddCity == ""{
		return 400, "Debe de especificar el City"
	}

	if t.AddPhone == ""{
		return 400, "Debe de especificar el Phone"
	}

	if t.AddPostalCode == ""{
		return 400, "Debe de especificar el PostalCode"
	}

	err = database.AddAddressDatabase(t, User)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el registro del Address apra el ID de Usuario " + User + " > " + err.Error()
	}

	return 200, "Insert Address Ok"
}