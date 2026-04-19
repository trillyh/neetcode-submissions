class Solution:

    def encode(self, strs: List[str]) -> str:
        self.words_len = []
        encode_string = ""
        for word in strs:
            encode_string += word
            self.words_len.append(len(word))
        return encode_string
            
    def decode(self, s: str) -> List[str]:
        decode_list: List[str] = []
        next_word_start = 0
        next_word_stop = 0
        for word_len in self.words_len:
            next_word_stop = next_word_start + word_len
            decode_list.append(s[next_word_start: next_word_stop])
            next_word_start = next_word_stop
        return decode_list
            

            

            
