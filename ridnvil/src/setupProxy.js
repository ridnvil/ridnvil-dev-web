const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/api', // This is the endpoint to proxy
        createProxyMiddleware({
            target: 'http://localhost:3001', // Backend server URL
            changeOrigin: true, // Changes the origin of the host header to the target URL
            pathRewrite: { '^/api': '' }, // Optional: Remove `/api` prefix when forwarding the request
        })
    );
};