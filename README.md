# 複数のプログラミング言語で配列の宣言、初期化するサンプルコード
このプロジェクトでは、Go言語、TypeScript、PHP、Rust それぞれの配列の宣言や初期化の方法をサンプルコードとしてまとめてみました。  
## サンプルコード一覧
- main.go：Go言語  
- main.ts：TypeScript
- main.php：PHP  
- main.rs：Rust  
## サンプル実行環境の起動方法
サンプルコードは、Dockerで動かせるようにしています。  
事前に Docker はセットアップしておいてください。  
Github から本リポジトリのファイルを取得して以下を実行してください。  
```sh
  ####
  # Mac ならターミナルで以下実行 (あじ太郎の Mac はインテル CPU です..)
  # Windows なら WSL で以下実行 (あじ太郎の WSL 環境は /etc/os-release を参照したら Ubuntu 22.04.2 LTS でした。アップデートせねば)
  ####

  # ファイルを配置したディレクトリへ
$ cd declaring_arrays

  # コンテナ起動
$ docker compose up      # 初回はビルドするので、けっこう時間がかかると思います、気長にお待ちください。。

  ####
  # 以下は別のタブ、または別のウィンドウにて
  ####

$ docker ps
  # arrays_ubuntu のステータスが Up になっていたらOK
  # コンテナに入ります (シェルを実行)
$ docker exec -it arrays_ubuntu bash
```
## サンプルコード実行方法
**Go言語**  
```sh
  # docker exec -it 〜 をしたウィンドウで
$ go run main.go
```
**TypeScript**  
```sh
  # docker exec -it 〜 をしたウィンドウで
$ node -r esbuild-register main.ts
```
**PHP**  
```sh
  # docker exec -it 〜 をしたウィンドウで
$ php main.php
```
**Rust**  
```sh
  # docker exec -it 〜 をしたウィンドウで
$ rustc main.rs  # コンパイル
$ ./main  # 実行
```
