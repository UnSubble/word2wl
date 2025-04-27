# Word2WL - Wordlist Generator

Word2WL is a powerful wordlist generator that creates word variations based on a given dataset and keyword. It allows you to mutate words in various ways, useful for tasks like password cracking, security assessments, and generating wordlists for penetration testing.

## Features

- Generate wordlist variations by mutating words from a given dataset.
- Support for multiple mutation levels (from basic to advanced).
- Ability to use special characters in the mutations.
- Support for recursive mutation to expand wordlists further.
- Concurrent processing with configurable threads for faster generation.
- Output to files, with gzip compression for large outputs.

## Installation

You can clone this repository and build the project locally.

```bash
git clone https://github.com/unsubble/word2wl
cd word2wl
go build -o word2wl .
```

## Usage

### Basic Usage:

To generate a wordlist from a dataset with a given keyword:
```bash
./word2wl -d "/path/to/dataset.txt" -k "yourkeyword" -o output.txt
```
### Command-line Flags:

- `-d, --dataset` (Required): Path to the dataset file (e.g., rockyou.txt).
    
- `-k, --keyword` (Required): Keyword to inject into the dataset.
    
- `-l, --level` (Optional): Mutation power level (1 = basic, 5 = advanced).
    
- `-t, --threads` (Optional): Number of concurrent threads (default: 4).
    
- `-b, --batch-size` (Optional): Number of words per batch (default: 100).
    
- `-o, --output-file` (Optional): Output file path (default: stdout).
    
- `-r, --recursive` (Optional): Recursively mutate generated words.
    
- `-R, --recursive-level` (Optional): Number of recursive mutation cycles (default: same as `--level`).
    
- `-v, --verbose` (Optional): Enable verbose output.
    
- `-s, --special` (Optional): Special characters that should be reserved and tokenized separately.
    

### Example:
```bash
./word2wl -d "/home/user/rockyou.txt" -k "example" -o output.txt -l 2 -t 8 -v
```

This command will generate a wordlist based on the `rockyou.txt` dataset, using the keyword "example" with mutation level 2, and output the results to `output.txt` with 8 threads and verbose output.