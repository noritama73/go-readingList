# go-readinglist

## Overview

Vue&Go製の読書メモアプリです．Dockerで動きます．読んだ本や論文，WEBサイトについてタイトル，メモ，参考リンク，タグを設定できます．

## Set up

```
$ make start
```

## Containers

| コンテナ名 |ポート | 説明 |
|:-----------|------------|:------------|
| gorl-frontend       | 8089 | フロントが動きます        |
| gorl-backend     | 8088 | バックエンドが動きます      |
| db       | 3306 | MySQL(mariadb)が動きます        |
| db_test         | 3307 | テスト用のデータベースです          |
