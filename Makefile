lint:
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run --config .golangci.yml ./...
	@echo "✅ Success!! No lint issues found!"
fmt:
	@echo "🧹 Formatting Go files..."
	@goimports -w .
	@echo "✅ Code formatted with goimports!"
test:
	@echo "🧪 Running go tests..."
	go test ./...
	@echo "✅ Tests passed!"