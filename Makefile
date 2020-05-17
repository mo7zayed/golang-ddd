GCC=go

serve:
	$(GCC) run main.go

dev:
	fresh
migrate:
	$(GCC) run infrastructure/migrations/main.go
seed:
	$(GCC) run infrastructure/seeds/**/main.go