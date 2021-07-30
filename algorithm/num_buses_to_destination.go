package algorithm

func NumBusesToDestination(routes [][]int, source int, target int) int {
	queue := make([]int, 0)
	graph := make(map[int][]int)
	nums := make([]int, len(routes))
	for i := range routes {
		for _, j := range routes[i] {
			if graph[j] == nil {
				graph[j] = make([]int, 1)
				graph[j][0] = i
			} else {
				graph[j] = append(graph[j], i)
			}
		}
	}
	if graph[source] == nil || graph[target] == nil {
		return -1
	} else if source == target {
		return 0
	}
	for _, i := range graph[source] {
		queue = append(queue, i)
		nums[i] = 1
	}
	graph[source] = nil
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			for _, i := range routes[queue[0]] {
				if i == target {
					return nums[queue[0]]
				} else if graph[i] != nil {
					for _, j := range graph[i] {
						if nums[j] == 0 {
							nums[j] = nums[queue[0]] + 1
							queue = append(queue, j)
						}
					}
					graph[i] = nil
				}
			}
			queue = queue[1:]
			size--
		}
	}
	return -1
}
