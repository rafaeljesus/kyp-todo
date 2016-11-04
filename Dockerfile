FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-todo /kyp-todo
ENV KYP_TODO_DB="postgres://postgres:@docker/kyp_todo_dev?sslmode=disable"
ENV KYP_TODO_PORT="3001"
ENV KYP_SECRET_KEY="c91267c27a8599ca0480ea505487d052e3b63a1dd39819db853225a518200399"
ENTRYPOINT ["/kyp-todo"]
