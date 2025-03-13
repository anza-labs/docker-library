# anza-labs docker library

[![GitHub License](https://img.shields.io/github/license/anza-labs/docker-library)][license]
[![GitHub issues](https://img.shields.io/github/issues/anza-labs/docker-library)](https://github.com/anza-labs/docker-library/issues)

## Usage

1. Building image:
    ```sh
    make build-zig \
        BUILD_ARGS='VERSION=0.14.0' \
        REPOSITORY="localhost:5005/${USER}" \
        TAG='0.14.0' \
        PLATFORM="$(sed -n 's/^# platforms=\(.*\)$/\1/p' ./library/zig/Dockerfile)"
    ```

2. Pushing image:
    ```sh
    make push-zig \
        REPOSITORY="localhost:5005/${USER}" \
        TAG='0.14.0' \
        PLATFORM="$(sed -n 's/^# platforms=\(.*\)$/\1/p' ./library/zig/Dockerfile)"
    ```

3. Generating multi-arch manifest:
    ```sh
    make manifest-zig \
        REPOSITORY="localhost:5005/${USER}" \
        TAG='0.14.0' \
        PLATFORM="$(sed -n 's/^# platforms=\(.*\)$/\1/p' ./library/zig/Dockerfile)"
    ```

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

## License

This project is licensed under the [Apache-2.0][license].

<!-- Resources -->

[license]: https://github.com/anza-labs/docker-library/blob/main/LICENSE
