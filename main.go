package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name string `yaml:"name"`
	Pass string `yaml:"pass"`
	// uintは整数型
	Age uint `yaml:"age"`
}

func main() {
	user := User{}
	// ファイルを読み込んでbyteで読み込む
	b, err := os.ReadFile("local.yaml")
	if err != nil {
		// NOTE: 本当であればエラーをラップした方が良いが、今回はパス
		log.Print(err)
		return
	}
	// byteで読み込まれたものをデコードし、第二引数の構造体に値を入れ込む
	yaml.Unmarshal(b, &user)
	log.Print(user.Name)
	log.Print(user.Pass)
	log.Print(user.Age)
}
