package offer46;

class Solution {
//    public int translateNum(int num) {
//        String s = String.valueOf(num);
//        int len = s.length();
//        if (len < 2) {
//            return len;
//        }
//
//        char[] charArray = s.toCharArray();
//        int[] dp = new int[len + 1];
//        dp[0] = 1;
//        dp[1] = 1;
//        for (int i = 1; i < len; i++) {
//            dp[i + 1] = dp[i];
//            int cur = 10 * (charArray[i - 1] - '0') + (charArray[i] - '0');
//            if (cur > 9 && cur < 26) {
//                dp[i + 1] += dp[i - 1] + dp[i];
//            }
//        }
//        return dp[len];
//    }

    // 滚动数组
//    public int translateNum(int num) {
//        String s = String.valueOf(num);
//
//        int p = 0, q = 0, r = 1;
//        for (int i = 0; i < s.length(); i++) {
//            p = q;
//            q = r;
//            r = 0;
//            r += q;
//            if (i == 0) {
//                continue;
//            }
//            String pre = s.substring(i - 1, i + 1);
//            if (pre.compareTo("25") <= 0 && pre.compareTo("10") >= 0) {
//                r += p;
//            }
//        }
//
//        return r;
//    }

    public int translateNum(int num) {
        if (num < 10)
            return 1;

        int cur = num % 100;
        if (10 <= cur && cur <= 25)
            return translateNum(num / 10) + translateNum(num / 100);
        else
            return translateNum(num / 10);
    }

    public static void main(String[] args) {
        System.out.println(new Solution().translateNum(123123));
        System.out.println(new Solution().translateNum(12));
        System.out.println(new Solution().translateNum(31));
    }
}