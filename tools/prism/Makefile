.PHONY: build

build: build/prism-bake.min.js

# Requires esbuild:
# npm install --global esbuild
build/prism-bake.min.js: prism-bake.js
	mkdir -p build
	esbuild --sourcemap --minify prism-bake.js --outfile=build/prism-bake.min.js
