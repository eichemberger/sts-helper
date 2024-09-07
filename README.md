# sts-helper

`sts-helper` is a command-line tool built in Go that allows users to manage AWS Security Token Service (STS) operations, such as assuming roles and getting caller identity. This tool is built using the Cobra library.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/sts-cli.git
   cd sts-helper
   ```

2. **Build the executable:**

   Make sure you have Go installed on your system. Then, run:

   ```bash
   go build -o sts
   ```

3. **Run the executable:**

   You can run the executable using:

   ```bash
   ./sts
   ```

## Usage

`sts-cli` provides two main commands: `assume` and `whoami`.

### Commands

#### 1. `assume`

Assumes an AWS role and prints the credentials of the assumed role. By default, the credentials are copied to the clipboard.

```bash
sts assume --role <role-arn> --session-name <session-name> --duration <duration> --copy=<true|false>
```

##### Flags:

- `--role` (required): The ARN of the role to assume.
- `--session-name`: The name for the assumed role session. Defaults to `"assumed-role"`.
- `--duration`: Duration for which the credentials should be valid, in seconds. Defaults to `3600` seconds (1 hour).
- `--copy`: Copy the credentials to the clipboard. Defaults to `true`. Set to `false` to disable copying.

##### Example:

```bash
sts assume --role arn:aws:iam::123456789012:role/example-role --session-name my-session --duration 1800 --copy=false
```

#### 2. `whoami`

Gets the caller identity of the current user.

```bash
sts whoami
```

### Examples

- **Getting the caller identity:**

  ```bash
  sts whoami
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to fork this project, create a feature branch, and submit a pull request! We welcome contributions of all kinds.

### Explanation:

- **Installation**: Steps to build and run the application.
- **Usage**: Detailed explanation of each command and its flags.
- **Examples**: Shows example usage to help users understand how to use the commands.
- **License and Contributing**: Standard sections for open-source projects.

Feel free to customize it further based on your project details!

This readme was written by an AI. 