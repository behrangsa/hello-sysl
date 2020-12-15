SYSLGO_SYSL=./specs/hello.sysl
SYSLGO_PACKAGES=hello
SYSLGO_APP.hello = Hello

-include local.mk
include codegen.mk
