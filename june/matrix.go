package june

import "sort"

/*
# Two City Scheduling
# https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3349/


There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.



Example 1:

Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.


Note:

1 <= costs.length <= 100
It is guaranteed that costs.length is even.
1 <= costs[i][0], costs[i][1] <= 1000
*/
func TwoCitySchedCost(costs [][]int) int {
	return twoCitySchedCost(costs)
}

// 思路：将所有的人派去B地，然后将一半的人派去A地，则AB之间的差值从小到大排序，前半部分就是B去A消费最少的金额
func twoCitySchedCost(costs [][]int) int {
	var res int
	clen := len(costs)
	diff := make([]int, clen)
	for i, cost := range costs {
		res += cost[1]
		diff[i] = cost[0] - cost[1]
	}
	sort.Ints(diff)
	for i := 0; i < clen/2; i++ {
		res += diff[i]
	}
	return res
}

/*
# Queue Reconstruction by Height
# https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3352/

Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k),
where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h.
Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.


Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/
func ReconstructQueue(people [][]int) [][]int {
	return reconstructQueue(people)
}
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	var res [][]int
	for _, p := range people {
		if p[1] >= len(res) {
			res = append(res, p)
		} else {
			host := res[p[1]:]
			tmp := append([][]int{p}, host...)
			res = append(res[:p[1]], tmp...)
		}
	}
	return res
}
