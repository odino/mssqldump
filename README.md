# mssqldump

Dump a MSSQL query into a TSV / JSON file.

## Usage

Download the [binary](https://github.com/odino/mssqldump/releases), specify a connection string as an environment variable (`MSSQL_CONN`)
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

It also supports output in newline delimited JSON -- BigQuery anyone?

```
./mssqldump -q "SELECT Name, 1 as ID, RAND() as thing from sys.Databases" -f json
{"ID":1,"Name":"master","thing":0.6214426254735921}
{"ID":1,"Name":"tempdb","thing":0.6214426254735921}
{"ID":1,"Name":"model","thing":0.6214426254735921}
{"ID":1,"Name":"msdb","thing":0.6214426254735921}
{"ID":1,"Name":"test","thing":0.6214426254735921}
```


If your query is `v`, `mssqldump` will print out its version information:

```
./mssqldump -qv
1.0.0
```
