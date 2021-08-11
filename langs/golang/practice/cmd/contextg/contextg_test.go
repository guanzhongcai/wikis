package contextg

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go watch(ctx, 1)
	go watch(ctx, 2)

	time.Sleep(2 * time.Second)
	t.Log("ctx err is: ", ctx.Err())
	// sleep 2秒: context canceled
	// sleep 4秒: context deadline exceeded

	if ctx.Err() == nil {
		t.Log("活可能干完了，因为没有触发超时，就做完了！")
		t.Log("也可能没干完，因为超时还没到，没有调用过cancel()，调用过cancel，ctx.Err()会有error信息！")
	}
}

func watch(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("监控停止了! id=%d\n", id)
			// todo: 此处应该做清理操作，然后退出goroutine，释放资源
			return
		default:
			fmt.Printf("监控中：id=%d ..\n", id)
			time.Sleep(time.Second)
		}
	}
}
