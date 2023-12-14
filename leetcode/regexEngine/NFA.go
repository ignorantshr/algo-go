package regexEngine

// convert tokens to an epsilon-NFA
// Thompson’s construction

type stateType uint8

const (
	_ stateType = iota
	entrypoint
	final

	epsilon uint8 = 0 // empty character
)

type state struct {
	types       stateType
	transitions map[uint8][]*state
}

// 将每个 token 的 NFA 通过 epsilon 串联起来
func toNFA(ctx *parseContext) *state {
	startState, endState := tokenToNFA(&ctx.tokens[0])
	startState.printStateHierarchy()

	for i := 1; i < len(ctx.tokens); i++ {
		startNext, endNext := tokenToNFA(&ctx.tokens[i])
		endState.transitions[epsilon] = append(
			endState.transitions[epsilon],
			startNext,
		)
		startNext.printStateHierarchy()
		endState = endNext
	}

	start := &state{
		types:       entrypoint,
		transitions: map[uint8][]*state{epsilon: {startState}},
	}

	end := &state{
		types:       final,
		transitions: map[uint8][]*state{},
	}

	endState.transitions[epsilon] = append(
		endState.transitions[epsilon],
		end,
	)

	return start
}

func tokenToNFA(t *token) (*state, *state) {
	start := &state{
		transitions: map[uint8][]*state{},
	}
	end := &state{
		transitions: map[uint8][]*state{},
	}

	switch t.tokenType {
	case literal:
		start.transitions[t.value.(uint8)] = []*state{end}
	case or:
		values := t.value.([]token)
		s1, e1 := tokenToNFA(&values[0]) // left
		s2, e2 := tokenToNFA(&values[1]) // right

		start.transitions[epsilon] = []*state{s1, s2}
		e1.transitions[epsilon] = []*state{end}
		e2.transitions[epsilon] = []*state{end}
	case bracket:
		for ch := range t.value.(map[byte]bool) {
			start.transitions[ch] = []*state{end}
		}
	case group, groupUncaptured: // 内部 token 串联起来
		tokens := t.value.([]token)
		start, end = tokenToNFA(&tokens[0])
		for i := 1; i < len(tokens); i++ {
			startNext, endNext := tokenToNFA(&tokens[i])
			end.transitions[epsilon] = append(end.transitions[epsilon], startNext)
			end = endNext
		}
	case repeat:
		payload := t.value.(repeatPayload)

		if payload.min == 0 { // 代表无需 A 即可直接可以到终态
			start.transitions[epsilon] = []*state{end}
		}

		var copyCount int // 重复的 A 的数量 start->A->A->A->end
		if payload.max == repeatInfinity {
			if payload.min == 0 {
				copyCount = 1
			} else {
				copyCount = payload.min
			}
		} else {
			copyCount = payload.max
		}

		from, to := tokenToNFA(&payload.token)
		start.transitions[epsilon] = append(start.transitions[epsilon], from) // start->A
		for i := 2; i <= copyCount; i++ {
			s, e := tokenToNFA(&payload.token)

			to.transitions[epsilon] = append(to.transitions[epsilon], s) // A->A

			from = s // 指向链尾
			to = e

			if i > payload.min { // max-min
				s.transitions[epsilon] = append(s.transitions[epsilon], end) // A->end
			}
		}

		to.transitions[epsilon] = append(to.transitions[epsilon], end)

		if payload.max == repeatInfinity { // 代表 A 可无限重复
			end.transitions[epsilon] = append(end.transitions[epsilon], from)
		}
	default:
		panic("unknown type of token")
	}

	return start, end
}
