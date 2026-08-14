import request from '@/api/request';

//获取访客列表
export async function getClientList(ucode) {
    const res = await request.get('/getClientList?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//获取选中访客的聊天记录
export async function getChat(ucode, uuid) {
    const res = await request.get('/getChat?ucode=' + ucode + '&uuid=' + uuid);
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//保存聊天信息
export async function saveChat(newMsg) {
    const res = await request.post('/saveChat', newMsg);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//获取访客资料
export async function getClientInfo(ucode, uuid) {
    const res = await request.get('/getClientInfo?ucode=' + ucode + '&uuid=' + uuid);
    if (res.data.code === 0) {
        return res.data.data[0];
    }
    return Promise.reject(new Error(res.data.message));
}

//获取访客资料
export async function delClient(ucode, ids) {
    const res = await request.get('/delClient?ucode=' + ucode + '&ids=' + ids);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//获取访问统计
export async function getStatistic(ucode, ids) {
    const res = await request.get('/getStatistic?ucode=' + ucode);
    if (res.data.code === 0) {
        return res.data.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//选中的访客有消息接收更新数据库未读为0
export async function updateIsReadNum(uuid) {
    const res = await request.get('/updateIsReadNum?uuid=' + uuid);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}
