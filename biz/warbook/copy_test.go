package warbook

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestName(t *testing.T) {
	for i := 1; i <= 79; i++ {
		ioutil.WriteFile(fmt.Sprintf("warbook_detail_type_%d.go", i), []byte(fmt.Sprintf(`package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/model/vo"
)

type WarBookDetailType_%d struct {
	
}

func (w *WarBookDetailType_%d) Handle(ctx context.Context,general *vo.BattleGeneral) {

}
`, i, i)), 0633)
	}
}
