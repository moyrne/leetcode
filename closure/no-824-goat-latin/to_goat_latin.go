package no_824_goat_latin

import "strings"

// https://leetcode-cn.com/problems/goat-latin/
/*
824. 山羊拉丁文
给你一个由若干单词组成的句子 sentence ，单词间由空格分隔。每个单词仅由大写和小写英文字母组成。
请你将句子转换为 “山羊拉丁文（Goat Latin）”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。山羊拉丁文的规则如下：

如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。
例如，单词 "apple" 变为 "applema" 。
如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。

例如，单词 "goat" 变为 "oatgma" 。
根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。

例如，在第一个单词后添加 "a" ，在第二个单词后添加 "aa" ，以此类推。
返回将 sentence 转换为山羊拉丁文后的句子。
*/

func toGoatLatin(sentence string) string {
	ans := &strings.Builder{}
	for i, cnt, n := 0, 1, len(sentence); i < n; i++ {
		if cnt > 1 {
			ans.WriteByte(' ')
		}
		start := i
		for i++; i < n && sentence[i] != ' '; i++ {
		}
		cnt++
		switch sentence[start] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ans.WriteString(sentence[start:i])
		default:
			ans.WriteString(sentence[start+1 : i])
			ans.WriteByte(sentence[start])
		}
		ans.WriteByte('m')
		ans.WriteString(strings.Repeat("a", cnt))
	}

	return ans.String()
}
