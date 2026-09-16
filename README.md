# Gastro Atlas

> An interactive atlas exploring the history, movement, and cultural journey of food around the world.

Gastro Atlas is a full-stack web application that visualizes how ingredients have moved across places and through history.

Instead of presenting food history as a static article, Gastro Atlas turns historical events into an interactive experience: users can explore an ingredient on a world map, move through its timeline, and inspect the historical events and sources associated with it.

The project is currently being developed around **tomato** as the first complete dataset, with the architecture designed to support additional ingredients in the future.

---

## Features

### 🌍 Interactive World Map

Explore historical events geographically through an interactive world map.

Events are associated with locations and can be displayed as points on the map, allowing users to follow how an ingredient's history spans different regions.

### ⏳ Historical Timeline

View an ingredient's history chronologically and explore significant events across different periods.

The timeline is generated from the same historical-event data used by the map, keeping the two views synchronized.

### 📖 Historical Event Information

Each event contains structured historical information, including:

* Title
* Description
* Time period
* Historical entity
* Location
* Sources
* Confidence level

This allows historical information to remain structured rather than being stored as a collection of unstructured articles.

### 🍅 Ingredient-Based Exploration

Ingredients serve as the main entry point into the atlas.

For example, the current tomato dataset contains historical events describing its origins, movement, introduction to different regions, and historical development.

### 🤖 AI Historian — Planned

Gastro Atlas is designed with a future AI historian in mind.

The planned system will use retrieval-augmented generation (RAG) to answer questions about food history using the project's structured historical data and source material.

The current development focus is on building a reliable historical dataset and visualization layer before introducing the AI interface.

---

## Architecture

Gastro Atlas consists of three main parts:

```text
                    ┌─────────────────────┐
                    │      Frontend       │
                    │ React + TypeScript  │
                    │                     │
                    │ Map / Timeline / UI │
                    └──────────┬──────────┘
                               │
                             HTTP
                               │
                    ┌──────────▼──────────┐
                    │       Backend       │
                    │       Go + Gin      │
                    │                     │
                    │     REST API        │
                    └──────────┬──────────┘
                               │
                           PostgreSQL
                               │
                    ┌──────────▼──────────┐
                    │      Historical     │
                    │        Data          │
                    │                     │
                    │ Ingredients / Events│
                    └─────────────────────┘
```

Historical data is initially prepared through a separate Python extraction pipeline, which can use local LLMs to transform source material into the project's structured event format.

---

## Tech Stack

### Frontend

* React
* TypeScript
* Vite
* Tailwind CSS
* shadcn/ui
* D3
* TopoJSON
* MapLibre
* Framer Motion
* React Router
* Axios

### Backend

* Go
* Gin
* PostgreSQL
* pgx
* sqlc

### Data & AI

* Python
* Pydantic
* Ollama
* Qwen3
* pgvector *(planned)*
* OpenAI API / RAG *(planned)*

### Development & Deployment

* Git / GitHub
* Docker
* Docker Compose

---

## Data Model

The core data structure is centered around historical events.

An ingredient contains a collection of historical events:

```text
Ingredient
├── name
├── description
└── events[]
    ├── title
    ├── description
    ├── time_period
    ├── entity
    ├── location
    ├── sources[]
    └── confidence
```

For example:

```json
{
  "title": "Tomatoes arrive in Europe",
  "description": "Tomatoes were introduced to Europe following...",
  "time_period": "16th century",
  "entity": "Spain",
  "location": "Spain",
  "sources": [
    "uvm-tomato-history"
  ],
  "confidence": "high"
}
```

The frontend derives map and timeline representations from these historical events rather than maintaining a separate list of places.

This keeps the historical event as the primary source of truth while allowing the same data to power multiple visualizations.

---

## Historical Data Pipeline

Historical information is prepared separately from the main web application.

The current extraction workflow is:

```text
Historical Sources
       │
       ▼
   Text Extraction
       │
       ▼
     Chunking
       │
       ▼
   Local LLM
   (Qwen3 8B)
       │
       ▼
 Structured JSON
       │
       ▼
 Historical Events
       │
       ▼
    PostgreSQL
```

The extraction pipeline uses Pydantic models to constrain the generated output to the project's schema.

This approach makes it possible to experiment with extracting historical events from larger collections of source material without manually creating every event.

---

## Project Structure

The repository is organized around the frontend, backend, and data pipeline.

```text
gastro-atlas/
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── data/
│   │   └── ...
│   └── ...
│
├── backend/
│   ├── app/
│   ├── api/
│   ├── core/
│   ├── db/
│   └── ...
│
├── extraction/
│   ├── ...
│   └── ...
│
├── docker-compose.yml
└── README.md
```

The exact structure may evolve as the project develops.

---

## Getting Started

### Prerequisites

Make sure the following are installed:

* Node.js
* npm
* Go
* PostgreSQL
* Docker *(recommended)*

For the historical-data extraction pipeline:

* Python 3.11+
* Ollama
* Qwen3

### Clone the repository

```bash
git clone <repository-url>
cd gastro-atlas
```

### Start the database

Using Docker Compose:

```bash
docker compose up -d
```

### Start the backend

```bash
cd backend
go run .
```

### Start the frontend

```bash
cd frontend
npm install
npm run dev
```

The development server will provide the local frontend URL in the terminal.

---

## Testing

The backend uses Go's standard testing framework.

Integration tests use **Testcontainers** to run against an isolated PostgreSQL environment.

Run the backend tests with:

```bash
go test ./...
```

---

## Roadmap

### Current

* [x] Interactive world map
* [x] Historical timeline
* [x] Ingredient-based data model
* [x] Historical event schema
* [x] PostgreSQL backend
* [x] Go REST API
* [x] Structured tomato dataset
* [x] Python historical-data extraction pipeline
* [x] Local LLM extraction experiments

### Planned

* [ ] Expand the dataset beyond tomato
* [ ] Improve historical source coverage
* [ ] Source-aware historical search
* [ ] Vector search with pgvector
* [ ] RAG-based AI historian
* [ ] Natural-language historical queries
* [ ] Connect AI responses to supporting historical sources

---

## Why Gastro Atlas?

Food history is often scattered across books, academic papers, museum archives, and historical records.

Gastro Atlas explores a different way of presenting that information: **treating food history as something that can be explored spatially and chronologically.**

The long-term goal is to combine structured historical data, interactive visualization, and retrieval-based AI into a single interface for exploring the history of food.

---

## Status

🚧 **In development**

The project is currently focused on the core visualization and historical-data infrastructure, with **tomato** serving as the initial dataset.

The AI historian is planned as a later layer on top of this foundation.
