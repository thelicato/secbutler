<h1 align="center">
	<img src="https://github.com/thelicato/secbutler/blob/main/docs/logo.png?raw=true" width="400">
</h1>

<h4 align="center">Essential utilities for pentester, bug-bounty hunters and security researchers</h4>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-license">License</a> •
</p>

---

`secbutler` is a utility tool made for **pentesters**, **bug-bounty hunters** and **security researchers** that contains all the most used and tedious stuff commonly used while performing cybersecurity activities (like installing sec-related tools, retrieving commands for revshells, serving common payloads, obtaining a working proxy, managing wordlists and so forth).

The goal is to obtain a tool that meets the requirements of the community, therefore suggestions and PRs are very welcome!

## ⚡ Features

- Generate a reverse shell command
- Obtain proxy
- Download & deploy common payloads
- Obtain reverse shell listener command
- Generate bash install script for common tools
- Generate bash download script for Wordlists
- Read common cheatsheets and payloads

## 📚 Usage

```
secbutler -h
```

This will display the help for the tool

```
                   __          __  __
   ________  _____/ /_  __  __/ /_/ /__  _____
  / ___/ _ \/ ___/ __ \/ / / / __/ / _ \/ ___/
 (__  )  __/ /__/ /_/ / /_/ / /_/ /  __/ /
/____/\___/\___/_.___/\__,_/\__/_/\___/_/

v0.2.1 - https://github.com/thelicato/secbutler

Essential utilities for pentester, bug-bounty hunters and security researchers

Usage:
  secbutler [flags]
  secbutler [command]

Available Commands:
  cheatsheet  Read common cheatsheets & payloads
  help        Help about any command
  listener    Obtain the command to start a reverse shell listener
  payloads    Obtain and serve common payloads
  proxy       Obtain a random proxy from FreeProxy
  revshell    Obtain the command for a reverse shell
  tools       Generate a install script for the most common cybersecurity tools
  version     Print the current version
  wordlists   Generate a download script for the most common wordlists

Flags:
  -h, --help   help for secbutler

Use "secbutler [command] --help" for more information about a command.


```

## 🚀 Installation

Run the following command to install the latest version:

```
go install github.com/thelicato/secbutler@latest
```

Or you can simply grab an executable from the [Releases](https://github.com/thelicato/secbutler/releases) page.

## 🪪 License

_secbutler_ is made with 🖤 and released under the [MIT LICENSE](https://github.com/thelicato/secbutler/blob/main/LICENSE).
