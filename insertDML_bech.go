package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/loov/hrtime"
)

// Player game data type
type Player struct {
	UsersId  uint64 `json:"UsersId"`
	ServerId uint64 `json:"ServerId"`
	Power    uint64 `json:"Power"`
	Charm    uint64 `json:"Charm"`
	Money    uint64 `json:"Money"`
	intimacy uint64 `json:"intimacy"`
}

// Create the dataClients
func createClients(db string) (context.Context, *spanner.Client) {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		fmt.Println(err)
	}

	return ctx, client
}

// 创建用户数据
func genPlayerList() []Player {
	players := []Player{}
	num := 800000
	// i : ServerId
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 250; j++ {
			player := Player{uint64(num), uint64(i), uint64(rand.Intn(3000)),
				uint64(rand.Intn(3000)), uint64(rand.Intn(10000)),
				uint64(rand.Intn(5000))}
			num++
			// fmt.Println(num, player)
			players = append(players, player)
		}
	}
	// fmt.Println(players)
	return players
}

// GetSQLList 数据量太大分成多个sql语句并返回多个sql语句的切片
// groupSize 按照个人需求进行变更，代表 每groupSize个记录合并成1条INSERT语句
// INSERT Singers (SingerId, FirstName, LastName) VALUES (" + strconv.Itoa(num) + ",'Melissas', 'Garcia')`
func GetSQList(PlayerList []Player, insertHeader string) []string {
	sqlList := []string{}
	sql := ""
	for i := 0; i < len(PlayerList); i++ {
		if i%25 == 0 {
			if sql != "" {
				//把上次拼接的SQL结果存储起来
				sqlList = append(sqlList, sql)
			}
			//重置SQL
			sql = insertHeader
		}
		sql = fmt.Sprintf("%s(%d,%d,%d,%d,%d,%d),",
			sql,
			PlayerList[i].UsersId,
			PlayerList[i].ServerId,
			PlayerList[i].Power,
			PlayerList[i].Charm,
			PlayerList[i].Money,
			PlayerList[i].intimacy,
		)
	}

	//把最后一次生成的SQL存储起来
	sqlList = append(sqlList, sql)
	return sqlList
}

func writeBatchUsingDML(ctx context.Context, client *spanner.Client, sql string) error {

	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {

		// generation Users data SQL
		stmt := spanner.Statement{
			SQL: sql,
		}
		// rowCount, err := txn.Update(ctx, stmt)
		_, err := txn.Update(ctx, stmt)
		if err != nil {
			return err
		}
		// fmt.Fprintf(os.Stdout, "%d record(s) inserted.\n", rowCount)
		return err
	})
	return err
	// fmt.Println(err)
}

func main() {
	// const insertHeader string = "INSERT Users (UsersId,ServerId,Power,Charm,Money,intimacy) VALUES"
	const insertHeader string = "INSERT Users_index (UsersId,ServerId,Power,Charm,Money,intimacy) VALUES"

	// Database name String
	databaseName := "projects/yunion-test-286209/instances/test-instance/databases/users-db"

	ctx, client := createClients(databaseName)
	defer client.Close()

	fmt.Println("----- writeBatchUsingDML Test: -----")
	playersList := genPlayerList()
	sqls := GetSQList(playersList, insertHeader)

	flag := true
	const numberOfExperiments = 100
	bench := hrtime.NewBenchmark(numberOfExperiments)
	for i := 0; i <= len(sqls); i++ {
		// 判断是否停止benchmark
		if flag != true {
			break
		}
		// 避免array 超出索引
		if i != len(sqls) {
			s := strings.TrimRight(sqls[i], ",")
			err := writeBatchUsingDML(ctx, client, s)
			if err != nil {
				fmt.Println("Error info:", err)
				os.Exit(1)
			}
		}
		flag = bench.Next()
	}
	fmt.Println("Benchmark result:")
	fmt.Println(bench.Histogram(10))
}
