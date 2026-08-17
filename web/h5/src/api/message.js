import request from '@/api/request';

//提交留言信息
export async function AddMessage(data) {
    const res = await request.post('/AddMessage', data);
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}
