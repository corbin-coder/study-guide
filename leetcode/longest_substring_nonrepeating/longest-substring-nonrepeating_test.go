package longestsubstringnonrepeating

import "testing"

func TestLengthOfLongestSubsctring(t *testing.T) {
	data := "abcablbmnop"
	res := lengthOfLongestSubstring(data)
	t.Log(res)
}
