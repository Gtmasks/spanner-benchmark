package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/loov/hrtime"
	"google.golang.org/api/iterator"
)

// Create the dataClients
func createClients(db string) (context.Context, *spanner.Client) {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		fmt.Println(err)
	}

	return ctx, client
}

func query(ctx context.Context, client *spanner.Client, sql string) error {

	stmt := spanner.Statement{SQL: sql}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	// 遍历 iter 输出查询结果
	for {
		_, err := iter.Next()
		// row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		// 遍历输出结果
		// 	var UsersId, Power int64
		// 	if err := row.Columns(&UsersId, &Power); err != nil {
		// 		return err
		// 	}
		// 	fmt.Printf("%d %d\n", UsersId, Power)
	}
}

// get the live time leaderboard
func rankByServer(ctx context.Context, client *spanner.Client) error {

	rand.Seed(time.Now().Unix())
	var exampleServer = rand.Intn(10)
	stmt := spanner.Statement{
		SQL: `SELECT UsersId FROM Users
				WHERE ServerId = @serverInfo
				order by Power desc
				limit 10`,
		Params: map[string]interface{}{
			"serverInfo": exampleServer,
		},
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		// row, err := iter.Next()
		_, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		// var UsersId int64

		// if err := row.Columns(&UsersId); err != nil {
		// 	return err
		// }
		// fmt.Printf("UsersId: %d\n", UsersId)
	}

}

func main() {
	const numberOfExperiments = 100
	// 点查方式
	// sql := "SELECT UsersId,ServerId,Power,Charm,Money,intimacy FROM Users WHERE ServerId=1"
	// sql := "SELECT UsersId,ServerId,Power,Charm,Money,intimacy FROM Users_index WHERE  ServerId=1"

	// 按照ServerId查询排行榜单，
	// sql := "SELECT UsersId,Power FROM Users WHERE ServerId=1 ORDER BY Power DESC  LIMIT 10"
	// sql := "SELECT UsersId,Power FROM Users_index WHERE ServerId=11 ORDER BY Power DESC LIMIT 10"

	// Database name String
	databaseName := "projects/yunion-test-286209/instances/test-instance/databases/users-db"

	ctx, client := createClients(databaseName)
	defer client.Close()

	fmt.Println("----- QueryDML Benchmark Test: -----")
	bench := hrtime.NewBenchmark(numberOfExperiments)
	for bench.Next() {
		// query(ctx, client, sql)
		rankByServer(ctx, client)
	}
	fmt.Println(bench.Histogram(10))

}
