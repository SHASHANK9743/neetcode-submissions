class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hashmap = {}
        result = []
        for _, word in enumerate(strs):
            if len(word) not in hashmap:
                hashmap[len(word)] = {
                    tuple(sorted(word)): [word]
                }
            else:
                if tuple(sorted(word)) in hashmap[len(word)]:
                    hashmap[len(word)][tuple(sorted(word))].append(word)
                else:
                    hashmap[len(word)][tuple(sorted(word))] = [word]
        
        for _, value_1 in hashmap.items():
            for _, value_2 in value_1.items():
                result.append(value_2)
        return result