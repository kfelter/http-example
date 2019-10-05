echo "GET http://localhost:8080/status" | vegeta attack -duration=5s | tee results.bin | vegeta report
vegeta report -type=json results.bin > metrics.json
cat results.bin | vegeta plot > plot.html
cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"