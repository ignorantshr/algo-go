package regexEngine

import (
	"fmt"
	"strconv"
	"strings"
)

type tokenType uint8

const (
	group tokenType = iota
	bracket
	or
	repeat
	literal
	groupUncaptured

	repeatInfinity = -1
)

type repeatPayload struct {
	min int
	max int
	token
}

type token struct {
	tokenType
	value any
}

type parseContext struct {
	pos    int
	tokens []token
}

func parseTokens(regex string) *parseContext {
	ctx := &parseContext{0, make([]token, 0)}
	for ctx.pos < len(regex) {
		process(regex, ctx)
		ctx.pos++
	}
	return ctx
}

func process(regex string, ctx *parseContext) {
	switch regex[ctx.pos] {
	case '(':
		gctx := &parseContext{
			pos:    ctx.pos,
			tokens: []token{},
		}
		parseGroup(regex, gctx)
		ctx.tokens = append(ctx.tokens, token{
			group,
			gctx.tokens,
		})
	case '[':
		parseBracket(regex, ctx)
	case '*', '?', '+': // 这些都代表重复，例如 a* = {0,} a+={1,} a?={0,1}
		parseRepeate(regex, ctx)
	case '{':
		parseRepeateSpecified(regex, ctx)
	case '|':
		parseOr(regex, ctx)
	default:
		// literal
		ctx.tokens = append(ctx.tokens, token{
			literal,
			regex[ctx.pos],
		})
	}
}

func parseGroup(regex string, ctx *parseContext) {
	ctx.pos++ // '('
	for regex[ctx.pos] != ')' {
		process(regex, ctx)
		ctx.pos++
	}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.pos++
	chars := make([]byte, 0)
	literalSet := make(map[byte]bool, 0)
	for regex[ctx.pos] != ']' {
		if regex[ctx.pos] == '-' {
			next := regex[ctx.pos+1]
			pre := chars[len(chars)-1]
			for i := pre; i <= next; i++ {
				literalSet[i] = true
			}
			chars = chars[:len(chars)-1]
			ctx.pos++
		} else {
			chars = append(chars, regex[ctx.pos])
		}
		ctx.pos++
	}

	for _, ch := range chars {
		literalSet[ch] = true
	}

	ctx.tokens = append(ctx.tokens, token{
		bracket,
		literalSet,
	})
}

func parseRepeate(regex string, ctx *parseContext) {
	var min, max int
	switch regex[ctx.pos] {
	case '*':
		min = 0
		max = repeatInfinity
	case '?':
		min = 0
		max = 1
	case '+':
		min = 1
		max = repeatInfinity
	}

	last := ctx.tokens[len(ctx.tokens)-1]  // 得知道这些重复的规则应用到谁身上
	ctx.tokens[len(ctx.tokens)-1] = token{ // 覆盖最后一个 token
		repeat,
		repeatPayload{
			min,
			max,
			last,
		},
	}
}

func parseRepeateSpecified(regex string, ctx *parseContext) {
	ctx.pos++ // skip '{'
	start := ctx.pos
	var min, max int

	for regex[ctx.pos] != '}' {
		ctx.pos++
	}

	pieces := strings.Split(regex[start:ctx.pos], ",")
	switch len(pieces) {
	case 1:
		if v, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err)
		} else {
			min = v
			max = v
		}
	case 2:
		if v, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err)
		} else {
			min = v
		}

		if pieces[1] == "" {
			max = repeatInfinity
		} else if v, err := strconv.Atoi(pieces[1]); err != nil {
			panic(err)
		} else {
			max = v
		}

	default:
		panic(fmt.Sprintf("%s is wrong", regex[start:ctx.pos]))
	}

	last := ctx.tokens[len(ctx.tokens)-1]  // 得知道这些重复的规则应用到谁身上
	ctx.tokens[len(ctx.tokens)-1] = token{ // 覆盖最后一个 token
		repeat,
		repeatPayload{
			min,
			max,
			last,
		},
	}
}

func parseOr(regex string, ctx *parseContext) {
	rctx := &parseContext{
		pos:    ctx.pos,
		tokens: []token{},
	}

	rctx.pos++
	for rctx.pos < len(regex) && regex[rctx.pos] != ')' {
		process(regex, rctx)
		rctx.pos++
	}

	right := token{
		groupUncaptured,
		rctx.tokens,
	}

	left := token{
		groupUncaptured,
		ctx.tokens,
	}

	ctx.pos = rctx.pos

	ctx.tokens = []token{{
		tokenType: or,
		value:     []token{left, right},
	}}
}
