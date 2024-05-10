package java_code.logic;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class uniqueNumber {

    public static List<Integer> FindUniqueNumbers(List<Integer> input){

        Map<Integer, Integer> statsMap = new HashMap<Integer, Integer>();

        // Count the frequency of each number
        for (int num : input){
            statsMap.put(num, statsMap.getOrDefault(num, 0)+1);
        }

        List<Integer> uniqueNumbers = new ArrayList<Integer>();

        // Iterate through the statistics map to find numbers occuring exactly once
        for (Map.Entry<Integer, Integer> entry : statsMap.entrySet()){
            if(entry.getValue() == 1){
                uniqueNumbers.add(entry.getKey());
            }
        }


        return uniqueNumbers;


    }
}
