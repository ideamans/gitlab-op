# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

GitLabアカウント管理CLIツール。クライアントとのGitLabリポジトリ共有を効率化するためのコマンドを提供する。

**主なコマンド:**
- `new-group` - グループ作成（ネストしたサブグループ対応）
- `invite` - ユーザーをグループに招待（存在しないユーザーは自動作成）
- `new-project` - グループ内にプロジェクト作成

## 開発コマンド

```bash
# ビルド
go build -o gitlab-op ./...

# または Makefile 経由
make

# ローカルインストール
make install

# テスト
go test ./...
```

## アーキテクチャ

シンプルな単一パッケージ構成のCLIツール。

**ファイル構成:**
- `main.go` - Cobraによるコマンド定義
- `app.go` - GitLab APIクライアントラッパー（`App`構造体）
- `config.go` - 認証設定の読み込み（`~/.gitlab-op/credentials`）
- `group.go` - グループ作成ロジック
- `invite.go` - ユーザー招待ロジック
- `project.go` - プロジェクト作成ロジック

**認証設定:**
```ini
# ~/.gitlab-op/credentials
[default]
url=https://my.gitlab.com/
token=<アクセストークン>
```

プロファイル切替: `GITLAB_OP_PROFILE` 環境変数

## 技術スタック

- Go 1.22+
- `github.com/spf13/cobra` - CLIフレームワーク
- `github.com/xanzy/go-gitlab` - GitLab APIクライアント
- `gopkg.in/ini.v1` - INIファイルパース

## リリース

バージョンタグ（`v*`）のプッシュでGitHub Actionsが起動し、GoReleaserによりクロスプラットフォームバイナリが自動生成される。
