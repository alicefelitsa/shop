<template>
  <div class="contact-page">
    <!-- Page Header -->
    <section class="page-header">
      <div class="page-header-bg" style="background-image: url('https://images.unsplash.com/photo-1550572017-edd951b55104?w=1920&h=500&fit=crop')"></div>
      <div class="page-header-overlay"></div>
      <div class="container page-header-content">
        <h1 class="page-title">Contact Us</h1>
        <p class="page-breadcrumb">
          <router-link to="/">Home</router-link>
          <span>/</span>
          <span>Contact</span>
        </p>
      </div>
    </section>

    <!-- Contact Content -->
    <section class="section contact-section">
      <div class="container">
        <div class="contact-grid">
          <!-- Contact Form -->
          <div class="contact-form-wrapper">
            <h2 class="form-title">Send Us a Message</h2>
            <p class="form-subtitle">
              Fill out the form below and we'll get back to you within 4-5 hours.
            </p>

            <form class="contact-form" novalidate @submit.prevent="submitForm">
              <div class="form-row">
                <div class="form-group">
                  <label class="form-label">Full Name <span class="required">*</span></label>
                  <input
                    v-model="form.name"
                    type="text"
                    class="form-input"
                    :class="{ 'is-invalid': errors.name }"
                    placeholder="Your full name"
                    @input="clearError('name')"
                  />
                  <transition name="fade-in">
                    <p v-if="errors.name" class="field-error">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
                      {{ errors.name }}
                    </p>
                  </transition>
                </div>
                <div class="form-group">
                  <label class="form-label">Email Address <span class="required">*</span></label>
                  <input
                    v-model="form.email"
                    type="email"
                    class="form-input"
                    :class="{ 'is-invalid': errors.email }"
                    placeholder="your@email.com"
                    @input="clearError('email')"
                  />
                  <transition name="fade-in">
                    <p v-if="errors.email" class="field-error">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
                      {{ errors.email }}
                    </p>
                  </transition>
                </div>
              </div>

              <div class="form-group">
                <label class="form-label">Subject <span class="required">*</span></label>
                <select
                  v-model="form.subject"
                  class="form-input form-select"
                  :class="{ 'is-invalid': errors.subject }"
                  @change="clearError('subject')"
                >
                  <option value="" disabled>Select a subject</option>
                  <option value="general">General Inquiry</option>
                  <option value="order">Order Inquiry</option>
                  <option value="wholesale">Wholesale / Bulk Order</option>
                  <option value="support">Technical Support</option>
                  <option value="partnership">Partnership</option>
                  <option value="other">Other</option>
                </select>
                <transition name="fade-in">
                  <p v-if="errors.subject" class="field-error">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
                    {{ errors.subject }}
                  </p>
                </transition>
              </div>

              <div class="form-group">
                <label class="form-label">Message <span class="required">*</span></label>
                <textarea
                  v-model="form.message"
                  class="form-input form-textarea"
                  :class="{ 'is-invalid': errors.message }"
                  placeholder="Please tell us how we can help you, and leave your contact information so we can assist you as soon as possible..."
                  rows="6"
                  @input="clearError('message')"
                ></textarea>
                <transition name="fade-in">
                  <p v-if="errors.message" class="field-error">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
                    {{ errors.message }}
                  </p>
                </transition>
              </div>

              <button type="submit" class="btn btn-primary btn-lg submit-btn" :disabled="submitting">
                <span v-if="!submitting">
                  Send Message
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="vertical-align:middle;margin-left:4px;"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z"/></svg>
                </span>
                <span v-else>Sending...</span>
              </button>
            </form>
          </div>

          <!-- Contact Info Sidebar -->
          <div class="contact-info">
            <div class="info-card">
              <h3 class="info-title">Get in Touch</h3>
              <p class="info-desc">
                Have questions about our products or need assistance with an order?
                We're here to help.
              </p>

              <div class="info-items">
                <div class="info-item">
                  <div class="info-icon">📧</div>
                  <div class="info-content">
                    <span class="info-label">Email</span>
                    <a v-if="contactInfo.email" :href="'mailto:' + contactInfo.email" class="info-value">{{ contactInfo.email }}</a>
                    <span v-else class="info-value">None</span>
                  </div>
                </div>

                <div class="info-item">
                  <div class="info-icon">📞</div>
                  <div class="info-content">
                    <span class="info-label">Phone</span>
                    <a v-if="contactInfo.phone" :href="phoneHref" class="info-value">{{ contactInfo.phone }}</a>
                    <span v-else class="info-value">None</span>
                  </div>
                </div>

                <div class="info-item">
                  <div class="info-icon">📍</div>
                  <div class="info-content">
                    <span class="info-label">Address</span>
                    <span class="info-value">{{ contactInfo.address || 'None' }}</span>
                  </div>
                </div>

                <div class="info-item">
                  <div class="info-icon">🕐</div>
                  <div class="info-content">
                    <span class="info-label">Business Hours</span>
                    <span class="info-value">{{ contactInfo.business_hours || 'None' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Payment Methods -->
            <div class="info-card payment-card">
              <h3 class="info-title">Payment Methods</h3>
              <div class="payment-methods">
                <div class="payment-item">
                  <span class="payment-icon">🏦</span>
                  <span>Bank Transfer (EUR, USD, AUD, CAD)</span>
                </div>
                <div class="payment-item">
                  <span class="payment-icon">💳</span>
                  <span>PayPal</span>
                </div>
                <div class="payment-item">
                  <span class="payment-icon">₿</span>
                  <span>Bitcoin, USDT, ETH, USDC, XMR</span>
                </div>
              </div>
            </div>

            <!-- Social -->
            <div class="info-card social-card">
              <h3 class="info-title">Join Our Community</h3>
              <div class="social-grid">
                <a href="#" class="social-btn">
                  <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
                  <span>Facebook</span>
                </a>
                <a href="#" class="social-btn">
                  <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M20.317 4.37a19.791 19.791 0 00-4.885-1.515.074.074 0 00-.079.037c-.21.375-.444.864-.608 1.25a18.27 18.27 0 00-5.487 0 12.64 12.64 0 00-.617-1.25.077.077 0 00-.079-.037A19.736 19.736 0 003.677 4.37a.07.07 0 00-.032.027C.533 9.046-.32 13.58.099 18.057a.082.082 0 00.031.057 19.9 19.9 0 005.993 3.03.078.078 0 00.084-.028c.462-.63.874-1.295 1.226-1.994a.076.076 0 00-.041-.106 13.107 13.107 0 01-1.872-.892.077.077 0 01-.008-.128 10.2 10.2 0 00.372-.292.074.074 0 01.077-.01c3.928 1.793 8.18 1.793 12.062 0a.074.074 0 01.078.01c.12.098.246.198.373.292a.077.077 0 01-.006.127 12.299 12.299 0 01-1.873.892.077.077 0 00-.041.107c.36.698.772 1.362 1.225 1.993a.076.076 0 00.084.028 19.839 19.839 0 006.002-3.03.077.077 0 00.032-.054c.5-5.177-.838-9.674-3.549-13.66a.061.061 0 00-.031-.03zM8.02 15.33c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.956 2.418-2.157 2.418zm7.975 0c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.947 2.418-2.157 2.418z"/></svg>
                  <span>Discord</span>
                </a>
                <a href="#" class="social-btn">
                  <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
                  <span>Instagram</span>
                </a>
                <a href="#" class="social-btn">
                  <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/></svg>
                  <span>Twitter / X</span>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- FAQ Section -->
    <section class="section faq-section">
      <div class="container">
        <h2 class="section-title">Frequently Asked Questions</h2>
        <p class="section-subtitle">Find quick answers to common questions below.</p>
        <div class="faq-list">
          <div
            v-for="(faq, idx) in faqs"
            :key="idx"
            class="faq-item"
            :class="{ open: openFaq === idx }"
          >
            <button class="faq-question" @click="toggleFaq(idx)">
              <span>{{ faq.q }}</span>
              <span class="faq-toggle">{{ openFaq === idx ? '−' : '+' }}</span>
            </button>
            <transition name="faq-expand">
              <div v-show="openFaq === idx" class="faq-answer">
                <p>{{ faq.a }}</p>
              </div>
            </transition>
          </div>
        </div>
      </div>
    </section>

    <!-- 提交结果弹窗 -->
    <transition name="modal">
      <div v-if="resultModal.show" class="modal-overlay" @click.self="closeResultModal">
        <div class="modal-card">
          <button type="button" class="modal-close" aria-label="Close" @click="closeResultModal">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M18 6L6 18M6 6l12 12"/></svg>
          </button>
          <div class="modal-icon" :class="resultModal.type">
            <svg v-if="resultModal.type === 'success'" width="34" height="34" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M20 6L9 17l-5-5"/></svg>
            <svg v-else width="34" height="34" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M12 7v6"/><path d="M12 16.5v.5"/></svg>
          </div>
          <h3 class="modal-title">{{ resultModal.type === 'success' ? 'Message Sent!' : 'Submission Failed' }}</h3>
          <p class="modal-text">{{ resultModal.text }}</p>
          <button type="button" class="btn btn-primary modal-btn" @click="closeResultModal">OK</button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import {AddMessage} from '@/api/message'
import {GetContactInfo} from '@/api/contact'

export default {
  name: 'ContactPage',
  data() {
    return {
      // 联系方式，由后台配置，未设置时显示 None
      contactInfo: {
        email: '',
        phone: '',
        address: '',
        business_hours: ''
      },
      form: {
        name: '',
        email: '',
        subject: '',
        message: ''
      },
      submitting: false,
      // 提交结果弹窗：type 为 success / error
      resultModal: {
        show: false,
        type: 'success',
        text: ''
      },
      // 各字段校验错误提示
      errors: {
        name: '',
        email: '',
        subject: '',
        message: ''
      },
      openFaq: null,
      faqs: [
        {
          q: 'How do I place an order?',
          a: 'You can place an order directly through our website by adding products to your cart, or contact us via email at info@hkroids.com and we will create a custom order for you.'
        },
        {
          q: 'What payment methods do you accept?',
          a: 'We accept bank transfers (EUR, USD, AUD, CAD), PayPal, and cryptocurrencies including Bitcoin, USDT, ETH, USDC, and Monero (XMR).'
        },
        {
          q: 'What is the shipping time?',
          a: 'We offer next-day shipping. Express delivery takes 6-8 days, while standard shipping takes 10-18 days depending on your location.'
        },
        {
          q: 'Do you have a reship policy?',
          a: 'Yes, we offer one free reshipment per order if there are any delivery issues. Your satisfaction is our priority.'
        },
        {
          q: 'Can I get a refund?',
          a: 'Refunds are available for orders that have not yet been shipped. Once a tracking number has been issued and the package is in transit, refunds cannot be processed.'
        },
        {
          q: 'How can I verify product quality?',
          a: 'Every product comes with a Certificate of Analysis (COA) and HPLC testing reports. We maintain ISO 9001 certification and test all batches for purity above 99%.'
        }
      ]
    }
  },
  computed: {
    // 电话拨打链接：去掉非数字字符
    phoneHref() {
      return 'tel:' + (this.contactInfo.phone || '').replace(/[^\d+]/g, '')
    }
  },
  created() {
    // 从后台加载联系方式配置
    this.fetchContactInfo()
  },
  methods: {
    // 加载联系方式配置，失败或未设置时保持为空，页面显示 None
    fetchContactInfo() {
      GetContactInfo().then(list => {
        if (list && list.length) {
          const row = list[0]
          this.contactInfo = {
            email: row.email || '',
            phone: row.phone || '',
            address: row.address || '',
            business_hours: row.business_hours || ''
          }
        }
      }).catch(() => {
      })
    },
    // 表单校验，不通过时在对应字段下方显示提示
    validateForm() {
      this.errors = { name: '', email: '', subject: '', message: '' }
      if (!this.form.name.trim()) {
        this.errors.name = 'Please enter your full name'
      }
      if (!this.form.email.trim()) {
        this.errors.email = 'Please enter your email address'
      } else if (!/^[\w.+-]+@[\w-]+(\.[\w-]+)+$/.test(this.form.email.trim())) {
        this.errors.email = 'Please enter a valid email address'
      }
      if (!this.form.subject) {
        this.errors.subject = 'Please select a subject'
      }
      if (!this.form.message.trim()) {
        this.errors.message = 'Please enter your message'
      }
      return !Object.values(this.errors).some(Boolean)
    },
    // 用户修改字段时清除该字段的错误提示
    clearError(field) {
      if (this.errors[field]) {
        this.errors[field] = ''
      }
    },
    async submitForm() {
      if (!this.validateForm()) {
        return
      }
      this.submitting = true
      try {
        // 提交留言到后端，存入 message 表
        await AddMessage({
          name: this.form.name,
          email: this.form.email,
          subject: this.form.subject,
          content: this.form.message
        })
        this.showResultModal('success', "Your message has been sent successfully! We'll respond within 4-5 hours.")
        this.form = { name: '', email: '', subject: '', message: '' }
      } catch (e) {
        this.showResultModal('error', e.message || 'Submission failed, please try again later.')
      } finally {
        this.submitting = false
      }
    },
    // 弹出提交结果弹窗
    showResultModal(type, text) {
      this.resultModal = { show: true, type, text }
    },
    // 关闭提交结果弹窗
    closeResultModal() {
      this.resultModal.show = false
    },
    toggleFaq(idx) {
      this.openFaq = this.openFaq === idx ? null : idx
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

/* ===== Contact Section ===== */
.contact-grid {
  display: grid;
  grid-template-columns: 1.4fr 1fr;
  gap: 48px;
  align-items: flex-start;
}

.form-title {
  font-size: 1.6rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 8px;
}

.form-subtitle {
  font-size: 0.95rem;
  color: var(--text-secondary);
  margin-bottom: 32px;
}

.contact-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.required {
  color: var(--danger);
}

.form-input {
  padding: 12px 16px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  font-size: 0.95rem;
  background: var(--bg-white);
  transition: all 0.25s ease;
  outline: none;
  color: var(--text-primary);
}

.form-input:focus {
  border-color: var(--primary);
  box-shadow: 0 0 0 3px rgba(26, 54, 93, 0.1);
}

.form-input::placeholder {
  color: var(--text-light);
}

.form-select {
  cursor: pointer;
  appearance: auto;
}

.form-textarea {
  resize: vertical;
  min-height: 120px;
}

.submit-btn {
  align-self: flex-start;
  min-width: 200px;
}

.submit-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

/* ===== 提交结果弹窗 ===== */
.modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  background: rgba(15, 36, 64, 0.5);
  backdrop-filter: blur(4px);
}

.modal-card {
  position: relative;
  width: 100%;
  max-width: 400px;
  padding: 44px 32px 32px;
  background: var(--bg-white);
  border-radius: var(--radius-lg);
  box-shadow: 0 20px 50px rgba(15, 36, 64, 0.25);
  text-align: center;
}

.modal-close {
  position: absolute;
  top: 14px;
  right: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 50%;
  background: transparent;
  color: var(--text-light);
  cursor: pointer;
  transition: all 0.2s ease;
}

.modal-close:hover {
  background: var(--bg-gray);
  color: var(--text-primary);
}

.modal-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 72px;
  height: 72px;
  margin: 0 auto 20px;
  border-radius: 50%;
}

.modal-icon.success {
  background: rgba(56, 161, 105, 0.12);
  color: var(--success);
  box-shadow: 0 0 0 10px rgba(56, 161, 105, 0.06);
}

.modal-icon.error {
  background: rgba(229, 62, 62, 0.1);
  color: var(--danger);
  box-shadow: 0 0 0 10px rgba(229, 62, 62, 0.06);
}

.modal-title {
  font-size: 1.3rem;
  font-weight: 800;
  color: var(--primary);
  margin-bottom: 10px;
}

.modal-text {
  font-size: 0.92rem;
  color: var(--text-secondary);
  line-height: 1.65;
  margin-bottom: 24px;
}

.modal-btn {
  min-width: 160px;
}

/* 弹窗出入场动画 */
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.25s ease;
}

.modal-enter-active .modal-card,
.modal-leave-active .modal-card {
  transition: transform 0.25s ease;
}

.modal-enter,
.modal-leave-to {
  opacity: 0;
}

.modal-enter .modal-card {
  transform: scale(0.9) translateY(12px);
}

.modal-leave-to .modal-card {
  transform: scale(0.95);
}

/* 字段校验错误提示 */
.field-error {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-top: 8px;
  font-size: 0.84rem;
  color: var(--danger);
  font-weight: 500;
}

.field-error svg {
  flex-shrink: 0;
}

.form-input.is-invalid {
  border-color: var(--danger);
}

.form-input.is-invalid:focus {
  border-color: var(--danger);
  box-shadow: 0 0 0 3px rgba(229, 62, 62, 0.1);
}

/* 提示淡入淡出动画 */
.fade-in-enter-active,
.fade-in-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}

.fade-in-enter,
.fade-in-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}

/* ===== Contact Info Cards ===== */
.contact-info {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.info-card {
  background: var(--bg-light);
  padding: 32px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-light);
}

.info-title {
  font-size: 1.15rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 12px;
}

.info-desc {
  font-size: 0.9rem;
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 24px;
}

.info-items {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-item {
  display: flex;
  gap: 14px;
  align-items: flex-start;
}

.info-icon {
  font-size: 1.3rem;
  flex-shrink: 0;
  width: 32px;
  text-align: center;
  padding-top: 2px;
}

.info-content {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.info-label {
  font-size: 0.78rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.8px;
  color: var(--text-light);
}

.info-value {
  font-size: 0.92rem;
  color: var(--text-primary);
  font-weight: 500;
}

a.info-value:hover {
  color: var(--accent-dark);
}

/* Payment */
.payment-methods {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.payment-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 0.9rem;
  color: var(--text-secondary);
}

.payment-icon {
  font-size: 1.2rem;
}

/* Social */
.social-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.social-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  border-radius: var(--radius-sm);
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  font-size: 0.85rem;
  font-weight: 500;
  color: var(--text-secondary);
  transition: all 0.25s ease;
}

.social-btn:hover {
  border-color: var(--primary);
  color: var(--primary);
  transform: translateY(-1px);
}

.social-btn svg {
  flex-shrink: 0;
}

/* ===== FAQ Section ===== */
.faq-section {
  background: var(--bg-light);
}

.faq-list {
  max-width: 800px;
  margin: 0 auto;
}

.faq-item {
  background: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  margin-bottom: 12px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.faq-item.open {
  border-color: var(--primary);
  box-shadow: var(--shadow-sm);
}

.faq-question {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 18px 24px;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
  text-align: left;
  transition: color 0.2s;
}

.faq-question:hover {
  color: var(--primary);
}

.faq-toggle {
  font-size: 1.4rem;
  color: var(--accent);
  font-weight: 300;
  flex-shrink: 0;
  margin-left: 16px;
}

.faq-answer {
  padding: 0 24px 18px;
  overflow: hidden;
}

.faq-answer p {
  font-size: 0.9rem;
  color: var(--text-secondary);
  line-height: 1.7;
}

.faq-expand-enter-active,
.faq-expand-leave-active {
  transition: all 0.3s ease;
}

.faq-expand-enter,
.faq-expand-leave-to {
  opacity: 0;
  max-height: 0;
  padding-bottom: 0;
}

/* 压缩各区块上下内边距，减少衔接处空白 */
.contact-section,
.faq-section {
  padding-top: 36px;
  padding-bottom: 36px;
}

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .page-header { height: 260px; }
  .page-title { font-size: 2.2rem; }

  .contact-grid {
    grid-template-columns: 1fr;
    gap: 40px;
  }
}

@media (max-width: 767px) {
  .page-header { height: 160px; padding-top: var(--header-height); }
  .page-title { font-size: 1.5rem; }
  .page-header-content { padding-top: 0; }

  .contact-section {
    padding-top: 20px;
    padding-bottom: 20px;
  }

  .faq-section {
    padding-top: 24px;
    padding-bottom: 24px;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .info-card {
    padding: 24px;
  }

  .social-grid {
    grid-template-columns: 1fr;
  }

  .faq-question {
    padding: 16px 20px;
    font-size: 0.9rem;
  }

  .faq-answer {
    padding: 0 20px 16px;
  }
}
</style>
