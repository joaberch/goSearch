# goSearch  
[![Go Report Card](https://goreportcard.com/badge/github.com/joaberch/goSearch)](https://goreportcard.com/report/github.com/joaberch/goSearch)  
A lightweight CLI search engine in Go that recursively searches words across files from a specified folder using an in-memory or a compressed XML inverted index.

---


## Features
- Fast recursive search for words across files and folders.  
- Work in-memory with optional XML persistence for even faster search.  
- Registering custom indexes to search through huge amounts of data in a short time.  
- Recursive parsing of the file tree with filter.  
- Pure Go CLI with **0** external dependencies.

---


## How to use
### Windows
1. **Download the latest release from *[Github Releases](https://github.com/joaberch/goSearch/releases)***.
2. **Extract the ZIP archive**.
3. **Run ``setup-windows.bat`` as an administrator**.

This will :
- Move the binary to a ``utils`` folder in the user directory.
- Add the ``utils`` folder to the system ``PATH`` variable so ``gosearch`` can be run everywhere.

> Manual setup :
> - Move the binary to a specific folder
> - Add that folder to the system ``PATH``

### Linux / MacOS
1. **Download the latest release from *[Github Releases](https://github.com/joaberch/goSearch/releases)***.
2. Extract the TAR archive.
3. Run the setup script : `bash setup-linux.sh`.

This will :
- Move the binary to ``~/utils``.
- Add the utils folder to the system ``PATH`` using ``.bashrc``.

> Manual setup :
> - Move the binary to ``~/utils``
> - Add ``export PATH="$PATH:$HOME/utils"`` to .bashrc or .zshrc

---

## Usage

### Help & Version

```bash
gosearch -h                         # Show help
gosearch -v                         # Show version
```

### Search from current directory

```bash
gosearch hello                      # Search for 'hello' recursively
```
### Save index to XML
```bash
gosearch -s                         # Save index of current directory
gosearch -s path/to                 # Save index of specified path
```
Saved to : ``~/Desktop/utils/index/<folder>.xml.gz``
### Use saved index
```bash
gosearch -u temp hello              # Search 'hello' in the index file 'temp.xml'
gosearch hello -m exact -u temp     # Same with exact match mode
```
### Output Example
```bash
Found 2 file(s) for "hello":
  C:\Users\you\Documents\project\file1.txt
  C:\Users\you\Documents\project\notes.txt
```

### Match modes

| **Mode**   | **Description**               | **Example**             | **Output**          |
|------------|-------------------------------|-------------------------|---------------------|
| *contains* | Default. Matches partial words | gosearch banan          | 'banane' & 'banana' |
| *exact*    | Matches whole words only      | gosearch -m exact banan | 'banan' only          |


---


## Requirements
- Go 1.25 or higher
- Windows, Linux or macOS
- Admin rights (optional for setup scripts)

---


## Roadmap
- [x] In-memory indexing  
- [x] XML index saving/loading
- [x] Compressing index 
- [ ] Stemming and lemmatization in Normalization  
- [x] Regex search  
- [x] Line numbers
- [x] Cross-platform support (for Linux and macOS)

---

## Contributing
Pull requests are welcome!

---


## License
This project is licensed under the MIT License.  
See [LICENSE](./LICENSE) for more information.
