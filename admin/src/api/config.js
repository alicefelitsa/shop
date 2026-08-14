import request from '@/api/request';

//获取系统配置
export async function getConfigs() {
    const res = await request.get('/getConfigs');
    if (res.data.code === 0) {
        return res.data.data[0];
    }
    return Promise.reject(new Error(res.data.message));
}