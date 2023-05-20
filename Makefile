# With this you can run 'make test' to test all of the packages,
# or you can run "'make test PKG='./structs'" to test a specific package
test:
ifdef PKG
	@echo "Running tests in package $(PKG)"
	@ginkgo $(PKG)
else
	@find . -type f -name "*_test.go" -exec dirname {} \; | sort | uniq | xargs -I {} sh -c 'echo "Running tests in {}"; cd {} && ginkgo'
endif
