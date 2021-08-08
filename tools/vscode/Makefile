VERSION=$(shell grep package.json -e '"version":' | grep -oe '\d\+\.\d\+\.\d\+')

source=$(shell find . -iname '*.js' -or -iname '*.json' | grep -v node_modules)

.PHONY: package

package: bake-lsp-$(VERSION).vsix

bake-lsp-$(VERSION).vsix: $(source)
	yarn package