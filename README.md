# storage

a tool use free CDN store file

## how it works

many website provider api to store image on CDN without authority check, this tool store file pieces as image on CDN. 

## usage

```bash
# upload
storage /path/to/file # upload file to CDN, generated pieces mate for download
# download
storage /path/to/file.meta
# download url
storage https://dxkite.cn/meta/jetbrains-agent-latest.zip.meta
# download meta uri
storage storage://meta?dl=aHR0cHM6Ly9keGtpdGUuY24vbWV0YS9qZXRicmFpbnMtYWdlbnQtbGF0ZXN0LnppcC5tZXRh
```

## supported CDN node 

- ali: alibaba CDN **invalid**