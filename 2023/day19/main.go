package main

import (
	"bytes"
	"strconv"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type WorkflowRule struct {
	HasCond bool
	CondVar int
	CondOp  int
	CondVal int
	ThenVar []byte
}

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	var workflows = make(map[string][]WorkflowRule)

	var start int
	for i, line := range lines {
		if len(line) == 0 {
			start = i + 1
			break
		}

		// px{a<2006:qkq,m>2090:A,rfg}
		// name{condVarCondOpCondVal:thenVar,thenVar}
		// name before {
		name := string(bytes.Split(line, []byte{'{'})[0])
		betweenBraces := bytes.Split(bytes.Split(line, []byte{'{'})[1], []byte{'}'})[0]

		split := bytes.Split(betweenBraces, []byte{','})

		var workflowRules []WorkflowRule
		for _, rule := range split {
			// a<2006:qkq
			// condVarCondOpCondVal:thenVar

			var workflowRule WorkflowRule

			// Check if it has : in it
			if bytes.Contains(rule, []byte{':'}) {
				// condVarCondOpCondVal:thenVar
				condVar := int(rule[0]-'a') + 1
				condOp := int(rule[1])
				condVal := rule[2:bytes.IndexByte(rule, ':')]
				thenVar := rule[bytes.IndexByte(rule, ':')+1:]

				workflowRule.HasCond = true
				workflowRule.CondVar = condVar
				workflowRule.CondOp = condOp
				workflowRule.CondVal, _ = strconv.Atoi(string(condVal))
				workflowRule.ThenVar = thenVar
			} else {
				// thenVar
				workflowRule.ThenVar = rule
			}

			workflowRules = append(workflowRules, workflowRule)
		}
		workflows[name] = workflowRules
	}

	sum := 0
	for i := start; i < len(lines); i++ {
		// {x=787,m=2655,a=1222,s=2876}
		line := lines[i]

		betweenBraces := bytes.Split(bytes.Split(line, []byte{'{'})[1], []byte{'}'})[0]
		split := bytes.Split(betweenBraces, []byte{','})

		vars := make(map[int]int)

		for _, varVal := range split {
			// x=787
			varVar := int(varVal[0]-'a') + 1
			varVal := varVal[2:]

			vars[varVar], _ = strconv.Atoi(string(varVal))
		}

		// Check if this passes the workflow
		wf := workflows["in"]
		for {
			var then string
			for _, rule := range wf {
				if rule.HasCond {
					// Check if the condition is true
					if rule.CondOp == '<' {
						if vars[rule.CondVar] < rule.CondVal {
							// Then we pass this rule
							then = string(rule.ThenVar)
							break
						}
					} else if rule.CondOp == '>' {
						if vars[rule.CondVar] > rule.CondVal {
							// Then we pass this rule
							then = string(rule.ThenVar)
							break
						}
					}
				} else {
					// Then we pass this rule
					then = string(rule.ThenVar)
					break
				}
			}
			if then == "A" {
				for _, val := range vars {
					sum += val
				}
				break
			} else if then == "R" {
				break
			}
			wf = workflows[then]
		}
	}

	return sum
}

var workflows = make(map[string][]WorkflowRule)

// lefts is values left in each variable (x, m, a, s)
func dfs(name string, left map[int][2]int) int {
	if name == "R" {
		return 0
	} else if name == "A" {
		c := 1
		for _, v := range left {
			c *= (v[1] - v[0] + 1)
		}

		return c
	}

	wf := workflows[name]

	total := 0
	for _, workflowRule := range wf {
		if workflowRule.HasCond {
			// Check how many values pass this condition and how many don't
			v := left[workflowRule.CondVar]

			// Do the case where we pass the condition
			if workflowRule.CondOp == '<' {
				newV := [2]int{v[0], workflowRule.CondVal - 1}
				newLeft := make(map[int][2]int)
				for k, v := range left {
					newLeft[k] = v
				}
				newLeft[workflowRule.CondVar] = newV

				total += dfs(string(workflowRule.ThenVar), newLeft)
			} else if workflowRule.CondOp == '>' {
				newV := [2]int{workflowRule.CondVal + 1, v[1]}
				newLeft := make(map[int][2]int)
				for k, v := range left {
					newLeft[k] = v
				}
				newLeft[workflowRule.CondVar] = newV

				total += dfs(string(workflowRule.ThenVar), newLeft)
			}

			// If we don't pass the condition, we modify the lefts
			if workflowRule.CondOp == '<' {
				newV := [2]int{workflowRule.CondVal, v[1]}
				left[workflowRule.CondVar] = newV
			} else if workflowRule.CondOp == '>' {
				newV := [2]int{v[0], workflowRule.CondVal}
				left[workflowRule.CondVar] = newV
			}
		} else {
			total += dfs(string(workflowRule.ThenVar), left)
		}
	}

	return total
}

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		// px{a<2006:qkq,m>2090:A,rfg}
		// name{condVarCondOpCondVal:thenVar,thenVar}
		// name before {
		name := string(bytes.Split(line, []byte{'{'})[0])
		betweenBraces := bytes.Split(bytes.Split(line, []byte{'{'})[1], []byte{'}'})[0]

		split := bytes.Split(betweenBraces, []byte{','})

		var workflowRules []WorkflowRule
		for _, rule := range split {
			// a<2006:qkq
			// condVarCondOpCondVal:thenVar

			var workflowRule WorkflowRule

			// Check if it has : in it
			if bytes.Contains(rule, []byte{':'}) {
				// condVarCondOpCondVal:thenVar
				condVar := int(rule[0]-'a') + 1
				condOp := int(rule[1])
				condVal := rule[2:bytes.IndexByte(rule, ':')]
				thenVar := rule[bytes.IndexByte(rule, ':')+1:]

				workflowRule.HasCond = true
				workflowRule.CondVar = condVar
				workflowRule.CondOp = condOp
				workflowRule.CondVal, _ = strconv.Atoi(string(condVal))
				workflowRule.ThenVar = thenVar
			} else {
				// thenVar
				workflowRule.ThenVar = rule
			}

			workflowRules = append(workflowRules, workflowRule)
		}
		workflows[name] = workflowRules
	}

	total := dfs("in", map[int][2]int{1: {1, 4000}, 13: {1, 4000}, 19: {1, 4000}, 24: {1, 4000}})

	// s := 1
	// for _, v := range total {
	// 	s *= (v[1] - v[0] + 1)
	// }

	return total
}
