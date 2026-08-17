import request from '@/api/request';

//获取系统配置
export async function GetConfigSetting() {
    const res = await request.get('/GetConfigSetting');
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存系统配置
export async function SaveConfigSetting(data) {
    const res = await request.post('/SaveConfigSetting', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
