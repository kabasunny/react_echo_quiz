package main

import (
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// フルパスで指定されたディレクトリの静的ファイルを提供
	buildPath := filepath.Join("..", "build")
	staticPath := filepath.Join(buildPath, "static")
	indexPath := filepath.Join(buildPath, "index.html")

	// 静的ファイルサーバーを設定
	e.Static("/static", staticPath)

	// HTMLファイルのパスを指定して、ルートへのリクエストに対してindex.htmlを返す
	e.File("/", indexPath)

	// サーバーの起動
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
