curl http://localhost:8888/

curl -X POST http://localhost:8888/signup \
     -H "Content-Type: application/json" \
     -d '{"userId":"test_user","password":"pass"}'

curl -X POST http://localhost:8888/login \
     -H "Content-Type: application/json" \
     -d '{"userId":"test_user","password":"pass"}'