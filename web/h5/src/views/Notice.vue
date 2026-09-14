<template>
  <div class="notice-page">
    <div class="notice-content container">
      <div class="notice-icon">{{ icon }}</div>
      <h1 class="notice-title">{{ title }}</h1>
      <p class="notice-desc">{{ desc }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AccessNotice',
  computed: {
    // 后台配置的访问方式（由路由守卫通过 query 传入）
    mode() {
      return this.$route.query.mode || ''
    },
    icon() {
      if (this.mode === 'h5') return '📱'
      if (this.mode === 'pc') return '💻'
      return '🔒'
    },
    title() {
      if (this.mode === 'h5') return 'Mobile Devices Only'
      if (this.mode === 'pc') return 'Desktop Only'
      return 'Access Restricted'
    },
    desc() {
      if (this.mode === 'h5') return 'This site is currently accessible on mobile devices only. Please open it on your phone or tablet.'
      if (this.mode === 'pc') return 'This site is currently accessible on desktop computers only. Please open it on your PC browser.'
      return 'Access to this site is currently restricted. Please try again later.'
    }
  }
}
</script>

<style scoped>
.notice-page {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--header-height) 0 24px;
}

.notice-content {
  width: 100%;
  text-align: center;
  padding: 0 20px;
}

.notice-icon {
  font-size: 5rem;
  line-height: 1;
  margin-bottom: 20px;
}

.notice-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 12px;
}

.notice-desc {
  font-size: 1.05rem;
  color: var(--text-secondary);
  max-width: 480px;
  margin: 0 auto;
  line-height: 1.7;
}

/* 移动端内容较高、居中后易贴边，补充上下留白 */
@media (max-width: 767px) {
  .notice-page {
    padding: calc(var(--header-height) + 40px) 0 48px;
  }
}
</style>
