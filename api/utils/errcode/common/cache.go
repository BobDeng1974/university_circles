package common

import "sync"

var defaultSyncPool *SyncPool

func init() {
	defaultSyncPool = NewSyncPool(
		16,      // 2^3  byte
		2097152, // 2^21 byte
		2,
	)
}

// Get 按大小获取合适的缓存块
func Get(size int) []byte {
	return defaultSyncPool.get(size)
}

// Put 退还缓存块
func Put(mem []byte) {
	defaultSyncPool.put(mem)
}

// SyncPool is a sync.Pool base slab allocation memory pool
type SyncPool struct {
	slot     []sync.Pool
	slotSize []int // 桶对应缓存大小
	slotNum  int   // 桶个数
	minSize  int
	maxSize  int
}

// NewSyncPool 初始化缓存桶
func NewSyncPool(minSize, maxSize, factor int) *SyncPool {
	slotNum := 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		slotNum++
	}

	pool := &SyncPool{
		slot:     make([]sync.Pool, slotNum),
		slotSize: make([]int, slotNum),
		slotNum:  slotNum,
		minSize:  minSize,
		maxSize:  maxSize,
	}

	chunkSize := minSize
	for i := 0; i < slotNum; i++ {
		pool.slotSize[i] = chunkSize
		pool.slot[i].New = func(size int) func() interface{} {
			return func() interface{} {
				buf := make([]byte, size)
				return &buf
			}
		}(chunkSize)
		chunkSize *= factor
	}

	return pool
}

func (pool *SyncPool) get(size int) []byte {
	idx := pool.getSlot(size)
	if idx < 0 { // 没有对应桶,直接堆上申请
		return make([]byte, 0, size)
	}

	mem := pool.slot[idx].Get().(*[]byte)
	return (*mem)[0:size]
}

func (pool *SyncPool) put(mem []byte) {
	idx := pool.getSlot(cap(mem))
	if idx < 0 { // 没有对应桶,由gc回收
		return
	}

	pool.slot[idx].Put(&mem)
}

func (pool *SyncPool) getSlot(size int) int {
	if size == 0 || size > pool.maxSize {
		return -1
	}

	for i := 0; i < pool.slotNum; i++ {
		if pool.slotSize[i] >= size {
			return i
		}
	}

	return -1
}
