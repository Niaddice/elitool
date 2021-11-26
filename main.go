package main

import (
	"io/ioutil"
	"os"
	"time"
)

func main() {
	os.Mkdir("./file", os.ModePerm)
	t := time.Now()
	for i := 0; i < 10; i++ {
		format := t.Format("20060102")
		os.Mkdir("./file/"+format, os.ModePerm)
		os.Create("./file/" + format + "/" + format + ".txt")
		t = time.Now().AddDate(0, -2, 0-i)
	}

	dir, _ := ioutil.ReadDir("./file")
	for _, v := range dir {
		dirTime, _ := time.Parse("20060102", v.Name())
		if time.Now().Sub(dirTime).Hours()/24 > 100 {
			os.RemoveAll("./file/" + v.Name())
		}
	}
}
