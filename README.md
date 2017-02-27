# maekawa
Maekawa is a client fotr AWS CloudWatch Events has idempotence.

Japanese version is under the English version.
(日本語のREADMEが英語の下にあります)

## how to install
Download latest binary for your archtecture from releaase page.

https://github.com/unasuke/maekawa/releases

And put that in your `$PATH` directory.

Or, run this.

```shell
$ go get github.com/unasuke/maekawa
```

## usage
### 1. install awscli and run `aws configure`
- [Installing the AWS Command Line Interface - AWS Command Line Interface](http://docs.aws.amazon.com/cli/latest/userguide/installing.html)
- [Configuring the AWS Command Line Interface - AWS Command Line Interface](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)

### 2. edit `config.yml`
```yaml
rules:
  -
    description: Execute every 5 minutes
    event_pattern:
    name: EveryFiveMinutes
    role_arn:
    schedule_expression: "cron(0/5 * * * ? *)"
    state: ENABLED
    targets:
      -
        arn: arn:aws:lambda:ap-northeast-1:000000000000:function:some-lambda-function
        id: 05c99eea-4059-4294-8583-bb2d00df4226
        input:
        input_path:
  -
    name: TestEvent
    schedule_expression: "rate(2 day)"
    state: ENABLED
    targets:
      -
        arn: arn:aws:lambda:ap-northeast-1:000000000000:function:another-lambda-function
        id: 4c4d91c5-5f8b-4d5f-b28f-3fac49c835ad
        input: |
          { "input": "json input" }
        input_path:
      -
        arn: arn:aws:lambda:ap-northeast-1:00000000000:function:another-lambda-function
        id: 3bde837b-3093-487c-9328-a677e05ae75d
        input: |
          { "input": "some input" }
        input_path:
```

### 3. run maekawa
```shell
$ maekawa --region=ap-northeast-1 --apply --dry-run # check what will changes
$ maekawa --region=ap-northeast-1 --apply # apply config.yml
```

## options
### --region
Specify aws region

### --apply
Apply your config to AWS CloudWatch Events.

### --dry-run
Show what will changes without apply.

### --file, -f
Specify config file path. (default `config.yml`)

### --help, -h
Show options.

-----------------------------------

# maekawa
maekawaは冪等性を持つAWS CloudWatch Eventsのclientです。

## インストール方法
最新版リリースの、適したアーキテクチャ向けのバイナリをreleaseページからダウンロードしてください。

https://github.com/unasuke/maekawa/releases

その後、ダウンロードしたバイナリを`$PATH`の通っている場所に置いてください。

もしくは、以下のコマンドを実行してください。

```shell
$ go get github.com/unasuke/maekawa
```

## 使い方
### 1. awscliのインストールと `aws configure`の実行
- [Installing the AWS Command Line Interface - AWS Command Line Interface](http://docs.aws.amazon.com/cli/latest/userguide/installing.html)
- [Configuring the AWS Command Line Interface - AWS Command Line Interface](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)

### 2. `config.yml`を編集する
```yaml
rules:
  -
    description: Execute every 5 minutes
    event_pattern:
    name: EveryFiveMinutes
    role_arn:
    schedule_expression: "cron(0/5 * * * ? *)"
    state: ENABLED
    targets:
      -
        arn: arn:aws:lambda:ap-northeast-1:000000000000:function:some-lambda-function
        id: 05c99eea-4059-4294-8583-bb2d00df4226
        input:
        input_path:
  -
    name: TestEvent
    schedule_expression: "rate(2 day)"
    state: ENABLED
    targets:
      -
        arn: arn:aws:lambda:ap-northeast-1:000000000000:function:another-lambda-function
        id: 4c4d91c5-5f8b-4d5f-b28f-3fac49c835ad
        input: |
          { "input": "json input" }
        input_path:
      -
        arn: arn:aws:lambda:ap-northeast-1:00000000000:function:another-lambda-function
        id: 3bde837b-3093-487c-9328-a677e05ae75d
        input: |
          { "input": "some input" }
        input_path:
```

### 3. maekawaを実行する
```shell
$ maekawa --region=ap-northeast-1 --apply --dry-run # 変更部分の確認
$ maekawa --region=ap-northeast-1 --apply # config.ymlの内容を適用
```

## オプション
### --region
AWS regionを指定します。

### --apply
configに記述された内容をAWS CloudWatch Eventsに適用します。

### --dry-run
適用をせずに、何が変更されるかを確認します。

### --file, -f
configファイルのpathを指定します。 (デフォルトは `config.yml`です)

### --help, -h
オプションの一覧を表示します。
