/*
Package bob provides answers from a lackadaisical teenager, Bob, to a purist linguist.

1) Bob answers 'Sure.' if you ask him a question, such as "How are you?".
2) He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
3) He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
4) He says 'Fine. Be that way!' if you address him without actually saying anything.
5) He answers 'Whatever.' to anything else.
*/
package bob

import (
	"strings"
	"regexp"
)

// Hey replies Bob with proper
func Hey(remark string) string {
	reg1, _ := regexp.Compile(`\s+`)
	reg2, _ := regexp.Compile("[^a-zA-Z0-9]+")
	hasStrings, _ := regexp.Match(`[a-zA-Z]`, []byte(remark))
	remark = reg1.ReplaceAllString(remark, "")
	
	if strings.ToUpper(remark) == remark && hasStrings {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		} else if matched, _ := regexp.Match(`[a-zA-Z]`, []byte(remark)); matched {
			return "Whoa, chill out!"
		}
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if reg2.ReplaceAllString(remark, "") == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
