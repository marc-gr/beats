# elastic_amcache_application_file

Windows Amcache application file inventory from the InventoryApplicationFile registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_application_file table provides detailed information about
application files tracked in the Windows Amcache InventoryApplicationFile hive.
This includes executable files, DLLs, and other binaries associated with installed
applications.

This data is valuable for:
- Detailed file-level software inventory
- Tracking file versions and updates
- Identifying unsigned or suspicious binaries
- Forensic analysis of executed files
- Threat hunting and malware detection

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `program_id` | `TEXT` | Associated program identifier (includes SHA1 hash) |
| `file_id` | `TEXT` | Unique file identifier |
| `lower_case_long_path` | `TEXT` | Full lowercase file path |
| `name` | `TEXT` | File name |
| `original_file_name` | `TEXT` | Original file name from PE metadata |
| `publisher` | `TEXT` | File publisher from version information |
| `version` | `TEXT` | File version string |
| `bin_file_version` | `TEXT` | Binary file version |
| `binary_type` | `TEXT` | Binary type (e.g., pe32, pe64) |
| `product_name` | `TEXT` | Product name from version information |
| `product_version` | `TEXT` | Product version string |
| `link_date` | `TEXT` | PE linker timestamp |
| `bin_product_version` | `TEXT` | Binary product version |
| `size` | `BIGINT` | File size in bytes |
| `language` | `BIGINT` | File language code |
| `usn` | `BIGINT` | NTFS Update Sequence Number |
| `appx_package_full_name` | `TEXT` | Full name of the APPX package |
| `is_os_component` | `TEXT` | Whether the file is an OS component |
| `appx_package_relative_id` | `TEXT` | Relative identifier within the APPX package |
| `sha1` | `TEXT` | SHA1 hash extracted from program_id |

## Examples

### List all application files with details

```sql
SELECT name, version, publisher, lower_case_long_path, size
FROM elastic_amcache_application_file
ORDER BY timestamp DESC
LIMIT 100;
```

### Find files by specific publisher

```sql
SELECT name, product_name, version, lower_case_long_path, sha1
FROM elastic_amcache_application_file
WHERE publisher LIKE '%Microsoft%'
ORDER BY name;
```

### Identify large executable files

```sql
SELECT name, size, lower_case_long_path, sha1
FROM elastic_amcache_application_file
WHERE size > 10485760
ORDER BY size DESC;
```

### Find files modified recently

```sql
SELECT name, lower_case_long_path, timestamp, sha1
FROM elastic_amcache_application_file
WHERE timestamp > (strftime('%s', 'now') - 86400 * 7)
ORDER BY timestamp DESC;
```

### Detect potential unsigned binaries

```sql
SELECT name, lower_case_long_path, publisher, sha1
FROM elastic_amcache_application_file
WHERE publisher = '' OR publisher IS NULL
ORDER BY timestamp DESC;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Provides file-level details for applications tracked in Amcache
- The SHA1 hash is automatically extracted from the program_id field
- USN (Update Sequence Number) helps track NTFS file system changes
- Data is cached for performance
- Only available on Windows platforms

## Related Tables

- `elastic_amcache_application`
- `elastic_amcache_application_shortcut`
- `file`
- `hash`
- `authenticode`

