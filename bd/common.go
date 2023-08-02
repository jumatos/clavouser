package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alejandroM/Clavo-Ecommerce/clavouser/models"
	"github.com/alejandroM/Clavo-Ecommerce/clavouser/secretm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/aws/aws-sdk-go-v2/aws"
	//"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var SecretModel models.SecretRDSJSon
var err error
var Db *sql.DB

func ReadSecrect() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("Mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("conexion exitosa de la DB")
	return nil
}

func ConnStr(claves models.SecretRDSJSon) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "clavo"
	dsn := fmt.Sprintf("%s: %s@tcp(%s)/%s?allowClearTextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
