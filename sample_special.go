package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"log"
	"strconv"
)


/* JSONデコード用に構造体を定義 */
type Person struct {
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
}



func (person *Person) NewPerson() {
	/* 氏名の出力 */
	fmt.Printf("氏名 : %s\n", person.Name)

	/* 性別の出力 */
	if person.Gender == "male" {
		fmt.Println("性別 : 男性")
	} else {
		fmt.Println("性別 : 女性")
	}

	/* 年齢の出力 */
	age := person.BirthDay()
	fmt.Println(age)
}

func (person *Person) BirthDay() string {
	const layout string = "2006-01-02"
	birthdayJST, error := time.ParseInLocation(layout, person.Birthday, time.Local)
	if error != nil {
		log.Fatal(error)
	}

	timeNow, _ := strconv.Atoi(time.Now().Format("20060102"))
	personBirthday, _ := strconv.Atoi(birthdayJST.Format("20060102"))
	
	personAge := int((timeNow - personBirthday) / 10000)
	return fmt.Sprintf("年齢 : %d", personAge)
}

/* 現日本時刻を取得するためにlocationを定義 */
var location string = "Asia/Tokyo"

/* 現日本時刻を取得するためにtime.Localの変更処理 */
func init() {
	loc, err := time.LoadLocation(location)
    if err != nil {
        loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	//APIからデータ取得
	const url = "*************"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	defer response.Body.Close()

	// 取得したデータをbyte型に格納
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// JSONをパースする。
	var person Person
	if err := json.Unmarshal(body, &person); err != nil {
		log.Fatal(err)
	}

	// 出力処理
	person.NewPerson()

}
