<template>
  <div class="home-page">
    <!-- Hero Banner Carousel -->
    <section class="hero-banner">
      <div class="banner-slides">
        <transition-group name="banner-fade" tag="div" class="slides-wrapper">
          <div
              v-for="(banner, index) in banners"
              v-show="currentSlide === index"
              :key="banner.id"
              class="banner-slide"
              :style="{ backgroundImage: `url(${banner.image})` }"
          >
            <div class="banner-overlay"></div>
            <div class="banner-content container">
              <h1 class="banner-title">{{ banner.title }}</h1>
              <p class="banner-subtitle">{{ banner.subtitle }}</p>
              <router-link :to="banner.link" class="btn btn-accent btn-lg banner-cta">
                {{ banner.cta }}
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M5 12h14M12 5l7 7-7 7"/>
                </svg>
              </router-link>
            </div>
          </div>
        </transition-group>
      </div>
      <!-- Banner Controls -->
      <div class="banner-dots">
        <button
            v-for="(banner, index) in banners"
            :key="'dot-' + banner.id"
            class="dot"
            :class="{ active: currentSlide === index }"
            @click="goToSlide(index)"
            :aria-label="'Go to slide ' + (index + 1)"
        ></button>
      </div>
      <button class="banner-arrow arrow-left" @click="prevSlide" aria-label="Previous">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M15 18l-6-6 6-6"/>
        </svg>
      </button>
      <button class="banner-arrow arrow-right" @click="nextSlide" aria-label="Next">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M9 18l6-6-6-6"/>
        </svg>
      </button>
    </section>

    <!-- Company Introduction -->
    <section class="section intro-section">
      <div class="container">
        <div class="intro-grid">
          <div class="intro-text fade-in-up">
            <span class="section-label">About HKRoids</span>
            <h2 class="intro-title">Your Trusted Research Peptides Wholesale Supplier</h2>
            <p class="intro-desc">
              We specialize in providing scientists, researchers, and institutions with premium-grade
              research peptides designed for cutting-edge studies in biotechnology, pharmacology,
              and medical research. Our commitment to purity, reliability, and innovation ensures
              that you receive products that meet the highest industry standards.
            </p>
            <div class="intro-stats">
              <div class="stat-item">
                <span class="stat-number">10+</span>
                <span class="stat-label">Years Experience</span>
              </div>
              <div class="stat-item">
                <span class="stat-number">99%+</span>
                <span class="stat-label">Purity Tested</span>
              </div>
              <div class="stat-item">
                <span class="stat-number">5000+</span>
                <span class="stat-label">Happy Clients</span>
              </div>
            </div>
            <router-link to="/about" class="btn btn-outline">
              Learn More About Us
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M5 12h14M12 5l7 7-7 7"/>
              </svg>
            </router-link>
          </div>
          <div class="intro-image">
            <div class="intro-image-fallback">
              <svg class="fallback-art" viewBox="0 0 400 300" preserveAspectRatio="xMidYMid slice" aria-hidden="true">
                <!-- Corner decorative hexagons -->
                <polygon points="360,-15 321,8 321,53 360,75 399,53 399,8" fill="none" stroke="rgba(255,255,255,0.1)"
                         stroke-width="2"/>
                <polygon points="35,240 9,255 9,285 35,300 61,285 61,255" fill="none" stroke="rgba(255,255,255,0.1)"
                         stroke-width="2"/>
                <!-- Scattered dots -->
                <circle cx="70" cy="60" r="3" fill="rgba(255,255,255,0.3)"/>
                <circle cx="330" cy="250" r="3" fill="rgba(255,255,255,0.3)"/>
                <circle cx="370" cy="180" r="2.5" fill="rgba(201,168,76,0.5)"/>
                <circle cx="40" cy="140" r="2.5" fill="rgba(201,168,76,0.5)"/>
                <!-- Molecular chain bonds -->
                <g stroke="rgba(255,255,255,0.4)" stroke-width="2">
                  <line x1="200" y1="92" x2="200" y2="52"/>
                  <line x1="250" y1="121" x2="302" y2="95"/>
                  <line x1="302" y1="95" x2="348" y2="112"/>
                  <line x1="150" y1="179" x2="98" y2="205"/>
                  <line x1="98" y1="205" x2="52" y2="188"/>
                  <line x1="200" y1="208" x2="200" y2="248"/>
                  <line x1="150" y1="121" x2="95" y2="90"/>
                  <line x1="250" y1="179" x2="308" y2="212"/>
                </g>
                <!-- Central benzene ring -->
                <polygon points="200,92 150,121 150,179 200,208 250,179 250,121" fill="rgba(255,255,255,0.04)"
                         stroke="rgba(255,255,255,0.6)" stroke-width="2.5"/>
                <polygon points="200,112 167,131 167,169 200,188 233,169 233,131" fill="none"
                         stroke="rgba(255,255,255,0.2)" stroke-width="1.5"/>
                <!-- Chain atoms -->
                <circle cx="200" cy="48" r="10" fill="#c9a84c"/>
                <circle cx="302" cy="95" r="7" fill="rgba(255,255,255,0.85)"/>
                <circle cx="348" cy="112" r="11" fill="#c9a84c"/>
                <circle cx="98" cy="205" r="7" fill="rgba(255,255,255,0.85)"/>
                <circle cx="52" cy="188" r="11" fill="#c9a84c"/>
                <circle cx="200" cy="252" r="9" fill="rgba(255,255,255,0.5)"/>
                <circle cx="92" cy="88" r="9" fill="none" stroke="rgba(255,255,255,0.6)" stroke-width="2"/>
                <circle cx="312" cy="215" r="9" fill="none" stroke="rgba(255,255,255,0.6)" stroke-width="2"/>
                <!-- Benzene ring vertex atoms -->
                <circle cx="200" cy="92" r="4" fill="rgba(255,255,255,0.9)"/>
                <circle cx="150" cy="121" r="4" fill="rgba(255,255,255,0.9)"/>
                <circle cx="150" cy="179" r="4" fill="rgba(255,255,255,0.9)"/>
                <circle cx="200" cy="208" r="4" fill="rgba(255,255,255,0.9)"/>
                <circle cx="250" cy="179" r="4" fill="rgba(255,255,255,0.9)"/>
                <circle cx="250" cy="121" r="4" fill="rgba(255,255,255,0.9)"/>
              </svg>
            </div>
            <div class="intro-image-badge">
              <span class="badge-number">10+</span>
              <span class="badge-text">Years of<br/>Excellence</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Products -->
    <section class="section products-section">
      <div class="container">
        <h2 class="section-title">Featured Products</h2>
        <p class="section-subtitle">
          Explore our most popular research peptides and laboratory products, trusted by researchers worldwide.
        </p>
        <div class="products-grid">
          <ProductCard
              v-for="product in featuredProducts"
              :key="product.id"
              :product="product"
          />
        </div>
        <div class="section-cta">
          <router-link to="/products" class="btn btn-primary btn-lg">
            View All Products
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M5 12h14M12 5l7 7-7 7"/>
            </svg>
          </router-link>
        </div>
      </div>
    </section>

    <!-- Brand Advantages -->
    <section class="section advantages-section">
      <div class="container">
        <h2 class="section-title">Why Choose Us</h2>
        <p class="section-subtitle">
          We are committed to delivering excellence in every product and service.
        </p>
        <div class="advantages-grid">
          <div v-for="(adv, idx) in advantages" :key="idx" class="advantage-card">
            <div class="advantage-icon" :style="{ background: adv.bgColor }">
              <span>{{ adv.icon }}</span>
            </div>
            <h3 class="advantage-title">{{ adv.title }}</h3>
            <p class="advantage-desc">{{ adv.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="cta-section">
      <div class="cta-bg"
           style="background-image: url('https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=1920&h=600&fit=crop')"></div>
      <div class="cta-overlay"></div>
      <div class="container cta-content">
        <h2 class="cta-title">Ready to Start Your Research?</h2>
        <p class="cta-desc">
          Contact us today for bulk pricing, custom orders, or any inquiries.
          Our team is ready to assist you within 4-5 hours.
        </p>
        <div class="cta-buttons">
          <router-link to="/contact" class="btn btn-accent btn-lg">Contact Us Now</router-link>
          <router-link to="/products" class="btn btn-lg" style="color:#fff;border:2px solid rgba(255,255,255,0.4);">
            Browse Products
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import ProductCard from '../components/ProductCard.vue'
import {GetProduct} from '@/api/product'

export default {
  name: 'HomePage',
  components: {ProductCard},
  data() {
    return {
      banners:[
        {
          id: 1,
          title: 'Premium Research Peptides',
          subtitle: 'Trusted by researchers worldwide. Lab-tested purity above 99%.',
          image: require('@/assets/b1.jpg'),
          cta: 'Explore Products',
          link: '/products'
        },
        {
          id: 2,
          title: 'Cutting-Edge Lab Equipment',
          subtitle: 'Professional-grade machinery for peptide synthesis and lyophilization.',
          image: require('@/assets/b2.jpg'),
          cta: 'View Equipment',
          link: '/products?category=equipment'
        },
        {
          id: 3,
          title: 'Wholesale & Bulk Orders',
          subtitle: 'Competitive pricing for research institutions and laboratories.',
          image: require('@/assets/b3.jpg'),
          cta: 'Contact Us',
          link: '/contact'
        }
      ],
      products: [],
      currentSlide: 0,
      slideInterval: null,
      advantages: [
        {
          icon: '🚀',
          title: 'Fast Delivery',
          description: 'Next-day shipping with 6-8 days express international delivery. We ensure your research never stops.',
          bgColor: 'rgba(66, 153, 225, 0.1)'
        },
        {
          icon: '🔬',
          title: 'Lab-Tested Quality',
          description: 'Every batch undergoes rigorous HPLC testing with purity above 99%. COA and test reports included.',
          bgColor: 'rgba(72, 187, 120, 0.1)'
        },
        {
          icon: '💰',
          title: 'Competitive Pricing',
          description: 'Wholesale pricing for bulk orders. Special discounts available for research institutions and universities.',
          bgColor: 'rgba(237, 137, 54, 0.1)'
        },
        {
          icon: '🛡️',
          title: 'Secure Payments',
          description: 'Multiple payment options including bank transfer, PayPal, and cryptocurrency for your convenience.',
          bgColor: 'rgba(159, 122, 234, 0.1)'
        },
        {
          icon: '📦',
          title: 'Free Reshipping',
          description: 'We offer one free reship guarantee on every order. Your satisfaction and security are our priority.',
          bgColor: 'rgba(237, 100, 166, 0.1)'
        },
        {
          icon: '🤝',
          title: 'Expert Support',
          description: 'Our knowledgeable team responds within 4-5 hours. Get professional guidance for your research needs.',
          bgColor: 'rgba(56, 178, 172, 0.1)'
        }
      ]
    }
  },
  computed: {
    featuredProducts() {
      return this.products.slice(0, 8)
    }
  },
  created() {
    // 从后端接口加载产品数据
    GetProduct().then(res => {
      this.products = res.productData || []
    }).catch(() => {
    })
  },
  mounted() {
    this.startAutoSlide()
  },
  beforeDestroy() {
    this.stopAutoSlide()
  },
  methods: {
    startAutoSlide() {
      this.slideInterval = setInterval(() => {
        this.nextSlide()
      }, 5000)
    },
    stopAutoSlide() {
      if (this.slideInterval) clearInterval(this.slideInterval)
    },
    nextSlide() {
      this.currentSlide = (this.currentSlide + 1) % this.banners.length
    },
    prevSlide() {
      this.currentSlide = (this.currentSlide - 1 + this.banners.length) % this.banners.length
    },
    goToSlide(index) {
      this.currentSlide = index
      this.stopAutoSlide()
      this.startAutoSlide()
    }
  }
}
</script>

<style scoped>
/* ===== Hero Banner ===== */
.hero-banner {
  position: relative;
  height: 60vh;
  min-height: 460px;
  max-height: 620px;
  overflow: hidden;
}

.banner-slides {
  height: 100%;
}

.slides-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
}

.banner-slide {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
}

.banner-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(15, 36, 64, 0.85) 0%, rgba(26, 54, 93, 0.6) 100%);
}

.banner-content {
  position: relative;
  z-index: 2;
  max-width: 680px;
  padding-top: var(--header-height);
}

.banner-title {
  font-size: 3.2rem;
  font-weight: 800;
  color: #fff;
  line-height: 1.15;
  margin-bottom: 20px;
  letter-spacing: -0.5px;
}

.banner-subtitle {
  font-size: 1.2rem;
  color: rgba(255, 255, 255, 0.85);
  line-height: 1.6;
  margin-bottom: 32px;
  max-width: 540px;
}

.banner-cta {
  font-size: 1rem;
}

/* Banner Dots */
.banner-dots {
  position: absolute;
  bottom: 32px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 10px;
  z-index: 10;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.4);
  border: 2px solid transparent;
  transition: all 0.3s ease;
  cursor: pointer;
}

.dot.active {
  background: var(--accent);
  border-color: #fff;
  transform: scale(1.15);
}

/* Banner Arrows */
.banner-arrow {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 10;
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.banner-arrow:hover {
  background: rgba(255, 255, 255, 0.3);
}

.arrow-left {
  left: 24px;
}

.arrow-right {
  right: 24px;
}

/* Banner Transition */
.banner-fade-enter-active,
.banner-fade-leave-active {
  transition: opacity 0.8s ease;
}

.banner-fade-enter,
.banner-fade-leave-to {
  opacity: 0;
}

/* ===== Intro Section ===== */
.intro-section {
  background: var(--bg-white);
  padding-top: 32px;
  /* 底部预留徽章（bottom: -20px）溢出空间 */
  padding-bottom: 40px;
}

.intro-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: center;
}

.section-label {
  display: inline-block;
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  color: var(--accent-dark);
  margin-bottom: 12px;
}

.intro-title {
  font-size: 2.2rem;
  font-weight: 800;
  color: var(--primary);
  line-height: 1.2;
  margin-bottom: 20px;
}

.intro-desc {
  font-size: 1.02rem;
  color: var(--text-secondary);
  line-height: 1.8;
  margin-bottom: 32px;
}

.intro-stats {
  display: flex;
  gap: 40px;
  margin-bottom: 32px;
}

.stat-item {
  display: flex;
  flex-direction: column;
}

.stat-number {
  font-size: 2rem;
  font-weight: 800;
  color: var(--primary);
  line-height: 1;
}

.stat-label {
  font-size: 0.85rem;
  color: var(--text-light);
  margin-top: 4px;
}

.intro-image {
  position: relative;
  height: 440px;
}

.intro-image-fallback {
  position: absolute;
  inset: 0;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  background: radial-gradient(circle at 78% 22%, rgba(255, 255, 255, 0.09) 0%, transparent 45%),
  radial-gradient(circle at 15% 85%, rgba(201, 168, 76, 0.12) 0%, transparent 40%),
  linear-gradient(135deg, var(--primary) 0%, var(--primary-dark) 100%);
  overflow: hidden;
}

.fallback-art {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
}

.intro-image-badge {
  position: absolute;
  bottom: -20px;
  left: -20px;
  background: linear-gradient(135deg, var(--accent), var(--accent-light));
  color: var(--primary-dark);
  padding: 20px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  gap: 10px;
  box-shadow: var(--shadow-md);
}

.badge-number {
  font-size: 2.2rem;
  font-weight: 800;
  line-height: 1;
}

.badge-text {
  font-size: 0.8rem;
  font-weight: 600;
  line-height: 1.3;
}

/* ===== Products Section ===== */
.products-section {
  background: var(--bg-light);
  padding-top: 36px;
  padding-bottom: 32px;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  margin-bottom: 32px;
}

.section-cta {
  text-align: center;
}

/* ===== Advantages Section ===== */
.advantages-section {
  background: var(--bg-white);
  padding-top: 36px;
}

.advantages-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}

.advantage-card {
  padding: 36px 28px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
  transition: all 0.35s ease;
  text-align: center;
}

.advantage-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
  border-color: transparent;
}

.advantage-icon {
  width: 64px;
  height: 64px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  font-size: 1.8rem;
}

.advantage-title {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 10px;
}

.advantage-desc {
  font-size: 0.9rem;
  color: var(--text-secondary);
  line-height: 1.65;
}

/* ===== CTA Section ===== */
.cta-section {
  position: relative;
  padding: 64px 0;
  overflow: hidden;
}

.cta-bg {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
}

.cta-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(15, 36, 64, 0.9) 0%, rgba(26, 54, 93, 0.85) 100%);
}

.cta-content {
  position: relative;
  z-index: 2;
  text-align: center;
}

.cta-title {
  font-size: 2.4rem;
  font-weight: 800;
  color: #fff;
  margin-bottom: 16px;
}

.cta-desc {
  font-size: 1.1rem;
  color: rgba(255, 255, 255, 0.8);
  max-width: 560px;
  margin: 0 auto 36px;
  line-height: 1.7;
}

.cta-buttons {
  display: flex;
  justify-content: center;
  gap: 16px;
  flex-wrap: wrap;
}

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .hero-banner {
    height: 56vh;
    max-height: none;
  }

  .banner-title {
    font-size: 2.4rem;
  }

  .intro-grid {
    gap: 40px;
  }

  .products-grid {
    grid-template-columns: repeat(3, 1fr);
  }

  .advantages-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .cta-bg {
    background-attachment: scroll;
  }
}

@media (max-width: 767px) {
  .hero-banner {
    height: 38vh;
    min-height: 340px;
  }

  .banner-title {
    font-size: 1.75rem;
  }

  .banner-subtitle {
    font-size: 1rem;
    margin-bottom: 24px;
  }

  .banner-arrow {
    display: none;
  }

  .intro-grid {
    grid-template-columns: 1fr;
    gap: 32px;
  }

  .intro-image {
    order: -1;
    height: 220px;
  }

  .intro-stats {
    gap: 24px;
    margin-bottom: 24px;
  }

  .stat-number {
    font-size: 1.6rem;
  }

  .products-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 14px;
    margin-bottom: 24px;
  }

  .advantages-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .advantage-card {
    padding: 24px 20px;
  }

  .cta-section {
    padding: 48px 0;
  }

  .cta-bg {
    background-attachment: scroll;
  }

  .cta-title {
    font-size: 1.6rem;
  }

  .cta-buttons {
    flex-direction: column;
    align-items: center;
  }

  /* 压缩各区块上下内边距，减少衔接处空白 */
  .intro-section {
    padding-top: 20px;
    padding-bottom: 28px;
  }

  .products-section {
    padding-top: 24px;
    padding-bottom: 20px;
  }

  .advantages-section {
    padding-top: 24px;
  }
}
</style>
