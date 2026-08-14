import request from '@/api/request';

export async function getQrcodeUrl(ucode) {
    const res = await request.get('/getQrcodeUrl?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}