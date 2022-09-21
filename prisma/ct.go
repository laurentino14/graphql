package prisma

import (
	"github.com/laurentino14/graphql/prisma/client"
)

func Ct() error {
	ct := client.NewClient()
	if err := ct.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := ct.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	return nil
}
