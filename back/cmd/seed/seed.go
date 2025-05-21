package main

import (
	"back/pkg"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	pkg.LoadEnv()

	if len(os.Args) < 2 {
		log.Fatal("引数が不足しています。引数には実行したいファイル名を指定してください。")
	}

	arg := os.Args[1]
	dir := "infrastructure/database/seeder"
	cmd := exec.Command("go", "run", fmt.Sprintf("%s/%s.go", dir, arg))

	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("コマンド実行に失敗しました: %v\n", err)
	}
}
