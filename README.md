<div align="center">
  <h1>MultCheck</h1>
  <br/>

  <p><i>MultCheck is a open-source, and easy-to-use malware AV test, created by <a href="https://infosec.exchange/@Pengrey">@Pengrey</a>.</i></p>
  <br />
  
</div>

>:warning: **This project is still in development, and is not ready for use.**

MultCheck is a malware-analysis tool that identify the exact position and details of malicious content in binary files using external Anti-Virus scanners.

Takes a binary as input, splits it until it pinpoints that exact bytes that the target engine will flag on and prints them to the screen. This can be helpful in confirming AV scan results and furthering investigation when analyzing malware.
<div id="original"></div>

>  idea from [DefenderCheck](https://github.com/matterpreter/DefenderCheck), [ThreatCheck](https://github.com/rasta-mouse/ThreatCheck) and [ThreatCheck](https://github.com/PACHAKUTlQ/ThreatCheck)

## Installation
-  Run `go build` under the root directory of the project.

-  Or directly run the compiled binaries in [Releases](https://github.com/MultSec/MultCheck/releases).

```bash
$ go build -o multcheck main.go
# Build for Windows
$ GOOS=windows go build -o multcheck.exe main.go
```

## Usage
MultCheck accepts a payload to scan with Windows Defender as default:
`./multcheck <target_file>`

Different built-in scanners can be also be used:
`./multcheck -scanner winDef <target_file>`

For custom scanners a configuration file should be provided:
`./multcheck -scanner C:\path\to\config\file.toml <target_file>`

An example of a configuration file is the following:
```toml
[scanner]
name: "AV name"
cmd: "Command (with full PATH) for scanning the target file. Use {{file}} as the file name to be scanned."
out: "A string present in positive detection but not in negative"
```

## Demo

```powershell
PS > ./multcheck -scanner winDef D:/fakepath/mimikatz.exe
Result: Detected (Static)

0004fb99  31 1c  31 20  31 24  31 28  31 2c  31 30  31 34  31 38  1·1 1$1(1,101418
0004fba9  31 3c  31 40  31 44  31 48  31 4c  31 50  31 54  31 58  1<1@1D1H1L1P1T1X
0004fbb9  31 5c  31 60  31 64  31 68  31 6c  31 70  31 74  31 78  1\1`1d1h1l1p1t1x
0004fbc9  31 7c  31 80  31 88  31 8c  31 90  31 94  31 98  31 9c  1|1·1·1·1·1·1·1·
0004fbd9  31 00  00 00  00 00  00 00  00 00  00 00  00 00  00 00  1···············
0004fbe9  00 00  00 00  00 00  00 00  00 00  00 00  00 00  00 00  ················
0004fbf9  00 00  00 00  00 00  00 52  61 72  21 1a  07 01  00 69  ·······Rar!····i
0004fc09  39 b1  8a 0c  01 05  08 00  07 01  01 b1  9d a1  80 00  9···············
0004fc19  c7 01  ba 7b  2e 02  03 0b  ed 9c  a1 80  00 04  80 dc  ···{.···········
0004fc29  d2 80  00 20  38 c7  f0 6f  80 23  00 0c  6d 69  6d 69  ··· 8··o·#··mimi
0004fc39  6b 61  74 7a  2e 65  78 65  0a 03  02 09  61 a3  ba 3e  katz.exe····a··>
0004fc49  cc d8  01 8d  61 ec  5a 70  86 54  54 32  24 50  60 58  ····a·Zp·TT2$P`X
0004fc59  77 bb  97 60  58 17  25 d8  10 58  75 d9  8c 71  8a 0b  w··`X·%··Xu··q··
0004fc69  91 60  bc 91  82 b7  64 b8  17 09  70 22  c1 08  a0 a4  ·`····d···p"····
0004fc79  50 46  3d 38  c6 17  02 08  2a 0d  82 96  ed bc  5c c4  PF=8····*·····\·
0004fc89  ce 61  9c e6  66 71  c7 3a  5e 67  1c c5  73 1e  80 17  ·a··fq·:^g··s···

PS >
```