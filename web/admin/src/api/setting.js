import request from '@/api/request';

//获取客服配置参数
export async function getUserConfigData() {
    const res = await request.get('/getUserConfigData');
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存客服配置
export async function saveUserConfig(data) {
    const res = await request.post('/saveUserConfig', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}