FROM golang:1.25 as builder

COPY . /sorty
WORKDIR /sorty/mayhem

ENV GOFLAGS=-mod=mod

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN apt update && apt install -y clang

RUN go mod init fuzz_sorty_module && \
    go mod edit -require github.com/jfcg/sorty/v2@v2.0.0 && \
    go mod edit -replace "github.com/jfcg/sorty/v2@v2.0.0=../" && \
    go get github.com/dvyukov/go-fuzz/go-fuzz-dep && \
    go get github.com/AdaLogics/go-fuzz-headers && \
    go mod tidy

RUN go-fuzz-build -libfuzzer -o fuzz_sorty.a . && \
    clang -fsanitize=fuzzer fuzz_sorty.a -o fuzz_sorty.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /sorty/mayhem/fuzz_sorty.libfuzzer /
