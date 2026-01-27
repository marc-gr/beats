# elastic_amcache_applications_view

View joining amcache application and file tables with UNION to show all applications and orphaned files

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_applications_view provides a comprehensive unified view of
Windows Amcache application and file data. It combines records from both the
InventoryApplication and InventoryApplicationFile hives using UNION ALL to ensure
all data is captured.

The view has two parts:
1. Applications with their associated files (LEFT JOIN from application table)
2. Orphaned files with no matching application (anti-join to find files without apps)

This approach ensures:
- All applications are represented (even those without file entries)
- All files are represented (even orphaned files without app entries)
- No data is lost from either table
- File data is prioritized when available via COALESCE

This view is valuable for:
- Complete application and file inventory
- Identifying orphaned files that may indicate partial uninstalls
- Correlating applications with their binaries
- Security analysis of application components
- Forensic investigation of installed software

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time in RFC3339 format |
| `program_id` | `TEXT` | Program identifier (anchor for joining) |
| `file_id` | `TEXT` | File identifier (from application_file table) |
| `lower_case_long_path` | `TEXT` | Full lowercase file path |
| `name` | `TEXT` | Application or file name |
| `original_file_name` | `TEXT` | Original file name from PE metadata |
| `publisher` | `TEXT` | Publisher or vendor |
| `version` | `TEXT` | Version string |
| `bin_file_version` | `TEXT` | Binary file version |
| `binary_type` | `TEXT` | Binary type (e.g., pe32, pe64) |
| `product_name` | `TEXT` | Product name from version information |
| `product_version` | `TEXT` | Product version string |
| `link_date` | `TEXT` | PE linker timestamp |
| `bin_product_version` | `TEXT` | Binary product version |
| `size` | `BIGINT` | File size in bytes |
| `language` | `BIGINT` | Language code |
| `usn` | `BIGINT` | NTFS Update Sequence Number |
| `appx_package_full_name` | `TEXT` | Full name of the APPX package |
| `is_os_component` | `TEXT` | Whether the file is an OS component |
| `appx_package_relative_id` | `TEXT` | Relative identifier within the APPX package |
| `file_sha1` | `TEXT` | SHA1 hash from file table |
| `program_instance_id` | `TEXT` | Program instance identifier (from application table) |
| `install_date` | `TEXT` | Installation date |
| `source` | `TEXT` | Installation source |
| `root_dir_path` | `TEXT` | Root directory path of the application |
| `hidden_arp` | `BIGINT` | Whether hidden in Add/Remove Programs |
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
| `user_sid` | `TEXT` | User security identifier |
| `app_sha1` | `TEXT` | SHA1 hash from application table |

## Required Tables

This view requires the following tables to be available:

- `elastic_amcache_application`
- `elastic_amcache_application_file`

## View Definition

```sql
CREATE VIEW elastic_amcache_applications_view AS
SELECT
  COALESCE(file.timestamp, app.timestamp) AS timestamp,
  COALESCE(file.date_time, app.date_time) AS date_time,
  app.program_id,
  file.file_id,
  file.lower_case_long_path,
  COALESCE(file.name, app.name) AS name,
  file.original_file_name,
  COALESCE(file.publisher, app.publisher) AS publisher,
  COALESCE(file.version, app.version) AS version,
  file.bin_file_version,
  file.binary_type,
  file.product_name,
  file.product_version,
  file.link_date,
  file.bin_product_version,
  file.size,
  COALESCE(file.language, app.language) AS language,
  file.usn,
  file.appx_package_full_name,
  file.is_os_component,
  file.appx_package_relative_id,
  file.sha1 AS file_sha1,
  app.program_instance_id,
  app.install_date,
  app.source,
  app.root_dir_path,
  app.hidden_arp,
  app.uninstall_string,
  app.registry_key_path,
  app.store_app_type,
  app.inbox_modern_app,
  app.manifest_path,
  app.package_full_name,
  app.msi_package_code,
  app.msi_product_code,
  app.msi_install_date,
  app.bundle_manifest_path,
  app.user_sid,
  app.sha1 AS app_sha1
FROM elastic_amcache_application AS app
LEFT JOIN elastic_amcache_application_file AS file 
  ON app.program_id = file.program_id

UNION ALL

SELECT
  file.timestamp,
  file.date_time,
  file.program_id,
  file.file_id,
  file.lower_case_long_path,
  file.name,
  file.original_file_name,
  file.publisher,
  file.version,
  file.bin_file_version,
  file.binary_type,
  file.product_name,
  file.product_version,
  file.link_date,
  file.bin_product_version,
  file.size,
  file.language,
  file.usn,
  file.appx_package_full_name,
  file.is_os_component,
  file.appx_package_relative_id,
  file.sha1 AS file_sha1,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL,
  NULL AS app_sha1
FROM elastic_amcache_application_file AS file
LEFT JOIN elastic_amcache_application AS app 
  ON file.program_id = app.program_id
WHERE app.program_id IS NULL;
```

## Examples

### List all applications with file information

```sql
SELECT name, version, publisher, file_id, lower_case_long_path
FROM elastic_amcache_applications_view
WHERE file_id IS NOT NULL
ORDER BY name;
```

### Find applications without file records

```sql
SELECT name, version, publisher, program_id
FROM elastic_amcache_applications_view
WHERE file_id IS NULL
ORDER BY name;
```

### Find orphaned files without application records

```sql
SELECT name, lower_case_long_path, file_sha1, program_id
FROM elastic_amcache_applications_view
WHERE program_instance_id IS NULL AND file_id IS NOT NULL
ORDER BY timestamp DESC;
```

### Analyze file sizes by application

```sql
SELECT name, version, COUNT(*) as file_count, SUM(size) as total_size
FROM elastic_amcache_applications_view
WHERE file_id IS NOT NULL
GROUP BY name, version
ORDER BY total_size DESC;
```

### Find recently added applications or files

```sql
SELECT name, version, publisher, file_id, timestamp
FROM elastic_amcache_applications_view
WHERE timestamp > (strftime('%s', 'now') - 86400 * 7)
ORDER BY timestamp DESC;
```

## Notes

- This view combines elastic_amcache_application and elastic_amcache_application_file
- Uses UNION ALL to include all applications and all files
- Part 1 (LEFT JOIN) shows applications with their associated files
- Part 2 (anti-join) shows orphaned files without matching applications
- COALESCE prioritizes file data over application data when both exist
- program_id is the key field used for joining the two tables
- Orphaned files can indicate incomplete installations or uninstalls
- Only available on Windows platforms
- View is automatically created when the osquery extension starts

## Related Tables

- `elastic_amcache_application`
- `elastic_amcache_application_file`
- `elastic_amcache_application_shortcut`
- `programs`
- `file`

