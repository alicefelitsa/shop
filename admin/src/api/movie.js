import request from '@/api/request';

//获取影片
export async function GetMovie(params) {
    const res = await request.get('/GetMovie', {params});
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}

//添加影片
export async function AddMovie(data) {
    const res = await request.post('/AddMovie', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//修改影片
export async function SaveMovie(data) {
    const res = await request.post('/SaveMovie', data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//删除影片
export async function DelMovie(data) {
    const res = await request.get('/DelMovie?ids=' + data);
    if (res.data.code === 0) {
        return res.data.message;
    }
    return Promise.reject(new Error(res.data.message));
}

//获取播放影片
export async function GetPlayVideo(data) {
    const res = await request.post('/GetPlayVideo', data);
    if (res.data.code === 0) {
        return res.data.episode;
    }
    return Promise.reject(new Error(res.data.message));
}