package smdb

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/sirupsen/logrus"
)

type Creds struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

func (c *Creds) DSL(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&sql_mode=ansi", c.Username, c.Password, c.Host, c.Port, database)
}

func New(secret, database string) (*xray.DB, error) {
	creds, err := GetCreds(secret)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Connected to %s:********@tcp(%s:%d)/%s", creds.Username, creds.Host, creds.Port, database)
	return xray.SQL(creds.Engine, creds.DSL(database))
}

func GetCreds(secret string) (*Creds, error) {
	ses := session.New()

	ses.Config.Region = aws.String("ap-southeast-2")

	svc := secretsmanager.New(ses)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secret),
	}

	result, err := svc.GetSecretValue(input)

	if err != nil {
		return nil, err
	}
	var c Creds

	err = json.Unmarshal([]byte(*result.SecretString), &c)

	return &c, err
}
