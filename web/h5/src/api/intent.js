import request from '@/api/request';

//提交购物意向（平台无支付，仅记录客户想购买的商品供后台报价）
export async function AddCartIntent(data) {
    const res = await request.post('/AddCartIntent', data);
    if (res.data.code === 0) {
        return res.data;
    }
    return Promise.reject(new Error(res.data.message));
}
