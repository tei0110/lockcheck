package parser

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestComposerParser_Parse(t *testing.T) {
	t.Run("名前とバージョンを取り出せる", func(t *testing.T) {
		lock := `{
			"packages": [
				{"name": "laravel/framework", "version": "11.9.2"},
				{"name": "guzzlehttp/guzzle", "version": "7.8.1"}
			],
			"packages-dev": []
		}`

		got, err := ComposerParser{}.Parse(strings.NewReader(lock))
		if err != nil {
			t.Fatalf("エラーが返った: %v", err)
		}

		want := []Package{
			{Ecosystem: "Packagist", Name: "laravel/framework", Version: "11.9.2"},
			{Ecosystem: "Packagist", Name: "guzzlehttp/guzzle", Version: "7.8.1"},
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v\nwant %+v", got, want)
		}
	})

	t.Run("packages-dev のパッケージは Dev が true になる", func(t *testing.T) {
		lock := `{
			"packages": [
				{"name": "laravel/framework", "version": "11.9.2"}
			],
			"packages-dev": [
				{"name": "phpunit/phpunit", "version": "11.2.5"}
			]
		}`

		got, err := ComposerParser{}.Parse(strings.NewReader(lock))
		if err != nil {
			t.Fatalf("エラーが返った: %v", err)
		}

		want := []Package{
			{Ecosystem: "Packagist", Name: "laravel/framework", Version: "11.9.2", Dev: false},
			{Ecosystem: "Packagist", Name: "phpunit/phpunit", Version: "11.2.5", Dev: true},
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v\nwant %+v", got, want)
		}
	})

	t.Run("バージョンの先頭の v を取り除く", func(t *testing.T) {
		lock := `{
			"packages": [
				{"name": "symfony/console", "version": "v7.1.2"}
			],
			"packages-dev": []
		}`

		got, err := ComposerParser{}.Parse(strings.NewReader(lock))
		if err != nil {
			t.Fatalf("エラーが返った: %v", err)
		}

		want := []Package{
			{Ecosystem: "Packagist", Name: "symfony/console", Version: "7.1.2"},
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v\nwant %+v", got, want)
		}
	})

	t.Run("パッケージが0件でもエラーにしない", func(t *testing.T) {
		lock := `{"packages": [], "packages-dev": []}`

		got, err := ComposerParser{}.Parse(strings.NewReader(lock))
		if err != nil {
			t.Fatalf("エラーが返った: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("0件のはずが %d 件返った: %+v", len(got), got)
		}
	})

	t.Run("JSON として壊れていたらエラーを返す", func(t *testing.T) {
		lock := `{"packages": [`

		_, err := ComposerParser{}.Parse(strings.NewReader(lock))
		if err == nil {
			t.Error("エラーが返るはずが、nil だった")
		}
	})

	t.Run("実際の形に近いファイルを読める", func(t *testing.T) {
		f, err := os.Open("testdata/composer.lock")
		if err != nil {
			t.Fatalf("テスト用のファイルを開けない: %v", err)
		}
		defer f.Close()

		got, err := ComposerParser{}.Parse(f)
		if err != nil {
			t.Fatalf("エラーが返った: %v", err)
		}

		want := []Package{
			{Ecosystem: "Packagist", Name: "guzzlehttp/guzzle", Version: "7.8.1"},
			{Ecosystem: "Packagist", Name: "laravel/framework", Version: "11.9.2"},
			{Ecosystem: "Packagist", Name: "symfony/console", Version: "7.1.2"},
			{Ecosystem: "Packagist", Name: "phpunit/phpunit", Version: "11.2.5", Dev: true},
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v\nwant %+v", got, want)
		}
	})
}
