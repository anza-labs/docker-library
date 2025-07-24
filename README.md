# anza-labs docker library

[![GitHub License](https://img.shields.io/github/license/anza-labs/docker-library)][license]
[![GitHub issues](https://img.shields.io/github/issues/anza-labs/docker-library)](https://github.com/anza-labs/docker-library/issues)

## Usage

1. Building image:
    ```sh
    mage build library/zig 0.14.1 arm64 image
    ```

2. Pushing image:
    ```sh
    mage push library/zig 0.14.1 arm64 image
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
