package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	var x, y, oldX, oldY int
	var positionLabelX [129]int
	var positionLabelY [129]int
	var k [129]float64
	var r [129]float64
	var bits [129][52]uint8

	fmt.Println("Poruszaj myszka")

	file, err := os.OpenFile("bits.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 129; {
		time.Sleep(20 * time.Millisecond)
		x, y = robotgo.GetMousePos()
		if x == oldX && y == oldY {
			continue
		} else {
			oldX = x
			positionLabelX[i] = x
			oldY = y
			positionLabelY[i] = y
		}
		i++
	}

	/////////////////////////////////////////////// obliczenia
	for i := 0; i < 128; i++ {
		if positionLabelX[i] == positionLabelX[i+1] {
			k[i] = 1.57
		} else {
			k[i] = math.Atan(math.Abs(float64(positionLabelY[i+1]-positionLabelY[i])) / math.Abs(float64(positionLabelX[i+1]-positionLabelX[i])))
		}

		r[i] = (k[i] / 1.57)
		//fmt.Println("r = ", r[i])

		//fmt.Printf("%b", int64(math.Float64bits(r[i])))
		//fmt.Println(strconv.FormatInt(int64(math.Float64bits(r[i])), 2))
		fmt.Println()
		temp := math.Float64bits(r[i])
		//fmt.Println("shift r= ", strconv.FormatInt(int64(temp), 2))
		//fmt.Println("////////////////////")
		for j := uint(0); j < 52; j++ {
			//fmt.Println(strconv.FormatInt(int64((1<<j)>>j), 2))
			bits[i][j] = uint8((temp & (1 << j)) >> j)
			ioutil.WriteFile("bits.txt", []byte(bits[i][:]), 0755)
			//_, err := file.Write(bits[i][:])
			/* 			if err != nil {
				log.Fatal(err)
			} */
		}
		fmt.Println(bits[i])

		//fmt.Println("////////////////////")
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
