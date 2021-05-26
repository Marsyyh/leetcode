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

        if (num1.size() == 0 || num2.size() == 0 || num1 == "0" || num2 == "0") {
            return result;
        }
        
        // loop from r to l
        for (int i = num1.size()-1; i >= 0; --i) {
            string cur = "";
            int nextBit = 0;
            int num1Bit = num1[i] - '0';

            // num1 1 bit time all num2
            for (int j = num2.size()-1; j >= 0; --j) {
                int num2Bit = num2[j] - '0';
                int curRe = num1Bit * num2Bit + nextBit;
                nextBit = curRe / 10;
                char curChar = (curRe % 10) + '0';
                cur = curChar + cur;
            }

            if (nextBit != 0) {
                char curChar = nextBit + '0';
                cur = curChar + cur;
            }

            for (int j = num1.size()-1; j > i; --j) {
                cur = cur + '0';
            }

            mergeString(result, cur);
        }

        return result;
    }

    void mergeString(string& num1, string num2) {
        int i = 0;
        int next = 0;
        string result = "";
        while (i < num1.size() && i < num2.size()) {
            int cur = (num1[num1.size()-1-i] - '0') + (num2[num2.size()-1-i] - '0') + next;
            if (cur >= 10) {
                next = 1;
                cur = cur - 10;
            } else {
                next = 0;
            }
            char curChar = cur + '0';
            result = curChar + result;
            i++;
        }

        while (i < num1.size()) {
            if (next != 0) {
                int cur = (num1[num1.size()-1-i] - '0') + next;
                if (cur >= 10) {
                    next = 1;
                    cur = cur - 10;
                } else {
                    next = 0;
                }
                char curChar = cur + '0';
                result = curChar + result;
                i++;
                continue;
            }
            result = (num1[num1.size()-1-i]) + result;
            i++;
        }

        while (i < num2.size()) {
            if (next != 0) {
                int cur = (num2[num2.size()-1-i] - '0') + next;
                if (cur >= 10) {
                    next = 1;
                    cur = cur - 10;
                } else {
                    next = 0;
                }
                char curChar = cur + '0';
                result = curChar + result;
                i++;
                continue;
            }
            result = (num2[num2.size()-1-i]) + result;
            i++;
        }

        if (next != 0) {
            result = '1' + result;
        }
        num1 = result;
    }
};
// @lc code=end

