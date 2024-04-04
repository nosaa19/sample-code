package java_code.logic;

public class maxSumSeq {
    
    // https://www.scaler.com/topics/maximum-sum-increasing-subsequence/
    public static int MaxSumSequenceArray(int[] arr, int i, int prev, int sum){
        // base case: nothing is remaining
        if ( i == arr.length) {
            return sum;
        }
        
        // If the element is greater than the previous element
        int incl = sum;
        if (arr[i] > prev) {
            incl = MaxSumSequenceArray(arr, i + 1, arr[i], sum + arr[i]);
        }
 
        // Call the recursive function without the current element.
        int excl = MaxSumSequenceArray(arr, i + 1, prev, sum);
 
        // return the maximum
        return Integer.max(incl, excl);
    }

    // https://medium.com/the-research-nest/solving-the-maximum-subarray-sum-a-super-simplified-explanation-34042ffd3fd7
    public static int MaxSumSubsequenceArray(int[] arr){
        int currentSum = arr[0];
        int maxSum = currentSum;
        int n = arr.length;
        for (int i = 1; i < n; i++) {
            currentSum = Math.max(arr[i], currentSum + arr[i]);
            maxSum = Math.max(maxSum, currentSum);
        }
        return maxSum;
    }
}
