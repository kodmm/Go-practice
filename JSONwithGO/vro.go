package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)



/* JSONデコード用に構造体定義 */

type VividInfo struct {
	Color *string `json:"color"`
	Weapon string `json:"weapon"`

}

type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Birthday *string `json:"birthday"`       //空文字が初期値なのか意味のある空文字なのか区別がつかないのでポインタで定義する
	Vivid VividInfo `json:"vivid_info"`
	
}

func (p *Person) GetInfo() string {
	if p.Vivid.Color != nil{
		return fmt.Sprintf("%d : %s (%s)", p.Id, p.Name, *p.Vivid.Color)
	} else {
		return fmt.Sprintf("%d : %s", p.Id, p.Name)
	}
}

func main() {
	//JSONファイル読み込み
	bytes, err := ioutil.ReadFile("vro.json")
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", bytes)
	//JSONデコード
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}
	
	//デコードしたデータを表示
	for _, p := range persons {
		fmt.Println(p.GetInfo())
		
	}

}