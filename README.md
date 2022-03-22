# go-readinglist

## Overview

Vue&Go製の読書メモアプリです．Dockerで動きます．読んだ本や論文，WEBサイトについてタイトル，メモ，参考リンク，タグを設定できます．

![image](https://user-images.githubusercontent.com/62528538/159502558-8ebecb3b-ea91-4ac2-9db2-733523ec1720.png)

![image](https://user-images.githubusercontent.com/62528538/159502826-c6c57a25-0f6b-4691-93c8-37f71f09bd54.png)

![image](https://user-images.githubusercontent.com/62528538/159503018-f099ba18-6ec0-45c6-8a2c-957a51f88889.png)




## Set up

```
$ cp .env.sample .env
$ make start
```

## Containers

| コンテナ名 |ポート | 説明 |
|:-----------|:------------:|:------------|
| gorl-frontend       | 8089 | フロントが動きます        |
| gorl-backend     | 8088 | バックエンドが動きます      |
| db       | 3306 | MySQL(mariadb)が動きます        |
| db_test         | 3307 | テスト用のデータベースです          |
