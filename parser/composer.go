package parser

import (
	"errors"
	"io"
)

// EcosystemPackagist は、Composer のパッケージを表す OSV での呼び方です。
const EcosystemPackagist = "Packagist"

// ComposerParser は、composer.lock を読みます。
type ComposerParser struct{}

// ComposerParser が Parser を満たすことを、コンパイル時に確かめます。
var _ Parser = ComposerParser{}

// Parse は、composer.lock の中身からパッケージの一覧を取り出します。
//
// 返す順番は、packages（本番用）のあとに packages-dev（開発用）です。
// それぞれの中では、ファイルに書かれた順番を保ちます。
//
// TODO(段階1): ここを実装する。composer_test.go のテストがすべて通れば完成。
//
// 手順の目安:
//  1. composer.lock の形に合わせた struct を作る
//     （必要なのは "packages" と "packages-dev"、その中の "name" と "version" だけ）
//  2. encoding/json の json.NewDecoder(r).Decode(&...) で読み込む
//  3. 読み込んだものを []Package に詰め替える
//     （Ecosystem は EcosystemPackagist、Version は先頭の "v" を取り除く）
func (ComposerParser) Parse(r io.Reader) ([]Package, error) {
	return nil, errors.New("まだ実装していません")
}
