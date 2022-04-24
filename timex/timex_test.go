package timex

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	fmt.Println(Parse("2021/12/19 8:36:29", "2006-01-02 15:04:05").Int64())
}
