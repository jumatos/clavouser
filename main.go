package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/alejandroM/Clavo-Ecommerce/clavouser/awsgo"
	"github.com/alejandroM/Clavo-Ecommerce/clavouser/bd"
	"github.com/alejandroM/Clavo-Ecommerce/clavouser/models"
	"github.com/aws/aws-lambda-go/events"        // go get github.com/aws/aws-lambda-go/events
	lambda "github.com/aws/aws-lambda-go/lambda" // go get github.com/aws/aws-lambda-go/lambda
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. DEbe enviar 'SecretName'")
		err := errors.New("error en los parametros debe enviar SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.USerEmail = att
			fmt.Println("Email = " + datos.USerEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecrect()
	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
