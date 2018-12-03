RSS Feed Concourse Resource
===========================

Retrieves and parses an Atom/RSS 2.0 XML feed from an arbitrary
URL, and splits out each syndicated post into files on-disk.

Resource Type Configuration
---------------------------

```yaml
resource_types:
- name: rss
  type: docker-image
  source:
    repository: starkandwayne/rss-resource
    tag:        latest
```

Source Configuration
--------------------

```yaml
resources:
- name: my-blog
  type: rss
  source:
    url: https://example.com/rss.xml
```

The following `source` properties are defined:

- `url` - The URL of the RSS feed to consume (_required_).
- `skip_tls_verify` - Skip verification of remote TLS
  certificates.  This is highly discouraged, but allows you to
  work around real-world problems like subject name mismatch and
  certificate expiration.
