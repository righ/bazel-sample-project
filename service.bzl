load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

def service_image(name, **kwargs):
    pkg_tar(
        name = "service_tar",
        srcs = ["@grpc_health_probe//file"],
        mode = "0o755",
        package_dir = "/bin",
        visibility = ["//visibility:public"],
    )

    container_image(
        name = "service_image",
        base = "@go_image_base//image",
        tars = [":service_tar"],
        visibility = ["//visibility:public"],
    )

    go_image(
        name = "container",
        base = ":service_image",
        embed = ["//" + name + ":go_default_library"],
        goarch = "amd64",
        goos = "linux",
        visibility = ["//visibility:public"],
        **kwargs
    )

    container_push(
        name = "push",
        format = "Docker",
        image = ":service_image",
        registry = "docker.io",
        repository = "righm9/" + name,
        tag = "{BUILD_TIMESTAMP}",
    )
