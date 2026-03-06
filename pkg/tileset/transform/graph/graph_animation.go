package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// GraphAnimation
// Animation
// channels 字段定义了动画数据作用的目标节点和属性路径（例如weights）。
// samplers 字段定义了关键帧的时间点和对应的权重值，这里的output对应的Accessor的数据就是关键帧最终要达到的权重值数组
//
// 在动画播放过程中，根据当前时间点，从关键帧数据中插值计算出当前的权重值
// 插值方法可以是线性插值或样条插值，具体取决于动画的interpolation属性
// 最终顶点数据=原始顶点数据+∑(权重×目标顶点数据)
//
// node和mesh的weights字段定义了每个变形目标的初始权重，优先使用node.weights
// weights 通常是一个[0, 1]的浮点数数组，每个元素对应一个PrimitiveTarget的权重
// 一个角色可能有“微笑”和“皱眉”两个变形目标。若weights为[0.5, 0]，则模型初始状态是“半笑”，第二个目标未激活。
// 一个动画可以同时控制多个目标，一个人可以同时微笑和皱眉
// 初始权重可以算是一个动画的中间状态。但当前模型处于哪个状态是未知的，所以一般不需要初始权重，而是在运动时设置初始权重
// 比如说一个角色正在走路，突然要跑，需要判断它的状态，先播放一段从走路到跑的中间动画，是快速迈开左脚还是右脚，这个中间动画需要包含走路的状态
type GraphAnimation struct {
	Graph *Graph
	*gltf.Animation
}

// Index
func (m *GraphAnimation) Index() int {
	return slices.Index(m.Graph.Animations, m.Animation)
}

// IsUsed
func (m *GraphAnimation) IsUsed() bool {
	return true
}

// Dispose
func (m *GraphAnimation) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Animations = slices.Delete(m.Graph.Animations, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphAnimation) IsEmptyExtras() bool {
	return m.Extras == nil
}
