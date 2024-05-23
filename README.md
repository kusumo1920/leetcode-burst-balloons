# Leetcode burst balloons

## Pseudocode of dynamic programming template

```
function dp(dp_state, memo_dict) {
    // check if we have seen this dp_state
    if dp_state in memo_dict
        return memo_dict[dp_state]
    
    // base case (a case that we know the answer for already) such as dp_state is empty
    if dp_state is the base cases
        return things like 0 or null
    
    calculate dp(dp_state) from dp(other_state)
    
    save dp_state and the result into memo_dict
}

function answerToProblem(input) {
    return dp(start_state, empty_memo_dict)
}
```

## Pseudocode for maxCoinsSolution1
```
// return maximum coins obtainable if we burst all balloons in `nums`
function dp(dp_state, memo_dict) {
    // check if have we seen this dp_state
    if nums in memo_dict
        return memo_dict[dp_state]
    
    // base case
    if nums is empty
        return 0
    
    max_coins = 0
    for i in 1 ... nums.length - 2: // this is due to nums contain fake balloons
        // burst nums[i]
        gain = nums[i-1] * nums[i] * nums[i+1]
        // burst the remaining balloons
        remaining = dp(nums without nums[i])
        max_coins = max(max_coins, gain + remaining)
    
    save dp_state and the result into memo_dict
    return max_coins
}

function maxCoin(nums) {
    nums = [1] + nums + [1] // add fake balloons
    return dp(nums, empty_memo_dict)
}
```