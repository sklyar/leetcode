package goal_parser_interpretation

import "strings"

// PS... I know the solution is overengineered

func interpret(command string) string {
	r := newReplacer(command)

	state := parseGoal
	for state != nil {
		state = state(r)
	}

	return r.str()
}

type replacer struct {
	r *strings.Reader
	b strings.Builder
}

func newReplacer(command string) *replacer {
	return &replacer{
		r: strings.NewReader(command),
		b: strings.Builder{},
	}
}

type stateFn func(*replacer) stateFn

func (r *replacer) next() (rune, bool) {
	ch, _, err := r.r.ReadRune()
	if err != nil {
		return 0, false
	}
	return ch, true
}

func (r *replacer) peek() (rune, bool) {
	ch, ok := r.next()
	if !ok {
		return 0, false
	}

	if err := r.r.UnreadRune(); err != nil {
		return 0, false
	}

	return ch, true
}

func (r *replacer) write(str string) {
	r.b.WriteString(str)
}

func (r *replacer) str() string {
	return r.b.String()
}

func parseGoal(r *replacer) stateFn {
	for {
		ch, ok := r.next()
		if !ok {
			return nil
		}

		switch ch {
		case 'G':
			r.write("G")
		case '(':
			ch, ok := r.peek()
			if !ok {
				return nil
			}

			if ch == ')' {
				return parseO
			} else if ch == 'a' {
				return parseAl
			}
		}
	}
}

func parseO(r *replacer) stateFn {
	ch, ok := r.next()
	if !ok {
		return nil
	}

	if ch != ')' {
		return nil
	}

	r.write("o")

	return parseGoal
}

func parseAl(r *replacer) stateFn {
	var str string
	for i := 0; i < 3; i++ {
		ch, ok := r.next()
		if !ok {
			return nil
		}

		str += string(ch)
	}

	if str != "al)" {
		return nil
	}

	r.write("al")

	return parseGoal
}
