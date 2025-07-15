# Flights Database API

A simple Node.js API that serves aircraft mapping data, designed to be deployed on Vercel.

## Project Structure

```
flights-database/
├── api/
│   └── aircraft.js          # Main API handler
├── data/
│   └── aircraft_mappings.json # JSON database (for reference)
├── package.json             # Node.js dependencies
├── index.js                 # Local development server
├── vercel.json             # Vercel deployment config
├── index.html               # API documentation
├── test.html               # Browser test interface
├── deploy.sh               # Deployment script
└── README.md               # This file
```

## Database Schema

The API contains aircraft mappings with the following structure:

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

1. Install Node.js (version 16 or later)
2. Install dependencies:
```bash
npm install
```

3. Run the development server:
```bash
npm start
```

The API will be available at `http://localhost:3000`

## Deployment to Vercel

1. Install Vercel CLI:
```bash
npm i -g vercel
```

2. Deploy to Vercel:
```bash
vercel --prod
```

Or use the deployment script:
```bash
./deploy.sh
```

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
- ✅ Built-in aircraft mappings database
- ✅ Vercel deployment ready
- ✅ Error handling and validation
- ✅ Clean API response format
- ✅ Local development server
- ✅ Static file serving

## Available Aircraft

- **IC5302**: ATR 72 (AT7)
- **AI101**: Boeing 787-8 (B788)
- **6E1234**: Airbus A320 (A320)
- **SG123**: Boeing 737-800 (B738)

## Technology Stack

- **Runtime**: Node.js
- **Framework**: Express.js
- **Deployment**: Vercel
- **Database**: In-memory JSON (can be extended to use external database) 