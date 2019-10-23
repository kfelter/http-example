DIR=results$(date +%s)
mkdir $DIR
(
    cd $DIR
    cat ../targets.txt | vegeta attack -rate=100 -duration=20s | tee results.bin | vegeta report
    vegeta report -type=json results.bin > metrics.json
    cat results.bin | vegeta plot > plot.html
    cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
)
