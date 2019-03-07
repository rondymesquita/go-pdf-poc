cd pdf
# go test -v -bench=. -benchtime=10s -benchmem
# echo ""
# echo "*** 10pages Sequence 1"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=sequence -file=10pages.pdf

# echo ""
# echo "*** 10pages Parallel 1"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=parallel -file=10pages.pdf -maxGoroutines=1

# echo ""
# echo "*** 10pages Parallel 2"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=parallel -file=10pages.pdf -maxGoroutines=2

# echo ""
# echo "*** 10pages Parallel 3"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=parallel -file=10pages.pdf -maxGoroutines=3

# echo ""
# echo "*** 10pages Parallel 4"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=parallel -file=10pages.pdf -maxGoroutines=4

echo ""
echo "*** 10pages Parallel 5"
go test -bench=. -benchtime=10x -count=10 -benchmem -memprofile memprofile.out -cpuprofile profile.out -mode=parallel -file=10pages.pdf -maxGoroutines=5

# echo ""
# echo "*** 10pages Parallel 10"
# go test -bench=. -benchtime=3x -count=3 -benchmem -mode=parallel -file=10pages.pdf -maxGoroutines=10

echo "*** Benchmarks done"
cd ..
