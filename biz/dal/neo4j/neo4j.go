package neo4j

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

var Neo4jDriver neo4j.DriverWithContext

func InitNeo4j() {
	//通过bolt协议连接到Neo4j
	var err error
	Neo4jDriver, err = neo4j.NewDriverWithContext(fmt.Sprintf("neo4j://%s", conf.GetConfig().Neo4j.DbUri),
		neo4j.BasicAuth(conf.GetConfig().Neo4j.UserName, conf.GetConfig().Neo4j.Password, ""))
	if err != nil {
		panic(any(err))
	}
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
		panic(any(fmt.Sprintf("Failed to create Neo4j driver :%v", err)))
	}
	hlog.CtxInfof(context.Background(), "connected to Neo4j success")
}
