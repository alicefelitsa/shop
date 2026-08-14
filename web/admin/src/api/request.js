/**
 * axios 实例
 */
import axios from 'axios';
import router from "@/router";
import {Message} from "element-ui";

const service = axios.create({
    baseURL: apiUrl,
    timeout: 30000,
    /*headers: {
        'Content-Type': 'multipart/form-data'
    },*/
});

/**
 * 添加请求拦截器
 */
service.interceptors.request.use(request => {
        request.headers['Authorization'] = localStorage.getItem("token");
        return request;
    }, (error) => {
        return Promise.reject(error);
    }
);

/**
 * 添加响应拦截器
 */
service.interceptors.response.use(response => {
        // 登录过期处理
        if (response.data?.code === 401) {
            localStorage.removeItem('token')
            localStorage.removeItem('avatar')
            localStorage.removeItem("account");
            if (router.currentRoute.path !== '/login') {
                router.push({path: '/login'});
            }
        }
        return response;
    }, (error) => {
        if (!error.response) {
            if (error.code === 'ECONNABORTED' && error.message.includes('timeout')) {
                Message.error('请求超时，请检查网络连接');
            } else if (error.message.includes('ERR_CONNECTION_REFUSED') || error.message.includes('Network Error')) {
                Message.error('无法连接到服务器，请检查网络或稍后重试');
            } else {
                Message.error('网络连接异常，请检查网络设置');
            }
        }
        return Promise.reject(error);
    }
);

export default service;
