package main

import (
	"cephalon-ent/common"
	"cephalon-ent/internal/db"
	"cephalon-ent/pkg/cep_ent/user"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	//_, err := db.DB.User.Create().SetPhone("18925537107").Save(ctx)
	//if err != nil {
	//	logrus.Errorf("%+v", err)
	//	return
	//}

	users, err := db.DB.User.Query().Where(user.DeletedAt(common.ZeroTime)).All(ctx)
	if err != nil {
		logrus.Errorf("%+v", err)
	}
	fmt.Printf("%+v", users)
}
