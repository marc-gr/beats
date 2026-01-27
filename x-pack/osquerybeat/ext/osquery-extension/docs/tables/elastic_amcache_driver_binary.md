# elastic_amcache_driver_binary

Windows Amcache driver binary inventory from the InventoryDriverBinary registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_driver_binary table provides detailed information about
driver binaries tracked in the Windows Amcache registry. This includes kernel-mode
and user-mode drivers, their versions, signing status, and metadata.

This data is valuable for:
- Driver inventory and version management
- Detecting unsigned or suspicious drivers
- Security monitoring for rootkits and malicious drivers
- Compliance verification for approved drivers
- Troubleshooting driver compatibility issues
- Forensic analysis of driver installations

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `driver_name` | `TEXT` | Name of the driver binary |
| `inf` | `TEXT` | INF file used for driver installation |
| `driver_version` | `TEXT` | Driver version string |
| `product` | `TEXT` | Product name from driver metadata |
| `product_version` | `TEXT` | Product version string |
| `wdf_version` | `TEXT` | Windows Driver Framework (WDF) version |
| `driver_company` | `TEXT` | Company that developed the driver |
| `driver_package_strong_name` | `TEXT` | Strong name of the driver package |
| `service` | `TEXT` | Windows service name for the driver |
| `driver_in_box` | `TEXT` | Whether the driver is included in Windows |
| `driver_signed` | `TEXT` | Whether the driver is digitally signed |
| `driver_is_kernel_mode` | `TEXT` | Whether the driver runs in kernel mode |
| `driver_id` | `TEXT` | Unique driver identifier |
| `driver_last_write_time` | `TEXT` | Last write time from driver metadata |
| `driver_type` | `BIGINT` | Driver type code |
| `driver_time_stamp` | `BIGINT` | PE header timestamp |
| `driver_check_sum` | `BIGINT` | PE header checksum |
| `image_size` | `BIGINT` | Size of the driver image in bytes |

## Examples

### List all driver binaries

```sql
SELECT driver_name, driver_version, driver_company, driver_signed
FROM elastic_amcache_driver_binary
ORDER BY timestamp DESC;
```

### Find unsigned drivers

```sql
SELECT driver_name, driver_company, product, driver_is_kernel_mode
FROM elastic_amcache_driver_binary
WHERE driver_signed = '0' OR driver_signed = 'False'
ORDER BY timestamp DESC;
```

### Find kernel-mode drivers

```sql
SELECT driver_name, driver_company, driver_version, driver_signed
FROM elastic_amcache_driver_binary
WHERE driver_is_kernel_mode = '1' OR driver_is_kernel_mode = 'True'
ORDER BY driver_name;
```

### Find third-party drivers (not inbox)

```sql
SELECT driver_name, driver_company, driver_version, inf
FROM elastic_amcache_driver_binary
WHERE driver_in_box = '0' OR driver_in_box = 'False'
ORDER BY driver_company, driver_name;
```

### Find recently installed drivers

```sql
SELECT driver_name, driver_company, driver_version, date_time
FROM elastic_amcache_driver_binary
WHERE timestamp > (strftime('%s', 'now') - 86400 * 30)
ORDER BY timestamp DESC;
```

### List drivers by company

```sql
SELECT driver_company, COUNT(*) as driver_count,
       SUM(CASE WHEN driver_signed = '1' THEN 1 ELSE 0 END) as signed_count
FROM elastic_amcache_driver_binary
GROUP BY driver_company
ORDER BY driver_count DESC;
```

### Find large driver binaries

```sql
SELECT driver_name, driver_company, image_size, driver_is_kernel_mode
FROM elastic_amcache_driver_binary
WHERE image_size > 1048576
ORDER BY image_size DESC;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Tracks all driver binaries installed on the system
- driver_signed indicates if the driver has a valid digital signature
- driver_is_kernel_mode distinguishes kernel vs user-mode drivers
- driver_in_box indicates if the driver ships with Windows
- Unsigned or suspicious drivers may indicate malware or rootkits
- WDF (Windows Driver Framework) version indicates driver architecture
- Data is cached for performance
- Only available on Windows platforms

## Related Tables

- `elastic_amcache_driver_package`
- `elastic_amcache_device_pnp`
- `drivers`
- `authenticode`
- `file`

