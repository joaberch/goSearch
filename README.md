# goSearch  
[![Go Report Card](https://goreportcard.com/badge/github.com/joaberch/goSearch)](https://goreportcard.com/report/github.com/joaberch/goSearch)  
A lightweight CLI search engine built in Go. Search words recursively across files from a specified folder using an in-memory or an XML inverted index.

---


## Features
- Quick search for words in files and folders.  
- Work in-memory with optional XML persistance.  
- Registering custom indexes to search through huge amounts of data in a short time.  
- Recursive parsing of the file tree with filter.  
- Simple CLI without direct dependencies.

---


## How to use
### Windows
1. **Download the release from *[Github Releases](https://github.com/joaberch/goSearch/releases)***
2. **Extract the archive (ZIP file)**
3. **Run setup-windows.bat as an administrator**

It will :
- move the executable to a utils folder in the user directory
- Add the path to the utils folder to the system PATH so gosearch can be run from everywhere

> If you can't run the batch file, you can manually :
> - Move the binary to a specific folder
> - Add that folder to the system PATH

### Linux
1. **Download the release from *[Github Releases](https://github.com/joaberch/goSearch/releases)***
2. Extract the archive (TAR)
3. Run the setup script : `bash setup-linux.sh`

It will :
- move the executable to a utils folder in the user directory `~/utils`
- Add the utils folder to the system path using .bashrc

> If you can't run the shell file, you can manually :
> - Move the binary to ``~\utils``
> - Add ``export PATH="$PATH:$HOME/utils"`` to .bashrc or .zshrc

---

## Usage

### Help & Version

```
gosearch -h  # Show help
gosearch -v  # Show version
```

### Search in current directory (and subfolders)

```
gosearch hello
```
### Save index to XML, in Desktop/utils/index/name.yml
```
gosearch -s            # Save index of current directory
gosearch -s path/to    # Save index of specified path
```
### Use saved index
```
gosearch -u hello temp  # Search for 'hello' in the index named 'temp.xml'
```
### Output Example
```
Found 2 file(s) for "hello":
  C:\Users\you\Documents\project\file1.txt
  C:\Users\you\Documents\project\notes.txt
```

---


## Requirements
- Go 1.25 or higher<  
- Windows (tested on amd64)  
- Admin rights (optional for setup script)

---


## Roadmap
- [x] In-memory indexing  
- [x] XML index saving/loading
- [x] Compressing index
- [ ] Refactor tags/flags  
- [ ] Stemming and lemmatization in Normalization  
- [ ] Regex search  
- [x] Cross-platform support (for Linux and macOS)

---


## Contributing
Pull requests are welcome!

---


## License
This project is licensed under the MIT License.  
See LICENSE for more information.
