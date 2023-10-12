func findPeakElement(nums []int) int {
    n :=len(nums)
    if n==1{
        return 0
    }
    if nums[0]>nums[1]{
        return 0
    }
    if nums[n-1]>nums[n-2]{
        return n-1
    }
    low,high :=1,n-2

    for low<=high{
        mid:=(low+high)/2
        if nums[mid]>nums[mid-1] && nums[mid]>nums[mid+1]{
            return mid
        } else if nums[mid] > nums[mid-1]{
            low=low+1
        } else {
            high = high-1
        }
    }
    return -1
}
