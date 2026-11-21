class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        min_length = 200
        shortest_word = ""
        for _, word in enumerate(strs):
            if len(word) < min_length:
                min_lenght = len(word)
                shortest_word = word
        i = len(word)
        
        while i>0:
            sub_char = shortest_word[:i]
            is_longest = True
            for _, word in enumerate(strs):
                if word.startswith(sub_char):
                    pass
                else:
                    is_longest = False
            if is_longest:
                return sub_char

            i-=1
        return ""