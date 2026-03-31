# Valid Anagram - https://neetcode.io/problems/is-anagram/question?list=neetcode150
# Problem : Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

# Goal - 
'''
    1. check if a1==a2
    2. Loop over anagrams and map
    3 Check if maps equal
'''
# Time complexity is O(n) since we iterate through both strings and compare hash maps. Space complexity is O(n) because we store character counts.”

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        letter_count_s = {}
        letter_count_t = {}

        for i in s:
            if i in letter_count_s:
                 letter_count_s[i] += 1
            else:
                 letter_count_s[i] = 1
    
        for i in t:
            if i in letter_count_t:
                 letter_count_t[i] += 1
            else:
                 letter_count_t[i] = 1

        if letter_count_s == letter_count_t:
            return True

        return False