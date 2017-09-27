package main

import (
	"encoding/csv"
	"fmt"
	"github.com/NeuronEvolution/pkg"
	"os"
	"strings"
)

func value(v string) string {
	return "'" + v + "'"
}

func main() {
	f, err := os.Open("./szA.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	var s = ""
	s += "insert into `fin-stock`.stock (stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date," +
		"company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en," +
		"create_time,update_time)\n values \n"

	for i := 1; i < len(records); i++ {
		var v []string
		l := records[i]

		v = append(v, value(l[0]))
		v = append(v, value("szA"))
		v = append(v, value(l[3]))
		v = append(v, value(l[4]))
		v = append(v, value(""))
		v = append(v, value(l[5]))
		v = append(v, value(""))
		v = append(v, value(""))
		v = append(v, value(l[9]))
		v = append(v, value(l[8]))
		v = append(v, value(l[7]))
		v = append(v, value(""))
		v = append(v, value(l[6]))
		v = append(v, value(""))
		v = append(v, "CURRENT_TIMESTAMP")
		v = append(v, "CURRENT_TIMESTAMP")

		if i == len(records)-1 {
			s += "(" + strings.Join(v, ",") + ")\n"
		} else {
			s += "(" + strings.Join(v, ",") + "),\n"
		}
	}

	pkg.NewFile("szA.sql", []byte(s))
}
