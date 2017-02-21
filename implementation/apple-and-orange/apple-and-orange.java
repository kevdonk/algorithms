import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {
    
    static boolean between(int x, int left, int right) {
        return x >= left && x <= right;
    }

    public static void main(String[] args) {
        int applesHit = 0;
        int orangesHit = 0;
        
        Scanner in = new Scanner(System.in);
        int s = in.nextInt();
        int t = in.nextInt();
        int a = in.nextInt();
        int b = in.nextInt();
        int m = in.nextInt();
        int n = in.nextInt();
        int appleMin = s - a;
        int appleMax = t - a;
        int orangeMin = s - b;
        int orangeMax = t - b;

        int distance;
        for(int i=0; i < m; i++){
            distance = in.nextInt();
            if (between(distance, appleMin, appleMax)) {
                applesHit += 1;
            }
        }

        for(int i=0; i < n; i++){
            distance = in.nextInt();
            if (between(distance, orangeMin, orangeMax)) {
                orangesHit += 1;
            }
        }
        
        System.out.println(applesHit);
        System.out.println(orangesHit);
    }
}