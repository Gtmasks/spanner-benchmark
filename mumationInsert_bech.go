package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

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

func write(ctx context.Context, client *spanner.Client, table string, PlayerList []Player, i int) error {

	UsersColumns := []string{"UsersId", "ServerId", "Power", "Charm", "Money", "intimacy"}

	m := []*spanner.Mutation{
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i].UsersId, 10),
				strconv.FormatUint(PlayerList[i].ServerId, 10), strconv.FormatUint(PlayerList[i].Power, 10), strconv.FormatUint(PlayerList[i].Charm, 10),
				strconv.FormatUint(PlayerList[i].Money, 10), strconv.FormatUint(PlayerList[i].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+1].UsersId, 10),
				strconv.FormatUint(PlayerList[i+1].ServerId, 10), strconv.FormatUint(PlayerList[i+1].Power, 10), strconv.FormatUint(PlayerList[i+1].Charm, 10),
				strconv.FormatUint(PlayerList[i+1].Money, 10), strconv.FormatUint(PlayerList[i+1].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+2].UsersId, 10),
				strconv.FormatUint(PlayerList[i+2].ServerId, 10), strconv.FormatUint(PlayerList[i+2].Power, 10), strconv.FormatUint(PlayerList[i+2].Charm, 10),
				strconv.FormatUint(PlayerList[i+2].Money, 10), strconv.FormatUint(PlayerList[i+2].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+3].UsersId, 10),
				strconv.FormatUint(PlayerList[i+3].ServerId, 10), strconv.FormatUint(PlayerList[i+3].Power, 10), strconv.FormatUint(PlayerList[i+3].Charm, 10),
				strconv.FormatUint(PlayerList[i+3].Money, 10), strconv.FormatUint(PlayerList[i+3].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+4].UsersId, 10),
				strconv.FormatUint(PlayerList[i+4].ServerId, 10), strconv.FormatUint(PlayerList[i+4].Power, 10), strconv.FormatUint(PlayerList[i+4].Charm, 10),
				strconv.FormatUint(PlayerList[i+4].Money, 10), strconv.FormatUint(PlayerList[i+4].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+5].UsersId, 10),
				strconv.FormatUint(PlayerList[i+5].ServerId, 10), strconv.FormatUint(PlayerList[i+5].Power, 10), strconv.FormatUint(PlayerList[i+5].Charm, 10),
				strconv.FormatUint(PlayerList[i+5].Money, 10), strconv.FormatUint(PlayerList[i+5].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+6].UsersId, 10),
				strconv.FormatUint(PlayerList[i+6].ServerId, 10), strconv.FormatUint(PlayerList[i+6].Power, 10), strconv.FormatUint(PlayerList[i+6].Charm, 10),
				strconv.FormatUint(PlayerList[i+6].Money, 10), strconv.FormatUint(PlayerList[i+6].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+7].UsersId, 10),
				strconv.FormatUint(PlayerList[i+7].ServerId, 10), strconv.FormatUint(PlayerList[i+7].Power, 10), strconv.FormatUint(PlayerList[i+7].Charm, 10),
				strconv.FormatUint(PlayerList[i+7].Money, 10), strconv.FormatUint(PlayerList[i+7].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+8].UsersId, 10),
				strconv.FormatUint(PlayerList[i+8].ServerId, 10), strconv.FormatUint(PlayerList[i+8].Power, 10), strconv.FormatUint(PlayerList[i+8].Charm, 10),
				strconv.FormatUint(PlayerList[i+8].Money, 10), strconv.FormatUint(PlayerList[i+8].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+9].UsersId, 10),
				strconv.FormatUint(PlayerList[i+9].ServerId, 10), strconv.FormatUint(PlayerList[i+9].Power, 10), strconv.FormatUint(PlayerList[i+9].Charm, 10),
				strconv.FormatUint(PlayerList[i+9].Money, 10), strconv.FormatUint(PlayerList[i+9].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+10].UsersId, 10),
				strconv.FormatUint(PlayerList[i+10].ServerId, 10), strconv.FormatUint(PlayerList[i+10].Power, 10), strconv.FormatUint(PlayerList[i+10].Charm, 10),
				strconv.FormatUint(PlayerList[i+10].Money, 10), strconv.FormatUint(PlayerList[i+10].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+11].UsersId, 10),
				strconv.FormatUint(PlayerList[i+11].ServerId, 10), strconv.FormatUint(PlayerList[i+11].Power, 10), strconv.FormatUint(PlayerList[i+11].Charm, 10),
				strconv.FormatUint(PlayerList[i+11].Money, 10), strconv.FormatUint(PlayerList[i+11].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+12].UsersId, 10),
				strconv.FormatUint(PlayerList[i+12].ServerId, 10), strconv.FormatUint(PlayerList[i+12].Power, 10), strconv.FormatUint(PlayerList[i+12].Charm, 10),
				strconv.FormatUint(PlayerList[i+12].Money, 10), strconv.FormatUint(PlayerList[i+12].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+13].UsersId, 10),
				strconv.FormatUint(PlayerList[i+13].ServerId, 10), strconv.FormatUint(PlayerList[i+13].Power, 10), strconv.FormatUint(PlayerList[i+13].Charm, 10),
				strconv.FormatUint(PlayerList[i+13].Money, 10), strconv.FormatUint(PlayerList[i+13].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+14].UsersId, 10),
				strconv.FormatUint(PlayerList[i+14].ServerId, 10), strconv.FormatUint(PlayerList[i+14].Power, 10), strconv.FormatUint(PlayerList[i+14].Charm, 10),
				strconv.FormatUint(PlayerList[i+14].Money, 10), strconv.FormatUint(PlayerList[i+14].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+15].UsersId, 10),
				strconv.FormatUint(PlayerList[i+15].ServerId, 10), strconv.FormatUint(PlayerList[i+15].Power, 10), strconv.FormatUint(PlayerList[i+15].Charm, 10),
				strconv.FormatUint(PlayerList[i+15].Money, 10), strconv.FormatUint(PlayerList[i+15].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+16].UsersId, 10),
				strconv.FormatUint(PlayerList[i+16].ServerId, 10), strconv.FormatUint(PlayerList[i+16].Power, 10), strconv.FormatUint(PlayerList[i+16].Charm, 10),
				strconv.FormatUint(PlayerList[i+16].Money, 10), strconv.FormatUint(PlayerList[i+16].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+17].UsersId, 10),
				strconv.FormatUint(PlayerList[i+17].ServerId, 10), strconv.FormatUint(PlayerList[i+17].Power, 10), strconv.FormatUint(PlayerList[i+17].Charm, 10),
				strconv.FormatUint(PlayerList[i+17].Money, 10), strconv.FormatUint(PlayerList[i+17].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+18].UsersId, 10),
				strconv.FormatUint(PlayerList[i+18].ServerId, 10), strconv.FormatUint(PlayerList[i+18].Power, 10), strconv.FormatUint(PlayerList[i+18].Charm, 10),
				strconv.FormatUint(PlayerList[i+18].Money, 10), strconv.FormatUint(PlayerList[i+18].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+19].UsersId, 10),
				strconv.FormatUint(PlayerList[i+19].ServerId, 10), strconv.FormatUint(PlayerList[i+19].Power, 10), strconv.FormatUint(PlayerList[i+19].Charm, 10),
				strconv.FormatUint(PlayerList[i+19].Money, 10), strconv.FormatUint(PlayerList[i+19].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+20].UsersId, 10),
				strconv.FormatUint(PlayerList[i+20].ServerId, 10), strconv.FormatUint(PlayerList[i+20].Power, 10), strconv.FormatUint(PlayerList[i+20].Charm, 10),
				strconv.FormatUint(PlayerList[i+20].Money, 10), strconv.FormatUint(PlayerList[i+20].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+21].UsersId, 10),
				strconv.FormatUint(PlayerList[i+21].ServerId, 10), strconv.FormatUint(PlayerList[i+21].Power, 10), strconv.FormatUint(PlayerList[i+21].Charm, 10),
				strconv.FormatUint(PlayerList[i+21].Money, 10), strconv.FormatUint(PlayerList[i+21].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+22].UsersId, 10),
				strconv.FormatUint(PlayerList[i+22].ServerId, 10), strconv.FormatUint(PlayerList[i+22].Power, 10), strconv.FormatUint(PlayerList[i+22].Charm, 10),
				strconv.FormatUint(PlayerList[i+22].Money, 10), strconv.FormatUint(PlayerList[i+22].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+23].UsersId, 10),
				strconv.FormatUint(PlayerList[i+23].ServerId, 10), strconv.FormatUint(PlayerList[i+23].Power, 10), strconv.FormatUint(PlayerList[i+23].Charm, 10),
				strconv.FormatUint(PlayerList[i+23].Money, 10), strconv.FormatUint(PlayerList[i+23].intimacy, 10)}),
		spanner.InsertOrUpdate(table, UsersColumns,
			[]interface{}{strconv.FormatUint(PlayerList[i+24].UsersId, 10),
				strconv.FormatUint(PlayerList[i+24].ServerId, 10), strconv.FormatUint(PlayerList[i+24].Power, 10), strconv.FormatUint(PlayerList[i+24].Charm, 10),
				strconv.FormatUint(PlayerList[i+24].Money, 10), strconv.FormatUint(PlayerList[i+24].intimacy, 10)}),
	}
	_, err := client.Apply(ctx, m)
	return err
}

func main() {
	// Database name String
	databaseName := "projects/yunion-test-286209/instances/test-instance/databases/users-db"

	players := genPlayerList()
	// fmt.Println(len(players))

	ctx, client := createClients(databaseName)
	defer client.Close()

	flag := true
	const numberOfExperiments = 100
	bench := hrtime.NewBenchmark(numberOfExperiments)
	for i := 0; i <= len(players); i += 25 {
		// 判断是否停止benchmark
		if flag != true {
			break
		}
		// 避免array 超出索引
		if i != len(players) {
			err := write(ctx, client, "Users", players, i)
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
