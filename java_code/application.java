package java_code;

import java.util.Arrays;
import java.util.List;

import java_code.logic.maxSumSeq;
import java_code.logic.uniqueNumber;

/**
 * main
 */
public class application {

    public static void main(String[] args) {
        int[] arr = { 2, -1, 1, 5};
        System.out.println("MaxSumSequenceArray: "+maxSumSeq.MaxSumSequenceArray(arr, 0, Integer.MIN_VALUE, 0));
        System.out.println("MaxSumSubsequenceArray: "+ maxSumSeq.MaxSumSubsequenceArray(arr));

        List<Integer> arr01 = Arrays.asList(1, 7, 8, 1, 9, 3, 8);
        System.out.println("UniqueArray"+ uniqueNumber.FindUniqueNumbers(arr01));
    }
}