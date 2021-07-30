package algorithm

func FrogPosition(n int, edges [][]int, t int, target int) float64 {
	var cnt float64
	child := make([][]int, n+1)
	step := make([]int, n+1)
	top := make([]int, n+1)
	for _, edge := range edges {
		if step[edge[0]] != 0 {
			step[edge[1]] += step[edge[0]] + 1
			top[edge[1]] = top[edge[0]]
		} else {
			step[edge[1]]++
			top[edge[1]] = edge[0]
		}
		if child[edge[1]] != nil {
			if child[top[edge[1]]] != nil {
				child[top[edge[1]]] = append(child[top[edge[1]]], child[edge[1]]...)
			} else {
				child[top[edge[1]]] = append([]int(nil), child[edge[1]]...)
			}
			for _, c := range child[edge[1]] {
				step[c] += step[edge[1]]
				top[c] = top[edge[1]]
			}
		}
	}
	if step[target] > t {
		return 0
	}
	for _, s := range step {
		if s <= t {
			cnt = cnt + 1.0
		}
	}
	return 1.0 / cnt
}
