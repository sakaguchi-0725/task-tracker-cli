# Task Tracker CLI

## アプリ概要
Task Tracker CLIは、シンプルなタスク管理を行うためのコマンドラインアプリケーションです。<br>
タスクの作成、更新、削除、一覧表示などの基本的な機能を提供し、JSONファイルをデータベースとして使用しています。<br>

## 使用技術
言語: Golang<br>
ライブラリ:<br>
・Cobra<br>
・Gomock<br>
・Testify<br>

## コマンド一覧
### タスク一覧取得
```shell
task-tracker list
```

出力例<br>
```shell
+--------------------------------------+-------+-------------+-------------+
|                  ID                  | TITLE | DESCRIPTION |   STATUS    |
+--------------------------------------+-------+-------------+-------------+
| b9aed851-3768-4a97-bed0-0877e6fdf1e6 | Task1 |             | not-started |
| 2559759f-c6a5-4e94-8312-85badfc6fce4 | Task2 |             | not-started |
| 3de64b7e-a08c-47de-bbc9-2d9fc07b436a | Task3 |             | not-started |
+--------------------------------------+-------+-------------+-------------+
```

### タスク作成
```shell
task-tracker create -t "タスク名" -d "詳細"
```

### タスク更新
```shell
task-tracker update -i "タスクID" -t "タスク名" -d "詳細" -s "in-progress"
```
ステータスは以下です。<br>
・not-started<br>
・in-progress<br>
・done<br>

### タスク削除
```shell
task-tracker delete -i "タスクID"
```