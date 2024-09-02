package dataservice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golangfiles/model"
	"net/http"
)

func CreateCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var customer model.Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {

		return err
	}

	getproduct := "Select quantity from product where product = ? "
	var pro int
	p := customer.Product
	err1 := db.QueryRow(getproduct, p).Scan(&pro)

	if pro >= customer.Quantity && customer.Quantity > 0 {

		query := "Insert INTo customer (name,phone,email,product,quantity) VALUES (?,?,?,?,?) "

		_, err := db.Exec(query, customer.Name, customer.Phone, customer.Email, customer.Product, customer.Quantity)
		if err != nil {
			return err
		}

		c := pro - customer.Quantity
		query2 := "Update product  SET quantity=? Where  product=?"

		_, err3 := db.Exec(query2, c, customer.Product)

		if err3 != nil {
			return err
		}

	} else {
		fmt.Println("Requested quantity is out of range")
	}

	if err1 != nil {
		return err1

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
	return nil
}
