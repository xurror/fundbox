module getting-to-go

go 1.20

replace getting-to-go/config => ./config

replace getting-to-go/controller => ./controller

replace getting-to-go/graph => ./graph

replace getting-to-go/graph/generated => ./graph/generated

replace getting-to-go/graph/resolver => ./graph/resolver

replace getting-to-go/model => ./model

replace getting-to-go/server => ./server

replace getting-to-go/server/middleware => ./server/middleware

replace getting-to-go/service => ./service

replace getting-to-go/util => ./util

require (
	github.com/99designs/gqlgen v0.17.32
	github.com/appleboy/gin-jwt/v2 v2.9.1
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-contrib/location v0.0.2
	github.com/gin-contrib/requestid v0.0.6
	github.com/gin-contrib/secure v0.0.1
	github.com/gin-contrib/timeout v0.0.3
	github.com/gin-gonic/gin v1.9.0
	github.com/go-playground/validator/v10 v10.11.2
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.7
	github.com/vektah/gqlparser/v2 v2.5.2-0.20230422221642-25e09f9d292d
	golang.org/x/crypto v0.6.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
