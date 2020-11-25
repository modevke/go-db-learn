package database

import(
	"database/sql"
	"context"

	_ "github.com/lib/pq"
)


func createDbConnection(ctx context.Context){

	// db, err := sql.Open("postgres", "user=postgres dbname=user_db sslmode=verify-full")
	// if err != nil{
	// 	panic(err)
	// }

}