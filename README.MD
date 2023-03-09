# Universal healthcheck from DevOps for DevOps :D

This module is a part of [healthcheck & wait](https://github.com/hypolas/healthcheck)

This can check:
- API response
- http code status

# Configuration

Export this environnement variables:

| Variable | Default value |  Description|   Required |   Set on |  More inf |
|---|---|---|---|---|---|
| `HYPOLAS_HEALTHCHECK_TYPE`  | *""*  | Type of healthcheck  | **X** | **http** |
| `HYPOLAS_HEALTHCHECK_HTTP_URL`  | *""*   | URL   | **X** |  |
| `HYPOLAS_HEALTHCHECK_HTTP_RESPONSES`  | *""*  | Hostname/IP of target   | | | (1)
| `HYPOLAS_HEALTHCHECK_HTTP_JSON`  | *""*  | Flatten JSON PATH   | | | (2)
| `HYPOLAS_HEALTHCHECK_HTTP_EXPECTED`  | *""*  | Response of JSON PATH   | | | (3)
| `HYPOLAS_HEALTHCHECK_HTTP_HEADERS`  | *""*  | Add Headers flatten   | | | (4)
| `HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT`  | *""*  | Set a Timout  | 
| `HYPOLAS_LOGS_FILE`  | *""*  | Path of logs file  | **X**

# Particularité => Les réponses aux API

## Exemple

Les APIs au format REST (json) sont prisent en charge. Vous pouvez allez chercher la valeur d'une réponse en profondeur.

Prenons cette réponse de l'API:

```json
{
  "hostname": "myhost",
  "services": [
    {
      "nom": "nginx",
      "status": "up"
    },
    {
      "nom": "httpd",
      "status": "down"
    }
  ]
}
```

HYPOLAS_HEALTHCHECK_HTTP_JSON est de type string flat avec comme séparateur un double "\_".

Je veux vérifier que la réponse du service "httpd" est "up". Pour cela je vais exporter 3 variables d'environnement:

```bash
export HYPOLAS_HEALTHCHECK_TYPE=http
export HYPOLAS_HEALTHCHECK_HTTP_JSON=services__1__status
export HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=up
```

Le service index 1 étant `httpd`, le résultat sera un Exit(1).
