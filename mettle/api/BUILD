go_library(
    name = "api",
    srcs = ["api.go"],
    deps = [
        "//grpcutil",
        "//third_party/go/gocloud",
        "//third_party/go:bytestream",
        "//third_party/go:cli-init",
        "//third_party/go:grpc",
        "//third_party/go:logging",
        "//third_party/go:longrunning",
        "//third_party/go:protobuf",
        "//third_party/go:remote-apis",
        "//third_party/go/prometheus",
        "//mettle/common",
    ],
    visibility = ["//mettle/..."],
)

go_test(
    name = "api_test",
    srcs = ["api_test.go"],
    deps = [
        ":api",
        "//grpcutil",
        "//third_party/go:grpc",
        "//third_party/go/gocloud",
        "//third_party/go:testify",
        "//third_party/go:longrunning",
        "//third_party/go:remote-apis",
        "//third_party/go:protobuf",
    ],
)
