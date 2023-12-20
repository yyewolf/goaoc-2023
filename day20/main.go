package main

import (
	"bytes"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Node struct {
	Type    int
	Targets [][]byte
	FMem    int
	CMem    map[string]int
}

func (r *Node) Equals(o *Node) bool {
	if r.Type != o.Type {
		return false
	}

	if len(r.Targets) != len(o.Targets) {
		return false
	}

	for i := 0; i < len(r.Targets); i++ {
		if !bytes.Equal(r.Targets[i], o.Targets[i]) {
			return false
		}
	}

	if r.FMem != o.FMem {
		return false
	}

	if len(r.CMem) != len(o.CMem) {
		return false
	}

	for k, v := range r.CMem {
		if o.CMem[k] != v {
			return false
		}
	}

	return true
}

var m = make(map[string]*Node)

const (
	TYPE_BROADCAST = 1
	TYPE_FLIPFLOP  = 2
	TYPE_CONJ      = 3
)

func doAction(from, to string, data int) (out []Pulse) {
	r := m[to]

	if r == nil {
		return
	}

	var outSignal int

	switch r.Type {
	case TYPE_BROADCAST:
		// broadcast
		outSignal = data
	case TYPE_FLIPFLOP:
		// flip flop
		// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
		// However, if a flip-flop module receives a low pulse, it flips between on and off.
		// If it was off, it turns on and sends a high pulse.
		// If it was on, it turns off and sends a low pulse.
		if data == 1 {
			return
		}
		if r.FMem == 0 {
			r.FMem = 1
		} else {
			r.FMem = 0
		}
		outSignal = r.FMem
	case TYPE_CONJ:
		// conjunction
		// Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
		// they initially default to remembering a low pulse for each input.
		// When a pulse is received, the conjunction module first updates its memory for that input.
		// Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.
		r.CMem[from] = data
		allHigh := true
		for _, k := range r.CMem {
			if k == 0 {
				allHigh = false
				break
			}
		}

		if !allHigh {
			outSignal = 1
		}
	}

	for _, target := range r.Targets {
		// doAction(string(target), data)
		out = append(out, Pulse{to, string(target), outSignal})
	}

	return
}

type Pulse struct {
	From string
	To   string
	Typ  int
}

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	var io = make(map[string][]string)
	m = make(map[string]*Node)

	for _, line := range lines {
		// split on " -> "
		parts := bytes.Split(line, []byte{' ', '-', '>', ' '})
		before := parts[0]
		targets := parts[1:]

		// split targets by ", "
		targets = bytes.Split(bytes.Join(targets, []byte{}), []byte{',', ' '})

		var name string
		r := &Node{
			Type:    0,
			Targets: targets,
		}

		if bytes.Equal(before, []byte{'b', 'r', 'o', 'a', 'd', 'c', 'a', 's', 't', 'e', 'r'}) {
			name = "broadcaster"
			r.Type = 1
		} else if before[0] == '%' {
			name = string(before[1:])
			r.Type = 2
		} else if before[0] == '&' {
			name = string(before[1:])
			r.Type = 3
			r.CMem = make(map[string]int)
		}

		for _, target := range targets {
			io[string(target)] = append(io[string(target)], name)
		}

		m[name] = r
	}

	for k, v := range m {
		if v.Type == 3 {
			for _, target := range io[k] {
				v.CMem[string(target)] = 0
			}
			m[k] = v
		}
	}

	var q []Pulse
	var highs int
	var lows int

	for btnPress := 0; btnPress < 1000; btnPress++ {
		q = append(q, Pulse{"button", "broadcaster", 0})

		for len(q) > 0 {
			msg := q[0]
			q = q[1:]

			if msg.Typ == 0 {
				lows++
			} else {
				highs++
			}

			msgs := doAction(msg.From, msg.To, msg.Typ)

			q = append(q, msgs...)
		}
	}

	return highs * lows
}

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	var io = make(map[string][]string)
	m = make(map[string]*Node)

	for _, line := range lines {
		// split on " -> "
		parts := bytes.Split(line, []byte{' ', '-', '>', ' '})
		before := parts[0]
		targets := parts[1:]

		// split targets by ", "
		targets = bytes.Split(bytes.Join(targets, []byte{}), []byte{',', ' '})

		var name string
		r := &Node{
			Type:    0,
			Targets: targets,
		}

		if bytes.Equal(before, []byte{'b', 'r', 'o', 'a', 'd', 'c', 'a', 's', 't', 'e', 'r'}) {
			name = "broadcaster"
			r.Type = 1
		} else if before[0] == '%' {
			name = string(before[1:])
			r.Type = 2
		} else if before[0] == '&' {
			name = string(before[1:])
			r.Type = 3
			r.CMem = make(map[string]int)
		}

		for _, target := range targets {
			io[string(target)] = append(io[string(target)], name)
		}

		m[name] = r
	}

	for k, v := range m {
		if v.Type == 3 {
			for _, target := range io[k] {
				v.CMem[string(target)] = 0
			}
			m[k] = v
		}
	}

	required := io["rx"][0]
	requiredConds := io[io["rx"][0]]

	var cycles = make(map[string]int)

	var q []Pulse

	for btnPress := 1; len(cycles) != len(requiredConds); btnPress++ {
		q = append(q, Pulse{"button", "broadcaster", 0})

		for len(q) > 0 {
			msg := q[0]
			q = q[1:]
			q = append(q, doAction(msg.From, msg.To, msg.Typ)...)

			for _, v := range requiredConds {
				_, ok := cycles[v]
				if !ok && m[required].CMem[v] == 1 {
					cycles[v] = btnPress
				}
			}
		}
	}

	// do lcm of all cycles
	arr := make([]int, 0, len(cycles))
	for _, v := range cycles {
		arr = append(arr, v)
	}

	return lcmOfArray(arr)
}

// Function to calculate the GCD of two numbers using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the LCM of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate the LCM of an array of integers
func lcmOfArray(arr []int) int {
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}

	return result
}
