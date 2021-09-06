package main


import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Money struct {
	CharCode string
	Name     string
	MaxValue float64
	MinValue float64
	SumValue float64
	DateMax  string
	DateMin  string
}

type Wallet struct {
	sync.Mutex
	Valute  map[string]Money
	counter float64
}

// format date day/month/year

func (w *Wallet) FirstDay(date string) string {

	url := "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=" + date
	doc, err := xmlquery.LoadURL(url)
	n := xmlquery.FindOne(doc, "//ValCurs").SelectElement("//ValCurs").InnerText()

	if err != nil || n == "Error in parameters" {
		fmt.Println("error", err)
		return date

	}
	for _, n := range xmlquery.Find(doc, "//ValCurs/Valute") {

		name := n.SelectElement("//Name").InnerText()
		v := strings.Replace(n.SelectElement("//Value").InnerText(), ",", ".", -1)
		numCode := n.SelectElement("//NumCode").InnerText()

		value, _ := strconv.ParseFloat(v, 32)

		valute := Money{Name: name, SumValue: value, MinValue: value, MaxValue: value, DateMax: date, DateMin: date}
		w.Valute[numCode] = valute

	}
	w.counter++
	date = lessDay(date)

	return date

}

func (w *Wallet) GetMoney(date string) {
	var min, max float64
	var dmax, dmin string

	date = w.FirstDay(date)

	for i := 0; i < 89; i++ {
		url := "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=" + date
		doc, err := xmlquery.LoadURL(url)
		n := strings.TrimSpace(xmlquery.FindOne(doc, "//ValCurs").SelectElement("//ValCurs").InnerText())

		if err != nil || n == "Error in parameters" {
			fmt.Printf( "информация за %s не получена \n", date)
			date = lessDay(date)
			continue
		}

		//v := xmlquery.FindOne(doc, "//Valute")
		for _, n := range xmlquery.Find(doc, "//ValCurs/Valute") {
			name := n.SelectElement("//Name").InnerText()
			v := strings.Replace(n.SelectElement("//Value").InnerText(), ",", ".", -1)
			numCode := n.SelectElement("//NumCode").InnerText()

			value, _ := strconv.ParseFloat(v, 32)




			sum := w.Valute[numCode].SumValue + value

			var valute Money

			if value > w.Valute[numCode].MaxValue {
				max = value
				dmax = date

			} else {
				max = w.Valute[numCode].MaxValue
				dmax = w.Valute[numCode].DateMax
			}

			if value < w.Valute[numCode].MinValue {
				min = value
				dmin = date
			} else {
				min = w.Valute[numCode].MinValue
				dmin = w.Valute[numCode].DateMin
			}

			valute = Money{Name: name, SumValue: sum, MinValue: min, MaxValue: max, DateMin: dmin, DateMax: dmax}

			w.Valute[numCode] = valute

		}

		w.counter++

		if err != nil {
			fmt.Println("err", err)
		}

		date = lessDay(date)

	}
}

func lessDay(s string) string {

	str := strings.Split(s, "/")

	day1, _ := strconv.Atoi(str[0])
	month1, _ := strconv.Atoi(str[1])
	year1, _ := strconv.Atoi(str[2])

	t := time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.UTC)
	r := t.Add(-24 * time.Hour)

	y, m, d := r.Date()
	year := strconv.Itoa(y)
	month := strconv.Itoa(int(m))
	day := strconv.Itoa(d)
	res := day + "/" + month + "/" + year
	return res

}
func main() {
	w := &Wallet{}
	m := make(map[string]Money)
	w.Valute = m
	w.GetMoney("10/11/2020")
	//fmt.Println(w.Valute)
	for _, v := range w.Valute {

		fmt.Println("Name is", v.Name)
		fmt.Println("Min is ", v.MinValue, " Date was", v.DateMin)
		fmt.Println("Max is", v.MaxValue, " Date was", v.DateMax)
		fmt.Println("Mid", v.SumValue/w.counter)
		fmt.Println()
	}

}

