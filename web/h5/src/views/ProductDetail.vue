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
              <img :src="product.album" :alt="product.name" />
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

            <!-- Price -->
            <div class="detail-price">
              <span class="price-range">{{ product.price }}</span>
            </div>

            <!-- Description -->
            <p class="detail-desc">{{ product.Introduction }}</p>

            <!-- Purity Badge -->
            <div v-if="product.purity" class="purity-highlight">
              <span class="purity-label">Purity</span>
              <span class="purity-value">{{ product.purity }}</span>
            </div>

            <!-- Actions -->
            <div class="detail-actions">
              <router-link to="/contact" class="btn btn-primary btn-lg add-cart-btn">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
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

          <!-- Specifications -->
          <div class="info-block" v-if="product.purity">
            <h3 class="info-block-title">Specifications</h3>
            <div class="specs-table">
              <div class="spec-row">
                <span class="spec-label">Purity</span>
                <span class="spec-value">{{ product.purity }}</span>
              </div>
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
import { GetProduct } from '../api/product'

export default {
  name: 'ProductDetail',
  components: { ProductCard },
  data() {
    return {
      products: [],
      // 接口加载中，避免加载完成前误显示 Product Not Found
      loading: true
    }
  },
  computed: {
    product() {
      const id = parseInt(this.$route.params.id)
      return this.products.find(p => p.id === id) || null
    },
    relatedProducts() {
      if (!this.product) return []
      return this.products
        .filter(p => p.category === this.product.category && p.id !== this.product.id)
        .slice(0, 4)
    },
    // 接口返回的详情描述按段落拆分
    detailParagraphs() {
      if (!this.product || !this.product.details) return []
      return this.product.details.split('\n').map(s => s.trim()).filter(Boolean)
    }
  },
  created() {
    // 从后端接口加载产品数据
    GetProduct().then(res => {
      this.products = res.productData || []
    }).catch(() => {}).finally(() => {
      this.loading = false
    })
  }
}
</script>

<style scoped>
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

.detail-desc {
  font-size: 0.95rem;
  color: var(--text-secondary);
  line-height: 1.75;
  margin-bottom: 20px;
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
  gap: 12px;
  margin-bottom: 32px;
  flex-wrap: wrap;
}

.add-cart-btn {
  flex: 1;
  min-width: 200px;
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

.specs-table {
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  overflow: hidden;
}

.spec-row {
  display: flex;
  border-bottom: 1px solid var(--border-light);
}

.spec-row:last-child {
  border-bottom: none;
}

.spec-label {
  width: 200px;
  flex-shrink: 0;
  padding: 14px 20px;
  font-size: 0.88rem;
  font-weight: 600;
  color: var(--text-primary);
  background: var(--bg-light);
  text-transform: capitalize;
}

.spec-value {
  flex: 1;
  padding: 14px 20px;
  font-size: 0.88rem;
  color: var(--text-secondary);
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
    padding: 12px 24px;
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

  .spec-row {
    flex-direction: column;
  }

  .spec-label {
    width: 100%;
    padding: 8px 14px;
    font-size: 0.82rem;
  }

  .spec-value {
    padding: 8px 14px;
    font-size: 0.82rem;
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
