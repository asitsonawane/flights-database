#!/bin/bash

echo "🚀 Deploying Flights Database API to Vercel..."

# Check if Vercel CLI is installed
if ! command -v vercel &> /dev/null; then
    echo "❌ Vercel CLI is not installed. Please install it first:"
    echo "npm i -g vercel"
    exit 1
fi

# Check if user is logged in to Vercel
if ! vercel whoami &> /dev/null; then
    echo "🔐 Please log in to Vercel first:"
    echo "vercel login"
    exit 1
fi

# Deploy to Vercel
echo "📦 Deploying..."
vercel --prod

echo "✅ Deployment complete!"
echo "🌐 Your API is now live at: https://your-app-name.vercel.app/api/aircraft"
echo ""
echo "📖 Usage examples:"
echo "GET all aircraft: curl https://your-app-name.vercel.app/api/aircraft"
echo "POST specific aircraft: curl -X POST https://your-app-name.vercel.app/api/aircraft -H 'Content-Type: application/json' -d '{\"flightCodes\": [\"IC5302\", \"AI101\"]}'" 