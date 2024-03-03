<div align="center">
  <h1>MultCheck</h1>
  <br/>

  <p><i>MultCheck is a open-source, and easy-to-use malware AV test, created by <a href="https://infosec.exchange/@Pengrey">@Pengrey</a>.</i></p>
  <br />
  
</div>

MultCheck is a malware-analysis tool that can be used to test the detection of a file by multiple AV engines.

It is designed to be easy to use, and to be able to test multiple AV engines. It is also designed to be easy to extend, and to be able to add custom AV engines.

## Installation
-  Run `go build` under the root directory of the project.

-  Or directly run the compiled binaries in [Releases](https://github.com/MultSec/MultCheck/releases).

```bash
$ cd src
# Build for Windows
## 64-bit
$ GOOS=windows GOARCH=amd64 go build -o ../bin/multcheck_x64.exe main.go

## 32-bit
$ GOOS=windows GOARCH=386 go build -o ../bin/multcheck_x32.exe main.go
```

## Demo

https://github.com/MultSec/MultCheck/assets/55480558/b3105f45-a2d2-42e2-b938-5388f0dc000a

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
The configuration file for custom scanners is a JSON file with the following structure:

```json
{
  "name": "AV name",
  "cmd": "Scan Program (with full PATH) for scanning the target file.",
  "args": "Scan arguments, use {{file}} as the file name to be scanned.",
  "out": "A string present in positive detection but not in negative"
}
```

## Example
```powershell
PS C:\Users\pengrey\Downloads> .\multcheck.exe -scanner .\windef.json C:\Users\pengrey\Downloads\mimikatz.exe
[>] Result: Malicious content found at offset: 00000121
00000000  d1 27 71 71 a9 b6 71 52  69 63 68 70 a9 b6 71 00  |.'qq..qRichp..q.|
00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 50  |...............P|
00000020  45 00 00 64 86 06 00 63  39 5a 5e 00 00 00 00 00  |E..d...c9Z^.....|
00000030  00 00 00 f0 00 22 00 0b  02 09 00 00 2c 0c 00 00  |....."......,...|


PS C:\Users\pengrey\Downloads> .\multcheck.exe -scanner .\windef.json C:\Users\pengrey\Downloads\Rubeus.exe
[>] Result: Malicious content found at offset: 00048e3d
00000000  65 74 5f 61 64 64 69 74  69 6f 6e 61 6c 5f 74 69  |et_additional_ti|
00000010  63 6b 65 74 73 00 67 65  74 5f 74 69 63 6b 65 74  |ckets.get_ticket|
00000020  73 00 73 65 74 5f 74 69  63 6b 65 74 73 00 53 79  |s.set_tickets.Sy|
00000030  73 74 65 6d 2e 4e 65 74  2e 53 6f 63 6b 65 74 73  |stem.Net.Sockets|


PS C:\Users\pengrey\Downloads> .\multcheck.exe -scanner .\windef.json C:\Users\pengrey\Downloads\multcheck.exe
[>] Result: Payload not detected.
PS C:\Users\pengrey\Downloads>
```

## License
This project is licensed under the GNU General Public License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
This project is inspired by the following projects:
- [DefenderCheck](https://github.com/matterpreter/DefenderCheck)
- [ThreatCheck](https://github.com/rasta-mouse/ThreatCheck)
- [ThreatCheck](https://github.com/PACHAKUTlQ/ThreatCheck)
- [GoCheck](https://github.com/gatariee/gocheck)
