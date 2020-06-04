wrk.method = "POST"
wrk.body   = '{"device_id":"a12345","account":"abcde056a","password":"129345","ver":666,"channel":999,"type":"777","pid":"56789"}'
wrk.headers["Content-Type"] = "application/json"

wrk -t2 -c100 -d10s -s /opt/workspace/src/microservice/jzapi/post.lua  --latency  "http://127.0.0.1:9908/user/login_account"


