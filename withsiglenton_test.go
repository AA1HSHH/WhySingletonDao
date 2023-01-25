package WhySingletonDao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestWithSigleton(t *testing.T) {
	if err := Init(); err != nil {
		panic("init wrong")
	}
	var wg sync.WaitGroup
	num := 100
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done()
			err := NewUserDaoInstance().InsertUser(strconv.Itoa(n), strconv.Itoa(n))
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
			user, err := NewUserDaoInstance().QueryUserById(strconv.Itoa(n))
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, user.Name, user.Value)

		}(i)
	}

	wg.Wait()

}
