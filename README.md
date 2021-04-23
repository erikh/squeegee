# Squeegee: so I scrape when you scrape (prometheus)

This is a simple tool to:

- Scrape JSON endpoints according to (get excited) a YAML syntax
- Parse them with [jsonquery](https://github.com/antchfx/jsonquery)
- Shove them into a prometheus gauge
- Serve that over port 8000

That's it, folks.

## Syntax

The syntax is simple, here's a complete example:

```yaml
---
metrics:
  watchers_count:
    url: https://api.github.com/repos/erikh/tftest
    headers:
      Authorization: "bearer <token>"
    query: /watchers_count
    interval: 1s
```

This will poll the watchers count for [erikh/tftest](https://github.com/erikh/tftest) and expose it as a prometheus metric called `watchers_count` for the scraped host.

## License

MIT

## Author

Erik Hollensbe <github@hollensbe.org>
