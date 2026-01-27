# elastic_amcache_device_pnp

Windows Amcache Plug and Play device inventory from the InventoryDevicePnp registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_device_pnp table provides detailed information about
Plug and Play (PnP) devices tracked in the Windows Amcache registry. This
includes hardware devices, their drivers, installation details, and current state.

This data is valuable for:
- Hardware inventory and asset management
- Driver version tracking and compliance
- Detecting unauthorized hardware (e.g., USB devices)
- Troubleshooting device and driver issues
- Security monitoring for malicious hardware
- Forensic analysis of connected devices

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `model` | `TEXT` | Device model name |
| `manufacturer` | `TEXT` | Device manufacturer name |
| `driver_name` | `TEXT` | Name of the installed driver |
| `parent_id` | `TEXT` | Parent device identifier |
| `matching_id` | `TEXT` | Hardware ID that matched during driver installation |
| `class` | `TEXT` | Device class (e.g., Display, Net, USB) |
| `class_guid` | `TEXT` | Device class GUID |
| `description` | `TEXT` | Device description |
| `enumerator` | `TEXT` | Device enumerator (e.g., PCI, USB, ROOT) |
| `service` | `TEXT` | Service name associated with the device |
| `install_state` | `TEXT` | Device installation state |
| `device_state` | `TEXT` | Current device state |
| `inf` | `TEXT` | INF file used for device installation |
| `driver_ver_date` | `TEXT` | Driver version date |
| `install_date` | `TEXT` | Device installation date |
| `first_install_date` | `TEXT` | First installation date of the device |
| `driver_package_strong_name` | `TEXT` | Strong name of the driver package |
| `driver_ver_version` | `TEXT` | Driver version string |
| `container_id` | `TEXT` | Container ID for the device |
| `problem_code` | `TEXT` | Device problem code if any |
| `provider` | `TEXT` | Driver provider |
| `driver_id` | `TEXT` | Driver identifier |
| `bus_reported_description` | `TEXT` | Description reported by the bus driver |
| `hw_id` | `TEXT` | Hardware ID |
| `extended_infs` | `TEXT` | Extended INF files |
| `compid` | `TEXT` | Compatible ID |
| `stack_id` | `TEXT` | Device stack identifier |
| `upper_class_filters` | `TEXT` | Upper class filter drivers |
| `lower_class_filters` | `TEXT` | Lower class filter drivers |
| `upper_filters` | `TEXT` | Upper filter drivers |
| `lower_filters` | `TEXT` | Lower filter drivers |
| `device_interface_classes` | `TEXT` | Device interface class GUIDs |
| `location_paths` | `TEXT` | Device location paths |

## Examples

### List all PnP devices

```sql
SELECT model, manufacturer, class, driver_name, device_state
FROM elastic_amcache_device_pnp
ORDER BY install_date DESC;
```

### Find USB devices

```sql
SELECT model, manufacturer, description, install_date, hw_id
FROM elastic_amcache_device_pnp
WHERE enumerator = 'USB' OR class = 'USB'
ORDER BY install_date DESC;
```

### Find devices with problems

```sql
SELECT model, manufacturer, description, problem_code, device_state
FROM elastic_amcache_device_pnp
WHERE problem_code IS NOT NULL AND problem_code != ''
ORDER BY timestamp DESC;
```

### Track network devices

```sql
SELECT model, manufacturer, description, driver_name, driver_ver_version
FROM elastic_amcache_device_pnp
WHERE class = 'Net'
ORDER BY model;
```

### Find recently installed devices

```sql
SELECT model, manufacturer, class, install_date, first_install_date
FROM elastic_amcache_device_pnp
WHERE timestamp > (strftime('%s', 'now') - 86400 * 7)
ORDER BY timestamp DESC;
```

### List devices by manufacturer

```sql
SELECT manufacturer, COUNT(*) as device_count, 
       GROUP_CONCAT(DISTINCT class) as classes
FROM elastic_amcache_device_pnp
GROUP BY manufacturer
ORDER BY device_count DESC;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Tracks all PnP devices that have been connected to the system
- Useful for detecting rogue hardware and unauthorized USB devices
- Problem codes indicate device errors or issues
- Filter drivers can indicate security software or potentially malicious drivers
- Data is cached for performance
- Only available on Windows platforms

## Related Tables

- `elastic_amcache_driver_binary`
- `elastic_amcache_driver_package`
- `drivers`
- `device_file`
- `usb_devices`

