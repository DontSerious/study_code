package offer48;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

public class Solution {
    public int lengthOfLongestSubstring(String s) {
        char[] sa = s.toCharArray();
        int n = s.length(), left = 0, ans = 0;
        int[] cnt = new int[128];

        for (int right = 0; right < n; ++right) {
            char c = sa[right];
            ++cnt[c];
            while (cnt[c] > 1)
                --cnt[sa[left++]];
            ans = Math.max(ans, right - left + 1);
        }

        return ans;
    }

    public static void main(String[] args) {
        System.out.println(new Solution().lengthOfLongestSubstring("pwwkew"));
    }
}
