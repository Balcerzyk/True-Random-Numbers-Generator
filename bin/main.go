	package main

	import (
		"fmt"
		"math"
		"time"
		"bytes"
		"encoding/binary"
		"os"
		"log"
		"github.com/go-vgo/robotgo"

		"gonum.org/v1/plot"
		"gonum.org/v1/plot/plotter"
		"gonum.org/v1/plot/vg"
	)

	func main() {
		var x, y, oldX, oldY int
		var positionLabelX [129]int
		var positionLabelY [129]int
		var k [129]float64
		var r [129]float64
		var bits [129][52]uint8
		var nums [832]uint8

		fmt.Println("GO GO mouseRanger")

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

		for i := uint(0); i < 128; i++ {
			if positionLabelX[i] == positionLabelX[i+1] {
				k[i] = math.Pi / 2;
			} else {
				k[i] = math.Atan2(math.Abs(float64(positionLabelY[i+1]-positionLabelY[i])), math.Abs(float64(positionLabelX[i+1]-positionLabelX[i])))
			}
			r[i] = (k[i] / ( math.Pi / 2 ))
			temp := math.Float64bits(r[i])
			for j := uint(0); j < 52; j++ {

				num := ( i* 52 + j ) / 8;
				temp2 := ( i*52 + j ) % 8;
				bits[i][j] = uint8((temp & (1 << j)) >> j)
				nums[num] = nums[num] | ( bits[i][j] << temp2 );
				fmt.Print(bits[i][j]);


			}
			//fmt.Println(bits[i])
		}

		v := make(plotter.Values, 832)
		for k := range nums {
			//fmt.Println(nums[k]);
			v[k] = float64(nums[k]);
		}

		// Make a plot and set its title.
		p, err := plot.New()
		if err != nil {
			panic(err)
		}
		p.Title.Text = "Histogram"

		// Create a histogram of our values drawn
		// from the standard normal.
		h, err := plotter.NewHist(v, 255)
		if err != nil {
			panic(err)
		}

		h.Normalize(1)
		p.Add(h)
		//fmt.Println(h.Bins[2].Weight);
		fmt.Println(h.Bins);
		fmt.Println(len(h.Bins));
		
		// Save the plot to a PNG file.
		if err := p.Save(64*vg.Inch, 64*vg.Inch, "hist.png"); err != nil {
			panic(err)
		}
		//entropia
		var result float64;

		result=0;
		for _ , el := range h.Bins {
			prob := el.Weight;
			if prob == 0 {
				continue;
			}
			result += prob*math.Log2(prob);
		}
		result=-result;
		fmt.Println();
		fmt.Println(result);

		buf := new(bytes.Buffer)
		bin_err := binary.Write(buf, binary.LittleEndian, nums)
		if bin_err != nil {
			fmt.Println("binary.Write failed:", bin_err)
		}
		fmt.Printf("% x", buf.Bytes())

		file, file_err := os.Create("test.bin")
		defer file.Close()
		if file_err != nil {
			log.Fatal(file_err)
		}
	
		_, file_err = file.Write(buf.Bytes())

		if file_err != nil {
			log.Fatal(err)
		}

	}
