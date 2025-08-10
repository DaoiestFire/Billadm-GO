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
     * GET 请求
     * @param {string} url - 请求的 URL
     * @param {Object} params - URL 查询参数
     * @returns {Promise} - 返回一个 Promise，resolve 时携带响应数据
     */
    get(url, params = {}) {
        return apiClient.get(url, {params})
            .then(response => response.data) // 直接返回响应体中的 data
            .catch(error => {
                throw error;
            });
    },

    /**
     * POST 请求
     * @param {string} url - 请求的 URL
     * @param {Object} data - 请求体数据
     * @returns {Promise} - 返回一个 Promise，resolve 时携带响应数据
     */
    post(url, data = {}) {
        return apiClient.post(url, data)
            .then(response => response.data)
            .catch(error => {
                throw error;
            });
    },

    /**
     * PUT 请求
     * @param {string} url - 请求的 URL
     * @param {Object} data - 请求体数据
     * @returns {Promise} - 返回一个 Promise，resolve 时携带响应数据
     */
    put(url, data = {}) {
        return apiClient.put(url, data)
            .then(response => response.data)
            .catch(error => {
                throw error;
            });
    },

    /**
     * DELETE 请求
     * @param {string} url - 请求的 URL
     * @param {Object} params - URL 查询参数 (可选)
     * @returns {Promise} - 返回一个 Promise，resolve 时携带响应数据
     */
    delete(url, params = {}) {
        return apiClient.delete(url, {params})
            .then(response => response.data)
            .catch(error => {
                throw error;
            });
    }
};

export default api;