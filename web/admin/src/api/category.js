import request from '@/api/request';

//获取分类列表
export async function GetCategory(params) {
    const res = await request.get('/GetCategory', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//添加分类
export async function AddCategory(data) {
    const res = await request.post('/AddCategory', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//修改分类
export async function SaveCategory(data) {
    const res = await request.post('/SaveCategory', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除分类
export async function DelCategory(data) {
    const res = await request.get('/DelCategory?ids=' + data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
