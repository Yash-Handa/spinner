# spinner

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c0c8fe5903194ebcbe781a8c3966d249)](https://app.codacy.com/manual/yashhanda7/spinner?utm_source=github.com&utm_medium=referral&utm_content=Yash-Handa/spinner&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/Yash-Handa/spinner)](https://goreportcard.com/report/github.com/Yash-Handa/spinner)
[![codecov](https://codecov.io/gh/Yash-Handa/spinner/branch/master/graph/badge.svg)](https://codecov.io/gh/Yash-Handa/spinner)
![Go](https://github.com/Yash-Handa/spinner/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Yash-Handa/spinner)](https://pkg.go.dev/github.com/Yash-Handa/spinner)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/Yash-Handa/spinner.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/Yash-Handa/spinner/alerts/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/Yash-Handa/spinner.svg)](https://github.com/Yash-Handa/spinner)
[![GitHub release](https://img.shields.io/github/release/Yash-Handa/spinner.svg)](https://GitHub.com/Yash-Handa/spinner)

Platform independent Go module to print spinners on Terminal/ cmd

## Overview / Features

- Have 80+ spinners, Segregated into ASCII(id: 1-999) and Unicode(id >= 1000) spinners

-  Custom (user-defined) spinners

- Support rich color rendering output for the spinners (in both 16bit colors and Hexcodes)

- Colors compatible with Windows system.

- Universal and Stable API method

- Made with concurrency handling in mind

## Documentation

Complete documentation of the API is hosted at [pkg.go.dev](https://pkg.go.dev/github.com/Yash-Handa/spinner)

## Spinners

Below Tables show spinners with their IDs. Use it as reference

### ASCII Spinners (ID 1 to 999)

<table>
  <tr>
    <td><img src="/.github/gifs/1-80.gif" alt="ID: 1 Inter: 80ms"></td>
    <td><img src="/.github/gifs/2-140.gif" alt="ID: 2 Inter: 140ms"></td>
    <td><img src="/.github/gifs/3-100.gif" alt="ID: 3 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1 Inter: 80ms</th>
    <th>ID: 2 Inter: 140ms</th>
    <th>ID: 3 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/4-100.gif" alt="ID: 4 Inter: 100ms"></td>
    <td><img src="/.github/gifs/5-120.gif" alt="ID: 5 Inter: 120ms"></td>
    <td><img src="/.github/gifs/6-120.gif" alt="ID: 6 Inter: 120ms"></td>
  </tr>
  <tr>
    <th>ID: 4 Inter: 100ms</th>
    <th>ID: 5 Inter: 120ms</th>
    <th>ID: 6 Inter: 120ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/7-100.gif" alt="ID: 7 Inter: 100ms"></td>
    <td><img src="/.github/gifs/8-120.gif" alt="ID: 8 Inter: 120ms"></td>
    <td><img src="/.github/gifs/9-100.gif" alt="ID: 9 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 7 Inter: 100ms</th>
    <th>ID: 8 Inter: 120ms</th>
    <th>ID: 9 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/10-80.gif" alt="ID: 10 Inter: 80ms"></td>
    <td><img src="/.github/gifs/11-80.gif" alt="ID: 11 Inter: 80ms"></td>
    <td><img src="/.github/gifs/12-200.gif" alt="ID: 12 Inter: 200ms"></td>
  </tr>
  <tr>
    <th>ID: 10 Inter: 80ms</th>
    <th>ID: 11 Inter: 80ms</th>
    <th>ID: 12 Inter: 200ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/13-120.gif" alt="ID: 13 Inter: 120ms"></td>
    <td><img src="/.github/gifs/14-80.gif" alt="ID: 14 Inter: 80ms"></td>
    <td><img src="/.github/gifs/15-120.gif" alt="ID: 15 Inter: 120ms"></td>
  </tr>
  <tr>
    <th>ID: 13 Inter: 120ms</th>
    <th>ID: 14 Inter: 80ms</th>
    <th>ID: 15 Inter: 120ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/16-100.gif" alt="ID: 16 Inter: 100ms"></td>
    <td><img src="/.github/gifs/17-80.gif" alt="ID: 17 Inter: 80ms"></td>
    <td><img src="/.github/gifs/18-180.gif" alt="ID: 18 Inter: 180ms"></td>
  </tr>
  <tr>
    <th>ID: 16 Inter: 100ms</th>
    <th>ID: 17 Inter: 80ms</th>
    <th>ID: 18 Inter: 180ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/19-150.gif" alt="ID: 19 Inter: 150ms"></td>
    <td><img src="/.github/gifs/20-120.gif" alt="ID: 20 Inter: 120ms"></td>
    <td><img src="/.github/gifs/21-100.gif" alt="ID: 21 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 19 Inter: 150ms</th>
    <th>ID: 20 Inter: 120ms</th>
    <th>ID: 21 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/22-100.gif" alt="ID: 22 Inter: 100ms"></td>
    <td><img src="/.github/gifs/23-100.gif" alt="ID: 23 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 22 Inter: 100ms</th>
    <th>ID: 23 Inter: 100ms</th>
  </tr>
</table>

### Unicode Spinners (ID >= 1000)

<table>
  <tr>
    <td><img src="/.github/gifs/1000-80.gif" alt="ID: 1000 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1001-80.gif" alt="ID: 1001 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1002-80.gif" alt="ID: 1002 Inter: 80ms"></td>
  </tr>
  <tr>
    <th>ID: 1000 Inter: 80ms</th>
    <th>ID: 1001 Inter: 80ms</th>
    <th>ID: 1002 Inter: 80ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1003-80.gif" alt="ID: 1003 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1004-100.gif" alt="ID: 1004 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1005-100.gif" alt="ID: 1005 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1003 Inter: 80ms</th>
    <th>ID: 1004 Inter: 100ms</th>
    <th>ID: 1005 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1006-100.gif" alt="ID: 1006 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1007-80.gif" alt="ID: 1007 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1008-100.gif" alt="ID: 1008 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1006 Inter: 100ms</th>
    <th>ID: 1007 Inter: 80ms</th>
    <th>ID: 1008 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1009-80.gif" alt="ID: 1009 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1010-140.gif" alt="ID: 1010 Inter: 140ms"></td>
    <td><img src="/.github/gifs/1011-120.gif" alt="ID: 1011 Inter: 120ms"></td>
  </tr>
  <tr>
    <th>ID: 1009 Inter: 80ms</th>
    <th>ID: 1010 Inter: 140ms</th>
    <th>ID: 1011 Inter: 120ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1012-100.gif" alt="ID: 1012 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1013-200.gif" alt="ID: 1013 Inter: 200ms"></td>
    <td><img src="/.github/gifs/1014-120.gif" alt="ID: 1014 Inter: 120ms"></td>
  </tr>
  <tr>
    <th>ID: 1012 Inter: 100ms</th>
    <th>ID: 1013 Inter: 200ms</th>
    <th>ID: 1014 Inter: 120ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1015-100.gif" alt="ID: 1015 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1016-180.gif" alt="ID: 1016 Inter: 180ms"></td>
    <td><img src="/.github/gifs/1017-120.gif" alt="ID: 1017 Inter: 120ms"></td>
  </tr>
  <tr>
    <th>ID: 1015 Inter: 100ms</th>
    <th>ID: 1016 Inter: 180ms</th>
    <th>ID: 1017 Inter: 120ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1018-120.gif" alt="ID: 1018 Inter: 120ms"></td>
    <td><img src="/.github/gifs/1019-100.gif" alt="ID: 1019 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1020-80.gif" alt="ID: 1020 Inter: 80ms"></td>
  </tr>
  <tr>
    <th>ID: 1018 Inter: 120ms</th>
    <th>ID: 1019 Inter: 100ms</th>
    <th>ID: 1020 Inter: 80ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1021-80.gif" alt="ID: 1021 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1022-80.gif" alt="ID: 1022 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1023-80.gif" alt="ID: 1023 Inter: 80ms"></td>
  </tr>
  <tr>
    <th>ID: 1021 Inter: 80ms</th>
    <th>ID: 1022 Inter: 80ms</th>
    <th>ID: 1023 Inter: 80ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1024-80.gif" alt="ID: 1024 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1025-100.gif" alt="ID: 1025 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1026-80.gif" alt="ID: 1026 Inter: 80ms"></td>
  </tr>
  <tr>
    <th>ID: 1024 Inter: 80ms</th>
    <th>ID: 1025 Inter: 100ms</th>
    <th>ID: 1026 Inter: 80ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1027-80.gif" alt="ID: 1027 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1028-80.gif" alt="ID: 1028 Inter: 80ms"></td>
    <td><img src="/.github/gifs/1029-80.gif" alt="ID: 1029 Inter: 80ms"></td>
  </tr>
  <tr>
    <th>ID: 1027 Inter: 80ms</th>
    <th>ID: 1028 Inter: 80ms</th>
    <th>ID: 1029 Inter: 80ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1030-100.gif" alt="ID: 1030 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1031-200.gif" alt="ID: 1031 Inter: 200ms"></td>
    <td><img src="/.github/gifs/1032-100.gif" alt="ID: 1032 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1030 Inter: 100ms</th>
    <th>ID: 1031 Inter: 200ms</th>
    <th>ID: 1032 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1033-100.gif" alt="ID: 1033 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1034-180.gif" alt="ID: 1034 Inter: 180ms"></td>
    <td><img src="/.github/gifs/1035-150.gif" alt="ID: 1035 Inter: 150ms"></td>
  </tr>
  <tr>
    <th>ID: 1033 Inter: 100ms</th>
    <th>ID: 1034 Inter: 180ms</th>
    <th>ID: 1035 Inter: 150ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1036-120.gif" alt="ID: 1036 Inter: 120ms"></td>
    <td><img src="/.github/gifs/1037-120.gif" alt="ID: 1037 Inter: 120ms"></td>
    <td><img src="/.github/gifs/1038-180.gif" alt="ID: 1038 Inter: 180ms"></td>
  </tr>
  <tr>
    <th>ID: 1036 Inter: 120ms</th>
    <th>ID: 1037 Inter: 120ms</th>
    <th>ID: 1038 Inter: 180ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1039-180.gif" alt="ID: 1039 Inter: 180ms"></td>
    <td><img src="/.github/gifs/1040-180.gif" alt="ID: 1040 Inter: 180ms"></td>
    <td><img src="/.github/gifs/1041-180.gif" alt="ID: 1041 Inter: 180ms"></td>
  </tr>
  <tr>
    <th>ID: 1039 Inter: 180ms</th>
    <th>ID: 1040 Inter: 180ms</th>
    <th>ID: 1041 Inter: 180ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1042-150.gif" alt="ID: 1042 Inter: 150ms"></td>
    <td><img src="/.github/gifs/1043-150.gif" alt="ID: 1043 Inter: 150ms"></td>
    <td><img src="/.github/gifs/1044-150.gif" alt="ID: 1044 Inter: 150ms"></td>
  </tr>
  <tr>
    <th>ID: 1042 Inter: 150ms</th>
    <th>ID: 1043 Inter: 150ms</th>
    <th>ID: 1044 Inter: 150ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1045-150.gif" alt="ID: 1045 Inter: 150ms"></td>
    <td><img src="/.github/gifs/1046-140.gif" alt="ID: 1046 Inter: 140ms"></td>
    <td><img src="/.github/gifs/1047-150.gif" alt="ID: 1047 Inter: 150ms"></td>
  </tr>
  <tr>
    <th>ID: 1045 Inter: 150ms</th>
    <th>ID: 1046 Inter: 140ms</th>
    <th>ID: 1047 Inter: 150ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1048-150.gif" alt="ID: 1048 Inter: 150ms"></td>
    <td><img src="/.github/gifs/1049-120.gif" alt="ID: 1049 Inter: 120ms"></td>
    <td><img src="/.github/gifs/1050-150.gif" alt="ID: 1050 Inter: 150ms"></td>
  </tr>
  <tr>
    <th>ID: 1048 Inter: 150ms</th>
    <th>ID: 1049 Inter: 120ms</th>
    <th>ID: 1050 Inter: 150ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1051-120.gif" alt="ID: 1051 Inter: 120ms"></td>
    <td><img src="/.github/gifs/1052-150.gif" alt="ID: 1052 Inter: 150ms"></td>
    <td><img src="/.github/gifs/1053-100.gif" alt="ID: 1053 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1051 Inter: 120ms</th>
    <th>ID: 1052 Inter: 150ms</th>
    <th>ID: 1053 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1054-100.gif" alt="ID: 1054 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1055-100.gif" alt="ID: 1055 Inter: 100ms"></td>
    <td><img src="/.github/gifs/1056-100.gif" alt="ID: 1056 Inter: 100ms"></td>
  </tr>
  <tr>
    <th>ID: 1054 Inter: 100ms</th>
    <th>ID: 1055 Inter: 100ms</th>
    <th>ID: 1056 Inter: 100ms</th>
  </tr>
  <tr>
    <td><img src="/.github/gifs/1057-200.gif" alt="ID: 1057 Inter: 200ms"></td>
    <td><img src="/.github/gifs/1058-180.gif" alt="ID: 1058 Inter: 180ms"></td>
  </tr>
  <tr>
    <th>ID: 1057 Inter: 200ms</th>
    <th>ID: 1058 Inter: 180ms</th>
  </tr>
</table>


## Contributions

The Project is Open Sourced under [MIT]("/LICENSE") License and will always welcomes Pull Request. Please read Contribtion.md.