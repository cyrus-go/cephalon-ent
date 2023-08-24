package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/internal/db"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

func main() {
	ctx := context.Background()

	//_, err := db.DB.User.Create().SetPhone("18925537107").Save(ctx)
	//if err != nil {
	//	logrus.Errorf("%+v", err)
	//	return
	//}

	users, err := db.DB.User.Query().Where(user.DeletedAt(common.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("%+v", err)
	}
	fmt.Printf("%+v\n", users)
}
