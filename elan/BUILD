go_binary(
    name = "elan",
    srcs = ["main.go"],
    deps = [
        "//elan/rpc",
        "//grpcutil",
        "//third_party/go:cli-init",
        "//third_party/go:http-admin",
        "//third_party/go/gocloud",
    ],
)

sh_cmd(
    name = "run_local",
    cmd = "mkdir -p plz-out/elan && exec $(out_location :elan) --host 127.0.0.1 --port 7777 -s file://\\\\$PWD/plz-out/elan --log_file plz-out/log/elan.log --admin_disabled --file_cache_path plz-out/elan_cache --file_cache_max_size 100M --token_file grpcutil/token.txt",
    srcs = [":elan"],
)
