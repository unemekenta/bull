# bull

## 設計
### DDD
参考: https://techblog.yahoo.co.jp/entry/2021011230061115/
https://taisablog.com/archives/go-ddd-domain-service

レイヤー名	| 責務 | そこに置かれるクラス例
--- | --- | ---
presentation(interfaces) | 表示における関心事を扱う。ルーティングやHTTPリクエストレスポンスやそのバリデーション、あるいはUIなど。 | Controller、Request(Form)、Response(Resource)など
application(usecase) | ドメインオブジェクトを利用してソフトウェアが行うべき仕事を表現する。 | Scenario(Usecase)、アプリケーションサービスなど
domain | 業務ロジックや業務ルールを表現する。	| ドメインモデル、ドメインサービス、Repositoryインターフェース、ドメインモデルのFactoryクラスなど
infrastructure | データの永続化処理や外部サービスとのやりとり(REST API/MQ)を行う。 | Repository、ORMapper、Adapter、Translatorなど

```
├── application
│   └── usecase
│       └── user_usecase.go
├── domain
│   ├── model
│   │   ├── errors.go
│   │   └── users.go
│   ├── repository
│   │   └── repository.go
│   └── service
│       └── user_service.go
├── go.mod
├── go.sum
├── gorm.db
├── infrastructure
│   ├── repository
│   │   ├── repository.go
│   │   └── user_repository.go
│   └── service
│       └── user_service.go
├── interfaces
│   ├── handler
│   │   └── user_handler.go
│   └── request
│   │   └── user_request.go
│   └── response
│       └── user_response.go
└── main.go
```
