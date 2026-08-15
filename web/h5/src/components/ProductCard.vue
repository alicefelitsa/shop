<template>
  <router-link :to="`/products/${product.id}`" class="product-card">
    <div class="card-image">
      <img
        v-lazy
        :data-src="product.thumbnail"
        :alt="product.name"
        src="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 400 400'%3E%3Crect fill='%23edf2f7' width='400' height='400'/%3E%3C/svg%3E"
        loading="lazy"
      />
      <div class="card-badges">
        <span v-if="product.onSale" class="badge badge-sale">Sale</span>
        <span v-if="product.bestSeller" class="badge badge-best">Best Seller</span>
      </div>
      <div class="card-overlay">
        <span class="view-btn">View Details</span>
      </div>
    </div>
    <div class="card-body">
      <span class="card-category">{{ categoryName }}</span>
      <h3 class="card-title">{{ product.name }}</h3>
      <div class="card-meta">
        <div class="card-rating">
          <span class="stars">
            <span v-for="i in 5" :key="i" class="star" :class="{ filled: i <= Math.round(product.rating) }">★</span>
          </span>
          <span class="rating-value">{{ Number(product.rating).toFixed(1) }}</span>
        </div>
        <span v-if="product.specs && product.specs.purity" class="card-purity">{{ product.specs.purity }} Purity</span>
      </div>
      <div class="card-price">
        <span class="price-range">${{ product.price.toFixed(2) }} – ${{ product.priceMax.toFixed(2) }}</span>
      </div>
    </div>
  </router-link>
</template>

<script>
import { categories } from '../data/products'

export default {
  name: 'ProductCard',
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  computed: {
    categoryName() {
      const cat = categories.find(c => c.id === this.product.category)
      return cat ? cat.name : this.product.category
    }
  }
}
</script>

<style scoped>
.product-card {
  display: flex;
  flex-direction: column;
  border-radius: var(--radius-md);
  overflow: hidden;
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  transition: all 0.35s ease;
  text-decoration: none;
  color: inherit;
}

.product-card:hover {
  transform: translateY(-6px);
  box-shadow: var(--shadow-lg);
  border-color: transparent;
}

/* Image */
.card-image {
  position: relative;
  aspect-ratio: 1 / 1;
  overflow: hidden;
  background: var(--bg-gray);
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.product-card:hover .card-image img {
  transform: scale(1.08);
}

/* Badges */
.card-badges {
  position: absolute;
  top: 12px;
  left: 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  z-index: 2;
}

.badge {
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 0.72rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.badge-sale {
  background: var(--danger);
  color: #fff;
}

.badge-best {
  background: var(--accent);
  color: var(--primary-dark);
}

/* Overlay */
.card-overlay {
  position: absolute;
  inset: 0;
  background: rgba(26, 54, 93, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.product-card:hover .card-overlay {
  opacity: 1;
}

.view-btn {
  padding: 10px 24px;
  background: #fff;
  color: var(--primary);
  border-radius: var(--radius-sm);
  font-weight: 600;
  font-size: 0.875rem;
  transform: translateY(10px);
  transition: all 0.3s ease;
}

.product-card:hover .view-btn {
  transform: translateY(0);
}

/* Card Body */
.card-body {
  padding: 18px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card-category {
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  color: var(--accent-dark);
  letter-spacing: 0.8px;
}

.card-title {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  margin-top: auto;
}

.card-rating {
  display: flex;
  align-items: center;
  gap: 4px;
}

.stars {
  display: flex;
  gap: 1px;
}

.star {
  color: #e2e8f0;
  font-size: 0.8rem;
}

.star.filled {
  color: var(--accent);
}

.rating-value {
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--text-primary);
}

.card-purity {
  font-size: 0.72rem;
  font-weight: 600;
  color: var(--success);
  background: rgba(56, 161, 105, 0.08);
  padding: 2px 8px;
  border-radius: 12px;
}

/* Price */
.card-price {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-top: 8px;
  border-top: 1px solid var(--border-light);
}

.price-range {
  font-size: 0.95rem;
  font-weight: 700;
  color: var(--primary);
}

@media (max-width: 767px) {
  .card-image {
    aspect-ratio: 4 / 3;
  }

  .card-badges {
    top: 8px;
    left: 8px;
    gap: 4px;
  }

  .badge {
    padding: 3px 8px;
    font-size: 0.65rem;
  }

  .card-body {
    padding: 10px;
    gap: 5px;
  }

  .card-category {
    font-size: 0.68rem;
    letter-spacing: 0.5px;
  }

  .card-title {
    font-size: 0.82rem;
    line-height: 1.3;
  }

  .card-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }

  .star {
    font-size: 0.72rem;
  }

  .rating-value {
    font-size: 0.7rem;
  }

  .card-purity {
    font-size: 0.65rem;
    padding: 1px 6px;
  }

  .card-price {
    padding-top: 6px;
  }

  .price-range {
    font-size: 0.82rem;
  }
}
</style>
