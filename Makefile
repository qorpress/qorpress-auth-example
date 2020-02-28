#---* Makefile *---#

# To do
# git rev
build:
	@docker build -t qorpress/qorpress-auth-example --no-cache .

run:
	@docker run -ti -p 9000:9000 -v $(PWD)/.config/gopress.yml:/opt/qor/.config/gopress.yml qorpress/qorpress-auth-example
