/*
 * @lc app=leetcode.cn id=43 lang=cpp
 *
 * [43] 字符串相乘
 */

// @lc code=start
/*
    thought:
        123 * 45 =  
                123 * 5
        --------------            
                    15
                   10
                   5
                   605
    ------------------ 
               123 * 4
        -------------- 
                   12
                   8
                  4
                  492
                   605
                  5525
    ------------------
*/
class Solution {
public:
    string multiply(string num1, string num2) {
        string result = "0";

        if (num1.size() == 0 || num2.size() == 0) {
            return result;
        }
        
        for (int i = 0; i < num1.size(); i++) {
            string cur = "";
            char nextBit = '0';
            int num1Bit = num1[i] - '0';
            for (int j = 0; j < num2.size(); j++) {
                int num2Bit = num2[j] - '0';
                int curRe = num1Bit * num2Bit + (nextBit - '0');
                nextBit = (curRe / 10) + '0';
                cur.append((curRe % 10)+'0');
            }
        }
    }
};
// @lc code=end

