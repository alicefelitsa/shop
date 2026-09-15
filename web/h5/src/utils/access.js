import Vue from 'vue'
import {GetSiteConfig} from '@/api/site'

// 展示状态：ready 表示访问校验已完成（完成前整页不渲染，避免头尾闪烁）；
// blank 表示访问受限且展示类型为 blank，App.vue 据此整页渲染空白，URL 保持当前路径
export const accessState = Vue.observable({ready: false, blank: false})

// 移动端 UA 判定
export function isMobileUA() {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini|Mobile/i.test(navigator.userAgent)
}

// 站点配置缓存，避免每次路由跳转重复请求
let configCache = null

// 获取后台配置的访问方式与展示类型，失败时回落 all/jump 放行
export async function getAccessConfig() {
  if (configCache) return configCache
  try {
    const data = await GetSiteConfig()
    configCache = {
      mode: (data && data.access_mode) || 'all',
      displayType: (data && data.display_type) || 'jump'
    }
  } catch (e) {
    configCache = {mode: 'all', displayType: 'jump'}
  }
  return configCache
}

// 校验当前客户端是否允许访问，返回 {allowed, mode, displayType}
export async function checkAccess() {
  const {mode, displayType} = await getAccessConfig()
  const mobile = isMobileUA()
  const allowed = mode === 'all' || (mode === 'h5' && mobile) || (mode === 'pc' && !mobile)
  return {allowed, mode, displayType}
}
