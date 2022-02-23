package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

// バージョン
var version = ""

func main() {

	// コマンド定義
	var v bool
	flag.BoolVar(&v, "v", false, "convini-cli version")

	flag.Parse()

	if v {
		fmt.Printf("convini-cli %s`\n", getVersion())
	}
}

/*
 * バージョン情報を取得します。
 * 実行時に以下のオプションを設定することでバージョンが表示されます。
 * go build -ldflags '-s -w -X main.version={表示するバージョンを指定}'
 */
func getVersion() string {
	if version != "" {
		// バージョン情報が設定されている場合
		return version
	}

	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return i.Main.Version
}
