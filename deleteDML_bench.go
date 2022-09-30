package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"cloud.google.com/go/spanner"
	"github.com/loov/hrtime"
)

//  Create the dataClients
func createClients(db string) (context.Context, *spanner.Client) {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		fmt.Println(err)
	}

	return ctx, client
}

// 批量删除2.5k 条记录
func deleteUsingDML(ctx context.Context, client *spanner.Client, count int, step int) error {

	s := "DELETE FROM Users_index WHERE UsersId >= " + strconv.Itoa(count-step) + " AND UsersId <= " + strconv.Itoa(count)
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{SQL: s}
		rowCount, err := txn.Update(ctx, stmt)
		// _, err := txn.Update(ctx, stmt)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%d record(s) deleted.\n", rowCount)
		return nil
	})
	return err
}

func main() {
	// 样本数量
	const numberOfExperiments = 100

	// Database name String
	databaseName := "projects/yunion-test-286209/instances/test-instance/databases/users-db"

	ctx, client := createClients(databaseName)
	defer client.Close()

	fmt.Println("----- DeleteDML Benchmark Test: -----")

	// 删除记录起始Id
	start := 802500
	// 删除记录 步长
	step := 25
	flag := true
	bench := hrtime.NewBenchmark(numberOfExperiments)
	for i := 0; i <= numberOfExperiments; i++ {
		// 判断是否停止benchmark
		if flag != true {
			break
		}
		// 避免array 超出索引
		if i != numberOfExperiments {
			err := deleteUsingDML(ctx, client, start, step)
			if err != nil {
				fmt.Printf("Error info:%s\n", err)
				os.Exit(1)
			}
			start -= step
		}
		flag = bench.Next()
	}
	fmt.Println("Benchmark result:")
	fmt.Println(bench.Histogram(10))
}
