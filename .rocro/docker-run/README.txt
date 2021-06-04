このディレクトリにあるファイル群は、
docker-run 機能のテスト (codelift#2284) 及びそのテスト環境の準備に使用するものです。

個々のファイルの説明:
    README.txt                      このファイル
    apache2.Dockerfile              docker-run で起動するコンテナを作る Dockerfile
    build-and-push.sh               上記コンテナをビルドして public repository にプッシュ
    create-public-repository.sh     テスト用 ecr-public repository を作成
    delete-public-repository.sh     テスト用 ecr-public repository を削除
    describe-public-repository.sh   テスト用 ecr-public repository の説明表示

テストの主要パラメータ:
    git repository                  bitbucker.org/tetrafolium/algebird
    git branch                      inspecode.docker-run
    aws zone                        us-east-1
    ecr-public repositoryName       docker-run-test
    docker imageTags                apache2
