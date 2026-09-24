.PHONY: run-load-test

K6_SCENARIO = tests/load/$(S).test.js
K6_INSTALLATION_GUIDE_PAGE=https://grafana.com/docs/learning-paths/run-first-k6-test/install-k6/

run-load-test:
	@test -n "$(S)" || { echo "Usage: make run-load-test S=<scenario> (e.g. S=simple-index)" >&2; exit 1; }
	@test -f "$(K6_SCENARIO)" || { echo "Scenario not found: $(K6_SCENARIO)" >&2; exit 1; }
	@command -v k6 >/dev/null 2>&1 || { echo "k6 is required. Install it: $(K6_INSTALLATION_GUIDE_PAGE)" >&2; exit 1; }
	k6 run "$(K6_SCENARIO)"
