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
1. Clone the repository
```
git clone https://github.com/joaberch/goSearch.git
cd goSearch
```
2.   Start the setup batch with admin right, in case you don't want to or can't you can manually do it:  
2.1  Build the project
```
go build
```
2.2  Add the exe to the system path to use it everywhere

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
- [ ] Refactor tags/flags  
- [ ] Stemming and lemmatization in Normalization  
- [ ] Regex search  
- [ ] Cross-platform support (for Linux and macOS)

---


## Contributing
Pull requests are welcome!

---


## License
This project is licensed under the MIT License.  
See LICENSE for more information.
