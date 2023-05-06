package offer12;

/*
	矩阵中的路径
	给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
	如果 word 存在于网格中，返回 true ；否则，返回 false 。

	单词必须按照字母顺序，通过相邻的单元格内的字母构成，
	其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

class Solution {
    public boolean exist(char[][] board, String word) {
        char[] words = word.toCharArray();

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (dfs(board, words, i, j, 0))
                    return true;
            }
        }

        return false;
    }

    boolean dfs(char[][] board, char[] word, int i, int j, int k) {
        // 退出条件
        if (i >= board.length || i < 0 || j >= board[0].length || j < 0 || board[i][j] != word[k])
            return false;
        // 当成功匹配的数量等于word的长度则代表匹配成功
        if (k == word.length - 1)
            return true;
        // 单字匹配成功将其赋值为不再可能被匹配的字符，防止重复
        board[i][j] = '\0';
        k++;

        // 深度遍历顺序为下上右左
        boolean res = dfs(board, word, i + 1, j, k) || dfs(board, word, i - 1, j, k) ||
                      dfs(board, word, i, j + 1, k) || dfs(board, word, i, j - 1, k);

        // 如果全部匹配失败，则回溯，重新回复数组至原来
        k--;
        board[i][j] = word[k];

        return res;
    }
}