## ps2elastic index

index prestashop products and categories to elasticsearch.

### Synopsis

index prestashop products and categories to elasticsearch

```
ps2elastic index [flags]
```

### Options

```
      --db-host string           database host (default "127.0.0.1")
      --db-name string           database name (default "prestashop")
      --db-pass string           database password
      --db-port string           datbase port (default "3306")
      --db-table-prefix string   database table prefix (default "ps_")
      --db-user string           database username (default "root")
      --dry-run                  dry run
      --es-host string           elastic host (default "127.0.0.1")
      --es-index string          elastic index name (default "prestashop")
      --es-port string           elastic port (default "9200")
  -h, --help                     help for index
```

### Options inherited from parent commands

```
  -b, --debug           debug mode.
  -g, --generate-docs   generate documentation.
  -v, --verbose         verbose output.
```

### SEE ALSO

* [ps2elastic](ps2elastic.md)	 - ps2elastic is an helper to load data from a prestashop database into elasticsearch.

###### Auto generated by spf13/cobra on 11-Feb-2021