![Go CI](https://github.com/gritsulyak/awesomeProject/actions/workflows/ci.yml/badge.svg)

make

docker compose up


ab -n 300 -c 10 http://127.0.0.1:10080/api/v1/satellite/moon

k6 run scripts/k6/script.js

$ echo "GET http://127.0.0.1:10080/api/v1/satellite/moon" | vegeta attack -rate=10000 -duration=5s | vegeta report 
Requests      [total, rate, throughput]         50000, 10000.39, 9999.95
Duration      [total, attack, wait]             5s, 5s, 218.436µs
Latencies     [min, mean, 50, 90, 95, 99, max]  156.628µs, 1.289ms, 493.524µs, 646.497µs, 1.041ms, 32.883ms, 97.569ms
Bytes In      [total, mean]                     800000, 16.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:50000  
Error Set: 

- current state