package main

import (
	"fmt"
	"slices"
)

type files struct {
	pos   int
	len   int
	value int
}

func (d *disk) removeFile(file files) {
	for pos, value := range *d {
		if value == file.value {
			(*d)[pos] = -1
		}
	}
}

func (d *disk) emptyDefragments(m diskMap) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	// keys = keys[1:]

	// empty space map with startPos and length
	emptySpace := make(map[int]int)
	var length int = 0
	for i := 0; i < len(*d); i++ {
		if (*d)[i] == -1 {
			length++
		} else {
			if length > 0 {
				emptySpace[i-length] = length
				length = 0
			}
		}
	}

	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		value := m[key]
		filePos := value.index
		fileLen := value.v

		spaces := make([]int, 0, len(emptySpace))
		for k := range emptySpace {
			spaces = append(spaces, k)
		}
		slices.Sort(spaces)

		spacePos := -1
		for _, space := range spaces {
			if emptySpace[space] >= fileLen {
				spacePos = space
				break
			}
		}
		spaceLen := emptySpace[spacePos]

		if spacePos == -1 || filePos < spacePos {
			continue
		}

		// Modify spaces map to reflect the changes

		// Delete the current empty space that's replaced by our moved file

		delete(emptySpace, spacePos)
		// If the file doesn't take up the whole empty space, add the remaining empty space
		if fileLen < spaceLen {
			emptySpace[spacePos+fileLen] = spaceLen - fileLen
		}
		// Add the empty space where the file previously was,
		// if there's an empty space after the file, merge them
		// otherwise, create a new empty space
		if _, f := emptySpace[filePos+fileLen]; f {
			emptySpace[filePos] = emptySpace[filePos] + emptySpace[filePos+fileLen]
			delete(emptySpace, filePos+fileLen)
		} else {
			emptySpace[filePos] = fileLen
		}

		// Place file in the map
		value.index = spacePos
		// Place file in the disk and
		for i := 0; i < fileLen; i++ {
			(*d)[spacePos+i] = key
			(*d)[filePos+i] = -1
		}
	}
}

func printDisk(d disk) {
	for _, v := range d {
		if v == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Print(v)
	}
	fmt.Println()
}

// find next empty space pos and length
// func (d *disk) findNextEmptyWithLength(startPos int) (int, int) {
// 	var pos int = -1
// 	var length int = 0

// 	for i := startPos; i < len(*d); i++ {
// 		if (*d)[i] == -1 {
// 			if pos == -1 {
// 				pos = i
// 			}
// 			length++
// 		} else {
// 			if pos != -1 {
// 				break
// 			}
// 		}
// 	}

// 	return pos, length
// }

var i int = 0

func doPartTwo(input []byte) int {
	var disk disk
	var diskMap diskMap = make(diskMap)

	for position, value := range input {
		isSpace := position%2 == 1

		valueInt := fastAtoi(value)

		var id int = -1 // space
		if !isSpace {
			id = getId(position)
			diskMap.preAdd(id, valueInt)
		}

		for i := 0; i < valueInt; i++ {
			disk = append(disk, id)
		}
	}

	diskMap.locate(disk)
	disk.emptyDefragments(diskMap)

	var checksum int = 0
	for i, value := range disk {
		if value == -1 {
			continue
		}
		checksum += i * value
	}

	return checksum
}
