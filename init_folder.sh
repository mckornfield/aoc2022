DAY_NUMBER=$1
if [[ -z "$DAY_NUMBER" ]]; then
    echo a day number is required
    exit 1
fi
DAY="day${DAY_NUMBER}"
mkdir $DAY
pushd $DAY
go mod init github.com/mckornfield/aoc2022/$DAY
go mod edit -replace github.com/mckornfield/aoc2022/shared=../shared
cp ../day_test_template.go ${DAY}_test.go
sed -i '' "s/day/${DAY}/g" ${DAY}_test.go
go mod tidy
touch ${DAY}_input.txt
echo "package ${DAY}" >>${DAY}.go
popd
