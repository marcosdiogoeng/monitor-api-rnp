# Network Monitoring

Este projeto realiza monitoramento de rede usando um agente em Go que coleta métricas de **ping** e **requisições HTTP**. Os dados são armazenados em um banco **PostgreSQL** e visualizados via **Grafana**.

---

## Componentes

- **Agente (Go)**: Executa testes de ping e HTTP periodicamente.
- **PostgreSQL**: Armazena os resultados dos testes.
- **Grafana**: Exibe dashboards com as métricas de rede.
- **Docker Compose**: Orquestra os serviços.

---

## Como executar

1. Suba os containers:
```bash
docker-compose up --build
```

3. Acesse o Grafana:
- URL: [http://localhost:3000](http://localhost:3000)
- Usuário: `admin`
- Senha: `admin`

---

## Fluxo de Dados

1. O agente coleta dados de `ping` e `HTTP` a cada 5 minutos
2. Os resultados são salvos no banco PostgreSQL
3. O Grafana consulta os dados via SQL e exibe nos dashboards

---

## Estrutura do Banco

### ping_results
| Campo       | Tipo     |
|-------------|----------|
| host        | TEXT     |
| packet_loss | REAL     |
| avg_latency | REAL     |
| created_at  | TIMESTAMP |

### http_results
| Campo       | Tipo     |
|-------------|----------|
| url         | TEXT     |
| status_code | INT      |
| duration_ms | INT      |
| created_at  | TIMESTAMP |

---

## Dashboards

- **Latência Média por Host**
- **Tempo de Carregamento por URL**
- **Códigos de Resposta HTTP**

Provisionamento automático via:
- `grafana/provisioning/dashboards.yml`
- `grafana/dashboards/network-monitor.json`

---

## Possíveis Melhorias

- Exportar métricas para Prometheus ou InfluxDB
- Configurar alertas no Grafana
- Enviar alertas


