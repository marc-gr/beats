# elastic_amcache_application

Windows Amcache application inventory from the InventoryApplication registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_application table provides access to Windows Amcache application
inventory data from the InventoryApplication registry hive. Amcache stores information
about installed applications, including metadata such as install dates, versions,
publishers, and file hashes.

This data is valuable for:
- Software inventory and license management
- Detecting unauthorized or malicious software installations
- Forensic analysis of application execution history
- Compliance monitoring and auditing

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `program_id` | `TEXT` | Unique program identifier (includes SHA1 hash) |
| `program_instance_id` | `TEXT` | Instance identifier for the program |
| `name` | `TEXT` | Application name |
| `version` | `TEXT` | Application version |
| `publisher` | `TEXT` | Application publisher or vendor |
| `language` | `BIGINT` | Application language code |
| `install_date` | `TEXT` | Installation date |
| `source` | `TEXT` | Installation source |
| `root_dir_path` | `TEXT` | Root directory path of the application |
| `hidden_arp` | `BIGINT` | Whether the application is hidden in Add/Remove Programs |
| `uninstall_string` | `TEXT` | Command to uninstall the application |
| `registry_key_path` | `TEXT` | Registry key path for the application |
| `store_app_type` | `TEXT` | Microsoft Store app type |
| `inbox_modern_app` | `TEXT` | Whether this is an inbox modern app |
| `manifest_path` | `TEXT` | Path to the application manifest |
| `package_full_name` | `TEXT` | Full package name for UWP apps |
| `msi_package_code` | `TEXT` | MSI package code |
| `msi_product_code` | `TEXT` | MSI product code |
| `msi_install_date` | `TEXT` | MSI installation date |
| `bundle_manifest_path` | `TEXT` | Path to bundle manifest |
| `user_sid` | `TEXT` | Security identifier of the user who installed the application |
| `sha1` | `TEXT` | SHA1 hash extracted from program_id |

## Examples

### List all installed applications

```sql
SELECT name, version, publisher, install_date
FROM elastic_amcache_application
ORDER BY install_date DESC;
```

### Find applications by publisher

```sql
SELECT name, version, install_date, sha1
FROM elastic_amcache_application
WHERE publisher LIKE '%Microsoft%'
ORDER BY name;
```

### Detect applications with suspicious characteristics

```sql
SELECT name, publisher, root_dir_path, sha1
FROM elastic_amcache_application
WHERE hidden_arp = 1 OR publisher = ''
ORDER BY timestamp DESC;
```

### Find recently installed applications

```sql
SELECT name, version, publisher, install_date, timestamp
FROM elastic_amcache_application
WHERE timestamp > (strftime('%s', 'now') - 86400 * 7)
ORDER BY timestamp DESC;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Requires appropriate permissions to access registry data
- The SHA1 hash is automatically extracted from the program_id field
- Data is cached for performance; the cache is refreshed periodically
- Only available on Windows platforms

## Related Tables

- `elastic_amcache_application_file`
- `elastic_amcache_application_shortcut`
- `programs`
- `registry`

