module ds_server

go 1.14

require (
	gitee.com/ha666/golibs v2019.420.1435+incompatible
	gitee.com/ha666/logs v2019.425.1401+incompatible
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aliyun/aliyun-oss-go-sdk v2.1.0+incompatible
	github.com/certifi/gocertifi v0.0.0-20200211180108-c7c1fbc02894 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-gonic/gin v1.6.2
	github.com/go-errors/errors v1.0.1
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-stack/stack v1.8.0
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/websocket v1.4.2
	github.com/holdno/snowFlakeByGo v0.0.0-20180510033652-d23f8a8cadd7
	github.com/iGoogle-ink/gopay v1.5.10
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/labstack/gommon v0.3.0
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/nsqio/go-nsq v1.0.7
	github.com/olivere/elastic v6.2.32+incompatible
	github.com/olivere/elastic/v7 v7.0.17
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.7.0
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron v1.2.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.4.2
	github.com/sohlich/elogrus v2.0.2+incompatible
	github.com/spf13/viper v1.7.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.5.1
	github.com/tebeka/strftime v0.1.4 // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/yangwenmai/ratelimit v0.0.0-20180104140304-44221c2292e1
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.25.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/redis.v5 v5.2.9 // indirect
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gopkg.in/yaml.v2 v2.2.8
)

replace ds_server => ../ds_server
