package balance

import (
	"github.com/shura1014/common/goerr"
	"github.com/shura1014/common/random"
	"sort"
	"sync"
)

type IWeight interface {
	GetWeight() int
}

type WeightNode interface {
	IWeight
	Node
}

type Weight[V WeightNode] struct {
	lock  sync.Mutex
	nodes []V
}

func NewWeight[V WeightNode]() *Weight[V] {
	return &Weight[V]{}
}

func (b *Weight[V]) Instance() (node V, err error) {
	index := random.Int(len(b.nodes) - 1)
	b.lock.Lock()
	defer b.lock.Unlock()
	for i := 0; i < len(b.nodes); i++ {
		node = b.nodes[(index+i)%len(b.nodes)]
		if node.IsValid() {
			return
		}
	}
	err = goerr.Text("无可用实例")
	return

}

func (b *Weight[V]) InitNodes(nodes ...V) {
	b.lock.Lock()
	maxCommonDivisor := MaxCommonDivisor(nodes...)
	for _, node := range nodes {
		weight := node.GetWeight() / maxCommonDivisor
		for i := 0; i < weight; i++ {
			b.nodes = append(b.nodes, node)
		}
	}
	b.lock.Unlock()
}

func MaxCommonDivisor[V IWeight](nodes ...V) int {
	var weightArray []int
	for _, node := range nodes {
		weightArray = append(weightArray, node.GetWeight())
	}
	sort.Ints(weightArray)
	if len(weightArray) > 0 {
		for i := weightArray[0]; i >= 2; i-- {
			isDivideExactly := false
			for _, w := range weightArray {
				isDivideExactly = w%i == 0
				if !isDivideExactly {
					break
				}
			}
			if isDivideExactly {
				return i
			}
		}
	}

	// 默认最大公约数
	return 1
}

func (b *Weight[V]) Name() string {
	return Weight_
}
