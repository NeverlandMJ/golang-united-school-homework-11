package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

type result []user

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func GetBatch(n int64, pool int64) (res result) {
	ids := []int64{}
	for i := int64(0); i < n; i++ {
		ids = append(ids, i)
	}
	size := int64(len(ids)) / pool
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := int64(0); i < pool; i++ {
		wg.Add(1)
		go func(ids []int64) {
			for _, id := range ids {
				user := getOne(id)
				mu.Lock()
				res = append(res, user)
				mu.Unlock()
			}
			defer wg.Done()
		}(ids[i*size : (i+1)*size])
	}
	wg.Wait()
	return res
}

