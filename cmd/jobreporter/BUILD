load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/aristorinjuang/microservices-go
gazelle(name = "gazelle")

go_library(
    name = "microservices-go_lib",
    srcs = ["main.go"],
    importpath = "github.com/aristorinjuang/microservices-go",
    visibility = ["//visibility:private"],
    deps = [
        "@org_uber_go_automaxprocs//:go_default_library",
        "@org_uber_go_zap//:go_default_library",
    ],
)

go_binary(
    name = "jobreporter",
    embed = [":microservices-go_lib"],
    visibility = ["//visibility:public"],
)
