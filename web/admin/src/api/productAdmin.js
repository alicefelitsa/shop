import request from '@/api/request';

//获取产品列表
export async function GetProductList(params) {
    const res = await request.get('/GetProductList', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//添加产品
export async function AddProduct(data) {
    const res = await request.post('/AddProduct', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//修改产品
export async function SaveProduct(data) {
    const res = await request.post('/SaveProduct', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除产品
export async function DelProduct(data) {
    const res = await request.get('/DelProduct?ids=' + data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
