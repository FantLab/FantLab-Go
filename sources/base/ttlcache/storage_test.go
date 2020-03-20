package ttlcache

import (
	"fantlab/base/assert"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_Usage(t *testing.T) {
	t.Run("simple_usage", func(t *testing.T) {
		now := time.Now()

		storage := NewWithExpireFunc(func(t time.Time) bool {
			return now.Sub(t) > 0
		})

		assert.True(t, storage.Len() == 0)

		storage.Set("key", "some", now.Add(1*time.Second), false)

		val, ok := storage.Get("key")

		assert.True(t, ok)
		assert.True(t, val == "some")
		assert.True(t, storage.Len() == 1)
	})

	t.Run("delete", func(t *testing.T) {
		now := time.Now()

		storage := NewWithExpireFunc(func(t time.Time) bool {
			return now.Sub(t) > 0
		})

		storage.Set("k1", "v1", now.Add(1*time.Second), false)
		storage.Set("k2", "v2", now.Add(1*time.Second), false)
		storage.Set("k3", "v3", now.Add(3*time.Second), false)

		assert.True(t, storage.Len() == 3)

		now = now.Add(2 * time.Second)

		storage.DeleteExpired()

		val, ok := storage.Get("k3")

		assert.True(t, ok)
		assert.True(t, val == "v3")
		assert.True(t, storage.Len() == 1)

		storage.DeleteAll()

		assert.True(t, storage.Len() == 0)

		storage.Set("k1", "v1", now.Add(1*time.Second), false)

		assert.True(t, storage.Len() == 1)

		storage.Delete("k1")

		assert.True(t, storage.Len() == 0)
	})

	t.Run("respect", func(t *testing.T) {
		now := time.Now()

		storage := NewWithExpireFunc(func(t time.Time) bool {
			return now.Sub(t) > 0
		})

		storage.Set("key", "some", now.Add(1*time.Second), true)

		val, ok := storage.Get("key")

		assert.True(t, ok)
		assert.True(t, val == "some")

		storage.Set("key", "some1", now.Add(1*time.Second), true)

		val, ok = storage.Get("key")

		assert.True(t, ok)
		assert.True(t, val == "some")

		now = now.Add(2 * time.Second)

		storage.Set("key", "some2", now.Add(1*time.Second), true)

		val, ok = storage.Get("key")

		assert.True(t, ok)
		assert.True(t, val == "some2")
	})
}

func Benchmark_ParallelStability(b *testing.B) {
	var lock sync.Mutex

	now := time.Now()

	storage := NewWithExpireFunc(func(t time.Time) bool {
		lock.Lock()
		defer lock.Unlock()

		return now.Sub(t) > 0
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := rand.Intn(10)

			if x%2 == 0 {
				storage.Set(strconv.Itoa(x), "value", now.Add(1*time.Second), true)
			} else {
				storage.Delete(strconv.Itoa(x))
			}

			if rand.Intn(50)%5 == 0 {
				lock.Lock()
				now = now.Add(2 * time.Second)
				lock.Unlock()

				storage.DeleteExpired()
			}
		}
	})
}
