package neo4j

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"testing"
)

func TestInitNeo4j(t *testing.T) {
	conf.InitConfig()
	InitNeo4j()
	mysql.InitMysql()

	_, err := neo4j.ExecuteQuery(context.Background(), Neo4jDriver,
		"match (t:team { id: $id, name: $name }) RETURN n",
		map[string]interface{}{
			"id":   1,
			"name": "Item 1",
		}, neo4j.EagerResultTransformer)
	if err != nil {
		fmt.Errorf("err:%v", err)
	}
}
