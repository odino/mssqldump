# mssqldump

Dump a MSSQL query into a TSV file.

## Usage

Download the binary, specify a connection string as an environment variable (`MSSQL_CONN`)
and run it specifying the query you wanna run (`-q`):

```
export MSSQL_CONN="sqlserver://sa:Password123\!@127.0.0.1:1433"

./mssqldump -q "SELECT Name, 1 as ID, RAND() as thing from sys.Databases"
master	1	0.4318099474883688
tempdb	1	0.4318099474883688
model	1	0.4318099474883688
msdb	1	0.4318099474883688
test	1	0.4318099474883688
```

If you want to include column names in the output, use the `-c` option:

```
./mssqldump -q "SELECT Name, 1 as ID, RAND() as thing from sys.Databases" -c
Name	ID	thing
master	1	0.23937146273025442
tempdb	1	0.23937146273025442
model	1	0.23937146273025442
msdb	1	0.23937146273025442
test	1	0.23937146273025442
```

If your query is `v`, `mssqldump` will print out its version information:

```
./mssqldump -qv
1.0.0
```
