// Estructuras para category
package entities

//Estructura para direcciones

type Address struct {
	AddId 			int  	`json:"addId"`
	AddTitle    	string  `json:"addTitle"`
	AddName			string  `json:"addName"`
	AddAddress     	string  `json:"addAddress"`
	AddCity      	string  `json:"addCity"`
	AddState     	string  `json:"addState"`
	AddPostalCode	string  `json:"addPostalCode"`
	AddPhone      	string  `json:"addPhone"`
}


//Estructura para products

type Product struct {
	ProdId         int     `json:"prodID"`
	ProdTitle      string  `json:"prodTitle"`
	ProdDescrition string  `json:"prodDescription"`
	ProdCreateAt   string  `json:"prodCreateAt"`
	ProdUpdated    string  `json:"prodUpdated"`
	ProdPrice      float64 `json:"prodPrice,omitempty"`
	ProdStock      int     `json:"prodStock"`
	ProdCategId    int     `json:"prodCategId"`
	ProdPath       string  `json:"prodPath"`
	ProdSearch     string  `json:"search,omitempty"`
	ProdCategPath  string  `json:"categPath,omitempty"`
}

//Nuevo modelo para Select Productos

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"` //Nuestro slice (arreglo) de productos
}
