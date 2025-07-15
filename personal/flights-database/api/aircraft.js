const fs = require('fs');
const path = require('path');

// Function to load aircraft database from JSON file
function loadAircraftDatabase() {
  try {
    const dataPath = path.join(process.cwd(), 'data', 'aircraft_mappings.json');
    const data = fs.readFileSync(dataPath, 'utf8');
    return JSON.parse(data);
  } catch (error) {
    console.error('Error loading aircraft database:', error);
    // Fallback to original database if file read fails
    return {
      "aircraft_mappings": {
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
        },
        "SG123": {
          "aircraftType": "B738",
          "displayName": "Boeing 737-800",
          "manufacturer": "Boeing"
        }
      }
    };
  }
}

module.exports = (req, res) => {
  // Set CORS headers
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');

  // Handle preflight requests
  if (req.method === 'OPTIONS') {
    res.status(200).end();
    return;
  }

  // Only allow GET and POST methods
  if (req.method !== 'GET' && req.method !== 'POST') {
    res.status(405).json({ error: 'Method not allowed' });
    return;
  }

  try {
    // Load the aircraft database
    const aircraftDatabase = loadAircraftDatabase();

    if (req.method === 'GET') {
      // For GET requests, return all aircraft mappings
      res.status(200).json({
        results: aircraftDatabase.aircraft_mappings
      });
      return;
    }

    // Handle POST request
    if (req.method === 'POST') {
      const { flightCodes } = req.body;

      if (!flightCodes || !Array.isArray(flightCodes)) {
        res.status(400).json({ error: 'Invalid request body. flightCodes array is required.' });
        return;
      }

      // Filter results based on requested flight codes
      const results = {};
      const missingFlights = [];
      
      flightCodes.forEach(flightCode => {
        if (aircraftDatabase.aircraft_mappings[flightCode]) {
          results[flightCode] = aircraftDatabase.aircraft_mappings[flightCode];
        } else {
          missingFlights.push(flightCode);
        }
      });

      const response = {
        results: results
      };

      // Add missing flights information if any
      if (missingFlights.length > 0) {
        response.missingFlights = missingFlights.map(flightCode => ({
          flightCode: flightCode,
          message: `Flight code '${flightCode}' not found in the database`
        }));
      }

      res.status(200).json(response);
      return;
    }
  } catch (error) {
    console.error('API Error:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
}; 