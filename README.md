# SAS API
Menyajikan data dari database aplikasi SAS dalam format JSON

## Usage
```
name: SAS API
version: 1.0.2
date: 29-09-2018

URL: https://yourdomain.tld/api/v1
Method: POST
parameters:
    action: pagu
    token: YourToken

    action: realisasi
    token: YourToken
    startDate: YYYY-MM-DD
    endDate: YYYY-MM-DD

Changes in version 1.0.2 (29-09-2018)
    realisasi:
    + add column kdkmpnen
    + add column kdskmpnen
    - remove nmitem

Changes in version 1.0.1 (24-09-2018)
    + add action: pagu
```

## Feedback
arief.karfianto@ksp.go.id
