import axios from 'axios';

// 创建 axios 实例
const apiClient = axios.create({
    baseURL: 'http://localhost:31943/api',
    timeout: 10000, // 请求超时时间 (毫秒)
    headers: {
        'Content-Type': 'application/json',
    },
});

// 创建并导出 api 对象
const api = {
    /**
     * POST 请求
     * @param {string} url - 请求的 URL
     * @param {Object} data - 请求体数据
     * @returns {Promise} - 返回一个 Promise，resolve 时携带响应数据
     */
    post(url, data = {}) {
        console.log(3)
        return apiClient.post(url, data)
            .then(response => response.data)
            .catch(error => {
                throw error;
            });
    },
};

export default api;