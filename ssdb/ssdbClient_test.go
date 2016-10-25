package ssdb

import (
	"fmt"
	"testing"
	"sync"
)

func TestSSDBClient_Hscan(t *testing.T) {
	ssdb := NewSSDBClient("localhost", "8888")
	err := ssdb.Connect()
	if err != nil {
		t.Error(err)
	} else {
		var wait sync.WaitGroup
		for i := 0; i < 100; i++ {
			wait.Add(1)
			go func() {
				result, err := ssdb.Hscan("ha", "", "", 10)
				result2, err := ssdb.Get("a")
				result3, err := ssdb.Hget("ha", "a")
				result4, err := ssdb.Set("b", "1111111111111111111")
				result5, err := ssdb.Hset("hb", "hello", "world")
				result6, err := ssdb.Hget("hb", "hello")

				if err != nil {
					t.Error(err)
				} else {
					fmt.Println(result, result2, result3, result4, result5, result6)
				}
				wait.Done()
			}()

		}
		wait.Wait()
	}
}