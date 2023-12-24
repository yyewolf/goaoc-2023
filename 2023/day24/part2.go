package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))

	// format: px py pz @ vx vy vz
	var nodes []*node
	for _, line := range lines {
		var n node
		_, err := fmt.Sscanf(string(line), "%d, %d, %d @ %d, %d, %d", &n.px, &n.py, &n.pz, &n.vx, &n.vy, &n.vz)
		if err != nil {
			panic(err)
		}

		nodes = append(nodes, &n)
	}

	// output to temp.z3
	f, err := os.Create("temp.z3")
	if err != nil {
		panic(err)
	}

	f.WriteString("(declare-const rx Real)\n")
	f.WriteString("(declare-const ry Real)\n")
	f.WriteString("(declare-const rz Real)\n")
	f.WriteString("(declare-const rvx Real)\n")
	f.WriteString("(declare-const rvy Real)\n")
	f.WriteString("(declare-const rvz Real)\n")

	// Must resolve this equations at the same time :
	for i, n := range nodes {
		// rx + rvx*t = px + vx*t
		// ry + rvy*t = py + vy*t
		// rz + rvz*t = pz + vz*t
		f.WriteString(fmt.Sprintf("(declare-const t%d Real)\n", i))
		f.WriteString(fmt.Sprintf("(assert (= (+ rx (* rvx t%d)) (+ %d (* %d t%d))))\n", i, n.px, n.vx, i))
		f.WriteString(fmt.Sprintf("(assert (= (+ ry (* rvy t%d)) (+ %d (* %d t%d))))\n", i, n.py, n.vy, i))
		f.WriteString(fmt.Sprintf("(assert (= (+ rz (* rvz t%d)) (+ %d (* %d t%d))))\n", i, n.pz, n.vz, i))
		if i > 1 {
			break
		}
	}

	f.WriteString("(check-sat)\n")
	f.WriteString("(get-value (rx ry rz))\n")

	f.Close()

	c := exec.Command("z3", "temp.z3")
	out, err := c.CombinedOutput()
	if err != nil {
		return 0
	}

	// Parse output to find rx, ry, rz

	rx, ry, rz := 0, 0, 0
	fmt.Sscanf(string(out), "sat\n((rx %d.0)\n (ry %d.0)\n (rz %d.0))", &rx, &ry, &rz)

	return rx + ry + rz
}
