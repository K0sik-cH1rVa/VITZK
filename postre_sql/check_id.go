package postre_sql

import (
	"context"
	"fmt"
	"time"
	"ychba/users"

	"github.com/jackc/pgx/v5"
)

func CheckID(ctx context.Context, conn *pgx.Conn) {
	tiki := time.NewTicker(5 * time.Minute)
	defer tiki.Stop()

	for range tiki.C {
		fmt.Println("Начинается игра")

		for _, u := range users.Users {
			err := InsertOrUpdate(ctx, conn, u.Id, u.Name, u.Age, u.WorkHours, u.Post)
			if err != nil {
				fmt.Println("Ошибка обработки ID", u.Id, err)
			}
		}
	}
}
