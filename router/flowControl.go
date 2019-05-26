package router

import "log"

type BucketLimit struct {
	count int // 定义最大流量
	bucket chan int // 定义令牌桶
}

func NewBucketLimit(cc int) *BucketLimit {
	return &BucketLimit{
		cc,
		make(chan int,cc),
	}
}

// 获得令牌
func (cl *BucketLimit) GetConn() bool {
	// 如果满了
	if len(cl.bucket) >= cl.count {
		log.Println("满了")
		return false
	}
	// 没有满 就颁发令牌
	cl.bucket <- 1
	return true
}

// 回收令牌
func (cl *BucketLimit) ReleaseConn() {
	i := <-cl.bucket
	log.Printf("New connction coming:%d\n",i)
}