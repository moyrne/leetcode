package _3508_Mid

type (
	Router struct {
		memoryLimit int

		queue            []packet
		packetSet        map[packet]struct{}
		destToTimestamps map[int][]int
	}
	packet struct {
		source      int
		destination int
		timestamp   int
	}
)

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:      memoryLimit,
		packetSet:        make(map[packet]struct{}),
		destToTimestamps: make(map[int][]int),
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	set := packet{source, destination, timestamp}

	if _, ok := r.packetSet[set]; ok {
		return false
	}

	if len(r.queue) >= r.memoryLimit {
		r.ForwardPacket()
	}

	r.queue = append(r.queue, set)
	r.packetSet[set] = struct{}{}
	r.destToTimestamps[destination] = append(r.destToTimestamps[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.queue) == 0 {
		return nil
	}

	val := r.queue[0]
	r.queue = r.queue[1:]
	delete(r.packetSet, val)

	r.destToTimestamps[val.destination] = r.destToTimestamps[val.destination][1:]
	if len(r.destToTimestamps[val.destination]) == 0 {
		delete(r.destToTimestamps, val.destination)
	}

	return []int{val.source, val.destination, val.timestamp}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	sts := r.destToTimestamps[destination]
	if len(sts) == 0 {
		return 0
	}

	return lowerBound(sts, endTime+1) - lowerBound(sts, startTime)
}

func lowerBound(nums []int, target int) int {
	right := len(nums) - 1

	for left := 0; left < right; {
		mid := (left + right) >> 1

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
