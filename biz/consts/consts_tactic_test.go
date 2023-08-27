package consts

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestTactic(t *testing.T) {
	//fmt.Printf("%d\n",TacticId(1))
	//fmt.Printf("%v",TacticId(1))
	fmt.Println(EyebrowedThrush)
	fmt.Println(cast.ToInt64(EyebrowedThrush))
	fmt.Println(int64(EyebrowedThrush))
}
