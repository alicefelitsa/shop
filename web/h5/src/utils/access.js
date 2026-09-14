import {GetSiteConfig} from '@/api/site'

// 移动端 UA 判定
export function isMobileUA() {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini|Mobile/i.test(navigator.userAgent)
}

// 访问方式缓存，避免每次路由跳转重复请求
let modeCache = ''

// 获取后台配置的访问方式（pc/h5/all），失败时回落 all 放行
export async function getAccessMode() {
  if (modeCache) return modeCache
  try {
    const data = await GetSiteConfig()
    modeCache = (data && data.access_mode) || 'all'
  } catch (e) {
    modeCache = 'all'
  }
  return modeCache
}

// 校验当前客户端是否允许访问，返回 {allowed, mode}
export async function checkAccess() {
  const mode = await getAccessMode()
  const mobile = isMobileUA()
  const allowed = mode === 'all' || (mode === 'h5' && mobile) || (mode === 'pc' && !mobile)
  return {allowed, mode}
}
