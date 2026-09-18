import request from '@/api/request';

//获取客户购物意向
export async function GetCartIntent(params) {
    const res = await request.get('/GetCartIntent', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除客户购物意向
export async function DelCartIntent(data) {
    const res = await request.get('/DelCartIntent?ids=' + data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
