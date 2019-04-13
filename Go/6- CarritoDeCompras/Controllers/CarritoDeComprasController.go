package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Product struct {
	Product int `json:"product"`
	Quantity int `json:"quantity"`
	User int `json:"user_id"`
}

type Sale struct {

	Sales []Product `json:sales`

}

type ProductSale struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	ImgPath string `json:"img_path"`
}
func main() {

	r := gin.Default()
	r.LoadHTMLGlob("Views/*.*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/buy",func( c *gin.Context){

		db, err := sql.Open("mysql", "root:12345@/carritodecompras")
		if err != nil {
			fmt.Println("error")
		}
		a := Sale{}

		b, err := c.GetRawData()
		if err != nil {
			panic(err)
		}
		json.Unmarshal(b, &a)
		fmt.Println("",len(a.Sales))
		for i := 0 ; i < len(a.Sales) ; i++ {

			insForm, err := db.Prepare("INSERT INTO sales(product, quantity, user_id ) VALUES(?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(a.Sales[i].Product, a.Sales[i].Quantity,a.Sales[i].User)

		}


	})
	r.GET("/products",func( c *gin.Context){

		data, error :=getJSON("SELECT * FROM product INNER JOIN warehouse ON product.id = warehouse.product_id WHERE warehouse.stock > 0")
		if error != nil {
			panic(error.Error())
		}


		c.JSON(200,gin.H{
			"code" : http.StatusOK,
			"products": data,// cast it to string before showing
		})

	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}

func getJSON(sqlString string) (string, error) {

	db, err := sql.Open("mysql", "root:12345@/carritodecompras")
	if err != nil {
		fmt.Println("error")
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}



