load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

http_archive(
    name = "bazel_gazelle",
    sha256 = "b8b6d75de6e4bf7c41b7737b183523085f56283f6db929b86c5e7e1f09cf59c9",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.31.1/bazel-gazelle-v0.31.1.tar.gz"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:repositories.bzl", "go_repositories")

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()

gazelle_dependencies()
