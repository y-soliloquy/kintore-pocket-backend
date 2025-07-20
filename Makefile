lint:
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run --config .golangci.yml ./...
	@echo "✅ Success!! No lint issues found!"
fmt:
	@echo "🧹 Formatting Go files..."
	@goimports -w .
	@echo "✅ Code formatted with goimports!"