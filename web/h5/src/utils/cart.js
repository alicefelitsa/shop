import Vue from 'vue'

// 购物车本地存储键
const STORAGE_KEY = 'utitanu_cart_v1'

function loadItems() {
  try {
    const raw = window.localStorage.getItem(STORAGE_KEY)
    const arr = raw ? JSON.parse(raw) : []
    return Array.isArray(arr) ? arr : []
  } catch (e) {
    return []
  }
}

function persist() {
  try {
    window.localStorage.setItem(STORAGE_KEY, JSON.stringify(cartState.items))
  } catch (e) {
    // 隐私模式等写入失败时忽略，购物车仅本次会话有效
  }
}

// 响应式购物车状态：items 元素形如
// { productId, name, image, itemNo, spec, price, qty }
export const cartState = Vue.observable({
  items: loadItems()
})

// 购物车操作集合（按 productId + itemNo 作为唯一键）
export const cart = {
  // 总件数（用于头部角标）
  count() {
    return cartState.items.reduce((n, it) => n + (it.qty || 0), 0)
  },
  // 加入购物车：同产品同货号则累加数量并刷新快照字段
  add(entry, qty) {
    const found = cartState.items.find(it => it.productId === entry.productId && it.itemNo === entry.itemNo)
    if (found) {
      found.qty = (found.qty || 0) + qty
      found.name = entry.name
      found.image = entry.image
      found.spec = entry.spec
      found.price = entry.price
    } else {
      cartState.items.push(Object.assign({}, entry, {qty}))
    }
    persist()
  },
  // 修改数量；<=0 视为移除
  setQty(productId, itemNo, qty) {
    if (qty <= 0) {
      cart.remove(productId, itemNo)
      return
    }
    const it = cartState.items.find(x => x.productId === productId && x.itemNo === itemNo)
    if (!it) return
    it.qty = qty
    persist()
  },
  remove(productId, itemNo) {
    const idx = cartState.items.findIndex(x => x.productId === productId && x.itemNo === itemNo)
    if (idx > -1) cartState.items.splice(idx, 1)
    persist()
  },
  clear() {
    cartState.items.splice(0, cartState.items.length)
    persist()
  }
}
