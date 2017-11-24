[aws資料](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2017-auto-scaling)   
[aws資料-2](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2016-elastic-load-balancing)   
[システム周り](./sunrise2017_infra.pdf)  
[アドテクサービスプラクティス](https://gist.github.com/katzchang/b97293ede9391822b21c42f7830b543a)  
## Balancing 
### ELB(ALB+CLB) 
AWSクラウド上のロードバランシングサービス
- スケーラブル
- 高い可用性
#### CLB
impやadseverを分けるときはドメインを変える必要がある
#### ALB
ドメインを変える必要がない
### AZ
Availability Zoneの略で同一リージョン内の独立したロケーションを指す

### docker
- 一台のサーバ上でdockerを扱うのは簡単
- 複数のクラスタ上で管理するのは困難
=> ___Amazon EC2 Container Service (ECS)___

### DSP
広告主の効果の最大化

## 今日から使えるアドテク

[アドテクサービスプラクティス](https://gist.github.com/katzchang/b97293ede9391822b21c42f7830b543a)
### 仮実装から始める
- 固定値
- 設定のインポート
- RDBの設定をエクスポート
- 入力フォーム
- 自動的な判断

### 計測する
[LT資料](https://speakerdeck.com/katzchang/gao-surupututodi-reitensinawebsabisuniokerupahuomansudui-ce-falsexian-shi)

### すばやくデプロイする
あらゆる繰り返しは省く

### YAGNI
本当に必要かどうか分からないようなものをあれこれ考えすぎて設計・実装したりするのはやめたほうがいい

