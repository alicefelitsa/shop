import request from '@/api/request';

//获取联系方式配置
export async function GetContactInfo() {
    const res = await request.get('/GetContactInfo');
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}
