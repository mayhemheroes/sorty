FROM golang:1.24 as builder

COPY . /sorty
WORKDIR /sorty/mayhem

ENV GOFLAGS=-mod=mod

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzz_sorty.a && \
    clang -fsanitize=fuzzer fuzz_sorty.a -o fuzz_sorty.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /sorty/mayhem/fuzz_sorty.libfuzzer /
