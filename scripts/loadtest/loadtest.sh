VEGETA="../../.dev/vegeta"

# run loadtest and download results
$VEGETA attack -duration=2m -rate=100 -format=http -targets="wallets_req.txt" | $VEGETA encode -output=wallets_result.json && \
cat wallets_result.json | $VEGETA report --type=hdrplot -output=report.plt