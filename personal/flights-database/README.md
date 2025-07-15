# Flights Database API

A simple Go API that serves aircraft mapping data from a JSON database, designed to be deployed on Vercel.

## Project Structure

```
flights-database/
├── api/
│   └── aircraft.go          # Main API handler
├── data/
│   └── aircraft_mappings.json # JSON database
├── go.mod                   # Go module file
├── vercel.json             # Vercel deployment config
└── README.md               # This file
```

## Database Schema

The JSON database contains aircraft mappings with the following structure:

```json
{
  "aircraft_mappings": {
    "FLIGHT_CODE": {
      "aircraftType": "ICAO_CODE",
      "displayName": "Human Readable Name",
      "manufacturer": "Manufacturer Name"
    }
  }
}
```

## API Endpoints

### GET /api/aircraft

Returns all aircraft mappings in the database.

**Response:**
```json
{
  "results": {
    "IC5302": {
      "aircraftType": "AT7",
      "displayName": "ATR 72",
      "manufacturer": "ATR"
    },
    "AI101": {
      "aircraftType": "B788",
      "displayName": "Boeing 787-8",
      "manufacturer": "Boeing"
    }
  }
}
```

### POST /api/aircraft

Returns aircraft mappings for specific flight codes.

**Request Body:**
```json
{
  "flightCodes": ["IC5302", "AI101", "6E1234"]
}
```

**Response:**
```json
{
  "results": {
    "IC5302": {
      "aircraftType": "AT7",
      "displayName": "ATR 72",
      "manufacturer": "ATR"
    },
    "AI101": {
      "aircraftType": "B788",
      "displayName": "Boeing 787-8",
      "manufacturer": "Boeing"
    },
    "6E1234": {
      "aircraftType": "A320",
      "displayName": "Airbus A320",
      "manufacturer": "Airbus"
    }
  }
}
```

## Local Development

1. Install Go (version 1.21 or later)
2. Clone this repository
3. Run the development server:

```bash
go run api/aircraft.go
```

## Deployment to Vercel

1. Install Vercel CLI:
```bash
npm i -g vercel
```

2. Deploy to Vercel:
```bash
vercel
```

3. Follow the prompts to configure your deployment

## Usage Examples

### Using curl

**Get all aircraft mappings:**
```bash
curl https://your-vercel-app.vercel.app/api/aircraft
```

**Get specific aircraft mappings:**
```bash
curl -X POST https://your-vercel-app.vercel.app/api/aircraft \
  -H "Content-Type: application/json" \
  -d '{"flightCodes": ["IC5302", "AI101"]}'
```

### Using JavaScript/Fetch

```javascript
// Get all aircraft mappings
const response = await fetch('https://your-vercel-app.vercel.app/api/aircraft');
const data = await response.json();
console.log(data.results);

// Get specific aircraft mappings
const response = await fetch('https://your-vercel-app.vercel.app/api/aircraft', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    flightCodes: ['IC5302', 'AI101']
  })
});
const data = await response.json();
console.log(data.results);
```

## Features

- ✅ CORS enabled for cross-origin requests
- ✅ Support for both GET and POST methods
- ✅ JSON database with aircraft mappings
- ✅ Vercel deployment ready
- ✅ Error handling and validation
- ✅ Clean API response format

## Adding New Aircraft

To add new aircraft mappings, simply edit the `data/aircraft_mappings.json` file and add new entries following the existing format. 