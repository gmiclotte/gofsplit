executable=gofsplit
touch $executable
rm $executable

# format go files
goimports -w -l -local "github.com/gmiclotte/gofsplit" .
gofmt -s -w -l .

# tidy the go modules
go mod tidy

# report possible issues
go vet ./...

# try building with race detection enabled
go build -race

# format config.toml, to get toml-sort: pip install toml-sort
# toml-sort --in-place config.toml
# toml-sort --in-place local_config.toml

# check which files were changed
git status

# list exported API
api="api-doc"
rm -f $api || true
for dir in */; do
  (
    cd "${dir}" || exit
    count=`ls -1 *.go 2>/dev/null | wc -l`
    if [ $count != 0 ]
    then
    go doc >> ../$api
    fi
  )
done
echo

# list todos
grep --color=auto TODO ./*.go ./*/*.go
echo

# create dependency graph
"$HOME/go/bin/goda" graph -cluster -short . | dot -Tsvg -o graph.svg

# run tests
# go test -covermode=count -coverprofile=count.out ./...
# go tool cover -html=count.out
# go test -benchtime=1x -cpuprofile cpu.prof -memprofile mem.prof -bench . ./service
# go test -race -v ./service
# for t in "cpu" "mem"; do
  # touch ${t}.svg
  # mv ${t}.svg ${t}.previous.svg
  # go tool pprof -svg -output ${t}.svg ${t}.prof
# done

# build without race detection
go build

# run the binary
# ./$executable local_config.toml
