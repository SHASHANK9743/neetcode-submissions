class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        else:
            s_char_arr = [0]*255
            t_char_arr = [0]*255

            for s_char, t_char in zip(s, t):
                s_char_arr[ord(s_char)] += 1
                t_char_arr[ord(t_char)] += 1 
            return s_char_arr == t_char_arr