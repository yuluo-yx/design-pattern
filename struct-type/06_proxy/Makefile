# 加入 go generated 来动态生成代理文件
go-gen:
	cd go && go run main.go static_proxy.go dynamic_proxy_reflect.go user_service.go dynamic_proxy_generator.go generate

# 执行时需要放开 main.go 的注释
go-gen-run:
	cd go && go run main.go static_proxy.go dynamic_proxy_reflect.go user_service.go proxy_generated.go dynamic_proxy_generator.go

go-gen-clear:
	cd go && rm -f proxy_generated.go

go-dev:
	cd go && go run main.go static_proxy.go dynamic_proxy_reflect.go user_service.go dynamic_proxy_generator.go

java-dev:
	cd java && javac -d . *.java && java main.Main

clear-java:
	@cd java && rm -rf main
