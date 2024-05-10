package java_code.logic;

public class nearestKios {
    /* 
        Given a n x n grid, find the nearest kiosk from any point. 
        There can be more than 1 kiosk, give the shortest distance for each node in the grid to all kiosks.
        A kiosk will have a value of k, otherwise a value of c will be seen.

        grid = [
            [c, c, c],
            [k, c, c],
            [c, c, k]
        ]

        output = [
            [1, 2, 2],
            [0, 1, 1],
            [1, 1, 0]
        ]
     */

     public static Integer[][] getShorthestDist(){

        Character[][] grid = {
            {'c', 'c', 'c'},
            {'k', 'c', 'c'},
            {'c', 'k', 'k'}
        };

        // Find coordinate of k

        // Calculate distance c to k

        Integer[][] result = new Integer[grid.length][grid.length];

        return result;
     } 
}
