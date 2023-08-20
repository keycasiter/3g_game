package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"testing"
)

func TestBattleExecute(t *testing.T) {
	ctx := context.Background()
	req := &app.RequestContext{}
	BattleExecute(ctx, req)
}
