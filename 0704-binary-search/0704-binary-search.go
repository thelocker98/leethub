func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    left := 0
    right := len(nums)-1
    center := 0

    for left <= right{
        center = (left+right)/2

        if nums[center] == target{
            return center
        }else if nums[center] > target {
            right = center -1
        }else{
            left = center +1
        }
    }
    return -1
}