<template>
  <div class="products-page">
    <!-- Page Header -->
    <section class="page-header">
      <div class="page-header-bg"></div>
      <div class="page-header-overlay"></div>
      <div class="container page-header-content">
        <h1 class="page-title">Our Products</h1>
        <p class="page-breadcrumb">
          <router-link to="/">Home</router-link>
          <span>/</span>
          <span>Products</span>
        </p>
      </div>
    </section>

    <!-- Products Section -->
    <section class="section products-section">
      <div class="container">
        <!-- Filters Bar -->
        <div class="filters-bar">
          <div class="filter-main">
            <div class="filter-categories">
              <button
                  class="filter-btn"
                  :class="{ active: activeCategory === 'all' }"
                  @click="setCategory('all')"
              >
                All Products
              </button>
              <!-- 当前选中的分类以可清除的胶囊展示，避免为它单独占一屏 -->
              <button
                  v-if="activeCategory !== 'all'"
                  class="filter-btn active"
                  @click="setCategory('all')"
              >
                {{ activeCategory }}
                <span class="clear-x">✕</span>
              </button>
              <button
                  class="filter-btn cat-toggle"
                  :class="{ open: showCategories }"
                  @click="showCategories = !showCategories"
              >
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M22 3H2l8 9.46V19l4 2v-8.54L22 3z"/>
                </svg>
                Categories
                <span class="cat-total">{{ categories.length }}</span>
                <span class="cat-caret" :class="{ open: showCategories }">▾</span>
              </button>
            </div>
            <!-- 分类面板：默认收起，展开后在限高卡片内滚动浏览，不再铺满首屏 -->
            <transition name="cat-panel">
              <div v-if="showCategories" class="category-panel">
                <div class="category-panel-scroll">
                  <button
                      v-for="cat in categories"
                      :key="cat.id"
                      class="cat-chip"
                      :class="{ active: activeCategory === cat.name }"
                      @click="pickCategory(cat.name)"
                  >
                    {{ cat.name }}
                    <span class="cat-count">{{ categoryCounts[cat.name] || 0 }}</span>
                  </button>
                </div>
              </div>
            </transition>
          </div>
          <div class="filter-search">
            <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                 stroke-width="2">
              <circle cx="11" cy="11" r="8"/>
              <path d="M21 21l-4.35-4.35"/>
            </svg>
            <input
                v-model="searchQuery"
                type="text"
                placeholder="Search products..."
                class="search-input"
            />
            <button v-if="searchQuery" class="search-clear" @click="searchQuery = ''">✕</button>
          </div>
        </div>

        <!-- Results Info -->
        <div v-if="!loading" class="results-info">
          <p class="results-count">
            Showing <strong>{{ pageStart }}–{{ pageEnd }}</strong> of <strong>{{ filteredProducts.length }}</strong>
            products
          </p>
          <div class="sort-control">
            <label>Sort by:</label>
            <select v-model="sortBy" class="sort-select">
              <option value="default">Default</option>
              <option value="name">Name A–Z</option>
              <!-- 价格排序选项已暂时隐藏
              <option value="price-asc">Price: Low to High</option>
              <option value="price-desc">Price: High to Low</option>
              -->
              <option value="rating">Highest Rated</option>
            </select>
          </div>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="loading-state">
          <div class="loading-spinner"></div>
          <p class="loading-text">Loading...</p>
        </div>

        <!-- Products Grid -->
        <div v-else-if="filteredProducts.length" class="products-grid">
          <ProductCard
              v-for="product in pagedProducts"
              :key="product.id"
              :product="product"
          />
        </div>

        <!-- Pagination -->
        <div v-if="!loading && totalPages > 1" class="pagination">
          <button class="page-btn" :disabled="currentPage === 1" @click="goToPage(currentPage - 1)">‹ Prev</button>
          <button
              v-for="page in totalPages"
              :key="page"
              class="page-num"
              :class="{ active: page === currentPage }"
              @click="goToPage(page)"
          >{{ page }}
          </button>
          <button class="page-btn" :disabled="currentPage === totalPages" @click="goToPage(currentPage + 1)">Next ›
          </button>
        </div>

        <!-- Empty State（不能用 v-else，中间隔了分页栏会配对错元素） -->
        <div v-if="!loading && !filteredProducts.length" class="empty-state">
          <div class="empty-icon">🔍</div>
          <h3>No products found</h3>
          <p>Try adjusting your search or filter criteria.</p>
          <button class="btn btn-primary" @click="resetFilters">Reset Filters</button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import ProductCard from '../components/ProductCard.vue'
import {GetProduct} from '@/api/product'

export default {
  name: 'ProductsPage',
  components: {ProductCard},
  data() {
    return {
      allProducts: [],
      categories: [],
      activeCategory: 'all',
      // 分类面板展开状态（默认收起，防止 68 个分类铺满首屏）
      showCategories: false,
      searchQuery: '',
      sortBy: 'default',
      // 接口加载中，避免加载完成前误显示 No products found
      loading: true,
      currentPage: 1,
      pageSize: 16
    }
  },
  computed: {
    // 各分类下的产品数量（分类面板胶囊右侧角标）
    categoryCounts() {
      const map = {}
      for (const p of this.allProducts) {
        map[p.category] = (map[p.category] || 0) + 1
      }
      return map
    },
    filteredProducts() {
      let result = [...this.allProducts]

      // Category filter（product.category 存储的是分类名称）
      if (this.activeCategory !== 'all') {
        result = result.filter(p => p.category === this.activeCategory)
      }

      // Search filter
      if (this.searchQuery.trim()) {
        const q = this.searchQuery.toLowerCase().trim()
        result = result.filter(p =>
            (p.name || '').toLowerCase().includes(q) ||
            (p.Introduction || '').toLowerCase().includes(q) ||
            (p.purity || '').toLowerCase().includes(q)
        )
      }

      // Sort
      switch (this.sortBy) {
        case 'name':
          result.sort((a, b) => (a.name || '').localeCompare(b.name || ''))
          break
        case 'price-asc':
          result.sort((a, b) => this.minPrice(a) - this.minPrice(b))
          break
        case 'price-desc':
          result.sort((a, b) => this.minPrice(b) - this.minPrice(a))
          break
        case 'rating':
          result.sort((a, b) => Number(b.level) - Number(a.level))
          break
      }

      return result
    },
    // 总页数
    totalPages() {
      return Math.ceil(this.filteredProducts.length / this.pageSize)
    },
    // 当前页的产品列表
    pagedProducts() {
      const start = (this.currentPage - 1) * this.pageSize
      return this.filteredProducts.slice(start, start + this.pageSize)
    },
    // 当前页显示区间（用于 Showing x–y）
    pageStart() {
      if (!this.filteredProducts.length) return 0
      return (this.currentPage - 1) * this.pageSize + 1
    },
    pageEnd() {
      return Math.min(this.currentPage * this.pageSize, this.filteredProducts.length)
    }
  },
  watch: {
    // 筛选、搜索、排序变化时回到第一页
    activeCategory() {
      this.currentPage = 1
    },
    searchQuery() {
      this.currentPage = 1
    },
    sortBy() {
      this.currentPage = 1
    }
  },
  created() {
    const cat = this.$route.query.category
    // 从后端接口加载产品与分类数据
    GetProduct().then(res => {
      this.allProducts = res.productData || []
      this.categories = res.categoryData || []
      // Read category from query params
      if (cat && this.categories.find(c => c.name === cat)) {
        this.activeCategory = cat
      }
    }).catch(() => {
    }).finally(() => {
      this.loading = false
    })
  },
  methods: {
    goToPage(page) {
      if (page < 1 || page > this.totalPages) return
      this.currentPage = page
      window.scrollTo({top: 0, behavior: 'smooth'})
    },
    // 从价格字符串（如 "$41.00 – $145.00"）中提取最低价用于排序
    minPrice(product) {
      const match = String(product.price || '').match(/\d+(?:\.\d+)?/)
      return match ? parseFloat(match[0]) : 0
    },
    setCategory(name) {
      this.activeCategory = name
      // Update URL query without navigation
      if (name === 'all') {
        this.$router.replace({query: {}}).catch(() => {
        })
      } else {
        this.$router.replace({query: {category: name}}).catch(() => {
        })
      }
    },
    // 从分类面板选中分类后收起面板
    pickCategory(name) {
      this.setCategory(name)
      this.showCategories = false
    },
    resetFilters() {
      this.activeCategory = 'all'
      this.searchQuery = ''
      this.sortBy = 'default'
      this.$router.replace({query: {}}).catch(() => {
      })
    }
  }
}
</script>

<style scoped>
/* ===== Page Header ===== */
.page-header {
  position: relative;
  height: 320px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  padding-top: var(--header-height);
}

.page-header-bg {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center;
  background-image: url('~@/assets/p1.jpg');
}

.page-header-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(15, 36, 64, 0.88) 0%, rgba(26, 54, 93, 0.75) 100%);
}

.page-header-content {
  position: relative;
  z-index: 2;
  text-align: center;
}

.page-title {
  font-size: 2.8rem;
  font-weight: 800;
  color: #fff;
  margin-bottom: 12px;
}

.page-breadcrumb {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  font-size: 0.95rem;
  color: rgba(255, 255, 255, 0.7);
}

.page-breadcrumb a {
  color: var(--accent);
}

.page-breadcrumb a:hover {
  color: var(--accent-light);
}

/* ===== Filters Bar ===== */
.filters-bar {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 20px;
  margin-bottom: 24px;
  flex-wrap: wrap;
}

.filter-main {
  flex: 1;
  min-width: 0;
}

.filter-categories {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.filter-btn {
  padding: 8px 20px;
  border-radius: 30px;
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-secondary);
  background: var(--bg-light);
  border: 1px solid var(--border-color);
  transition: all 0.25s ease;
  white-space: nowrap;
}

.filter-btn:hover {
  border-color: var(--primary);
  color: var(--primary);
}

.filter-btn.active {
  background: var(--primary);
  color: #fff;
  border-color: var(--primary);
}

.filter-btn .clear-x {
  margin-left: 6px;
  font-size: 0.72rem;
  opacity: 0.85;
}

/* Categories 下拉按钮 */
.cat-toggle {
  display: inline-flex;
  align-items: center;
  gap: 7px;
}

.cat-toggle.open {
  border-color: var(--primary);
  color: var(--primary);
}

.cat-total {
  font-size: 0.72rem;
  font-weight: 700;
  padding: 1px 7px;
  border-radius: 10px;
  background: rgba(15, 36, 64, 0.07);
  color: var(--text-secondary);
}

.cat-caret {
  font-size: 0.68rem;
  line-height: 1;
  transition: transform 0.2s ease;
}

.cat-caret.open {
  transform: rotate(180deg);
}

/* 分类面板：限高卡片内滚动，替代原先铺满首屏的分类墙 */
.category-panel {
  position: relative;
  margin-top: 12px;
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  box-shadow: 0 10px 30px rgba(15, 36, 64, 0.08);
  overflow: hidden;
}

.category-panel-scroll {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  padding: 16px 16px 22px;
  max-height: 264px;
  overflow-y: auto;
}

.category-panel-scroll::-webkit-scrollbar {
  width: 6px;
}

.category-panel-scroll::-webkit-scrollbar-thumb {
  background: var(--border-color);
  border-radius: 3px;
}

.category-panel-scroll::-webkit-scrollbar-thumb:hover {
  background: var(--text-light);
}

/* 底部渐隐提示可继续滚动 */
.category-panel::after {
  content: '';
  position: absolute;
  left: 0;
  right: 6px;
  bottom: 0;
  height: 26px;
  pointer-events: none;
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0), var(--bg-white));
}

.cat-chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 0.82rem;
  font-weight: 500;
  color: var(--text-secondary);
  background: var(--bg-light);
  border: 1px solid var(--border-color);
  transition: all 0.2s ease;
  white-space: nowrap;
}

.cat-chip:hover {
  border-color: var(--primary);
  color: var(--primary);
}

.cat-chip.active {
  background: var(--primary);
  border-color: var(--primary);
  color: #fff;
}

.cat-count {
  font-size: 0.68rem;
  font-weight: 700;
  padding: 1px 6px;
  border-radius: 9px;
  background: rgba(15, 36, 64, 0.06);
  color: var(--text-light);
}

.cat-chip.active .cat-count {
  background: rgba(255, 255, 255, 0.22);
  color: #fff;
}

/* 面板展开/收起动画（Vue2 transition 类名） */
.cat-panel-enter-active,
.cat-panel-leave-active {
  transition: opacity 0.22s ease, transform 0.22s ease;
}

.cat-panel-enter,
.cat-panel-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}

.filter-search {
  position: relative;
  min-width: 260px;
}

.search-icon {
  position: absolute;
  left: 14px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--text-light);
  pointer-events: none;
}

.search-input {
  width: 100%;
  padding: 10px 40px 10px 42px;
  border: 1px solid var(--border-color);
  border-radius: 30px;
  font-size: 0.9rem;
  background: var(--bg-white);
  transition: all 0.25s ease;
  outline: none;
}

.search-input:focus {
  border-color: var(--primary);
  box-shadow: 0 0 0 3px rgba(26, 54, 93, 0.1);
}

.search-clear {
  position: absolute;
  right: 14px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--text-light);
  font-size: 0.8rem;
  padding: 4px;
  line-height: 1;
}

.search-clear:hover {
  color: var(--danger);
}

/* ===== Results Info ===== */
.results-info {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
  flex-wrap: wrap;
  gap: 12px;
}

.results-count {
  font-size: 0.9rem;
  color: var(--text-secondary);
}

.sort-control {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.875rem;
  color: var(--text-secondary);
}

.sort-select {
  padding: 6px 12px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  background: var(--bg-white);
  font-size: 0.875rem;
  outline: none;
  cursor: pointer;
}

.sort-select:focus {
  border-color: var(--primary);
}

/* ===== Products Grid ===== */
.products-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

/* ===== Empty State ===== */
.empty-state {
  text-align: center;
  padding: 80px 20px;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 16px;
}

.empty-state h3 {
  font-size: 1.3rem;
  color: var(--primary);
  margin-bottom: 8px;
}

.empty-state p {
  color: var(--text-secondary);
  margin-bottom: 24px;
}

/* ===== Loading ===== */
.loading-state {
  text-align: center;
  padding: 80px 20px;
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

/* ===== Pagination ===== */
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-top: 40px;
  flex-wrap: wrap;
}

.page-num,
.page-btn {
  min-width: 40px;
  height: 40px;
  padding: 0 14px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  background: #fff;
  color: var(--text-secondary);
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.25s ease;
}

.page-num:hover,
.page-btn:hover:not(:disabled) {
  border-color: var(--primary);
  color: var(--primary);
}

.page-num.active {
  background: var(--primary);
  border-color: var(--primary);
  color: #fff;
}

.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

/* 压缩区块上下内边距，减少横幅与筛选栏、分页与页脚之间的空白 */
.products-section {
  padding-top: 32px;
  padding-bottom: 40px;
}

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .page-header {
    height: 260px;
  }

  .page-title {
    font-size: 2.2rem;
  }

  .products-grid {
    grid-template-columns: repeat(3, 1fr);
  }

  .filter-search {
    min-width: 200px;
  }
}

@media (max-width: 767px) {
  .page-header {
    height: 160px;
    padding-top: var(--header-height);
  }

  .page-title {
    font-size: 1.5rem;
    margin-bottom: 6px;
  }

  .page-breadcrumb {
    font-size: 0.82rem;
  }

  .page-header-content {
    padding-top: 0;
  }

  .products-section {
    padding-top: 20px;
    padding-bottom: 24px;
  }

  .filters-bar {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
    margin-bottom: 16px;
  }

  .filter-categories {
    flex-wrap: wrap;
    gap: 10px;
  }

  .category-panel-scroll {
    max-height: 220px;
    padding: 12px 12px 18px;
  }

  .cat-chip {
    padding: 5px 11px;
    font-size: 0.76rem;
  }

  .filter-btn {
    padding: 6px 12px;
    font-size: 0.78rem;
  }

  .filter-search {
    min-width: auto;
  }

  .search-input {
    padding: 9px 36px 9px 38px;
    font-size: 0.85rem;
  }

  .results-info {
    margin-bottom: 14px;
    gap: 8px;
  }

  .results-count {
    font-size: 0.82rem;
  }

  .sort-control {
    font-size: 0.82rem;
  }

  .sort-select {
    padding: 5px 10px;
    font-size: 0.82rem;
  }

  .products-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 10px;
  }

  .empty-state {
    padding: 48px 16px;
  }

  .empty-icon {
    font-size: 3rem;
  }

  .pagination {
    gap: 6px;
    margin-top: 24px;
  }

  .page-num,
  .page-btn {
    min-width: 34px;
    height: 34px;
    padding: 0 10px;
    font-size: 0.82rem;
  }
}
</style>
