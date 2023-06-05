.PHONY: secrets
secrets: secrets/postgres_password.txt secrets/localhost.pem

secrets/postgres_password.txt:
	@mkdir -p secrets
	@./bin/rand > secrets/postgres_password.txt

secrets/localhost.pem:
	@mkdir -p secrets
	@cd secrets && ../bin/gencert localhost
