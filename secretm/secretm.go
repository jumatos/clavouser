package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/alejandroM/Clavo-Ecommerce/clavouser/awsgo"
	"github.com/alejandroM/Clavo-Ecommerce/clavouser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJSon, error) {
	var datosSecret models.SecretRDSJSon
	fmt.Println(" > Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)

	fmt.Println(" > Lectura Secret OK " + nombreSecret)
	return datosSecret, nil
}
