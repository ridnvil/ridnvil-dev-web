const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/api', // This is the endpoint to proxy
        createProxyMiddleware({
            target: 'http://127.0.0.1:3001/api', // Backend server URL
            changeOrigin: true, // Changes the origin of the host header to the target URL
        })
    );
};