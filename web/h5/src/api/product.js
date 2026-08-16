import request from '@/api/request';

//获取产品与分类数据
export async function GetProduct(params) {
    const res = await request.get('/GetProduct', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}
