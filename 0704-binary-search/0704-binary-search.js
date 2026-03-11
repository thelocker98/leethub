/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    if (nums.length == 0){
        return -1;
    }
    var left =0;
    var right = nums.length-1;

    while (left <= right){
        var center = parseInt((left+right)/2);

        if(nums[center] == target){
            return center;
        }else if(nums[center] > target){
            right = center -1;
        }else{
            left = center +1
        }
    }
    return -1;
};