# GitLabアカウント管理ツール

クライアントとのGitLabリポジトリの共有に関するユースケースをCLIコマンドにした。

## 認証

AWS風に`~/.gitlab-op/credentials`ファイルを作成する。

```ini
[default]
url=https://my.gitlab.com/
token=<アクセストークン>
```

`GITLAB_OP_PROFILE`環境変数でセクションを切り替えることができる。

```bash
GITLAB_OP_PROFILE=another gitlab-op
```

## グループ追加

新しいクライアントグループを作成するときに使う。`/`で区切ってサブグループの作成も可能。

```bash
gitlab-op new-group clients/corp Corp株式会社
```

## ユーザーの招待

クライアントグループにユーザーを招待するときに使う。メールアドレスは複数指定できる。

メールアドレスに該当するユーザーがいない場合は作成され、パスワードリセットの案内が送信される。

```bash
gitlab-op invite clients/corp someone1@example.com sameone2@example.com
```

## プロジェクト追加

クライアントグループにプロジェクトを作成するときに使う。

```bash
gitlab-op new-project clients/corp/proj 新しいプロジェクト
```
