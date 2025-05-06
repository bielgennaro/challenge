# 🛰️ V3 Challenge API

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.24-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-Rekognition-FF9900?logo=amazon-aws&logoColor=white)

## 📋 Visão Geral

V3 Challenge API é uma API robusta desenvolvida para receber, validar e processar dados de telemetria de dispositivos IoT, incluindo informações de giroscópio, GPS e imagens. A API utiliza tecnologias modernas como Node.js, Docker e integração com AWS Rekognition para reconhecimento facial avançado.

## 🚀 Funcionalidades

- ✅ Recebimento e validação de dados de telemetria (giroscópio, GPS, fotos)
- 📦 Armazenamento consistente em banco de dados
- 🧠 Reconhecimento inteligente de imagens com AWS Rekognition
- 🔍 Verificação de similaridade entre fotos enviadas
- 🐳 Containerização completa com Docker
- 🧪 Cobertura completa de testes unitários

## 📚 Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/telemetry/gyroscope` | Recebe e processa dados do giroscópio |
| `POST` | `/telemetry/gps` | Recebe e processa dados de GPS |
| `POST` | `/telemetry/photo` | Recebe, processa e analisa fotos |

## 📐 Estrutura dos Dados ( Pode ser alterado futuramente, apenas exemplos)

### Giroscópio
```json
{
  "deviceId": "string",
  "timestamp": "ISO8601",
  "x": "number",
  "y": "number",
  "z": "number",
  "precision": "number"
}
```

### GPS
```json
{
  "deviceId": "string",
  "timestamp": "ISO8601",
  "latitude": "number",
  "longitude": "number",
  "altitude": "number",
  "accuracy": "number"
}
```

### Foto
```json
{
  "deviceId": "string",
  "timestamp": "ISO8601",
  "imageData": "base64",
  "format": "string",
  "metadata": {
    "location": {
      "latitude": "number",
      "longitude": "number"
    },
    "deviceInfo": "string"
  }
}
```

## 🖥️ Tecnologias Utilizadas

- **Backend**: Golang
- **Banco de Dados**: Postgres
- **Containerização**: Docker, Docker Compose
- **Testes**: Go testing, Testify
- **Reconhecimento de Imagem**: AWS Rekognition

## 🛠️ Configuração e Instalação

### Pré-requisitos

- Go (v1.24 ou superior)
- Docker e Docker Compose
- Conta AWS (para Rekognition)


## 📊 Resultados da Análise de Imagem

Quando uma foto é enviada ao endpoint `/telemetry/photo`, o sistema utiliza AWS Rekognition para:

1. Analisar a imagem quanto à qualidade e conteúdo
2. Comparar com imagens anteriores
3. Retornar resultados de reconhecimento

Exemplo de resposta (pode ser alterado futuramente):
```json
{
  "success": true,
  "photoId": "8a7b6c5d4e3f2g1h",
  "recognition": {
    "matched": true,
    "confidence": 98.7,
    "matchedPhotoIds": ["1a2b3c4d5e6f7g8h"]
  }
}
```

Desenvolvido com ❤️ para a V3.