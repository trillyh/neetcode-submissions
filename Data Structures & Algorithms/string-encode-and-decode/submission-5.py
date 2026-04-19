from typing import List

class Solution:
    def encode(self, strs: List[str]) -> str:
        self.words_len = []  # Defining self.words_len inside the method
        encode_string = ""
        for word in strs:
            encode_string += word
            self.words_len.append(len(word))
        return encode_string
            
    def decode(self, s: str) -> List[str]:
        decode_list: List[str] = []
        words_len_index = 0
        last_word_index = 0
        
        for i in range(len(self.words_len)):
            word_len = self.words_len[i]
            decode_list.append(s[last_word_index:last_word_index + word_len])
            last_word_index += word_len
            
        return decode_list
