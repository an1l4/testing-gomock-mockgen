# testing-gomock-mockgen
## steps
- mockgen -destination=mocks/mock_doer.go -package=mocks github.com/an1l4/testing-gomock-mockgen/doer Doer
- go test -v github.com/an1l4/testing-gomock-mockgen/user
