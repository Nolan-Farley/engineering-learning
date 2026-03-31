// Contains Duplicate - https://neetcode.io/problems/duplicate-integer/question?list=neetcode150
// PROBLEM: Given an array on nums, return true if any value app/* Given an intiger array nums, return true if any value appears more than once in the array, other wise return false

// GOAL - Aim for 0(n) time and 0(n) space, n eaquals the size of array input
// MEANING - it will grow in space and compleixty as same rate as array passed in
// Algo - Hash Map , its a have i seen this value before pattern

func hasDuplicateWorse{
    seen : map[int]bool{}

    for _, v := range nums{ 
        if seen[v] { //base case - grab value from map if exsits
            return true
        }
        seen[n] = true //assign latest value to map
    }
    return false 
}


// O(n^2) - for every element we are looping through every other element checking for match
func hasDuplicateWorse(nums []int) bool {
    for i := 0 ; i < len(nums) ; i++ {
		for j := i + 1 ; j < len(nums) ; j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
	}
    return false
}
