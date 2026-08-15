<template>
  <div class="products-page">
    <!-- Page Header -->
    <section class="page-header">
      <div class="page-header-bg" style="background-image: url('https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=1920&h=500&fit=crop')"></div>
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
          <div class="filter-categories">
            <button
              v-for="cat in categories"
              :key="cat.id"
              class="filter-btn"
              :class="{ active: activeCategory === cat.id }"
              @click="setCategory(cat.id)"
            >
              {{ cat.name }}
            </button>
          </div>
          <div class="filter-search">
            <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8"/><path d="M21 21l-4.35-4.35"/>
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
        <div class="results-info">
          <p class="results-count">
            Showing <strong>{{ filteredProducts.length }}</strong> of <strong>{{ allProducts.length }}</strong> products
          </p>
          <div class="sort-control">
            <label>Sort by:</label>
            <select v-model="sortBy" class="sort-select">
              <option value="default">Default</option>
              <option value="name">Name A–Z</option>
              <option value="price-asc">Price: Low to High</option>
              <option value="price-desc">Price: High to Low</option>
              <option value="rating">Highest Rated</option>
            </select>
          </div>
        </div>

        <!-- Products Grid -->
        <div v-if="filteredProducts.length" class="products-grid">
          <ProductCard
            v-for="product in filteredProducts"
            :key="product.id"
            :product="product"
          />
        </div>

        <!-- Empty State -->
        <div v-else class="empty-state">
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
import { products, categories } from '../data/products'

export default {
  name: 'ProductsPage',
  components: { ProductCard },
  data() {
    return {
      allProducts: products,
      categories,
      activeCategory: 'all',
      searchQuery: '',
      sortBy: 'default'
    }
  },
  computed: {
    filteredProducts() {
      let result = [...this.allProducts]

      // Category filter
      if (this.activeCategory !== 'all') {
        result = result.filter(p => p.category === this.activeCategory)
      }

      // Search filter
      if (this.searchQuery.trim()) {
        const q = this.searchQuery.toLowerCase().trim()
        result = result.filter(p =>
          p.name.toLowerCase().includes(q) ||
          p.description.toLowerCase().includes(q) ||
          (p.specs && p.specs.purity && p.specs.purity.toLowerCase().includes(q))
        )
      }

      // Sort
      switch (this.sortBy) {
        case 'name':
          result.sort((a, b) => a.name.localeCompare(b.name))
          break
        case 'price-asc':
          result.sort((a, b) => a.price - b.price)
          break
        case 'price-desc':
          result.sort((a, b) => b.price - a.price)
          break
        case 'rating':
          result.sort((a, b) => b.rating - a.rating)
          break
      }

      return result
    }
  },
  created() {
    // Read category from query params
    const cat = this.$route.query.category
    if (cat && categories.find(c => c.id === cat)) {
      this.activeCategory = cat
    }
  },
  methods: {
    setCategory(id) {
      this.activeCategory = id
      // Update URL query without navigation
      if (id === 'all') {
        this.$router.replace({ query: {} }).catch(() => {})
      } else {
        this.$router.replace({ query: { category: id } }).catch(() => {})
      }
    },
    resetFilters() {
      this.activeCategory = 'all'
      this.searchQuery = ''
      this.sortBy = 'default'
      this.$router.replace({ query: {} }).catch(() => {})
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
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  margin-bottom: 24px;
  flex-wrap: wrap;
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

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .page-header { height: 260px; }
  .page-title { font-size: 2.2rem; }

  .products-grid {
    grid-template-columns: repeat(3, 1fr);
  }

  .filter-search {
    min-width: 200px;
  }
}

@media (max-width: 767px) {
  .page-header { height: 160px; padding-top: var(--header-height); }
  .page-title { font-size: 1.5rem; margin-bottom: 6px; }
  .page-breadcrumb { font-size: 0.82rem; }
  .page-header-content { padding-top: 0; }

  .products-section {
    padding-top: 20px;
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
}
</style>
