import request from '@/api/request';

//获取打招呼列表
export async function getGreetList(ucode) {
    const res = await request.get('/getGreetList?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存打招呼
export async function saveGreet(ucode, data) {
    const res = await request.post('/saveGreet?ucode=' + ucode, data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//设置排序
export async function saveGreetSort(data) {
    const res = await request.post('/saveGreetSort', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//设置状态
export async function saveGreetStatus(data) {
    const res = await request.post('/saveGreetStatus', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除打招呼
export async function delGreet(data) {
    const res = await request.post('/delGreet', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}