package WhySingletonDao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestNormal(t *testing.T) {
	if err := Init(); err != nil {
		panic("init wrong")
	}
	num := 100
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done()
			err := InsertUser(strconv.Itoa(n), strconv.Itoa(n))
			if err != nil {
				fmt.Println(err)
			}

		}(i)
	}

	wg.Wait()

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done()
			user, err := QueryUserById(strconv.Itoa(n))
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, user.Name, user.Value)

		}(i)
	}

	wg.Wait()

}
