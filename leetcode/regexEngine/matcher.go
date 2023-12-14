package regexEngine

import "fmt"

const (
	startOfText byte = 1
	endOfText   byte = 2
)

var debug bool

type Matcher struct {
	*state
}

func Compile(expr string) *Matcher {
	m := &Matcher{}

	ctx := parseTokens(expr)
	m.state = toNFA(ctx)

	m.state.printStateHierarchy()
	return m
}

func (m *Matcher) MathchAll(s string) bool {
	return m.state.check(s, -1)
}

func (s *state) check(input string, pos int) bool {
	ch := getChar(input, pos)
	if ch == endOfText && s.types == final {
		// fmt.Printf("%c\n", ch)
		return true
	}

	if states := s.transitions[ch]; len(states) > 0 {
		nextState := states[0]
		// fmt.Printf("%c->", ch)
		if nextState.check(input, pos+1) {
			return true
		}
	}

	for _, state := range s.transitions[epsilon] {
		// fmt.Printf("%d ->", state.types)
		if state.check(input, pos) {
			return true
		}

		if ch == startOfText && state.check(input, pos+1) {
			return true
		}
	}

	return false
}

func getChar(input string, pos int) byte {
	if pos >= len(input) {
		return endOfText
	}
	if pos < 0 {
		return startOfText
	}

	return input[pos]
}

func (s *state) printStateHierarchy() {
	if !debug {
		return
	}

	visited := make(map[*state]bool)
	_printStateHierarchy(s, visited, "")
	println()
}

func _printStateHierarchy(s *state, visited map[*state]bool, indent string) {
	if visited[s] {
		fmt.Printf("%sType: %d [%p](already visited)\n", indent, s.types, s)
		return
	}

	visited[s] = true
	fmt.Printf("%sType: %d [%p]\n", indent, s.types, s)

	for input, nextStates := range s.transitions {
		fmt.Printf("%s- Input: %d\n", indent, input)
		for _, nextState := range nextStates {
			_printStateHierarchy(nextState, visited, indent+"  ")
		}
	}
}
