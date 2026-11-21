// Package parser は、lock ファイルを読んでパッケージの一覧を取り出します。
package parser

import "io"

// Package は、lock ファイルに書かれた1つのパッケージです。
type Package struct {
	// Ecosystem は、パッケージの配布元の種類です。
	// 脆弱性データベース OSV での呼び方に合わせます（"Packagist"、"npm" など）。
	Ecosystem string
	// Name は、パッケージ名です（例: "laravel/framework"）。
	Name string
	// Version は、先頭の "v" を取り除いたバージョンです（例: "11.9.2"）。
	Version string
	// Dev は、開発用の依存（composer.lock の packages-dev）なら true です。
	Dev bool
}

// Parser は、lock ファイルの中身からパッケージの一覧を取り出します。
// lock ファイルの種類ごとに、この interface を満たす型を1つ作ります。
//
// ファイルのパスではなく io.Reader を受け取るので、
// テストでは strings.NewReader で文字列をそのまま渡せます。
type Parser interface {
	Parse(r io.Reader) ([]Package, error)
}
