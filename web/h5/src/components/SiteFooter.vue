<template>
  <footer class="site-footer">
    <div class="footer-main container">
      <!-- Company Info -->
      <div class="footer-col footer-about">
        <router-link to="/" class="footer-logo">
          <span class="logo-icon">⬡</span>
          <span class="logo-text">HKR<span class="logo-accent">oids</span></span>
        </router-link>
        <p class="footer-desc">
          Your trusted supplier of premium research peptides and laboratory equipment.
          Over a decade of experience serving researchers and institutions worldwide.
        </p>
        <div class="social-links">
          <a href="#" class="social-link" aria-label="Facebook" title="Facebook">
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
          </a>
          <a href="#" class="social-link" aria-label="Twitter" title="Twitter">
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/></svg>
          </a>
          <a href="#" class="social-link" aria-label="Instagram" title="Instagram">
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
          </a>
          <a href="#" class="social-link" aria-label="Discord" title="Discord">
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M20.317 4.37a19.791 19.791 0 00-4.885-1.515.074.074 0 00-.079.037c-.21.375-.444.864-.608 1.25a18.27 18.27 0 00-5.487 0 12.64 12.64 0 00-.617-1.25.077.077 0 00-.079-.037A19.736 19.736 0 003.677 4.37a.07.07 0 00-.032.027C.533 9.046-.32 13.58.099 18.057a.082.082 0 00.031.057 19.9 19.9 0 005.993 3.03.078.078 0 00.084-.028c.462-.63.874-1.295 1.226-1.994a.076.076 0 00-.041-.106 13.107 13.107 0 01-1.872-.892.077.077 0 01-.008-.128 10.2 10.2 0 00.372-.292.074.074 0 01.077-.01c3.928 1.793 8.18 1.793 12.062 0a.074.074 0 01.078.01c.12.098.246.198.373.292a.077.077 0 01-.006.127 12.299 12.299 0 01-1.873.892.077.077 0 00-.041.107c.36.698.772 1.362 1.225 1.993a.076.076 0 00.084.028 19.839 19.839 0 006.002-3.03.077.077 0 00.032-.054c.5-5.177-.838-9.674-3.549-13.66a.061.061 0 00-.031-.03zM8.02 15.33c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.956 2.418-2.157 2.418zm7.975 0c-1.183 0-2.157-1.085-2.157-2.419 0-1.333.956-2.419 2.157-2.419 1.21 0 2.176 1.096 2.157 2.42 0 1.333-.947 2.418-2.157 2.418z"/></svg>
          </a>
        </div>
      </div>

      <!-- Quick Links -->
      <div class="footer-col">
        <h4 class="footer-heading">Quick Links</h4>
        <ul class="footer-links">
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/about">About Us</router-link></li>
          <li><router-link to="/products">Products</router-link></li>
          <li><router-link to="/contact">Contact Us</router-link></li>
        </ul>
      </div>

      <!-- Categories -->
      <div class="footer-col">
        <h4 class="footer-heading">Categories</h4>
        <ul class="footer-links">
          <li><router-link to="/products?category=peptides">Research Peptides</router-link></li>
          <li><router-link to="/products?category=raws">Peptide Raws</router-link></li>
          <li><router-link to="/products?category=blends">Peptide Blends</router-link></li>
          <li><router-link to="/products?category=equipment">Lab Equipment</router-link></li>
        </ul>
      </div>

      <!-- Contact Info -->
      <div class="footer-col">
        <h4 class="footer-heading">Contact Info</h4>
        <ul class="footer-contact">
          <li>
            <span class="contact-icon">📧</span>
            <a :href="'mailto:' + contactInfo.email">{{ contactInfo.email }}</a>
          </li>
          <li>
            <span class="contact-icon">📞</span>
            <a :href="phoneHref">{{ contactInfo.phone }}</a>
          </li>
          <li>
            <span class="contact-icon">📍</span>
            <span>{{ contactInfo.address }}</span>
          </li>
          <li>
            <span class="contact-icon">🕐</span>
            <span>{{ contactInfo.business_hours }}</span>
          </li>
        </ul>
      </div>
    </div>

    <!-- Footer Bottom -->
    <div class="footer-bottom">
      <div class="container footer-bottom-inner">
        <p>&copy; {{ currentYear }} HKRoids.com. All rights reserved.</p>
        <div class="footer-bottom-links">
          <a href="#">Privacy Policy</a>
          <a href="#">Terms of Service</a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import {GetContactInfo} from '@/api/contact'

export default {
  name: 'SiteFooter',
  data() {
    return {
      // 联系方式，默认值兼容接口未返回的场景，后台可配置
      contactInfo: {
        email: 'info@hkroids.com',
        phone: '+1 (234) 567-890',
        address: 'Hong Kong, China',
        business_hours: 'Mon – Sat: 9:00 AM – 6:00 PM'
      }
    }
  },
  computed: {
    currentYear() {
      return new Date().getFullYear()
    },
    // 电话拨打链接：去掉非数字字符
    phoneHref() {
      return 'tel:' + (this.contactInfo.phone || '').replace(/[^\d+]/g, '')
    }
  },
  created() {
    // 从后台加载联系方式配置，失败时保留默认值
    GetContactInfo().then(list => {
      if (list && list.length) {
        this.contactInfo = {...this.contactInfo, ...list[0]}
      }
    }).catch(() => {
    })
  }
}
</script>

<style scoped>
.site-footer {
  background: var(--bg-dark);
  color: #cbd5e0;
  margin-top: auto;
}

.footer-main {
  display: grid;
  grid-template-columns: 1.5fr 1fr 1fr 1.2fr;
  gap: 48px;
  padding: 64px var(--container-padding) 48px;
}

.footer-logo {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 1.4rem;
  font-weight: 800;
  color: #fff;
  margin-bottom: 16px;
}

.logo-icon {
  color: var(--accent);
}

.logo-accent {
  color: var(--accent);
}

.footer-desc {
  font-size: 0.9rem;
  line-height: 1.7;
  color: #a0aec0;
  margin-bottom: 20px;
}

.social-links {
  display: flex;
  gap: 10px;
}

.social-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.08);
  color: #a0aec0;
  transition: all 0.3s ease;
}

.social-link:hover {
  background: var(--accent);
  color: var(--primary-dark);
  transform: translateY(-2px);
}

.footer-heading {
  font-size: 1rem;
  font-weight: 700;
  color: #fff;
  margin-bottom: 20px;
  letter-spacing: 0.5px;
}

.footer-links li {
  margin-bottom: 12px;
}

.footer-links a {
  font-size: 0.9rem;
  color: #a0aec0;
  transition: all 0.2s ease;
}

.footer-links a:hover {
  color: var(--accent);
  padding-left: 4px;
}

.footer-contact li {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  margin-bottom: 14px;
  font-size: 0.9rem;
}

.footer-contact a {
  color: #a0aec0;
}

.footer-contact a:hover {
  color: var(--accent);
}

.contact-icon {
  flex-shrink: 0;
  font-size: 0.95rem;
}

/* Footer Bottom */
.footer-bottom {
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  padding: 20px 0;
}

.footer-bottom-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 0.85rem;
  color: #718096;
}

.footer-bottom-links {
  display: flex;
  gap: 24px;
}

.footer-bottom-links a {
  color: #718096;
}

.footer-bottom-links a:hover {
  color: var(--accent);
}

/* Responsive */
@media (max-width: 1023px) {
  .footer-main {
    grid-template-columns: 1fr 1fr;
    gap: 36px;
  }
}

@media (max-width: 767px) {
  .footer-main {
    grid-template-columns: 1fr;
    gap: 32px;
    padding: 48px var(--container-padding) 36px;
  }

  .footer-bottom-inner {
    flex-direction: column;
    gap: 12px;
    text-align: center;
  }
}
</style>
