package bd

import (
	"os"

	"github.com/demian63s/gambituser/models"
	"github.com/demian63s/gambituser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
