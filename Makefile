COVERDIR=$(CURDIR)/.cover
COVERAGEFILE=$(COVERDIR)/cover.out
COVERAGEREPORT=$(COVERDIR)/report.html

coverage: coverage-ci
	@sed -i -e "s|_$(CURDIR)/|./|g" "${COVERAGEFILE}"

coverage-ci:
	@mkdir -p $(COVERDIR)
	@ginkgo -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find . -type f -name *.coverprofile -exec grep -h -v "^mode:" {} >> "${COVERAGEFILE}" \; -exec rm -f {} \;

coverage-html:
	@go tool cover -html="${COVERAGEFILE}" -o $(COVERAGEREPORT)
	@xdg-open $(COVERAGEREPORT) 2> /dev/null > /dev/null

fmt:
	@go fmt ./...

test:
	@ginkgo --failFast ./...

test-watch:
	@ginkgo watch -cover -r ./...

vet:
	@go vet ./...


.PHONY: coverage coverage-ci coverage-html fmt test test-watch vet
