<template>
  <div class="header-wrapper">
    <header class="site-header" :class="{ scrolled: isScrolled, 'menu-open': menuOpen }">
      <div class="header-inner container">
        <!-- Logo -->
        <router-link to="/" class="logo" @click.native="closeMenu">
          <span class="logo-icon">⬡</span>
          <span class="logo-text">HKR<span class="logo-accent">oids</span></span>
        </router-link>

        <!-- Desktop Navigation -->
        <nav class="nav-desktop">
          <router-link
            v-for="item in navItems"
            :key="item.path"
            :to="item.path"
            class="nav-link"
            :class="{ active: isActive(item.path) }"
          >
            {{ item.label }}
          </router-link>
        </nav>

        <!-- Header Actions -->
        <div class="header-actions">
          <router-link to="/contact" class="btn btn-accent btn-sm header-cta" @click.native="closeMenu">
            Get Quote
          </router-link>

          <!-- Hamburger -->
          <button class="hamburger" :class="{ open: menuOpen }" @click="toggleMenu" aria-label="Toggle menu">
            <span></span>
            <span></span>
            <span></span>
          </button>
        </div>
      </div>
    </header>

    <!-- Mobile Navigation (outside header to avoid backdrop-filter containing block) -->
    <transition name="slide-down">
      <div v-if="menuOpen" class="nav-mobile">
        <nav class="nav-mobile-inner">
          <router-link
            v-for="item in navItems"
            :key="'m-' + item.path"
            :to="item.path"
            class="nav-mobile-link"
            :class="{ active: isActive(item.path) }"
            @click.native="closeMenu"
          >
            <span class="nav-mobile-icon">{{ item.icon }}</span>
            {{ item.label }}
          </router-link>
          <router-link to="/contact" class="btn btn-accent mobile-cta" @click.native="closeMenu">
            Get Discounted Quote
          </router-link>
        </nav>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'SiteHeader',
  data() {
    return {
      menuOpen: false,
      isScrolled: false,
      navItems: [
        { path: '/', label: 'Home', icon: '🏠' },
        { path: '/about', label: 'About Us', icon: '🏢' },
        { path: '/products', label: 'Products', icon: '🧪' },
        { path: '/contact', label: 'Contact', icon: '✉️' }
      ]
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  watch: {
    '$route'() {
      this.closeMenu()
    }
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 20
    },
    toggleMenu() {
      this.menuOpen = !this.menuOpen
      document.body.style.overflow = this.menuOpen ? 'hidden' : ''
    },
    closeMenu() {
      this.menuOpen = false
      document.body.style.overflow = ''
    },
    isActive(path) {
      if (path === '/') return this.$route.path === '/'
      return this.$route.path.startsWith(path)
    }
  }
}
</script>

<style scoped>
/* Wrapper */
.header-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.site-header {
  background: #fff;
  transition: box-shadow 0.3s ease, border-bottom-color 0.3s ease;
  height: var(--header-height);
  border-bottom: 1px solid transparent;
  transform: translateZ(0);
  backface-visibility: hidden;
}

.site-header.scrolled {
  box-shadow: var(--shadow-md);
  border-bottom-color: var(--border-light);
}

.header-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
}

/* Logo */
.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--primary);
  letter-spacing: -0.5px;
  z-index: 1001;
}

.logo-icon {
  font-size: 1.6rem;
  color: var(--accent);
}

.logo-accent {
  color: var(--accent);
}

/* Desktop Nav */
.nav-desktop {
  display: flex;
  align-items: center;
  gap: 6px;
}

.nav-link {
  padding: 8px 18px;
  font-size: 0.925rem;
  font-weight: 500;
  color: var(--text-secondary);
  border-radius: var(--radius-sm);
  transition: all 0.25s ease;
  position: relative;
}

.nav-link::after {
  content: '';
  position: absolute;
  bottom: 2px;
  left: 18px;
  right: 18px;
  height: 2px;
  background: var(--accent);
  border-radius: 1px;
  transform: scaleX(0);
  transition: transform 0.25s ease;
}

.nav-link:hover {
  color: var(--primary);
  background: var(--bg-light);
}

.nav-link.active {
  color: var(--primary);
  font-weight: 600;
}

.nav-link.active::after {
  transform: scaleX(1);
}

/* Header Actions */
.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
  z-index: 1001;
}

/* Hamburger */
.hamburger {
  display: none;
  flex-direction: column;
  justify-content: center;
  gap: 5px;
  width: 36px;
  height: 36px;
  padding: 6px;
  border-radius: var(--radius-sm);
  transition: background 0.2s;
}

.hamburger:hover {
  background: var(--bg-light);
}

.hamburger span {
  display: block;
  width: 100%;
  height: 2px;
  background: var(--text-primary);
  border-radius: 2px;
  transition: all 0.3s ease;
}

.hamburger.open span:nth-child(1) {
  transform: translateY(7px) rotate(45deg);
}

.hamburger.open span:nth-child(2) {
  opacity: 0;
}

.hamburger.open span:nth-child(3) {
  transform: translateY(-7px) rotate(-45deg);
}

/* Mobile Nav */
.nav-mobile {
  position: fixed;
  top: var(--header-height);
  left: 0;
  right: 0;
  bottom: 0;
  background: #fff;
  z-index: 999;
  overflow-y: auto;
}

.nav-mobile-inner {
  display: flex;
  flex-direction: column;
  padding: 24px;
  gap: 4px;
}

.nav-mobile-link {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 16px 20px;
  font-size: 1.05rem;
  font-weight: 500;
  color: var(--text-secondary);
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
}

.nav-mobile-link:hover {
  background: var(--bg-light);
  color: var(--primary);
}

.nav-mobile-link.active {
  background: var(--primary);
  color: #fff;
  font-weight: 700;
  box-shadow: inset 3px 0 0 var(--accent);
}

.nav-mobile-icon {
  font-size: 1.2rem;
  width: 28px;
  text-align: center;
}

.mobile-cta {
  margin-top: 16px;
  padding: 14px;
  text-align: center;
}

/* Transitions */
.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease;
}

.slide-down-enter,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Responsive */
@media (max-width: 767px) {
  .nav-desktop {
    display: none;
  }

  .header-cta {
    display: none;
  }

  .hamburger {
    display: flex;
  }
}

@media (min-width: 768px) {
  .nav-mobile {
    display: none !important;
  }
}
</style>
