<template>
  <div class="cart-page">
    <section class="section cart-section">
      <div class="container">
        <p class="page-breadcrumb">
          <router-link to="/">Home</router-link>
          <span>/</span>
          <span>Shopping Cart</span>
        </p>

        <h1 class="cart-title">Shopping Cart</h1>

        <!-- Submitted success state -->
        <div v-if="!items.length && submitted" class="cart-empty">
          <div class="cart-empty-icon">✅</div>
          <h2>Request Submitted</h2>
          <p>Thank you! Our team will review your cart and contact you with a discounted quote shortly.</p>
          <router-link to="/products" class="btn btn-primary">Continue Shopping</router-link>
        </div>

        <!-- Empty state -->
        <div v-else-if="!items.length" class="cart-empty">
          <div class="cart-empty-icon">🛒</div>
          <h2>Your cart is empty</h2>
          <p>Add products from any product page to see them here.</p>
          <router-link to="/products" class="btn btn-primary">Browse Products</router-link>
        </div>

        <!-- Cart list -->
        <template v-else>
          <div class="cart-list">
            <div v-for="it in items" :key="it.productId + '-' + it.itemNo" class="cart-item">
              <router-link :to="`/products/${it.productId}`" class="cart-item-image">
                <img :src="it.image" :alt="it.name"/>
              </router-link>

              <div class="cart-item-info">
                <router-link :to="`/products/${it.productId}`" class="cart-item-name">{{ it.name }}</router-link>
                <div class="cart-item-meta">
                  <span v-if="it.itemNo" class="cart-item-no">{{ it.itemNo }}</span>
                  <span v-if="it.spec" class="cart-item-spec">{{ it.spec }}</span>
                </div>
                <span v-if="it.price" class="cart-item-price">{{ it.price }}</span>
              </div>

              <div class="cart-item-qty">
                <div class="qty-stepper">
                  <button type="button" class="qty-btn" @click="dec(it)">−</button>
                  <span class="qty-value">{{ it.qty }}</span>
                  <button type="button" class="qty-btn" @click="inc(it)">+</button>
                </div>
              </div>

              <button type="button" class="cart-item-remove" title="Remove" @click="remove(it)">✕</button>
            </div>
          </div>

          <!-- Summary -->
          <div class="cart-summary">
            <div class="cart-summary-row">
              <span>Total items</span>
              <strong>{{ totalCount }}</strong>
            </div>
            <p class="cart-summary-note">
              Final pricing is provided in your quote. Submit your cart to get a discounted offer.
            </p>
            <div class="cart-summary-actions">
              <button type="button" class="btn btn-outline" @click="clearAll">Clear Cart</button>
              <router-link to="/products" class="btn btn-outline">Continue Shopping</router-link>
              <button type="button" class="btn btn-accent" @click="openSubmit">Get Discounted Quote</button>
            </div>
          </div>
        </template>
      </div>
    </section>

    <!-- Submit intent modal：提交购物意向（不跳留言页） -->
    <div v-if="submitOpen" class="intent-mask">
      <div class="intent-modal">
        <button type="button" class="intent-close" aria-label="Close" @click="closeSubmit">&times;</button>
        <h3 class="intent-title">Submit for a Quote</h3>
        <p class="intent-sub">We'll reply with a discounted offer. No payment needed.</p>

        <ul class="intent-items">
          <li v-for="it in items" :key="'m-' + it.productId + '-' + it.itemNo">
            <span class="intent-item-name">{{ it.name }}<em v-if="it.itemNo"> · {{ it.itemNo }}</em></span>
            <span class="intent-item-qty">× {{ it.qty }}</span>
          </li>
        </ul>

        <label class="intent-label">Name *</label>
        <input v-model.trim="form.name" class="intent-input" type="text" placeholder="Your name"/>
        <label class="intent-label">Email / WhatsApp *</label>
        <input v-model.trim="form.email" class="intent-input" type="text" placeholder="How can we reach you?"/>
        <label class="intent-label">Remark *</label>
        <textarea v-model.trim="form.remark" class="intent-input intent-textarea" rows="3"
                  placeholder="Anything we should know?"></textarea>

        <div class="intent-actions">
          <p v-if="submitError" class="intent-error">{{ submitError }}</p>
          <div class="intent-actions-row">
            <button type="button" class="btn btn-outline" @click="closeSubmit">Cancel</button>
            <button type="button" class="btn btn-accent" :disabled="submitting" @click="doSubmit">
              {{ submitting ? 'Submitting...' : 'Submit Request' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {cart, cartState} from '@/utils/cart'
import {AddCartIntent} from '@/api/intent'

export default {
  name: 'CartPage',
  data() {
    return {
      submitOpen: false,
      submitting: false,
      submitted: false,
      submitError: '',
      form: {name: '', email: '', remark: ''}
    }
  },
  computed: {
    items() {
      return cartState.items
    },
    totalCount() {
      return cartState.items.reduce((n, it) => n + (it.qty || 0), 0)
    }
  },
  methods: {
    inc(it) {
      cart.setQty(it.productId, it.itemNo, (it.qty || 0) + 1)
    },
    dec(it) {
      cart.setQty(it.productId, it.itemNo, (it.qty || 0) - 1)
    },
    remove(it) {
      cart.remove(it.productId, it.itemNo)
    },
    clearAll() {
      cart.clear()
      this.submitted = false
    },
    // 打开提交意向弹窗
    openSubmit() {
      this.submitError = ''
      this.submitOpen = true
    },
    closeSubmit() {
      if (this.submitting) return
      this.submitOpen = false
    },
    // 提交购物意向到后端（平台无支付，仅记录意向供后台跟进报价）
    async doSubmit() {
      if (this.submitting) return
      if (!this.form.name || !this.form.email || !this.form.remark) {
        this.submitError = 'Please fill in your name, contact and remark.'
        return
      }
      if (!this.items.length) {
        this.submitError = 'Your cart is empty.'
        return
      }
      this.submitting = true
      this.submitError = ''
      try {
        await AddCartIntent({
          name: this.form.name,
          email: this.form.email,
          remark: this.form.remark,
          items: this.items.map(it => ({
            productId: it.productId,
            name: it.name,
            itemNo: it.itemNo || '',
            spec: it.spec || '',
            price: it.price || '',
            qty: it.qty || 0
          }))
        })
        cart.clear()
        this.submitOpen = false
        this.submitted = true
        this.form = {name: '', email: '', remark: ''}
      } catch (e) {
        this.submitError = e.message || 'Submit failed, please try again.'
      } finally {
        this.submitting = false
      }
    }
  }
}
</script>

<style scoped>
.cart-section {
  padding-top: calc(var(--header-height) + 24px);
  padding-bottom: 48px;
}

.page-breadcrumb {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 0.9rem;
  color: var(--text-light);
  margin-bottom: 16px;
}

.page-breadcrumb a {
  color: var(--text-secondary);
}

.page-breadcrumb a:hover {
  color: var(--accent-dark);
}

.cart-title {
  font-size: 1.8rem;
  font-weight: 800;
  color: var(--primary);
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 2px solid var(--accent);
  display: inline-block;
}

/* Empty */
.cart-empty {
  text-align: center;
  padding: 64px 20px;
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
}

.cart-empty-icon {
  font-size: 3.5rem;
  margin-bottom: 12px;
}

.cart-empty h2 {
  font-size: 1.3rem;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.cart-empty p {
  color: var(--text-secondary);
  margin-bottom: 20px;
}

/* List */
.cart-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 24px;
}

.cart-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 16px;
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
}

.cart-item-image {
  flex: 0 0 72px;
  width: 72px;
  height: 72px;
  border-radius: var(--radius-sm);
  overflow: hidden;
  background: var(--bg-gray);
}

.cart-item-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.cart-item-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.cart-item-name {
  font-size: 0.98rem;
  font-weight: 600;
  color: var(--text-primary);
  text-decoration: none;
}

.cart-item-name:hover {
  color: var(--primary);
}

.cart-item-meta {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
}

.cart-item-no {
  display: inline-block;
  padding: 2px 9px;
  border-radius: 10px;
  font-size: 0.76rem;
  font-weight: 600;
  color: var(--primary);
  background: rgba(15, 36, 64, 0.06);
  border: 1px solid var(--border-color);
}

.cart-item-spec {
  font-size: 0.8rem;
  color: var(--text-secondary);
}

.cart-item-price {
  font-size: 0.9rem;
  font-weight: 700;
  color: var(--primary);
}

.cart-item-qty {
  flex: 0 0 auto;
}

/* 数量选择器（与详情页一致） */
.qty-stepper {
  display: flex;
  align-items: center;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  overflow: hidden;
  background: #fff;
}

.qty-btn {
  width: 34px;
  height: 34px;
  border: none;
  background: var(--bg-light);
  font-size: 1rem;
  color: var(--primary);
  cursor: pointer;
  transition: background 0.2s ease;
}

.qty-btn:hover {
  background: var(--border-light);
}

.qty-value {
  min-width: 38px;
  text-align: center;
  font-weight: 700;
  color: var(--text-primary);
}

.cart-item-remove {
  flex: 0 0 auto;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: var(--radius-sm);
  background: transparent;
  color: var(--text-light);
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.cart-item-remove:hover {
  background: rgba(229, 62, 62, 0.08);
  color: var(--danger);
}

/* Summary */
.cart-summary {
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: 20px;
}

.cart-summary-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 0.95rem;
  color: var(--text-secondary);
  margin-bottom: 10px;
}

.cart-summary-row strong {
  color: var(--text-primary);
  font-size: 1.05rem;
}

.cart-summary-note {
  font-size: 0.82rem;
  color: var(--text-light);
  margin-bottom: 16px;
}

.cart-summary-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.btn-outline {
  background: #fff;
  color: var(--text-secondary);
  border: 1px solid var(--border-color);
}

.btn-outline:hover {
  border-color: var(--primary);
  color: var(--primary);
}

/* Submit intent modal */
.intent-mask {
  position: fixed;
  inset: 0;
  background: rgba(15, 36, 64, 0.45);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding: 12vh 20px 20px;
  overflow-y: auto;
  z-index: 200;
}

.intent-modal {
  position: relative;
  width: 100%;
  max-width: 460px;
  max-height: 84vh;
  overflow-y: auto;
  background: var(--bg-white);
  border-radius: var(--radius-lg);
  padding: 24px;
  box-shadow: 0 18px 50px rgba(15, 36, 64, 0.25);
}

.intent-close {
  position: absolute;
  top: 12px;
  right: 12px;
  width: 30px;
  height: 30px;
  padding: 0;
  border: none;
  background: transparent;
  border-radius: 50%;
  font-size: 20px;
  line-height: 1;
  color: var(--text-light);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.intent-close:hover {
  background: rgba(15, 36, 64, 0.06);
  color: var(--primary);
}

.intent-title {
  font-size: 1.2rem;
  font-weight: 800;
  color: var(--primary);
  margin-bottom: 6px;
  padding-right: 30px;
}

.intent-sub {
  font-size: 0.85rem;
  color: var(--text-secondary);
  margin-bottom: 14px;
  line-height: 1.6;
}

.intent-items {
  list-style: none;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  padding: 10px 12px;
  margin-bottom: 14px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  max-height: 150px;
  overflow-y: auto;
}

.intent-items li {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  font-size: 0.85rem;
  color: var(--text-secondary);
}

.intent-item-name em {
  font-style: normal;
  color: var(--text-light);
}

.intent-item-qty {
  flex: 0 0 auto;
  font-weight: 700;
  color: var(--primary);
}

.intent-label {
  display: block;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--text-secondary);
  margin-bottom: 6px;
}

.intent-input {
  width: 100%;
  box-sizing: border-box;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  padding: 9px 12px;
  font-size: 0.9rem;
  color: var(--text-primary);
  margin-bottom: 12px;
  font-family: inherit;
}

.intent-input:focus {
  outline: none;
  border-color: var(--primary);
}

.intent-textarea {
  resize: vertical;
  margin-bottom: 21px;
}

.intent-error {
  font-size: 0.82rem;
  color: var(--danger);
  margin: 0 0 10px;
}

.intent-actions {
  position: sticky;
  bottom: 0;
  margin: 0 -24px -24px;
  padding: 14px 24px 0;
  background: var(--bg-white);
}

.intent-actions-row {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

@media (max-width: 767px) {
  .cart-item {
    flex-wrap: wrap;
    gap: 12px;
  }

  .cart-item-info {
    flex: 1 1 55%;
  }

  .cart-item-qty {
    margin-left: auto;
  }

  .cart-summary-actions {
    flex-direction: column;
  }

  .cart-summary-actions .btn {
    width: 100%;
    justify-content: center;
  }

  /* 提交弹窗：手机端改为底部抽屉式，内容滚动、按钮吸底 */
  .intent-mask {
    padding: 0;
    align-items: flex-end;
  }

  .intent-modal {
    max-width: none;
    max-height: 88vh;
    border-radius: 20px 20px 0 0;
    padding: 20px 18px 18px;
  }

  .intent-actions {
    margin: 0 -18px -18px;
    padding: 10px 18px 12px;
  }

  .intent-actions .btn {
    flex: 1;
    justify-content: center;
  }
}
</style>
