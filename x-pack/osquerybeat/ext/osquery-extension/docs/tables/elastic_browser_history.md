# elastic_browser_history

Query browser history from multiple browsers with a unified schema

## Platforms

- ✅ Linux
- ✅ macOS
- ✅ Windows

## Description

Query browser history from multiple browsers (Chrome, Edge, Firefox, Safari) with a unified schema.

Supports automatic discovery from standard profile locations on Linux, macOS, and Windows.
Can also query custom directories for forensics, backups, or mounted drives using the custom_data_dir constraint.

Returns one row per visit with detailed navigation information including timestamps, URLs, transition types,
and browser-specific metadata.

## Supported Browsers

- **Chrome**: Linux, macOS, Windows
- **Edge**: Linux, macOS, Windows  
- **Firefox**: Linux, macOS, Windows
- **Safari**: macOS only

## Auto-Discovery

The table automatically discovers browsers from standard user profile locations:

**Linux**:
- Chrome: `~/.config/google-chrome/`
- Edge: `~/.config/microsoft-edge/`
- Firefox: `~/.mozilla/firefox/`

**macOS**:
- Chrome: `~/Library/Application Support/Google/Chrome/`
- Edge: `~/Library/Application Support/Microsoft Edge/`
- Firefox: `~/Library/Application Support/Firefox/`
- Safari: `~/Library/Safari/`

**Windows**:
- Chrome: `%LOCALAPPDATA%\Google\Chrome\User Data\`
- Edge: `%LOCALAPPDATA%\Microsoft\Edge\User Data\`
- Firefox: `%APPDATA%\Mozilla\Firefox\`

## Custom Data Directories

Use the `custom_data_dir` constraint to query non-standard locations:

```sql
-- Query from backup
SELECT * FROM elastic_browser_history 
WHERE custom_data_dir = '/mnt/backup/Users/john/AppData/Local/Google';

-- Query with glob pattern
SELECT * FROM elastic_browser_history 
WHERE custom_data_dir GLOB '/forensics/users/*/Library/Application Support/Google';
```

## Field Distinction

**transition_type** = HOW the user navigated (navigation method)  
**visit_source** = WHERE the visit data originated (data provenance)

Example: `transition_type = 'TYPED'` and `visit_source = 'synced'` means the user typed the URL on another device.

## Browser-Specific Fields

- **Chromium fields** (`ch_*`): Only populated for Chrome, Edge, Brave
- **Firefox fields** (`ff_*`): Only populated for Firefox
- **Safari fields** (`sf_*`): Only populated for Safari
- Universal fields are always populated regardless of browser

## Schema

| Column | Type | Description |
|--------|------|-------------|
| `timestamp` | `BIGINT` | Unix timestamp in seconds since epoch (visit time) |
| `datetime` | `TEXT` | Human-readable datetime string in RFC3339 format (visit time) |
| `url_id` | `BIGINT` | Unique URL identifier within the browser profile |
| `scheme` | `TEXT` | URL scheme/protocol (e.g., "https", "http", "file", "ftp") |
| `domain` | `TEXT` | Registrable domain (eTLD+1) extracted from hostname (e.g., "github.com", "google.com", "example.co.uk") |
| `hostname` | `TEXT` | Full hostname from URL (e.g., "www.github.com", "mail.google.com") |
| `url` | `TEXT` | Full URL visited |
| `title` | `TEXT` | Page title |
| `browser` | `TEXT` | Browser name (chrome, edge, firefox, safari, brave) |
| `parser` | `TEXT` | Parser used to extract the history |
| `user` | `TEXT` | Username/profile owner |
| `profile_name` | `TEXT` | Browser profile name |
| `transition_type` | `TEXT` | Navigation method - how user reached this page (TYPED, LINK, BOOKMARK, RELOAD, REDIRECT, etc.) |
| `referring_url` | `TEXT` | URL that linked to this page (if navigated via link) |
| `visit_id` | `BIGINT` | Unique visit identifier |
| `from_visit_id` | `BIGINT` | Visit ID that led to this visit (navigation chain) |
| `visit_source` | `TEXT` | Data origin - where this visit data came from (browsed/local, synced, imported, extension) |
| `is_hidden` | `INTEGER` | Whether visit is hidden (1) or visible (0) |
| `history_path` | `TEXT` | Path to the browser history database file |
| `ch_visit_duration_ms` | `BIGINT` | Duration of visit in milliseconds (Chromium-based browsers only) |
| `ff_session_id` | `INTEGER` | Firefox session tracking identifier |
| `ff_frecency` | `INTEGER` | Firefox frecency score (frequency + recency algorithm) |
| `sf_domain_expansion` | `TEXT` | Safari domain classification/expansion |
| `sf_load_successful` | `INTEGER` | Whether page loaded successfully (1) or failed (0) |
| `custom_data_dir` | `TEXT` | Custom data directory path (for querying non-standard locations) |

## Examples

### Get all browser history

```sql
SELECT * FROM elastic_browser_history;
```

### Get history from specific browser

```sql
SELECT * FROM elastic_browser_history WHERE browser = 'chrome';
```

### Recent history (last 7 days)

```sql
SELECT url, title, browser, datetime 
FROM elastic_browser_history 
WHERE timestamp > (strftime('%s', 'now') - 604800)
ORDER BY timestamp DESC;
```

### Search for specific domains (registrable domain)

```sql
SELECT browser, profile_name, url, title, datetime
FROM elastic_browser_history
WHERE domain = 'github.com'
ORDER BY timestamp DESC;
```

### Most visited domains

```sql
SELECT domain, COUNT(*) as visits
FROM elastic_browser_history
WHERE domain != ''
GROUP BY domain
ORDER BY visits DESC
LIMIT 20;
```

### Find typed URLs (direct user navigation)

```sql
SELECT url, title, browser, datetime
FROM elastic_browser_history
WHERE transition_type LIKE '%TYPED%'
ORDER BY timestamp DESC;
```

### Query from custom directory

```sql
SELECT * FROM elastic_browser_history 
WHERE custom_data_dir = '/mnt/backup/Users/john/Library/Application Support/Google';
```

### Analyze navigation methods

```sql
SELECT transition_type, COUNT(*) as count
FROM elastic_browser_history
GROUP BY transition_type
ORDER BY count DESC;
```

### Most visited hostnames

```sql
SELECT hostname, COUNT(*) as visits
FROM elastic_browser_history
WHERE hostname != ''
GROUP BY hostname
ORDER BY visits DESC
LIMIT 20;
```

### Find non-HTTPS visits (security audit)

```sql
SELECT browser, url, title, datetime
FROM elastic_browser_history
WHERE scheme IN ('http', 'ftp')
ORDER BY timestamp DESC;
```

## Notes

- Automatically discovers browsers from standard user profile locations
- Supports Chrome, Edge, Firefox (all platforms) and Safari (macOS only)
- Use custom_data_dir constraint to query non-standard locations for forensics or backups
- Returns one row per visit with detailed navigation metadata
- Browser-specific fields (ch_*, ff_*, sf_*) only populated for respective browsers
- Requires read access to browser profile directories
- url_id is unique within a browser profile and used for grouping visits to the same URL
- Use timestamp filters to improve performance on large history databases
- transition_type shows HOW user navigated (method), visit_source shows WHERE data originated (provenance)
- domain field contains registrable domain (eTLD+1), hostname contains full hostname

## Related Tables

- `users`
- `logged_in_users`
- `processes`
- `file`
- `hash`

