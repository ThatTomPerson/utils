package main

import (
	"fmt"
	"os"

	"github.com/ThatTomPerson/utils/smdb"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 3 {
		logrus.Fatal("smdb <secret> <database>")
	}

	c, err := smdb.GetCreds(os.Args[1])
	if err != nil {
		logrus.Fatalf("smdb <secret> <database>\n%v", err)
	}

	fmt.Printf("%s://%s:%s@%s/%s", c.Engine, c.Username, c.Password, c.Host, os.Args[2])
}
