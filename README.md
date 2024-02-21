<div align="center">
  <h1>MultCheck</h1>
  <br/>

  <p><i>MultCheck is a open-source, and easy-to-use malware AV test, created by <a href="https://infosec.exchange/@Pengrey">@Pengrey</a>.</i></p>
  <br />
  
</div>

>:warning: **This project is still in development, and is not ready for use.**

MultCheck is a malware-analysis tool that can be used to test the detection of a file by multiple AV engines.

It is designed to be easy to use, and to be able to test multiple AV engines. It is also designed to be easy to extend, and to be able to add custom AV engines.

## Installation
-  Run `go build` under the root directory of the project.

-  Or directly run the compiled binaries in [Releases](https://github.com/MultSec/MultCheck/releases).

```bash
$ go build -o multcheck main.go
# Build for Windows
$ GOOS=windows go build -o multcheck.exe main.go
```

## Usage
MultCheck accepts a target file as an argument:
`./multcheck <target_file>`

Different built-in scanners can be used by specifying the `-scanner` flag:
`./multcheck -scanner <scanner_name> <target_file>`

Custom scanners can be added by creating a configuration file and providing the path to the file through the `-scanner` flag:
`./multcheck -scanner <path_to_config_file> <target_file>`

## Supported Scanners
- Windows Defender (winDef)

## Configuration
The configuration file for custom scanners is a TOML file with the following structure:

```toml
[scanner]
name: "AV name"
cmd: "Command (with full PATH) for scanning the target file. Use {{file}} as the file name to be scanned."
out: "A string present in positive detection but not in negative"
```

## Example
```bash
$ ./multcheck -scanner winDef test.exe
[>] Result: Payload not detected.

$
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
This project is inspired by the following projects:
- [DefenderCheck](https://github.com/matterpreter/DefenderCheck)
- [ThreatCheck](https://github.com/rasta-mouse/ThreatCheck)
- [ThreatCheck](https://github.com/PACHAKUTlQ/ThreatCheck)
