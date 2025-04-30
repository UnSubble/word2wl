# Word2WL - Wordlist Generator

Word2WL is a powerful wordlist generator for penetration testers and security professionals. It creates wordlist variations based on a provided dataset and a keyword, allowing deep and recursive mutation of words for use in password attacks and enumeration.

At its core, Word2WL supports mutation levels, recursive transformation, and special character injection, offering a high degree of customization. It’s multithreaded and supports gzip compression for efficient handling of large outputs.

## Installation

  To build from source:

```bash
git clone https://github.com/unsubble/word2wl

cd word2wl

go build -o word2wl .
```
## Usage
``` bash
./word2wl -h
```

Usage:

```bash
word2wl [flags]
```
  

#### Examples:
Generate a wordlist from a dataset and a keyword:

```bash
./word2wl -d "/home/user/rockyou.txt" -k "example" -o output.txt
```

  
Use recursive mutation with level 2 and verbose output:

```bash
./word2wl -d "dataset.txt" -k "admin" -l 2 -r -v -o result.txt
```

  

## Flags:

Flag Description

__`-d, --dataset (Required)`__ Path to the dataset file (e.g., rockyou.txt)

__`-k, --keyword (Required)`__ Keyword to inject into the dataset

__`-l, --level`__ Mutation power level (default: 1, max: 5)

__`-r, --recursive`__ Enable recursive mutation of generated words

__`-R, --recursive-level`__ Number of recursive cycles (default: same as --level)

__`-t, --threads`__ Number of concurrent threads (default: 4)

__`-b, --batch-size`__ Number of words per batch (default: 100)

__`-o, --output-file`__ Output file path (default: stdout)

__`-g, --generator`__ Select generator type (word or path, default: word)

__`-s, --special`__ Special characters to preserve/tokenize separately

__`-v, --verbose`__ Enable verbose output

__`-V, --version`__ Show version

__`-h, --help`__ Help message for Word2WL

## License
Word2WL is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.