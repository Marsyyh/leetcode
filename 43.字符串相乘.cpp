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
        
        if (num1.size() == 0 || num2.size() == 0) {
            return string("0");
        }
        
        for (int i = 0; i < num1.size(); i++) {
            
        }
    }
};
// @lc code=end

