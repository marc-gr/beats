# elastic_amcache_driver_package

Windows Amcache driver package inventory from the InventoryDriverPackage registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_driver_package table provides information about driver
packages tracked in the Windows Amcache registry. A driver package is a
collection of files and metadata required to install a device driver.

This data is valuable for:
- Driver package inventory and management
- Tracking driver updates and versions
- Security monitoring for malicious driver packages
- Compliance verification for certified drivers
- Troubleshooting driver installation issues
- Forensic analysis of driver package installations

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `class_guid` | `TEXT` | Device class GUID |
| `class` | `TEXT` | Device class name |
| `directory` | `TEXT` | Driver package directory path |
| `date` | `TEXT` | Driver package date |
| `version` | `TEXT` | Driver package version |
| `provider` | `TEXT` | Driver package provider/vendor |
| `submission_id` | `TEXT` | Windows Hardware Quality Labs (WHQL) submission ID |
| `driver_in_box` | `TEXT` | Whether the driver package is included in Windows |
| `inf` | `TEXT` | INF file name for the driver package |
| `flight_ids` | `TEXT` | Windows Insider flight identifiers |
| `recovery_ids` | `TEXT` | Recovery partition identifiers |
| `is_active` | `TEXT` | Whether the driver package is currently active |
| `hwids` | `TEXT` | Hardware IDs supported by this driver package |
| `sysfile` | `TEXT` | System files included in the driver package |

## Examples

### List all driver packages

```sql
SELECT class, provider, version, inf, is_active
FROM elastic_amcache_driver_package
ORDER BY timestamp DESC;
```

### Find active driver packages

```sql
SELECT class, provider, version, directory
FROM elastic_amcache_driver_package
WHERE is_active = '1' OR is_active = 'True'
ORDER BY class, provider;
```

### Find third-party driver packages

```sql
SELECT provider, class, version, inf, submission_id
FROM elastic_amcache_driver_package
WHERE driver_in_box = '0' OR driver_in_box = 'False'
ORDER BY provider, class;
```

### List driver packages by class

```sql
SELECT class, COUNT(*) as package_count,
       GROUP_CONCAT(DISTINCT provider) as providers
FROM elastic_amcache_driver_package
GROUP BY class
ORDER BY package_count DESC;
```

### Find recently installed driver packages

```sql
SELECT provider, class, version, date, date_time
FROM elastic_amcache_driver_package
WHERE timestamp > (strftime('%s', 'now') - 86400 * 30)
ORDER BY timestamp DESC;
```

### Find driver packages with WHQL certification

```sql
SELECT provider, class, version, submission_id, inf
FROM elastic_amcache_driver_package
WHERE submission_id IS NOT NULL AND submission_id != ''
ORDER BY provider;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Driver packages contain INF files and related driver binaries
- submission_id indicates WHQL (Windows Hardware Quality Labs) certification
- driver_in_box indicates if the package ships with Windows
- is_active shows whether the package is currently in use
- hwids lists the hardware IDs that the package supports
- Data is cached for performance
- Only available on Windows platforms

## Related Tables

- `elastic_amcache_driver_binary`
- `elastic_amcache_device_pnp`
- `drivers`
- `file`

