package main

// parse reads the file given in input and return the elves
// func parse(file string) ([]*Elf, error) {
// 	f, err := os.ReadFile(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	lines := strings.Split(string(f), "\n")
// 	currentElf := &Elf{}
// 	elves := []*Elf{}
// 	for _, line := range lines {
// 		if line == "" {
// 			elves = append(elves, currentElf)
// 			currentElf = &Elf{}
// 		} else {
// 			amount, err := strconv.Atoi(line)
// 			if err != nil {
// 				return nil, err
// 			}
// 			currentElf.Calories = append(currentElf.Calories, amount)
// 		}
// 	}
// 	return elves, nil
// }
