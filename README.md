<div align="center">
  <img width="125px" src="assets/MultSCheck.png" />
  <h1>MultCheck</h1>
  <br/>

  <p><i>MultCheck is a open-source, and easy-to-use malware AV test, created by <a href="https://infosec.exchange/@Pengrey">@Pengrey</a>.</i></p>
  <br />
  
</div>

MultCheck is a malware-analysis tool that can be used to test the detection of a file by multiple AV engines.

It is designed to be easy to use, and to be able to test multiple AV engines. It is also designed to be easy to extend, and to be able to add custom AV engines.

## Installation
-  Run `go build` under the root directory of the project.

-  Or directly run the compiled binaries in the bin directory.

```bash
$ cd src
# Build for Windows
## 64-bit
$ GOOS=windows GOARCH=amd64 go build -o ../bin/multcheck_64.exe

## 32-bit
$ GOOS=windows GOARCH=386 go build -o ../bin/multcheck_32.exe
```

## Demo
https://github.com/MultSec/MultCheck/assets/55480558/708c9f74-aef8-4950-bf09-e27987eddc17

# Documentation
For more information on how to use the MultLdr-cli, check the [documentation](https://multsec.github.io/docs/multcheck/)

## License
This project is licensed under the GNU General Public License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
This project is inspired by the following projects:
- [DefenderCheck](https://github.com/matterpreter/DefenderCheck)
- [ThreatCheck](https://github.com/rasta-mouse/ThreatCheck)
- [ThreatCheck](https://github.com/PACHAKUTlQ/ThreatCheck)
- [GoCheck](https://github.com/gatariee/gocheck)
