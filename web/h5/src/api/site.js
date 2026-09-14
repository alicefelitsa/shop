import request from '@/api/request';

//获取站点公开配置（访问方式等）
export async function GetSiteConfig() {
    const res = await request.get('/GetSiteConfig');
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}
