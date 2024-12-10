package main

func fastAtoi(buf byte) (n int) {
	return int(buf - '0')
}

func getId(pos int) int {
	return pos / 2
}

type disk []int
type diskMap map[int]*value

type value struct {
	v     int
	index int
}

func (d *diskMap) preAdd(key int, v int) {
	(*d)[key] = &value{
		v:     v,
		index: -1,
	}
}

func (d *diskMap) locate(l disk) {
	for idx, val := range l {
		v, found := (*d)[val]
		if found && val != -1 && v.index == -1 {
			v.index = idx
		}
	}
}

func (d *disk) empty() {
	// Calculate the new length of the disk after removing all empty spaces
	emptiedDiskLen := 0
	for _, value := range *d {
		if value != -1 {
			emptiedDiskLen++
		}
	}

	emptiedDisk := []int{}
	for _, value := range *d {
		if emptiedDiskLen == len(emptiedDisk) {
			break
		}

		if value == -1 {
			// This is space, we need to remove the last value from the disk and place it here
			lastValue, indexOfLastValue := d.getLastNonEmpty()
			*d = (*d)[:indexOfLastValue]
			emptiedDisk = append(emptiedDisk, lastValue)
		} else {
			emptiedDisk = append(emptiedDisk, value)
		}
	}

	*d = emptiedDisk
}

func (d *disk) getLastNonEmpty() (int, int) {
	var lastValue, indexOfLastValue int = -1, -1
	for i := len(*d) - 1; i >= 0; i-- {
		if (*d)[i] != -1 {
			lastValue, indexOfLastValue = (*d)[i], i
			break
		}
	}

	return lastValue, indexOfLastValue
}

func doPartOne(input []byte) int {
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
	disk.empty()

	var checksum int = 0
	for i, value := range disk {
		checksum += i * value
	}

	return checksum
}
