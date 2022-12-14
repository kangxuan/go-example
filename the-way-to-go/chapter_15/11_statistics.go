package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const form1 = `<html><body><form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></html></body>`
const error1 = `<p class="error">%s</p>`

var pageTop = ""
var pageBottom = ""

func main() {
	//编写一个网页程序，可以让用户输入一连串的数字。计算出这些数字的均值和中值，并且打印出来。
	http.HandleFunc("/", HomePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("服务启动失败", err)
	}
}

func HomePage(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := request.ParseForm()
	_, err = fmt.Fprint(w, pageTop, form1)
	if err != nil {
		return
	}
	if err != nil {
		_, err2 := fmt.Fprintf(w, error1, err)
		if err2 != nil {
			return
		}
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			_, err = fmt.Fprint(w, formatStats(stats))
			if err != nil {
				return
			}
		} else if message != "" {
			_, err = fmt.Fprintf(w, error1, message)
			if err != nil {
				return
			}
		}
	}
	_, err = fmt.Fprintf(w, pageBottom)
	if err != nil {
		return
	}
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	var text string
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		// 处理网页中输入中文符号
		if strings.Contains(slice[0], "&#65292") {
			text = strings.Replace(slice[0], "&#65292;", " ", -1)
		} else {
			text = strings.Replace(slice[0], ",", " ", -1)
		}

		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

func sum(numbers []float64) (total float64) {
	for _, X := range numbers {
		total += X
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
