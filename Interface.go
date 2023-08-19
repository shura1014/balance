package balance

import "github.com/shura1014/common/goerr"

const (
	RoundRobin_ = "roundrobin"
	Random_     = "random"
	Weight_     = "weight"
)

type Node interface {
	// IsValid 有效校验，如果节点无效应当继续选择下一个节点
	IsValid() bool
}

type Balance[V Node] interface {
	// Instance 获取一个实例
	Instance() (V, error)
	// InitNodes 初始化
	InitNodes(nodes ...V)
	Name() string
}

func GetBalance[V Node](b string) Balance[V] {
	switch b {
	case RoundRobin_:
		return &RoundRobin[V]{}
	case Random_:
		return &Random[V]{}
	case Weight_:
		// Go的泛型还不支持
		panic(goerr.Text("No support please use NewWeight()"))
		//return &Weight[WeightNode]{}
	default:
		return &RoundRobin[V]{}
	}
}
