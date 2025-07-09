package main

import (
	"back/infrastructure"
	"back/pkg"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	pkg.LoadEnv()

	if len(os.Args) < 2 {
		FatalLog()
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	dir := "infrastructure/database/migration"

	command := os.Args[1]

	orm := infrastructure.NewGorm()
	switch command {
	case "create":
		if len(os.Args) != 3 {
			FatalLog()
		}
		cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", dir, "-seq", fmt.Sprintf("create_%s_table", os.Args[2]))
		execCommand(cmd)
		log.Println("マイグレーションファイルが作成されました。")
	case "up":
		if len(os.Args) != 2 {
			FatalLog()
		}
		cmd := exec.Command("migrate", "-path", dir, "-database", dsn, "-verbose", "up")
		execCommand(cmd)
		log.Println("マイグレーションが完了しました。")
	case "down":
		if len(os.Args) != 2 {
			FatalLog()
		}
		cmd := exec.Command("migrate", "-path", dir, "-database", dsn, "-verbose", "down", "-all")
		execCommand(cmd)
		log.Println("ロールバックが完了しました。")
	case "reset":
		if len(os.Args) != 2 {
			FatalLog()
		}
		orm.Exec("DROP SCHEMA public CASCADE;")
		orm.Exec("CREATE SCHEMA public;")
		log.Println("リセットが完了しました。")
	case "refresh":
		if len(os.Args) != 2 {
			FatalLog()
		}
		orm.Exec("DROP SCHEMA public CASCADE;")
		orm.Exec("CREATE SCHEMA public;")
		cmd := exec.Command("migrate", "-path", dir, "-database", dsn, "-verbose", "up")
		execCommand(cmd)
		log.Println("リフレッシュが完了しました。")
	default:
		FatalLog()
	}
}

func execCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		FatalLog()
	}
}

func FatalLog() {
	log.Fatal(`無効なコマンドです。以下のいずれかのコマンドを指定してください:

  - ./migrate create [table_name]
  - ./migrate up
  - ./migrate down
  - ./migrate reset
  - ./migrate refresh`)
}
