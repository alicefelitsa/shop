import request from '@/api/request';

//获取客户留言
export async function GetMessage(params) {
    const res = await request.get('/GetMessage', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除客户留言
export async function DelMessage(data) {
    const res = await request.get('/DelMessage?ids=' + data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
