import request from '@/api/request';

//获取打招呼列表
export async function getAnswerList(ucode) {
    const res = await request.get('/getAnswerList?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存打招呼
export async function saveAnswer(ucode, data) {
    const res = await request.post('/saveAnswer?ucode=' + ucode, data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//设置排序
export async function saveAnswerSort(data) {
    const res = await request.post('/saveAnswerSort', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//设置状态
export async function saveAnswerStatus(data) {
    const res = await request.post('/saveAnswerStatus', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除打招呼
export async function delAnswer(data) {
    const res = await request.post('/delAnswer', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}