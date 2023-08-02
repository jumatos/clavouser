package bd

import (
	"fmt"

	"github.com/alejandroM/Clavo-Ecommerce/clavouser/models"
	"github.com/alejandroM/Clavo-Ecommerce/clavouser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID,User_DateAdd ) values ('" + sig.USerEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ejecucion Exitosa")
	return nil
}
