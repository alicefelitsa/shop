import request from '@/api/request';

//获取操作日志
export async function getLogs(ucode, ids) {
    const res = await request.get('/getLogs?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}