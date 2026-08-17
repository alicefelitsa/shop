import request from '@/api/request';

//获取联系方式配置
export async function GetContactSetting() {
    const res = await request.get('/GetContactSetting');
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存联系方式配置
export async function SaveContactSetting(data) {
    const res = await request.post('/SaveContactSetting', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
