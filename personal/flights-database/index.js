const express = require('express');
const cors = require('cors');
const aircraftHandler = require('./api/aircraft');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(cors());
app.use(express.json());

// Serve static files
app.use(express.static('.'));

// API routes
app.all('/api/aircraft', aircraftHandler);

// Root route - serve the documentation
app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

// Privacy Policy route
app.get('/privacy-policy', (req, res) => {
  res.sendFile(__dirname + '/privacy-policy.html');
});

// Start server
app.listen(PORT, () => {
  console.log('🚀 Flights Database API starting on http://localhost:' + PORT);
  console.log('📖 API endpoints:');
  console.log('   GET  /api/aircraft - Get all aircraft mappings');
  console.log('   POST /api/aircraft - Get specific aircraft mappings');
  console.log('   GET  / - API documentation');
});

module.exports = app; 