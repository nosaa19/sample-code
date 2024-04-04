package java_code;

import java_code.logic.maxSumSeq;

/**
 * main
 */
public class application {

    public static void main(String[] args) {
        int[] arr = { 2, -1, 1, 5};
        System.out.println("MaxSumSequenceArray: "+maxSumSeq.MaxSumSequenceArray(arr, 0, Integer.MIN_VALUE, 0));
        System.out.println("MaxSumSubsequenceArray: "+ maxSumSeq.MaxSumSubsequenceArray(arr));
    }
}