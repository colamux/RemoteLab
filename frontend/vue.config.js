module.exports = {
    // devServer: {
    //     port: 8080,
    //     host: "0.0.0.0",
    //     open: true
    // }
    devServer: {
        proxy: {
            '/api': {
                // 代理 /api => 'http://localhost:8000'
                target: 'http://localhost:8000',
                // 允许跨域
                changeOrigin: true,
                ws: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        }
    }
}

