load("@io_bazel_rules_go//go:def.bzl", "go_binary")

load("//:rules.bzl", "pkg_tar", "file_size")


go_binary(
    name = "mytar",
    srcs = ["mytar.go"],
)

pkg_tar(
    name = "archive",
    srcs = [
        "1.txt",
        "2.txt",
    ],
    package_dir = "/my/layout",
)

file_size(
    name = "archive_size",
    src = "archive"
)

# ---

pkg_tar(
    name = "archive_archive",
    srcs = [
        "archive",
    ],
)

pkg_tar(
    name = "empty_archive",
)

file_size(
    name = "mytar_size",
    src = "mytar"
)
