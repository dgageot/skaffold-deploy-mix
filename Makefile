.PHONY: test
test:
	skaffold build -q > build_result.json
	skaffold deploy --build-artifacts=build_result.json -p david
	skaffold deploy --tail --build-artifacts=build_result.json -p world

