package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	adr := flag.String("adr", "http://google.com", "a string, address")
	num := flag.Int("num", 5, "an integer, count of go's")
	tout := flag.Float64("tout", 500, "a float, timeOut in ms")

	flag.Parse()

	fmt.Println("\tФлаги:")
	fmt.Println("adr:", *adr)
	fmt.Println("num:", *num)
	fmt.Println("tout:", *tout)

	client := http.Client{
		Timeout: time.Duration(time.Duration(*tout) * time.Millisecond),
	}
	channel := make(chan float64)

	start := time.Now()

	for i := 0; i < *num; i++ {
		go func() {
			t := time.Now()
			resp, err := client.Get(*adr)
			if err != nil {
				//fmt.Println(err)
				channel <- -1
				return
			}
			defer resp.Body.Close()

			elapsed := time.Since(t).Seconds()
			//fmt.Println(elapsed)
			channel <- elapsed
		}()
	}

	slice := make([]float64, *num)
	for i := 0; i < *num; i++ {
		msg := <-channel
		slice[i] = msg
	}
	endtime := time.Since(start).Seconds()
	fmt.Println("\nВремя, за которое отработали все запросы (от первого до последнего):", endtime)

	fmt.Println("\n\tВремя всех вопросов:")
	var sum float64
	var count, max int
	// В этом цикле считается максимальное значение, сумма неотрицительных, а также вывод всех значений
	for i := 0; i < *num; i++ {
		fmt.Println(i, "|", slice[i])
		if slice[i] >= 0 {
			sum += slice[i]
			if slice[max] < slice[i] {
				max = i
			}
		} else {
			count++
		}
	}

	min := max
	//А этот цикл нужен для поиска минимального. В функцию решил не выводить..
	for i := 0; i < *num; i++ {
		if slice[min] > slice[i] {
			min = i
		}
	}

	fmt.Println("\nМаксимальное время возвращения ответа:", slice[max], ", под номером", max)
	fmt.Println("Минимальное время возвращения ответа:", slice[min], ", под номером", min)
	fmt.Println("Среднее время на запрос:", sum/float64(*num-count))
	fmt.Println("Количество ответов, которых не дождались / ошибок:", count)

	fmt.Scanln()
}

//timeout := time.After(time.Second * time.Duration(*tout))
//channel := make(chan float64, 5)
//channel <- time.Since(start).Seconds()
