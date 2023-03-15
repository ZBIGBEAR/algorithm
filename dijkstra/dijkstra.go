package dijkstra

/*
迪杰斯特拉算法:解决一个点到其他所有点的最短路径问题
*/

const (
	Max = 1 << 10
)

// 已知点的集合points和各个点之间的有向边长度edges，求从start点开始到剩余所有点的最短路径分别是多少
func Dijkstra(points []int, edges [][]int, start int) map[int]int {
	// 原始点集v
	v := make(map[int]struct{})
	for i := range points {
		v[points[i]] = struct{}{}
	}

	// 已经加入s的集合
	s := make(map[int]struct{})

	// 将第一个点point[start]从v集合移出，并加入s集合
	delete(v, start)
	s[start] = struct{}{}

	// 记录开始节点到其他各个节点最短距离
	paths := make(map[int]int)

	// 初始化paths
	for i := range edges[start] {
		paths[i] = edges[start][i]
	}

	for len(v) > 0 {
		// 只要v不为空，则一直循环
		// 1.找出paths中最小路径的点
		minPoint := 0
		for point, path := range paths {
			if _, ok := s[point]; ok {
				continue
			}
			if minPoint == 0 {
				minPoint = point
				continue
			}
			if path < paths[minPoint] {
				minPoint = point
			}
		}

		// 2.将minPoint点加入s集合
		s[minPoint] = struct{}{}

		// 3.将minPoint点从v集合中移除
		delete(v, minPoint)

		// 4.更新paths:如果start讲过点point到v中每个节点的距离更近则更是paths
		for point, _ := range v {
			newPath := paths[minPoint] + edges[minPoint][point]
			if newPath < paths[point] {
				paths[point] = newPath
			}
		}
	}

	return paths
}
