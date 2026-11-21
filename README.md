# lockcheck

lock ファイルを読んで、使っているパッケージに既知の脆弱性がないかを調べるコマンドラインツールです。
脆弱性の情報は [OSV](https://osv.dev/)（Google が運営する脆弱性データベース）に問い合わせます。

- 名前の `lockcheck` は仮のものです
- Go の学習を兼ねて作っています。同じ目的の本格的な道具としては、Google の [osv-scanner](https://github.com/google/osv-scanner) があります

## 使い方（完成時の予定）

```sh
lockcheck composer.lock            # 検査して、結果を表で表示する
lockcheck --no-dev composer.lock   # 開発用の依存を除いて検査する
lockcheck --json composer.lock     # 結果を JSON で出力する
```

脆弱性が1件でも見つかると、終了コード 1 で終わります。CI に組み込んで、見つかったらビルドを止められます。

## 進め方

テストを先に用意してあります。テストがすべて通るように中身を書いていきます。

| 段階 | 内容 | 状態 |
|---|---|---|
| 1 | `composer.lock` を読んで、パッケージの一覧を返す | テストあり・実装待ち |
| 2 | OSV にまとめて問い合わせ、脆弱性を受け取る | 未着手 |
| 3 | 結果を表で表示する。終了コード、`--no-dev`、`--json` | 未着手 |
| 4 | `package-lock.json`、`Cargo.lock`、`go.sum` に対応する | 未着手 |

### 段階1の進め方

```sh
go test ./parser/
```

最初は6件中5件が失敗します。`parser/composer.go` の `Parse` を実装して、すべて通れば完成です。
手順の目安は `Parse` のコメントに書いてあります。

## 構成

```
lockcheck/
├── main.go            引数を受け取り、各部品をつなぐだけ（段階3で追加）
├── parser/            lock ファイルを読んで、パッケージの一覧にする
│   ├── parser.go        Package 型と Parser interface
│   ├── composer.go      composer.lock の読み取り
│   └── testdata/        テスト用の lock ファイル
├── osv/               OSV への問い合わせ（段階2で追加）
└── report/            結果の表示（段階3で追加）
```

- **lock ファイルの種類ごとの違いは `parser` に閉じ込めます**。どの種類も、共通の `Package`（配布元の種類・名前・バージョン）に変換します。OSV への問い合わせと表示は、lock ファイルの種類を知りません
- **`Parser` は `io.Reader` を受け取ります**。テストではファイルを用意せず、文字列をそのまま渡せます
- **開発用の依存も標準で検査します**。開発用のパッケージが乗っ取られる被害も多いためです。外したいときは `--no-dev` を使います

## 必要なもの

- Go 1.24 以上
