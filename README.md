# normalize_csv
submission for Truss coding challenge

## About

This project accepts a csv on stdin, attempts to normalize the fields into a predictable format, and returns the results on stdout.
If it detects a malformed string in any field except Notes, it will drop the entire row and continue processing.
 
Assumptions:
- The first line will contain headers
- A malformed header should cancel the entire job
- The csv will contain only the following columns:
    - Timestamp
    - Address
    - ZIP
    - FullName
    - FooDuration
    - BarDuration
    - TotalDuration
    - Notes

## Install 

Please ensure you have go installed on your machine and you are using at least macOS 10.13 before continuing.

```Shell
git clone myrepo
cd myrepo
go build
```

## Use

For output to the command line:
```Shell
./normalize_csv < sample.csv
```

For output to a file:
```Shell
./normalize_csv < sample.csv
./normalize_csv < sample.csv > output.csv
```

If you encounter an issue please let me know at agaesser@gmail.com!
