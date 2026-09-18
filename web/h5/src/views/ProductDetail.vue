<template>
  <div class="detail-page" v-if="product">
    <!-- Product Detail -->
    <section class="section detail-section">
      <div class="container">
        <p class="page-breadcrumb">
          <router-link to="/">Home</router-link>
          <span>/</span>
          <router-link to="/products">Products</router-link>
          <span>/</span>
          <span>{{ product.name }}</span>
        </p>
        <div class="detail-grid">
          <!-- Image Gallery -->
          <div class="detail-gallery">
            <div class="gallery-main">
              <img :src="product.album" :alt="product.name"/>
            </div>
          </div>

          <!-- Product Info -->
          <div class="detail-info">
            <span class="detail-category">{{ product.category }}</span>
            <h1 class="detail-title">{{ product.name }}</h1>

            <!-- Rating -->
            <div class="detail-rating">
              <span class="stars">
                <span v-for="i in 5" :key="i" class="star" :class="{ filled: i <= Math.round(product.level) }">★</span>
              </span>
              <span class="rating-text">{{ Number(product.level).toFixed(1) }}</span>
            </div>

            <div class="detail-price">
              <span class="price-range">{{ product.price }}</span>
            </div>

            <!-- 货号选择（主信息区）：点击选中某个货号后才可加购 -->
            <div v-if="itemCodes.length" class="detail-items">
              <div class="item-chips">
                <span
                    v-for="(code, i) in itemCodes"
                    :key="i"
                    class="item-chip item-chip-selectable"
                    :class="{ selected: selectedItem === code }"
                    @click="selectItem(code)"
                >{{ code }}</span>
              </div>
            </div>

            <!-- Description -->
            <p class="detail-desc">{{ product.Introduction }}</p>

            <!-- Actions：需先选中货号才能加购；纯度与加购同一行 -->
            <div class="detail-actions">
              <div class="action-row">
                <div v-if="product.purity" class="purity-highlight">
                  <span class="purity-label">Purity</span>
                  <span class="purity-value">{{ product.purity }}</span>
                </div>
                <div class="qty-stepper">
                  <button type="button" class="qty-btn" :disabled="qty <= 1" @click="qty--">−</button>
                  <span class="qty-value">{{ qty }}</span>
                  <button type="button" class="qty-btn" @click="qty++">+</button>
                </div>
                <button
                    type="button"
                    class="btn btn-accent btn-lg add-cart-btn"
                    :disabled="itemCodes.length > 0 && !selectedItem"
                    @click="addToCart"
                >
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="9" cy="21" r="1"/>
                    <circle cx="20" cy="21" r="1"/>
                    <path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/>
                  </svg>
                  {{ addedFlash ? 'Added to Cart ✓' : 'Add to Cart' }}
                </button>
              </div>
              <p v-if="itemCodes.length && !selectedItem" class="select-hint">
                Please select an item number above before adding to cart.
              </p>
              <router-link to="/contact" class="btn btn-primary btn-lg quote-btn">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9 2-2 2z"/>
                  <path d="M22 6l-10 7L2 6"/>
                </svg>
                Get Discounted Quote
              </router-link>
            </div>

            <!-- Guarantees -->
            <div class="detail-guarantees">
              <div class="guarantee-item">
                <span class="guarantee-icon">🚚</span>
                <span>Fast international shipping (6-18 days)</span>
              </div>
              <div class="guarantee-item">
                <span class="guarantee-icon">🔬</span>
                <span>Lab-tested with COA included</span>
              </div>
              <div class="guarantee-item">
                <span class="guarantee-icon">📦</span>
                <span>Free reship guarantee</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Product Detailed Information -->
        <div class="detail-info-section">
          <h2 class="detail-info-title">Product Information</h2>

          <!-- Overview -->
          <div class="info-block">
            <h3 class="info-block-title">Product Overview</h3>
            <template v-if="detailParagraphs.length">
              <p v-for="(para, idx) in detailParagraphs" :key="idx" class="info-text">{{ para }}</p>
            </template>
            <p v-else class="info-text">{{ product.Introduction }}</p>
          </div>

          <!-- Item Numbers（货号汇总） -->
          <div class="info-block" v-if="itemCodes.length">
            <h3 class="info-block-title">Item Numbers</h3>
            <div class="item-chips item-chips-lg">
              <span v-for="(code, i) in itemCodes" :key="i" class="item-chip">{{ code }}</span>
            </div>
          </div>

          <!-- Specifications -->
          <div class="info-block" v-if="specRows.length || product.purity">
            <h3 class="info-block-title">Specifications (per kit)</h3>
            <div class="spec-scroll">
              <table class="spec-grid">
                <thead>
                <tr>
                  <th>Specification</th>
                  <th>No.</th>
                  <th>Price</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(row, idx) in specRows" :key="idx">
                  <td>{{ row.spec }}</td>
                  <td>{{ row.item || '—' }}</td>
                  <td class="spec-price">{{ row.price }}</td>
                </tr>
                <tr v-if="product.purity">
                  <td>Purity</td>
                  <td colspan="2">{{ product.purity }}</td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Closing divider -->
          <div class="info-divider" aria-hidden="true">
            <span class="info-divider-line"></span>
            <span class="info-divider-gem">⬡</span>
            <span class="info-divider-line"></span>
          </div>
        </div>

        <!-- Related Products -->
        <div class="related-section" v-if="relatedProducts.length">
          <h2 class="section-title">Related Products</h2>
          <div class="related-grid">
            <ProductCard
                v-for="p in relatedProducts"
                :key="p.id"
                :product="p"
            />
          </div>
        </div>
      </div>
    </section>

    <!-- 加购成功通知卡片 -->
    <transition name="toast-slide">
      <div v-if="cartToast" class="cart-toast">
        <span class="cart-toast-icon">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
            <path d="M20 6L9 17l-5-5"/>
          </svg>
        </span>
        <div class="cart-toast-text">
          <strong>Added to Cart</strong>
          <span>Please check your cart for the quote.</span>
        </div>
        <router-link to="/cart" class="cart-toast-link" @click.native="cartToast = false">View Cart</router-link>
      </div>
    </transition>
  </div>

  <!-- Loading -->
  <div v-else-if="loading" class="section loading-state">
    <div class="container" style="text-align:center;">
      <div class="loading-spinner"></div>
      <p class="loading-text">Loading...</p>
    </div>
  </div>

  <!-- Not Found -->
  <div v-else class="section not-found">
    <div class="container" style="text-align:center;">
      <div style="font-size:4rem;margin-bottom:16px;">😕</div>
      <h2>Product Not Found</h2>
      <p style="color:var(--text-secondary);margin:12px 0 24px;">The product you're looking for doesn't exist.</p>
      <router-link to="/products" class="btn btn-primary">Back to Products</router-link>
    </div>
  </div>
</template>

<script>
import ProductCard from '../components/ProductCard.vue'
import {GetProductDetail} from '@/api/product'
import {cart} from '@/utils/cart'

export default {
  name: 'ProductDetail',
  components: {ProductCard},
  data() {
    return {
      product: null,
      relatedProducts: [],
      // 接口加载中，避免加载完成前误显示 Product Not Found
      loading: true,
      // 当前选中的货号（未选中时禁止加购）
      selectedItem: '',
      // 加购数量
      qty: 1,
      // 加购成功提示闪现
      addedFlash: false,
      addedTimer: null,
      // 加购成功顶部提示条
      cartToast: false,
      toastTimer: null
    }
  },
  computed: {
    // 货号：item_no 逗号分隔字符串拆成数组，去空白与空值
    itemCodes() {
      if (!this.product) return []
      return (this.product.item_no || '').split(',').map(s => s.trim()).filter(Boolean)
    },
    // 选中货号对应的规格行（用于加购时记录规格与单价）
    selectedSpecRow() {
      if (!this.selectedItem) return null
      return this.specRows.find(r => r.item === this.selectedItem) || null
    },
    // 规格表格行：优先读独立字段 product.specs（后端 JSON 数组），缺失/解析失败时回退解析 details
    specRows() {
      if (!this.product) return []
      // 1) 独立 specs 字段（[{spec,item,price}...]）
      if (this.product.specs) {
        try {
          const arr = JSON.parse(this.product.specs)
          if (Array.isArray(arr) && arr.length) {
            return arr.map(r => ({
              spec: (r && r.spec) || '',
              item: (r && r.item) || '',
              price: (r && r.price) || ''
            }))
          }
        } catch (e) {
          // JSON 解析失败，回退到 details 解析
        }
      }
      // 2) 回退：兼容旧库或后台新录产品——从 details 首段解析 规格 | 货号 | 价格
      if (!this.product.details) return []
      const lines = this.product.details.split('\n').map(s => s.trim())
      const rows = []
      let started = false
      for (const line of lines) {
        if (!line) {
          if (started) break
          continue
        }
        if (!line.includes('|')) {
          // 规格行之前的表头（如 "Specifications (per kit):"）跳过；开始后再遇非规格行即结束
          if (started) break
          continue
        }
        started = true
        const parts = line.split('|').map(s => s.trim())
        if (parts.length >= 3) {
          rows.push({
            spec: parts[0],
            item: parts[1].replace(/^Item No\.?\s*:?\s*/i, ''),
            price: parts.slice(2).join(' | ')
          })
        } else {
          // 个别历史数据缺少货号列，仅两列
          rows.push({spec: parts[0], item: '', price: parts[1] || ''})
        }
      }
      return rows
    },
    // 接口返回的详情描述按段落拆分（规格块改由表格呈现，此处仅保留介绍段落）
    detailParagraphs() {
      if (!this.product || !this.product.details) return []
      const lines = this.product.details.split('\n').map(s => s.trim())
      const firstSpec = lines.findIndex(l => l.includes('|'))
      if (firstSpec === -1) return lines.filter(l => l && !/^specifications/i.test(l))
      const end = lines.findIndex((l, i) => i >= firstSpec && l === '')
      if (end === -1) return []
      return lines.slice(end).filter(Boolean)
    }
  },
  watch: {
    // 相关产品切换时组件被复用，created 不会重新执行，需监听路由参数重新请求
    '$route.params.id'() {
      this.fetchDetail()
    }
  },
  created() {
    // 从后端接口加载产品详情与相关产品
    this.fetchDetail()
  },
  beforeDestroy() {
    if (this.addedTimer) clearTimeout(this.addedTimer)
    if (this.toastTimer) clearTimeout(this.toastTimer)
  },
  methods: {
    // 点击货号标签：选中/取消选中
    selectItem(code) {
      this.selectedItem = this.selectedItem === code ? '' : code
    },
    // 加入购物车：有货号时必须先选中货号；无货号商品可直接加购
    addToCart() {
      if (!this.product) return
      if (this.itemCodes.length && !this.selectedItem) return
      cart.add({
        productId: this.product.id,
        name: this.product.name,
        image: this.product.album,
        itemNo: this.selectedItem,
        spec: this.selectedSpecRow ? this.selectedSpecRow.spec : '',
        price: this.selectedSpecRow ? this.selectedSpecRow.price : (this.product.price || '')
      }, this.qty)
      this.addedFlash = true
      if (this.addedTimer) clearTimeout(this.addedTimer)
      this.addedTimer = setTimeout(() => {
        this.addedFlash = false
      }, 1500)
      // 顶部提示条：已添加到购物车，请到购物车查看
      this.cartToast = true
      if (this.toastTimer) clearTimeout(this.toastTimer)
      this.toastTimer = setTimeout(() => {
        this.cartToast = false
      }, 5000)
    },
    fetchDetail() {
      const id = parseInt(this.$route.params.id)
      this.loading = true
      this.product = null
      this.relatedProducts = []
      // 切换产品时重置选货与数量
      this.selectedItem = ''
      this.qty = 1
      this.addedFlash = false
      GetProductDetail({id}).then(res => {
        const list = res.productData || []
        this.product = list.length ? list[0] : null
        if (this.product) {
          this.relatedProducts = (res.relatedProducts || []).filter(p => p.id !== this.product.id)
        }
      }).catch(() => {
      }).finally(() => {
        this.loading = false
      })
    }
  }
}
</script>

<style scoped>
/* 加购成功通知卡片：右下角悬浮 */
.cart-toast {
  position: fixed;
  right: 24px;
  bottom: 24px;
  z-index: 1200;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 18px;
  background: #fff;
  border: 1px solid var(--border-light);
  border-radius: var(--radius-md);
  box-shadow: 0 12px 32px rgba(15, 36, 64, 0.18);
}

.cart-toast-icon {
  flex: 0 0 34px;
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: linear-gradient(135deg, #34d399, #059669);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 10px rgba(5, 150, 105, 0.35);
}

.cart-toast-text {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.cart-toast-text strong {
  font-size: 0.95rem;
  font-weight: 700;
  color: var(--primary);
}

.cart-toast-text span {
  font-size: 0.78rem;
  color: var(--text-secondary);
}

.cart-toast-link {
  margin-left: 6px;
  padding-left: 14px;
  border-left: 1px solid var(--border-color);
  font-size: 0.82rem;
  font-weight: 700;
  color: var(--accent-dark);
  white-space: nowrap;
}

.cart-toast-link:hover {
  color: var(--primary);
}

.toast-slide-enter-active,
.toast-slide-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.toast-slide-enter,
.toast-slide-leave-to {
  opacity: 0;
  transform: translateY(16px);
}

@media (max-width: 767px) {
  .cart-toast {
    left: 16px;
    right: 16px;
    bottom: 16px;
  }
}

/* ===== Breadcrumb ===== */
.page-breadcrumb {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 0.9rem;
  color: var(--text-light);
  margin-bottom: 20px;
}

.page-breadcrumb a {
  color: var(--text-secondary);
}

.page-breadcrumb a:hover {
  color: var(--accent-dark);
}

/* ===== Detail Grid ===== */
.detail-section {
  padding-top: calc(var(--header-height) + 8px);
  padding-bottom: 32px;
}

.detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 56px;
  margin-bottom: 40px;
}

/* Gallery */
.gallery-main {
  position: relative;
  border-radius: var(--radius-lg);
  overflow: hidden;
  background: var(--bg-gray);
  aspect-ratio: 1 / 1;
}

.gallery-main img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sale-badge,
.best-badge {
  position: absolute;
  top: 16px;
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 0.78rem;
  font-weight: 700;
  text-transform: uppercase;
}

.sale-badge {
  left: 16px;
  background: var(--danger);
  color: #fff;
}

.best-badge {
  left: 16px;
  background: var(--accent);
  color: var(--primary-dark);
}

.sale-badge + .best-badge {
  left: auto;
  right: 16px;
}

.gallery-thumbs {
  display: flex;
  gap: 12px;
  margin-top: 16px;
}

.thumb-btn {
  width: 80px;
  height: 80px;
  border-radius: var(--radius-sm);
  overflow: hidden;
  border: 2px solid var(--border-color);
  transition: all 0.25s ease;
  cursor: pointer;
  padding: 0;
}

.thumb-btn.active,
.thumb-btn:hover {
  border-color: var(--primary);
}

.thumb-btn img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Detail Info */
.detail-category {
  display: inline-block;
  font-size: 0.78rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1.5px;
  color: var(--accent-dark);
  margin-bottom: 8px;
}

.detail-title {
  font-size: 1.8rem;
  font-weight: 800;
  color: var(--primary);
  line-height: 1.25;
  margin-bottom: 16px;
}

.detail-rating {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 20px;
}

.stars {
  display: flex;
  gap: 2px;
}

.star {
  color: #e2e8f0;
  font-size: 1rem;
}

.star.filled {
  color: var(--accent);
}

.rating-text {
  font-size: 0.88rem;
  color: var(--text-light);
}

.detail-price {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-light);
}

.price-range {
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--primary);
}

/* Item No. 货号标签 */
.detail-items {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-light);
}

.items-label {
  font-size: 0.78rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.8px;
  color: var(--text-secondary);
  white-space: nowrap;
  padding-top: 4px;
}

.item-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.item-chip {
  display: inline-block;
  padding: 3px 10px;
  border-radius: 12px;
  font-size: 0.78rem;
  font-weight: 600;
  color: var(--primary);
  background: rgba(15, 36, 64, 0.06);
  border: 1px solid var(--border-color);
}

.item-chips-lg .item-chip {
  font-size: 0.85rem;
  padding: 5px 13px;
}

/* 可选中货号标签 */
.item-chip-selectable {
  cursor: pointer;
  transition: all 0.2s ease;
  user-select: none;
}

.item-chip-selectable:hover {
  border-color: var(--primary);
  color: var(--primary);
}

.item-chip.selected {
  background: var(--primary);
  border-color: var(--primary);
  color: #fff;
}

.detail-desc {
  font-size: 0.95rem;
  color: var(--text-secondary);
  line-height: 1.75;
  margin-bottom: 20px;
  /* 简介最多显示 3 行，超出省略 */
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.purity-highlight {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: rgba(56, 161, 105, 0.08);
  padding: 8px 16px;
  border-radius: var(--radius-sm);
  margin-bottom: 24px;
}

.purity-label {
  font-size: 0.82rem;
  font-weight: 600;
  color: var(--success);
  text-transform: uppercase;
}

.purity-value {
  font-size: 1rem;
  font-weight: 800;
  color: var(--success);
}

/* Actions */
.detail-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 5px;
  margin-bottom: 12px;
}

.action-row {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
}

/* 纯度徽章与加购同一行：三者统一固定高度(48px)，比下方 Quote 按钮更高以突出加购行 */
.action-row .purity-highlight {
  margin-bottom: 0;
  height: 48px;
  padding: 0 16px;
  box-sizing: border-box;
  display: flex;
  align-items: center;
}

/* 数量选择器 */
.qty-stepper {
  display: flex;
  align-items: center;
  height: 48px;
  box-sizing: border-box;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  overflow: hidden;
  background: #fff;
}

.qty-btn {
  width: 42px;
  height: 100%;
  border: none;
  background: var(--bg-light);
  font-size: 1.1rem;
  color: var(--primary);
  cursor: pointer;
  transition: background 0.2s ease;
}

.qty-btn:hover:not(:disabled) {
  background: var(--border-light);
}

.qty-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.qty-value {
  min-width: 44px;
  text-align: center;
  font-weight: 700;
  color: var(--text-primary);
}

.add-cart-btn {
  flex: 1 1 160px;
  min-width: 0;
  height: 48px;
  padding: 0 16px;
  box-sizing: border-box;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.add-cart-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.quote-btn {
  width: 100%;
  height: 40px;
  padding: 0 20px;
  box-sizing: border-box;
  font-size: 0.95rem;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

/* 无提示文字时（加购行直接接 Quote 按钮）拉开两者距离，避免过于紧密 */
.action-row + .quote-btn {
  margin-top: 10px;
}

.select-hint {
  font-size: 0.8rem;
  color: var(--text-light);
  margin: 0;
}

/* Guarantees */
.detail-guarantees {
  border-top: 1px solid var(--border-light);
  padding-top: 24px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.guarantee-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 0.88rem;
  color: var(--text-secondary);
}

.guarantee-icon {
  font-size: 1.1rem;
}

/* ===== Detailed Information ===== */
.detail-info-section {
  margin-bottom: 48px;
}

.detail-info-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 2px solid var(--accent);
  display: inline-block;
}

.info-block {
  margin-bottom: 28px;
}

.info-block-title {
  font-size: 1.05rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 12px;
}

.info-text {
  font-size: 0.92rem;
  color: var(--text-secondary);
  line-height: 1.75;
  margin-bottom: 12px;
}

.info-text:last-child {
  margin-bottom: 0;
}

.info-divider {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-top: 36px;
}

.info-divider-line {
  flex: 1;
  height: 1px;
}

.info-divider-line:first-child {
  background: linear-gradient(to right, transparent, rgba(201, 168, 76, 0.55));
}

.info-divider-line:last-child {
  background: linear-gradient(to left, transparent, rgba(201, 168, 76, 0.55));
}

.info-divider-gem {
  font-size: 0.95rem;
  line-height: 1;
  color: var(--accent);
  transform: translateY(-1px);
}

.spec-scroll {
  overflow-x: auto;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  background: #fff;
}

.spec-grid {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.88rem;
}

.spec-grid th,
.spec-grid td {
  padding: 12px 18px;
  text-align: left;
  border-bottom: 1px solid var(--border-light);
  white-space: nowrap;
  color: var(--text-secondary);
}

.spec-grid thead th {
  background: var(--bg-light);
  color: var(--text-primary);
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.6px;
  text-transform: uppercase;
}

.spec-grid tbody td:first-child {
  color: var(--text-primary);
  font-weight: 600;
}

.spec-grid tbody tr:nth-child(even) {
  background: var(--bg-light);
}

.spec-grid tbody tr:last-child td {
  border-bottom: none;
}

.spec-grid .spec-price {
  font-weight: 700;
  color: var(--primary);
}

/* ===== Related ===== */
.related-section {
  margin-bottom: 16px;
}

.related-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  margin-top: 32px;
}

/* ===== Loading ===== */
.loading-state {
  padding-top: calc(var(--header-height) + 80px);
  padding-bottom: 120px;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  margin: 0 auto 16px;
  border: 3px solid var(--border-color);
  border-top-color: var(--primary);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

.loading-text {
  color: var(--text-secondary);
  font-size: 0.95rem;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .detail-grid {
    gap: 40px;
  }

  .related-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 767px) {
  .page-breadcrumb {
    font-size: 0.78rem;
    gap: 6px;
    margin-bottom: 14px;
  }

  .detail-section {
    padding-top: calc(var(--header-height) + 4px);
    padding-bottom: 16px;
  }

  .detail-grid {
    grid-template-columns: 1fr;
    gap: 16px;
    margin-bottom: 24px;
  }

  .gallery-main {
    aspect-ratio: 4 / 3;
  }

  .gallery-thumbs {
    margin-top: 10px;
  }

  .thumb-btn {
    width: 56px;
    height: 56px;
  }

  .detail-category {
    font-size: 0.72rem;
  }

  .detail-title {
    font-size: 1.15rem;
    margin-bottom: 10px;
  }

  .detail-rating {
    margin-bottom: 12px;
  }

  .rating-text {
    font-size: 0.82rem;
  }

  .detail-price {
    margin-bottom: 14px;
    padding-bottom: 14px;
  }

  .price-range {
    font-size: 1.2rem;
  }

  .detail-desc {
    font-size: 0.85rem;
    margin-bottom: 14px;
    line-height: 1.65;
  }

  .purity-highlight {
    padding: 6px 12px;
    margin-bottom: 16px;
  }

  .purity-label {
    font-size: 0.75rem;
  }

  .purity-value {
    font-size: 0.9rem;
  }

  .detail-actions {
    flex-direction: column;
    gap: 10px;
    margin-bottom: 24px;
  }

  .add-cart-btn {
    min-width: auto;
    padding: 0 16px;
    font-size: 0.95rem;
  }

  .detail-guarantees {
    padding-top: 16px;
    gap: 8px;
  }

  .guarantee-item {
    font-size: 0.82rem;
  }

  .detail-info-section {
    margin-bottom: 24px;
  }

  .detail-info-title {
    font-size: 1.1rem;
    margin-bottom: 16px;
    padding-bottom: 8px;
  }

  .info-block {
    margin-bottom: 20px;
  }

  .info-block-title {
    font-size: 0.95rem;
    margin-bottom: 8px;
  }

  .info-text {
    font-size: 0.85rem;
    margin-bottom: 8px;
  }

  .info-divider {
    gap: 10px;
    margin-top: 26px;
  }

  .spec-grid {
    font-size: 0.82rem;
  }

  .spec-grid th,
  .spec-grid td {
    padding: 9px 12px;
  }

  .related-section {
    margin-bottom: 16px;
  }

  .related-section .section-title {
    font-size: 1.2rem;
  }

  .related-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 10px;
    margin-top: 16px;
  }
}
</style>
