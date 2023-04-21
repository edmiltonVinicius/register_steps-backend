echo "GO GO GO, the tests will run..."
echo "================================================================"
echo ""

go test -coverprofile=./tmp/coverage.out  ./...
go tool cover -html=coverage.out -o ./tmp/coverage.html

echo ""
echo "================================================================"
echo "Finished, the coverage file has moved to the temp directory."

