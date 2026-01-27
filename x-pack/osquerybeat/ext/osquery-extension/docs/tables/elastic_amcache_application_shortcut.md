# elastic_amcache_application_shortcut

Windows Amcache application shortcuts from the InventoryApplicationShortcut registry hive

## Platforms

- ❌ Linux
- ❌ macOS
- ✅ Windows

## Description

The elastic_amcache_application_shortcut table provides information about
application shortcuts (.lnk files) tracked in the Windows Amcache registry.
This includes Start Menu shortcuts, Desktop shortcuts, and other application
launch points.

This data is valuable for:
- Tracking application launch mechanisms
- Identifying persistence mechanisms via shortcuts
- Forensic analysis of user activity
- Detecting malicious shortcuts
- Understanding application installation patterns

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Last write time of the registry key in UNIX epoch seconds |
| `date_time` | `TEXT` | Last write time of the registry key in RFC3339 format |
| `shortcut_path` | `TEXT` | Full path to the shortcut file (.lnk) |
| `shortcut_target_path` | `TEXT` | Target path that the shortcut points to |
| `shortcut_aumid` | `TEXT` | Application User Model ID (AUMID) for the shortcut |
| `shortcut_program_id` | `TEXT` | Program identifier associated with the shortcut |

## Examples

### List all application shortcuts

```sql
SELECT shortcut_path, shortcut_target_path, timestamp
FROM elastic_amcache_application_shortcut
ORDER BY timestamp DESC;
```

### Find shortcuts in Start Menu

```sql
SELECT shortcut_path, shortcut_target_path, shortcut_aumid
FROM elastic_amcache_application_shortcut
WHERE shortcut_path LIKE '%Start Menu%'
ORDER BY shortcut_path;
```

### Find shortcuts on Desktop

```sql
SELECT shortcut_path, shortcut_target_path, timestamp
FROM elastic_amcache_application_shortcut
WHERE shortcut_path LIKE '%Desktop%'
ORDER BY timestamp DESC;
```

### Find recently created shortcuts

```sql
SELECT shortcut_path, shortcut_target_path, date_time
FROM elastic_amcache_application_shortcut
WHERE timestamp > (strftime('%s', 'now') - 86400 * 7)
ORDER BY timestamp DESC;
```

### Find shortcuts pointing to specific locations

```sql
SELECT shortcut_path, shortcut_target_path, shortcut_program_id
FROM elastic_amcache_application_shortcut
WHERE shortcut_target_path LIKE '%System32%'
ORDER BY shortcut_path;
```

## Notes

- This table reads from the Windows Amcache.hve registry hive
- Tracks LNK (shortcut) files created and used by applications
- AUMID (Application User Model ID) is used for app identification in Windows
- Data is cached for performance
- Only available on Windows platforms
- Useful for detecting persistence mechanisms that use shortcuts

## Related Tables

- `elastic_amcache_application`
- `elastic_amcache_application_file`
- `file`
- `startup_items`

